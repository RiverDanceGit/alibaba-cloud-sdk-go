package ccc

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SignInGroup invokes the ccc.SignInGroup API synchronously
func (client *Client) SignInGroup(request *SignInGroupRequest) (response *SignInGroupResponse, err error) {
	response = CreateSignInGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SignInGroupWithChan invokes the ccc.SignInGroup API asynchronously
func (client *Client) SignInGroupWithChan(request *SignInGroupRequest) (<-chan *SignInGroupResponse, <-chan error) {
	responseChan := make(chan *SignInGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SignInGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SignInGroupWithCallback invokes the ccc.SignInGroup API asynchronously
func (client *Client) SignInGroupWithCallback(request *SignInGroupRequest, callback func(response *SignInGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SignInGroupResponse
		var err error
		defer close(result)
		response, err = client.SignInGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SignInGroupRequest is the request struct for api SignInGroup
type SignInGroupRequest struct {
	*requests.RpcRequest
	SignedSkillGroupIdList string `position:"Query" name:"SignedSkillGroupIdList"`
	UserId                 string `position:"Query" name:"UserId"`
	DeviceId               string `position:"Query" name:"DeviceId"`
	InstanceId             string `position:"Query" name:"InstanceId"`
}

// SignInGroupResponse is the response struct for api SignInGroup
type SignInGroupResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateSignInGroupRequest creates a request to invoke SignInGroup API
func CreateSignInGroupRequest() (request *SignInGroupRequest) {
	request = &SignInGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "SignInGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateSignInGroupResponse creates a response to parse from SignInGroup response
func CreateSignInGroupResponse() (response *SignInGroupResponse) {
	response = &SignInGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
