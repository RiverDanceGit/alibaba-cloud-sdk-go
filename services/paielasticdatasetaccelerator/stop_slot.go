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

// StopSlot invokes the paielasticdatasetaccelerator.StopSlot API synchronously
func (client *Client) StopSlot(request *StopSlotRequest) (response *StopSlotResponse, err error) {
	response = CreateStopSlotResponse()
	err = client.DoAction(request, response)
	return
}

// StopSlotWithChan invokes the paielasticdatasetaccelerator.StopSlot API asynchronously
func (client *Client) StopSlotWithChan(request *StopSlotRequest) (<-chan *StopSlotResponse, <-chan error) {
	responseChan := make(chan *StopSlotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopSlot(request)
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

// StopSlotWithCallback invokes the paielasticdatasetaccelerator.StopSlot API asynchronously
func (client *Client) StopSlotWithCallback(request *StopSlotRequest, callback func(response *StopSlotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopSlotResponse
		var err error
		defer close(result)
		response, err = client.StopSlot(request)
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

// StopSlotRequest is the request struct for api StopSlot
type StopSlotRequest struct {
	*requests.RoaRequest
	SlotId string `position:"Path" name:"SlotId"`
}

// StopSlotResponse is the response struct for api StopSlot
type StopSlotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopSlotRequest creates a request to invoke StopSlot API
func CreateStopSlotRequest() (request *StopSlotRequest) {
	request = &StopSlotRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "StopSlot", "/api/v1/slots/[SlotId]/action/stop", "datasetacc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopSlotResponse creates a response to parse from StopSlot response
func CreateStopSlotResponse() (response *StopSlotResponse) {
	response = &StopSlotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
