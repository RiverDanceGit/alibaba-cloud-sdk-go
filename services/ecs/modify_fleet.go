package ecs

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

// ModifyFleet invokes the ecs.ModifyFleet API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyfleet.html
func (client *Client) ModifyFleet(request *ModifyFleetRequest) (response *ModifyFleetResponse, err error) {
	response = CreateModifyFleetResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFleetWithChan invokes the ecs.ModifyFleet API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyfleet.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFleetWithChan(request *ModifyFleetRequest) (<-chan *ModifyFleetResponse, <-chan error) {
	responseChan := make(chan *ModifyFleetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFleet(request)
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

// ModifyFleetWithCallback invokes the ecs.ModifyFleet API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyfleet.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFleetWithCallback(request *ModifyFleetRequest, callback func(response *ModifyFleetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFleetResponse
		var err error
		defer close(result)
		response, err = client.ModifyFleet(request)
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

// ModifyFleetRequest is the request struct for api ModifyFleet
type ModifyFleetRequest struct {
	*requests.RpcRequest
}

// ModifyFleetResponse is the response struct for api ModifyFleet
type ModifyFleetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyFleetRequest creates a request to invoke ModifyFleet API
func CreateModifyFleetRequest() (request *ModifyFleetRequest) {
	request = &ModifyFleetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyFleet", "ecs", "openAPI")
	return
}

// CreateModifyFleetResponse creates a response to parse from ModifyFleet response
func CreateModifyFleetResponse() (response *ModifyFleetResponse) {
	response = &ModifyFleetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
