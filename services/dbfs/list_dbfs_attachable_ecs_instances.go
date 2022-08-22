package dbfs

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

// ListDbfsAttachableEcsInstances invokes the dbfs.ListDbfsAttachableEcsInstances API synchronously
func (client *Client) ListDbfsAttachableEcsInstances(request *ListDbfsAttachableEcsInstancesRequest) (response *ListDbfsAttachableEcsInstancesResponse, err error) {
	response = CreateListDbfsAttachableEcsInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDbfsAttachableEcsInstancesWithChan invokes the dbfs.ListDbfsAttachableEcsInstances API asynchronously
func (client *Client) ListDbfsAttachableEcsInstancesWithChan(request *ListDbfsAttachableEcsInstancesRequest) (<-chan *ListDbfsAttachableEcsInstancesResponse, <-chan error) {
	responseChan := make(chan *ListDbfsAttachableEcsInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDbfsAttachableEcsInstances(request)
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

// ListDbfsAttachableEcsInstancesWithCallback invokes the dbfs.ListDbfsAttachableEcsInstances API asynchronously
func (client *Client) ListDbfsAttachableEcsInstancesWithCallback(request *ListDbfsAttachableEcsInstancesRequest, callback func(response *ListDbfsAttachableEcsInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDbfsAttachableEcsInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListDbfsAttachableEcsInstances(request)
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

// ListDbfsAttachableEcsInstancesRequest is the request struct for api ListDbfsAttachableEcsInstances
type ListDbfsAttachableEcsInstancesRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListDbfsAttachableEcsInstancesResponse is the response struct for api ListDbfsAttachableEcsInstances
type ListDbfsAttachableEcsInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	TotalCount   int         `json:"TotalCount" xml:"TotalCount"`
	EcsLabelInfo []LabelInfo `json:"EcsLabelInfo" xml:"EcsLabelInfo"`
}

// CreateListDbfsAttachableEcsInstancesRequest creates a request to invoke ListDbfsAttachableEcsInstances API
func CreateListDbfsAttachableEcsInstancesRequest() (request *ListDbfsAttachableEcsInstancesRequest) {
	request = &ListDbfsAttachableEcsInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "ListDbfsAttachableEcsInstances", "dbfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDbfsAttachableEcsInstancesResponse creates a response to parse from ListDbfsAttachableEcsInstances response
func CreateListDbfsAttachableEcsInstancesResponse() (response *ListDbfsAttachableEcsInstancesResponse) {
	response = &ListDbfsAttachableEcsInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
