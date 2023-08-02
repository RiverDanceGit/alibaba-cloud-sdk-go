package avatar

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

// Data is a nested struct in avatar response
type Data struct {
	Name                 string               `json:"Name" xml:"Name"`
	TotalPage            int                  `json:"TotalPage" xml:"TotalPage"`
	TaskUuid             string               `json:"TaskUuid" xml:"TaskUuid"`
	TenantId             int64                `json:"TenantId" xml:"TenantId"`
	Process              string               `json:"Process" xml:"Process"`
	SessionId            string               `json:"SessionId" xml:"SessionId"`
	FailReason           string               `json:"FailReason" xml:"FailReason"`
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	PageNo               int                  `json:"PageNo" xml:"PageNo"`
	MakeStatus           string               `json:"MakeStatus" xml:"MakeStatus"`
	InstanceId           string               `json:"InstanceId" xml:"InstanceId"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	Description          string               `json:"Description" xml:"Description"`
	UniqueCode           string               `json:"UniqueCode" xml:"UniqueCode"`
	Type                 string               `json:"Type" xml:"Type"`
	GrabType             string               `json:"GrabType" xml:"GrabType"`
	Token                string               `json:"Token" xml:"Token"`
	ModelType            string               `json:"ModelType" xml:"ModelType"`
	Code                 string               `json:"Code" xml:"Code"`
	MakeFailReason       string               `json:"MakeFailReason" xml:"MakeFailReason"`
	IsCancel             bool                 `json:"IsCancel" xml:"IsCancel"`
	Image                string               `json:"Image" xml:"Image"`
	AvatarType           string               `json:"AvatarType" xml:"AvatarType"`
	ActionType           string               `json:"ActionType" xml:"ActionType"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	OutputText           string               `json:"OutputText" xml:"OutputText"`
	Portrait             string               `json:"Portrait" xml:"Portrait"`
	Status               string               `json:"Status" xml:"Status"`
	SupportedResolutions SupportedResolutions `json:"SupportedResolutions" xml:"SupportedResolutions"`
	Channel              Channel              `json:"Channel" xml:"Channel"`
	TaskResult           TaskResult           `json:"TaskResult" xml:"TaskResult"`
	List                 []ListItem           `json:"List" xml:"List"`
}
