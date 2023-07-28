package dts

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

// ModifyConsumptionTimestamp invokes the dts.ModifyConsumptionTimestamp API synchronously
func (client *Client) ModifyConsumptionTimestamp(request *ModifyConsumptionTimestampRequest) (response *ModifyConsumptionTimestampResponse, err error) {
	response = CreateModifyConsumptionTimestampResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyConsumptionTimestampWithChan invokes the dts.ModifyConsumptionTimestamp API asynchronously
func (client *Client) ModifyConsumptionTimestampWithChan(request *ModifyConsumptionTimestampRequest) (<-chan *ModifyConsumptionTimestampResponse, <-chan error) {
	responseChan := make(chan *ModifyConsumptionTimestampResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyConsumptionTimestamp(request)
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

// ModifyConsumptionTimestampWithCallback invokes the dts.ModifyConsumptionTimestamp API asynchronously
func (client *Client) ModifyConsumptionTimestampWithCallback(request *ModifyConsumptionTimestampRequest, callback func(response *ModifyConsumptionTimestampResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyConsumptionTimestampResponse
		var err error
		defer close(result)
		response, err = client.ModifyConsumptionTimestamp(request)
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

// ModifyConsumptionTimestampRequest is the request struct for api ModifyConsumptionTimestamp
type ModifyConsumptionTimestampRequest struct {
	*requests.RpcRequest
	SubscriptionInstanceId string `position:"Query" name:"SubscriptionInstanceId"`
	ConsumptionTimestamp   string `position:"Query" name:"ConsumptionTimestamp"`
	OwnerId                string `position:"Query" name:"OwnerId"`
}

// ModifyConsumptionTimestampResponse is the response struct for api ModifyConsumptionTimestamp
type ModifyConsumptionTimestampResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateModifyConsumptionTimestampRequest creates a request to invoke ModifyConsumptionTimestamp API
func CreateModifyConsumptionTimestampRequest() (request *ModifyConsumptionTimestampRequest) {
	request = &ModifyConsumptionTimestampRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2016-08-01", "ModifyConsumptionTimestamp", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyConsumptionTimestampResponse creates a response to parse from ModifyConsumptionTimestamp response
func CreateModifyConsumptionTimestampResponse() (response *ModifyConsumptionTimestampResponse) {
	response = &ModifyConsumptionTimestampResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
