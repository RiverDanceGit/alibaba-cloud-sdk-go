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

// ListIvrTrackingDetails invokes the ccc.ListIvrTrackingDetails API synchronously
func (client *Client) ListIvrTrackingDetails(request *ListIvrTrackingDetailsRequest) (response *ListIvrTrackingDetailsResponse, err error) {
	response = CreateListIvrTrackingDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListIvrTrackingDetailsWithChan invokes the ccc.ListIvrTrackingDetails API asynchronously
func (client *Client) ListIvrTrackingDetailsWithChan(request *ListIvrTrackingDetailsRequest) (<-chan *ListIvrTrackingDetailsResponse, <-chan error) {
	responseChan := make(chan *ListIvrTrackingDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIvrTrackingDetails(request)
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

// ListIvrTrackingDetailsWithCallback invokes the ccc.ListIvrTrackingDetails API asynchronously
func (client *Client) ListIvrTrackingDetailsWithCallback(request *ListIvrTrackingDetailsRequest, callback func(response *ListIvrTrackingDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIvrTrackingDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListIvrTrackingDetails(request)
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

// ListIvrTrackingDetailsRequest is the request struct for api ListIvrTrackingDetails
type ListIvrTrackingDetailsRequest struct {
	*requests.RpcRequest
	ContactId  string           `position:"Query" name:"ContactId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListIvrTrackingDetailsResponse is the response struct for api ListIvrTrackingDetails
type ListIvrTrackingDetailsResponse struct {
	*responses.BaseResponse
	Code           string                       `json:"Code" xml:"Code"`
	HttpStatusCode int                          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                       `json:"Message" xml:"Message"`
	RequestId      string                       `json:"RequestId" xml:"RequestId"`
	Data           DataInListIvrTrackingDetails `json:"Data" xml:"Data"`
}

// CreateListIvrTrackingDetailsRequest creates a request to invoke ListIvrTrackingDetails API
func CreateListIvrTrackingDetailsRequest() (request *ListIvrTrackingDetailsRequest) {
	request = &ListIvrTrackingDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListIvrTrackingDetails", "", "")
	request.Method = requests.POST
	return
}

// CreateListIvrTrackingDetailsResponse creates a response to parse from ListIvrTrackingDetails response
func CreateListIvrTrackingDetailsResponse() (response *ListIvrTrackingDetailsResponse) {
	response = &ListIvrTrackingDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
