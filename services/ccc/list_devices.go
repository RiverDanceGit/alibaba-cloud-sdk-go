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

// ListDevices invokes the ccc.ListDevices API synchronously
func (client *Client) ListDevices(request *ListDevicesRequest) (response *ListDevicesResponse, err error) {
	response = CreateListDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevicesWithChan invokes the ccc.ListDevices API asynchronously
func (client *Client) ListDevicesWithChan(request *ListDevicesRequest) (<-chan *ListDevicesResponse, <-chan error) {
	responseChan := make(chan *ListDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevices(request)
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

// ListDevicesWithCallback invokes the ccc.ListDevices API asynchronously
func (client *Client) ListDevicesWithCallback(request *ListDevicesRequest, callback func(response *ListDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevicesResponse
		var err error
		defer close(result)
		response, err = client.ListDevices(request)
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

// ListDevicesRequest is the request struct for api ListDevices
type ListDevicesRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListDevicesResponse is the response struct for api ListDevices
type ListDevicesResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           []Device `json:"Data" xml:"Data"`
}

// CreateListDevicesRequest creates a request to invoke ListDevices API
func CreateListDevicesRequest() (request *ListDevicesRequest) {
	request = &ListDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateListDevicesResponse creates a response to parse from ListDevices response
func CreateListDevicesResponse() (response *ListDevicesResponse) {
	response = &ListDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
