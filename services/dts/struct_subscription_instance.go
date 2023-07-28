package dts

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// SubscriptionInstance is a nested struct in dts response
type SubscriptionInstance struct {
	Status                   string                                            `json:"Status" xml:"Status"`
	BeginTimestamp           string                                            `json:"BeginTimestamp" xml:"BeginTimestamp"`
	ErrorMessage             string                                            `json:"ErrorMessage" xml:"ErrorMessage"`
	PayType                  string                                            `json:"PayType" xml:"PayType"`
	ConsumptionClient        string                                            `json:"ConsumptionClient" xml:"ConsumptionClient"`
	SubscriptionInstanceName string                                            `json:"SubscriptionInstanceName" xml:"SubscriptionInstanceName"`
	SubscriptionInstanceID   string                                            `json:"SubscriptionInstanceID" xml:"SubscriptionInstanceID"`
	EndTimestamp             string                                            `json:"EndTimestamp" xml:"EndTimestamp"`
	ConsumptionCheckpoint    string                                            `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	SourceEndpoint           SourceEndpoint                                    `json:"SourceEndpoint" xml:"SourceEndpoint"`
	SubscriptionDataType     SubscriptionDataType                              `json:"SubscriptionDataType" xml:"SubscriptionDataType"`
	SubscriptionObject       SubscriptionObjectInDescribeSubscriptionInstances `json:"SubscriptionObject" xml:"SubscriptionObject"`
}
