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

// RemovePhoneNumbers invokes the ccc.RemovePhoneNumbers API synchronously
func (client *Client) RemovePhoneNumbers(request *RemovePhoneNumbersRequest) (response *RemovePhoneNumbersResponse, err error) {
	response = CreateRemovePhoneNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// RemovePhoneNumbersWithChan invokes the ccc.RemovePhoneNumbers API asynchronously
func (client *Client) RemovePhoneNumbersWithChan(request *RemovePhoneNumbersRequest) (<-chan *RemovePhoneNumbersResponse, <-chan error) {
	responseChan := make(chan *RemovePhoneNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemovePhoneNumbers(request)
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

// RemovePhoneNumbersWithCallback invokes the ccc.RemovePhoneNumbers API asynchronously
func (client *Client) RemovePhoneNumbersWithCallback(request *RemovePhoneNumbersRequest, callback func(response *RemovePhoneNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemovePhoneNumbersResponse
		var err error
		defer close(result)
		response, err = client.RemovePhoneNumbers(request)
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

// RemovePhoneNumbersRequest is the request struct for api RemovePhoneNumbers
type RemovePhoneNumbersRequest struct {
	*requests.RpcRequest
	NumberList string `position:"Query" name:"NumberList"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// RemovePhoneNumbersResponse is the response struct for api RemovePhoneNumbers
type RemovePhoneNumbersResponse struct {
	*responses.BaseResponse
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	Data           string   `json:"Data" xml:"Data"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	FailureList    []string `json:"FailureList" xml:"FailureList"`
	Params         []string `json:"Params" xml:"Params"`
}

// CreateRemovePhoneNumbersRequest creates a request to invoke RemovePhoneNumbers API
func CreateRemovePhoneNumbersRequest() (request *RemovePhoneNumbersRequest) {
	request = &RemovePhoneNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "RemovePhoneNumbers", "", "")
	request.Method = requests.POST
	return
}

// CreateRemovePhoneNumbersResponse creates a response to parse from RemovePhoneNumbers response
func CreateRemovePhoneNumbersResponse() (response *RemovePhoneNumbersResponse) {
	response = &RemovePhoneNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
