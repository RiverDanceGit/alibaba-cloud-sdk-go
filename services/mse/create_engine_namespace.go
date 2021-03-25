package mse

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

// CreateEngineNamespace invokes the mse.CreateEngineNamespace API synchronously
func (client *Client) CreateEngineNamespace(request *CreateEngineNamespaceRequest) (response *CreateEngineNamespaceResponse, err error) {
	response = CreateCreateEngineNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEngineNamespaceWithChan invokes the mse.CreateEngineNamespace API asynchronously
func (client *Client) CreateEngineNamespaceWithChan(request *CreateEngineNamespaceRequest) (<-chan *CreateEngineNamespaceResponse, <-chan error) {
	responseChan := make(chan *CreateEngineNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEngineNamespace(request)
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

// CreateEngineNamespaceWithCallback invokes the mse.CreateEngineNamespace API asynchronously
func (client *Client) CreateEngineNamespaceWithCallback(request *CreateEngineNamespaceRequest, callback func(response *CreateEngineNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEngineNamespaceResponse
		var err error
		defer close(result)
		response, err = client.CreateEngineNamespace(request)
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

// CreateEngineNamespaceRequest is the request struct for api CreateEngineNamespace
type CreateEngineNamespaceRequest struct {
	*requests.RpcRequest
	ClusterId    string           `position:"Query" name:"ClusterId"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	ServiceCount requests.Integer `position:"Query" name:"ServiceCount"`
	Name         string           `position:"Query" name:"Name"`
	Desc         string           `position:"Query" name:"Desc"`
}

// CreateEngineNamespaceResponse is the response struct for api CreateEngineNamespace
type CreateEngineNamespaceResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateEngineNamespaceRequest creates a request to invoke CreateEngineNamespace API
func CreateCreateEngineNamespaceRequest() (request *CreateEngineNamespaceRequest) {
	request = &CreateEngineNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateEngineNamespace", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateEngineNamespaceResponse creates a response to parse from CreateEngineNamespace response
func CreateCreateEngineNamespaceResponse() (response *CreateEngineNamespaceResponse) {
	response = &CreateEngineNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
