package hitsdb

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

// Instance is a nested struct in hitsdb response
type Instance struct {
	InstanceId           string `json:"InstanceId" xml:"InstanceId"`
	InstanceAlias        string `json:"InstanceAlias" xml:"InstanceAlias"`
	InstanceDescription  string `json:"InstanceDescription" xml:"InstanceDescription"`
	UserId               string `json:"UserId" xml:"UserId"`
	RegionId             string `json:"RegionId" xml:"RegionId"`
	ZoneId               string `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus       string `json:"InstanceStatus" xml:"InstanceStatus"`
	ChargeType           string `json:"ChargeType" xml:"ChargeType"`
	NetworkType          string `json:"NetworkType" xml:"NetworkType"`
	GmtCreated           string `json:"GmtCreated" xml:"GmtCreated"`
	GmtExpire            string `json:"GmtExpire" xml:"GmtExpire"`
	InstanceClass        string `json:"InstanceClass" xml:"InstanceClass"`
	InstanceStorage      string `json:"InstanceStorage" xml:"InstanceStorage"`
	InstanceTps          string `json:"InstanceTps" xml:"InstanceTps"`
	LockMode             string `json:"LockMode" xml:"LockMode"`
	EngineType           string `json:"EngineType" xml:"EngineType"`
	MaxSeriesPerDatabase string `json:"MaxSeriesPerDatabase" xml:"MaxSeriesPerDatabase"`
	VpcId                string `json:"VpcId" xml:"VpcId"`
	VswitchId            string `json:"VswitchId" xml:"VswitchId"`
	Status               string `json:"Status" xml:"Status"`
	PaymentType          string `json:"PaymentType" xml:"PaymentType"`
	CreateTime           int64  `json:"CreateTime" xml:"CreateTime"`
	ExpiredTime          int64  `json:"ExpiredTime" xml:"ExpiredTime"`
}
