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

// ListSipCallRecords invokes the ccc.ListSipCallRecords API synchronously
func (client *Client) ListSipCallRecords(request *ListSipCallRecordsRequest) (response *ListSipCallRecordsResponse, err error) {
	response = CreateListSipCallRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSipCallRecordsWithChan invokes the ccc.ListSipCallRecords API asynchronously
func (client *Client) ListSipCallRecordsWithChan(request *ListSipCallRecordsRequest) (<-chan *ListSipCallRecordsResponse, <-chan error) {
	responseChan := make(chan *ListSipCallRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSipCallRecords(request)
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

// ListSipCallRecordsWithCallback invokes the ccc.ListSipCallRecords API asynchronously
func (client *Client) ListSipCallRecordsWithCallback(request *ListSipCallRecordsRequest, callback func(response *ListSipCallRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSipCallRecordsResponse
		var err error
		defer close(result)
		response, err = client.ListSipCallRecords(request)
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

// ListSipCallRecordsRequest is the request struct for api ListSipCallRecords
type ListSipCallRecordsRequest struct {
	*requests.RpcRequest
	InstanceId    string `position:"Query" name:"InstanceId"`
	ContactIdList string `position:"Query" name:"ContactIdList"`
}

// ListSipCallRecordsResponse is the response struct for api ListSipCallRecords
type ListSipCallRecordsResponse struct {
	*responses.BaseResponse
	Code           string           `json:"Code" xml:"Code"`
	HttpStatusCode int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string           `json:"Message" xml:"Message"`
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	Data           []SipCallRecords `json:"Data" xml:"Data"`
}

// CreateListSipCallRecordsRequest creates a request to invoke ListSipCallRecords API
func CreateListSipCallRecordsRequest() (request *ListSipCallRecordsRequest) {
	request = &ListSipCallRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListSipCallRecords", "", "")
	request.Method = requests.POST
	return
}

// CreateListSipCallRecordsResponse creates a response to parse from ListSipCallRecords response
func CreateListSipCallRecordsResponse() (response *ListSipCallRecordsResponse) {
	response = &ListSipCallRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
