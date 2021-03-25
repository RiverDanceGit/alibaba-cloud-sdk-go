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

// QueryBusinessLocations invokes the mse.QueryBusinessLocations API synchronously
func (client *Client) QueryBusinessLocations(request *QueryBusinessLocationsRequest) (response *QueryBusinessLocationsResponse, err error) {
	response = CreateQueryBusinessLocationsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBusinessLocationsWithChan invokes the mse.QueryBusinessLocations API asynchronously
func (client *Client) QueryBusinessLocationsWithChan(request *QueryBusinessLocationsRequest) (<-chan *QueryBusinessLocationsResponse, <-chan error) {
	responseChan := make(chan *QueryBusinessLocationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBusinessLocations(request)
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

// QueryBusinessLocationsWithCallback invokes the mse.QueryBusinessLocations API asynchronously
func (client *Client) QueryBusinessLocationsWithCallback(request *QueryBusinessLocationsRequest, callback func(response *QueryBusinessLocationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBusinessLocationsResponse
		var err error
		defer close(result)
		response, err = client.QueryBusinessLocations(request)
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

// QueryBusinessLocationsRequest is the request struct for api QueryBusinessLocations
type QueryBusinessLocationsRequest struct {
	*requests.RpcRequest
}

// QueryBusinessLocationsResponse is the response struct for api QueryBusinessLocations
type QueryBusinessLocationsResponse struct {
	*responses.BaseResponse
	RequestId string         `json:"RequestId" xml:"RequestId"`
	Success   string         `json:"Success" xml:"Success"`
	Message   string         `json:"Message" xml:"Message"`
	ErrorCode string         `json:"ErrorCode" xml:"ErrorCode"`
	Data      []LocationData `json:"Data" xml:"Data"`
}

// CreateQueryBusinessLocationsRequest creates a request to invoke QueryBusinessLocations API
func CreateQueryBusinessLocationsRequest() (request *QueryBusinessLocationsRequest) {
	request = &QueryBusinessLocationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryBusinessLocations", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryBusinessLocationsResponse creates a response to parse from QueryBusinessLocations response
func CreateQueryBusinessLocationsResponse() (response *QueryBusinessLocationsResponse) {
	response = &QueryBusinessLocationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
