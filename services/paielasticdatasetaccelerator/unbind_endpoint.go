package paielasticdatasetaccelerator

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

// UnbindEndpoint invokes the paielasticdatasetaccelerator.UnbindEndpoint API synchronously
func (client *Client) UnbindEndpoint(request *UnbindEndpointRequest) (response *UnbindEndpointResponse, err error) {
	response = CreateUnbindEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindEndpointWithChan invokes the paielasticdatasetaccelerator.UnbindEndpoint API asynchronously
func (client *Client) UnbindEndpointWithChan(request *UnbindEndpointRequest) (<-chan *UnbindEndpointResponse, <-chan error) {
	responseChan := make(chan *UnbindEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindEndpoint(request)
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

// UnbindEndpointWithCallback invokes the paielasticdatasetaccelerator.UnbindEndpoint API asynchronously
func (client *Client) UnbindEndpointWithCallback(request *UnbindEndpointRequest, callback func(response *UnbindEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindEndpointResponse
		var err error
		defer close(result)
		response, err = client.UnbindEndpoint(request)
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

// UnbindEndpointRequest is the request struct for api UnbindEndpoint
type UnbindEndpointRequest struct {
	*requests.RoaRequest
	EndpointId string `position:"Path" name:"EndpointId"`
	SlotId     string `position:"Path" name:"SlotId"`
}

// UnbindEndpointResponse is the response struct for api UnbindEndpoint
type UnbindEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindEndpointRequest creates a request to invoke UnbindEndpoint API
func CreateUnbindEndpointRequest() (request *UnbindEndpointRequest) {
	request = &UnbindEndpointRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "UnbindEndpoint", "/api/v1/endpoints/[EndpointId]/slots/[SlotId]", "datasetacc", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateUnbindEndpointResponse creates a response to parse from UnbindEndpoint response
func CreateUnbindEndpointResponse() (response *UnbindEndpointResponse) {
	response = &UnbindEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
