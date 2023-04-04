package events

// Autogenerated file. DO NOT MODIFY DIRECTLY!
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

var EventsMap = map[EventName]*EventSchema{
	"ExampleEvent": {
		Name:                "ExampleEvent",
		Reason:              "ExampleEvent",
		Action:              "ExampleEvent",
		Type:                EventTypeWarning,
		ReportingController: "controller",
		Message:             "ExampleEvent message.",
	},
	"ClusterNodeIpUpdated": {
		Name:                "ClusterNodeIpUpdated",
		Reason:              "ClusterNodeIpUpdated",
		Action:              "None",
		Type:                EventTypeNormal,
		ReportingController: "worker",
		Message:             "Successfully updated cluster as change detected in cluster nodes",
	},
	"ClusterNodeIpUpdateFail": {
		Name:                "ClusterNodeIpUpdateFail",
		Reason:              "ClusterNodeIpUpdateFail",
		Action:              "None",
		Type:                EventTypeWarning,
		ReportingController: "worker",
		Message:             "Failed to update node ip in cluster CR",
	},
}

var (
	EventExampleEvent EventName = "ExampleEvent"
)