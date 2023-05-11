package alinlp

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

// RequestDialogApi invokes the alinlp.RequestDialogApi API synchronously
func (client *Client) RequestDialogApi(request *RequestDialogApiRequest) (response *RequestDialogApiResponse, err error) {
	response = CreateRequestDialogApiResponse()
	err = client.DoAction(request, response)
	return
}

// RequestDialogApiWithChan invokes the alinlp.RequestDialogApi API asynchronously
func (client *Client) RequestDialogApiWithChan(request *RequestDialogApiRequest) (<-chan *RequestDialogApiResponse, <-chan error) {
	responseChan := make(chan *RequestDialogApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestDialogApi(request)
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

// RequestDialogApiWithCallback invokes the alinlp.RequestDialogApi API asynchronously
func (client *Client) RequestDialogApiWithCallback(request *RequestDialogApiRequest, callback func(response *RequestDialogApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestDialogApiResponse
		var err error
		defer close(result)
		response, err = client.RequestDialogApi(request)
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

// RequestDialogApiRequest is the request struct for api RequestDialogApi
type RequestDialogApiRequest struct {
	*requests.RpcRequest
	BotProfile  string `position:"Body" name:"BotProfile"`
	Query       string `position:"Body" name:"Query"`
	UserProfile string `position:"Body" name:"UserProfile"`
	History     string `position:"Body" name:"History"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
}

// RequestDialogApiResponse is the response struct for api RequestDialogApi
type RequestDialogApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateRequestDialogApiRequest creates a request to invoke RequestDialogApi API
func CreateRequestDialogApiRequest() (request *RequestDialogApiRequest) {
	request = &RequestDialogApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "RequestDialogApi", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRequestDialogApiResponse creates a response to parse from RequestDialogApi response
func CreateRequestDialogApiResponse() (response *RequestDialogApiResponse) {
	response = &RequestDialogApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
