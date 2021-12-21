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

// BargeInCall invokes the ccc.BargeInCall API synchronously
func (client *Client) BargeInCall(request *BargeInCallRequest) (response *BargeInCallResponse, err error) {
	response = CreateBargeInCallResponse()
	err = client.DoAction(request, response)
	return
}

// BargeInCallWithChan invokes the ccc.BargeInCall API asynchronously
func (client *Client) BargeInCallWithChan(request *BargeInCallRequest) (<-chan *BargeInCallResponse, <-chan error) {
	responseChan := make(chan *BargeInCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BargeInCall(request)
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

// BargeInCallWithCallback invokes the ccc.BargeInCall API asynchronously
func (client *Client) BargeInCallWithCallback(request *BargeInCallRequest, callback func(response *BargeInCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BargeInCallResponse
		var err error
		defer close(result)
		response, err = client.BargeInCall(request)
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

// BargeInCallRequest is the request struct for api BargeInCall
type BargeInCallRequest struct {
	*requests.RpcRequest
	UserId         string           `position:"Query" name:"UserId"`
	DeviceId       string           `position:"Query" name:"DeviceId"`
	BargedUserId   string           `position:"Query" name:"BargedUserId"`
	JobId          string           `position:"Query" name:"JobId"`
	TimeoutSeconds requests.Integer `position:"Query" name:"TimeoutSeconds"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
}

// BargeInCallResponse is the response struct for api BargeInCall
type BargeInCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateBargeInCallRequest creates a request to invoke BargeInCall API
func CreateBargeInCallRequest() (request *BargeInCallRequest) {
	request = &BargeInCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "BargeInCall", "", "")
	request.Method = requests.POST
	return
}

// CreateBargeInCallResponse creates a response to parse from BargeInCall response
func CreateBargeInCallResponse() (response *BargeInCallResponse) {
	response = &BargeInCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
