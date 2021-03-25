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

// ListNacosHistoryConfigs invokes the mse.ListNacosHistoryConfigs API synchronously
func (client *Client) ListNacosHistoryConfigs(request *ListNacosHistoryConfigsRequest) (response *ListNacosHistoryConfigsResponse, err error) {
	response = CreateListNacosHistoryConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// ListNacosHistoryConfigsWithChan invokes the mse.ListNacosHistoryConfigs API asynchronously
func (client *Client) ListNacosHistoryConfigsWithChan(request *ListNacosHistoryConfigsRequest) (<-chan *ListNacosHistoryConfigsResponse, <-chan error) {
	responseChan := make(chan *ListNacosHistoryConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNacosHistoryConfigs(request)
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

// ListNacosHistoryConfigsWithCallback invokes the mse.ListNacosHistoryConfigs API asynchronously
func (client *Client) ListNacosHistoryConfigsWithCallback(request *ListNacosHistoryConfigsRequest, callback func(response *ListNacosHistoryConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNacosHistoryConfigsResponse
		var err error
		defer close(result)
		response, err = client.ListNacosHistoryConfigs(request)
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

// ListNacosHistoryConfigsRequest is the request struct for api ListNacosHistoryConfigs
type ListNacosHistoryConfigsRequest struct {
	*requests.RpcRequest
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	DataId      string           `position:"Query" name:"DataId"`
	NamespaceId string           `position:"Query" name:"NamespaceId"`
	RequestPars string           `position:"Query" name:"RequestPars"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Group       string           `position:"Query" name:"Group"`
}

// ListNacosHistoryConfigsResponse is the response struct for api ListNacosHistoryConfigs
type ListNacosHistoryConfigsResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	Success      bool          `json:"Success" xml:"Success"`
	Message      string        `json:"Message" xml:"Message"`
	ErrorCode    string        `json:"ErrorCode" xml:"ErrorCode"`
	PageNumber   int           `json:"PageNumber" xml:"PageNumber"`
	PageSize     int           `json:"PageSize" xml:"PageSize"`
	TotalCount   int           `json:"TotalCount" xml:"TotalCount"`
	HttpCode     string        `json:"HttpCode" xml:"HttpCode"`
	HistoryItems []HistoryItem `json:"HistoryItems" xml:"HistoryItems"`
}

// CreateListNacosHistoryConfigsRequest creates a request to invoke ListNacosHistoryConfigs API
func CreateListNacosHistoryConfigsRequest() (request *ListNacosHistoryConfigsRequest) {
	request = &ListNacosHistoryConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListNacosHistoryConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateListNacosHistoryConfigsResponse creates a response to parse from ListNacosHistoryConfigs response
func CreateListNacosHistoryConfigsResponse() (response *ListNacosHistoryConfigsResponse) {
	response = &ListNacosHistoryConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
