/*
 *  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package slice

import (
	"context"
	"fmt"
	"log"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	retry "k8s.io/client-go/util/retry"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/go-logr/logr"
	controllerv1alpha1 "github.com/kubeslice/apis/pkg/controller/v1alpha1"
	"github.com/kubeslice/kubeslice-monitoring/pkg/events"
	"github.com/kubeslice/kubeslice-monitoring/pkg/metrics"
	kubeslicev1beta1 "github.com/kubeslice/worker-operator/api/v1beta1"
	"github.com/kubeslice/worker-operator/controllers"
	ossEvents "github.com/kubeslice/worker-operator/events"
	"github.com/kubeslice/worker-operator/pkg/logger"
	"github.com/kubeslice/worker-operator/pkg/manifest"
	"github.com/kubeslice/worker-operator/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
)

const VPC_NS_FMT = "%s-vpc-access-gw-system"

var (
	sliceFinalizer = "networking.kubeslice.io/slice-finalizer"
	controllerName = "sliceReconciler"
)

// SliceReconciler reconciles a Slice object
type SliceReconciler struct {
	client.Client
	EventRecorder           *events.EventRecorder
	Scheme                  *runtime.Scheme
	Log                     logr.Logger
	NetOpPods               []NetOpPod
	HubClient               HubClientProvider
	WorkerRouterClient      WorkerRouterClientProvider
	WorkerNetOpClient       WorkerNetOpClientProvider
	WorkerGatewayEdgeClient WorkerGatewayEdgeClientProvider

	// metrics
	gaugeAppPods *prometheus.GaugeVec
}

//+kubebuilder:rbac:groups=networking.kubeslice.io,resources=slices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.kubeslice.io,resources=slices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking.kubeslice.io,resources=slices/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.istio.io,resources=gateways,verbs=get;list;create;update;watch;delete
//+kubebuilder:rbac:groups=networking.istio.io,resources=serviceentries,verbs=get;list;create;update;watch;delete
//+kubebuilder:rbac:groups=networking.istio.io,resources=virtualservices,verbs=get;list;create;update;watch;delete
//+kubebuilder:webhook:path=/mutate-webhook,mutating=true,failurePolicy=fail,groups="";apps,resources=pods;deployments;statefulsets;daemonsets,verbs=create;update,versions=v1,name=webhook.kubeslice.io,admissionReviewVersions=v1,sideEffects=NoneOnDryRun
//+kubebuilder:webhook:path=/validate-webhook,mutating=false,failurePolicy=fail,groups="networking.kubeslice.io",resources=serviceexports,verbs=create;update,versions=v1beta1,name=webhook.kubeslice.io,admissionReviewVersions=v1,sideEffects=NoneOnDryRun
//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

func (r *SliceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("slice", req.NamespacedName)
	debugLog := log.V(1)
	ctx = logger.WithLogger(ctx, log)

	slice := &kubeslicev1beta1.Slice{}

	err := r.Get(ctx, req.NamespacedName, slice)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Return and don't requeue
			log.Info("Slice resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Slice")
		return ctrl.Result{}, err
	}

	log.Info("reconciling", "slice", slice.Name)

	*r.EventRecorder = (*r.EventRecorder).WithSlice(slice.Name)
	// label kubeslice-system namespace with kubeslice.io/inject=true label
	namespace := &corev1.Namespace{}
	err = r.Get(ctx, types.NamespacedName{Name: "kubeslice-system"}, namespace)
	if err != nil {
		return ctrl.Result{}, err
	}
	// A namespace might not have any labels attached to it. Directly accessing the label map
	// leads to a crash for such namespaces.
	// If the label map is nil, create one and use the setter api to attach it to the namespace.
	nsLabels := namespace.ObjectMeta.GetLabels()
	if nsLabels == nil {
		nsLabels = make(map[string]string)
	}
	if _, ok := nsLabels[InjectSidecarKey]; !ok {
		nsLabels[InjectSidecarKey] = "true"
		namespace.ObjectMeta.SetLabels(nsLabels)

		err = r.Update(ctx, namespace)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	// Examine DeletionTimestamp to determine if object is under deletion
	// The object is not being deleted, so if it does not have our finalizer,
	// then lets add the finalizer and update the object. This is equivalent
	// registering our finalizer.
	// The object is being deleted
	// send slice deletion event to netops
	//cleanup slice resources
	// Stop reconciliation as the item is being deleted
	requeue, result, err := r.handleSliceDeletion(slice, ctx, req)
	if requeue {
		return result, err
	}

	if slice.Status.SliceConfig == nil {
		err := fmt.Errorf("slice not reconciled from hub")
		log.Error(err, "Slice is not reconciled from hub yet, skipping reconciliation")
		return ctrl.Result{}, err
	}

	if slice.Status.SliceConfig.SliceOverlayNetworkDeploymentMode != controllerv1alpha1.NONET {
		if slice.Status.DNSIP == "" {
			requeue, result, err := r.handleDnsSvc(ctx, slice)
			if requeue {
				return result, err
			}
		}
	}

	res, err, requeue := r.ReconcileSliceNamespaces(ctx, slice)
	if requeue {
		debugLog.Info("Reconciling SliceNamespaces", "res", res, "err", err)
		return res, err
	}

	if slice.Status.SliceConfig.SliceOverlayNetworkDeploymentMode == controllerv1alpha1.NONET {
		debugLog.Info("No communication slice, skipping reconciliation of qos, netop, egw, router etc")
		// to support net to no-net switching write a function to delete network components if present
	} else {
		debugLog.Info("Slice with network, continue reconciliation of qos, netop, egw, router etc")
		// syncQoStoNetop, reconcile slice router, slicegw edge, ext gateways
		res, err, requeue := r.reconcileNetworkComponents(ctx, slice)
		if requeue {
			debugLog.Info("Retry reconciling SliceNetworkComponents", "res", res, "err", err)
			return res, err
		}
	}

	if isIngressConfigured(slice) {
		debugLog.Info("Installing ingress")
		err = manifest.InstallIngress(ctx, r.Client, slice)
		if err != nil {
			log.Error(err, "unable to install ingress")
			utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceIngressInstallFailed, controllerName)
			return ctrl.Result{}, nil
		}
	}

	res, err, requeue = r.ReconcileSliceRouter(ctx, slice)
	if err != nil {
		log.Error(err, "Failed to reconcile slice router")
	}
	if requeue {
		return res, err
	}

	res, err, requeue = r.ReconcileSliceGwEdge(ctx, slice)
	if err != nil {
		log.Error(err, "Slice Edge reconciliation failed")
		return res, err
	}
	if requeue {
		return ctrl.Result{
			Requeue: true,
		}, nil
	}

	debugLog.Info("reconciling app pods")
	res, err, requeue = r.ReconcileAppPod(ctx, slice)
	if err != nil {
		log.Error(err, "App pod reconciliation failed")
		return res, err
	}

	if requeue {
		// reconciliation success, update the app pod list in controller
		log.Info("updating app pod list in hub workersliceconfig status")
		sliceConfigName := slice.Name + "-" + controllers.ClusterName
		if err = r.HubClient.UpdateAppPodsList(ctx, sliceConfigName, slice.Status.AppPods); err != nil {
			log.Error(err, "Failed to update app pod list in hub")
			utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceAppPodsListUpdateFailed, controllerName)
			return ctrl.Result{}, err
		}
		return ctrl.Result{
			Requeue: true,
		}, nil
	}
	// expose the number of app pods metric of a slice
	r.exposeMetric(slice.Status.AppPods, slice)

	return ctrl.Result{
		RequeueAfter: controllers.ReconcileInterval,
	}, nil
}

func (r *SliceReconciler) reconcileNetworkComponents(ctx context.Context, slice *kubeslicev1beta1.Slice) (reconcile.Result, error, bool) {
	log := logger.FromContext(ctx)
	debugLog := log.V(1)

	debugLog.Info("Syncing slice QoS config with NetOp pods")
	err := r.SyncSliceQosProfileWithNetOp(ctx, slice)
	if err != nil {
		log.Error(err, "Failed to sync QoS profile with netop pods")
		utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceQoSProfileWithNetOpsSync, controllerName)
	} else {
		utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceUpdated, controllerName)
	}

	debugLog.Info("reconciling SliceRouter")
	res, err, requeue := r.ReconcileSliceRouter(ctx, slice)
	if err != nil {
		log.Error(err, "Failed to reconcile slice router")
		return res, err, true
	}
	if requeue {
		return res, nil, true
	}

	debugLog.Info("reconciling SliceGwEdge")
	res, err, requeue = r.ReconcileSliceGwEdge(ctx, slice)
	if err != nil {
		log.Error(err, "Slice Edge reconciliation failed")
		return res, err, true
	}
	if requeue {
		return res, nil, true
	}

	debugLog.Info("ExternalGatewayConfig", "obj", slice.Status.SliceConfig.ExternalGatewayConfig)
	if slice.Status.SliceConfig.ExternalGatewayConfig != nil &&
		slice.Status.SliceConfig.ExternalGatewayConfig.GatewayType == controllerv1alpha1.GATEWAY_TYPE_ISTIO {
		if isEgressConfigured(slice) {
			debugLog.Info("Installing istio egress")
			err = manifest.InstallEgress(ctx, r.Client, slice)
			if err != nil {
				log.Error(err, "unable to install egress")
				utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceEgressInstallFailed, controllerName)
				return ctrl.Result{}, err, false
			}
		}

		if isIngressConfigured(slice) {
			debugLog.Info("Installing istio ingress")
			err = manifest.InstallIngress(ctx, r.Client, slice)
			if err != nil {
				log.Error(err, "unable to install ingress")
				utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceIngressInstallFailed, controllerName)
				return ctrl.Result{}, err, false
			}
		}
	}

	return ctrl.Result{}, nil, false
}

func (r *SliceReconciler) exposeMetric(appPods []kubeslicev1beta1.AppPod, slice *kubeslicev1beta1.Slice) {
	// this extra check is needed when current app pods are zero and slice status has old app pods
	// then it doesn't goes to app pods loop hence it don't update the app pods metrics count to zero
	// Set no. of app pods in prometheus metrics
	if len(appPods) == 0 && len(slice.Status.AppPods) > 0 {
		for _, appPod := range slice.Status.AppPods {
			r.gaugeAppPods.WithLabelValues(slice.Name, appPod.PodNamespace).Set(0)
		}
	}
	mapAppPodsPerNamespace := make(map[string][]kubeslicev1beta1.AppPod)
	for _, appPod := range appPods {
		mapAppPodsPerNamespace[appPod.PodNamespace] = append(mapAppPodsPerNamespace[appPod.PodNamespace], appPod)
	}

	for namespace, pods := range mapAppPodsPerNamespace {
		r.gaugeAppPods.WithLabelValues(slice.Name, namespace).Set(float64(len(pods)))
	}
}

func isEgressConfigured(slice *kubeslicev1beta1.Slice) bool {
	return slice.Status.SliceConfig.ExternalGatewayConfig != nil && slice.Status.SliceConfig.ExternalGatewayConfig.Egress.Enabled
}
func isIngressConfigured(slice *kubeslicev1beta1.Slice) bool {
	return slice.Status.SliceConfig.ExternalGatewayConfig != nil && slice.Status.SliceConfig.ExternalGatewayConfig.Ingress.Enabled
}
func (r *SliceReconciler) handleDnsSvc(ctx context.Context, slice *kubeslicev1beta1.Slice) (bool, reconcile.Result, error) {
	log := logger.FromContext(ctx).WithName("slice-dns-svc")
	debugLog := log.V(1)
	log.Info("Finding DNS IP")
	svc := &corev1.Service{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: controllers.ControlPlaneNamespace,
		Name:      controllers.DNSDeploymentName,
	}, svc)

	if err != nil {
		if errors.IsNotFound(err) {
			debugLog.Info("DNS service not found in the cluster, probably coredns is not deployed for no-net slice; continuing")
		} else {
			log.Error(err, "Failed to get DNS Service")
			return true, ctrl.Result{}, err
		}
	} else {
		debugLog.Info("got dns service", "svc", svc)
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			// Fetch latest slice obj
			if getErr := r.Get(ctx, types.NamespacedName{Name: slice.Name, Namespace: controllers.ControlPlaneNamespace}, slice); getErr != nil {
				log.Error(err, "Unable to fetch slice during retry", "slice", slice.Name)
				return getErr
			}
			slice.Status.DNSIP = svc.Spec.ClusterIP
			err := r.Status().Update(ctx, slice)
			if err != nil {
				utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceUpdateFailed, controllerName)
				return err
			}
			utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceUpdated, controllerName)
			return nil
		})
		if err != nil {
			log.Error(err, "Failed to update Slice status for dns")
			return true, ctrl.Result{}, err
		}
		return true, ctrl.Result{}, nil
	}
	return false, reconcile.Result{}, nil
}

func (r *SliceReconciler) handleSliceDeletion(slice *kubeslicev1beta1.Slice, ctx context.Context, req reconcile.Request) (bool, reconcile.Result, error) {
	log := logger.FromContext(ctx).WithName("slice-deletion")
	// Examine DeletionTimestamp to determine if object is under deletion
	// The object is not being deleted, so if it does not have our finalizer,
	// then lets add the finalizer and update the object. This is equivalent
	// registering our finalizer.
	// The object is being deleted
	// send slice deletion event to netops
	//cleanup slice resources
	// Stop reconciliation as the item is being deleted
	if slice.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(slice, sliceFinalizer) {
			controllerutil.AddFinalizer(slice, sliceFinalizer)
			if err := r.Update(ctx, slice); err != nil {
				return true, ctrl.Result{}, err
			}
		}
	} else {
		if controllerutil.ContainsFinalizer(slice, sliceFinalizer) {
			log.Info("Deleting slice", "slice", slice.Name)
			if slice.Status.SliceConfig != nil &&
				slice.Status.SliceConfig.SliceOverlayNetworkDeploymentMode != controllerv1alpha1.NONET {
				err := r.SendSliceDeletionEventToNetOp(ctx, req.NamespacedName.Name, req.NamespacedName.Namespace)
				if err != nil {
					log.Error(err, "Failed to send slice deletetion event to netop")
				}
			}
			if err := r.cleanupSliceResources(ctx, slice); err != nil {
				log.Error(err, "error while deleting slice")
				utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceDeletionFailed, controllerName)
				return true, ctrl.Result{}, err
			}
			controllerutil.RemoveFinalizer(slice, sliceFinalizer)
			if err := r.Update(ctx, slice); err != nil {
				return true, ctrl.Result{}, err
			}
		}
		utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceDeleted, controllerName)
		return true, ctrl.Result{}, nil
	}
	return false, reconcile.Result{}, nil
}

// Setup SliceReconciler
// Initializes metrics and sets up with manager
func (r *SliceReconciler) Setup(mgr ctrl.Manager, mf metrics.MetricsFactory) error {
	gaugeAppPods := mf.NewGauge("app_pods", "App pods in slice", []string{"slice", "slice_namespace"})

	r.gaugeAppPods = gaugeAppPods

	return r.SetupWithManager(mgr)
}

// SetupWithManager sets up the controller with the Manager.
func (r *SliceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Create a label selector that matches based on the existence of a label key
	sliceSelector := labels.NewSelector()
	requirement, err := labels.NewRequirement(controllers.ApplicationNamespaceSelectorLabelKey, selection.Exists, nil)
	if err != nil {
		log.Fatalf("Error creating label requirement: %v", err)
	}
	sliceSelector = sliceSelector.Add(*requirement)
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubeslicev1beta1.Slice{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Pod{}).
		Owns(&kubeslicev1beta1.SliceGateway{}).
		Watches(
			&appsv1.Deployment{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, o client.Object) (recs []reconcile.Request) {
				log := logger.FromContext(ctx)
				debuglog := log.V(1)
				debuglog.Info("Triggered slice reconciler by", "type", reflect.TypeOf(o))
				sliceName := o.(*appsv1.Deployment).Labels[controllers.ApplicationNamespaceSelectorLabelKey]
				recs = append(recs, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Name:      sliceName,
						Namespace: controllers.ControlPlaneNamespace,
					},
				})
				debuglog.Info("Requeuing slice", "name", sliceName)
				return
			}),
			builder.WithPredicates(predicate.Funcs{
				CreateFunc: func(e event.CreateEvent) bool {
					return false
				},
				DeleteFunc: func(e event.DeleteEvent) bool {
					return sliceSelector.Matches(labels.Set(e.Object.GetLabels()))
				},
				UpdateFunc: func(e event.UpdateEvent) bool {
					if sliceSelector.Matches(labels.Set(e.ObjectOld.GetLabels())) {
						oldObj, ok := e.ObjectOld.(*appsv1.Deployment)
						if !ok {
							return false
						}
						newObj, ok := e.ObjectNew.(*appsv1.Deployment)
						if !ok {
							return false
						}
						// trigger in case of scale down
						if oldObj.Status.ReadyReplicas > newObj.Status.ReadyReplicas {
							return true
						}
					}
					return false
				},
				GenericFunc: func(e event.GenericEvent) bool {
					return false
				},
			}),
		).
		Watches(
			&appsv1.DaemonSet{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, o client.Object) (recs []reconcile.Request) {
				log := logger.FromContext(ctx)
				debuglog := log.V(1)
				debuglog.Info("Triggered slice reconciler by", "type", reflect.TypeOf(o))
				sliceName := o.(*appsv1.DaemonSet).Labels[controllers.ApplicationNamespaceSelectorLabelKey]
				recs = append(recs, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Name:      sliceName,
						Namespace: controllers.ControlPlaneNamespace,
					},
				})
				debuglog.Info("Requeuing slice", "name", sliceName)
				return
			}),
			builder.WithPredicates(predicate.Funcs{
				CreateFunc: func(e event.CreateEvent) bool {
					return false
				},
				DeleteFunc: func(e event.DeleteEvent) bool {
					return sliceSelector.Matches(labels.Set(e.Object.GetLabels()))
				},
				UpdateFunc: func(e event.UpdateEvent) bool {
					if sliceSelector.Matches(labels.Set(e.ObjectOld.GetLabels())) {
						oldObj, ok := e.ObjectOld.(*appsv1.DaemonSet)
						if !ok {
							return false
						}
						newObj, ok := e.ObjectNew.(*appsv1.DaemonSet)
						if !ok {
							return false
						}
						if oldObj.Status.NumberReady > newObj.Status.NumberReady {
							return true
						}
					}
					return false
				},
				GenericFunc: func(e event.GenericEvent) bool {
					return false
				},
			}),
		).
		Watches(
			&appsv1.StatefulSet{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, o client.Object) (recs []reconcile.Request) {
				log := logger.FromContext(ctx)
				debuglog := log.V(1)
				debuglog.Info("Triggered slice reconciler by", "type", reflect.TypeOf(o))
				sliceName := o.(*appsv1.StatefulSet).Labels[controllers.ApplicationNamespaceSelectorLabelKey]
				recs = append(recs, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Name:      sliceName,
						Namespace: controllers.ControlPlaneNamespace,
					},
				})
				debuglog.Info("Requeuing slice", "name", sliceName)
				return
			}),
			builder.WithPredicates(predicate.Funcs{
				CreateFunc: func(e event.CreateEvent) bool {
					return false
				},
				DeleteFunc: func(e event.DeleteEvent) bool {
					return sliceSelector.Matches(labels.Set(e.Object.GetLabels()))
				},
				UpdateFunc: func(e event.UpdateEvent) bool {
					if sliceSelector.Matches(labels.Set(e.ObjectOld.GetLabels())) {
						oldObj, ok := e.ObjectOld.(*appsv1.StatefulSet)
						if !ok {
							return false
						}
						newObj, ok := e.ObjectNew.(*appsv1.StatefulSet)
						if !ok {
							return false
						}
						if oldObj.Status.ReadyReplicas > newObj.Status.ReadyReplicas {
							return true
						}
					}
					return false
				},
				GenericFunc: func(e event.GenericEvent) bool {
					return false
				},
			}),
		).
		Complete(r)
}
