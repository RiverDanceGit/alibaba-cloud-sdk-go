package ecs

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

// DescribeFleets invokes the ecs.DescribeFleets API synchronously
// api document: https://help.aliyun.com/api/ecs/describefleets.html
func (client *Client) DescribeFleets(request *DescribeFleetsRequest) (response *DescribeFleetsResponse, err error) {
	response = CreateDescribeFleetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFleetsWithChan invokes the ecs.DescribeFleets API asynchronously
// api document: https://help.aliyun.com/api/ecs/describefleets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFleetsWithChan(request *DescribeFleetsRequest) (<-chan *DescribeFleetsResponse, <-chan error) {
	responseChan := make(chan *DescribeFleetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFleets(request)
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

// DescribeFleetsWithCallback invokes the ecs.DescribeFleets API asynchronously
// api document: https://help.aliyun.com/api/ecs/describefleets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFleetsWithCallback(request *DescribeFleetsRequest, callback func(response *DescribeFleetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFleetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFleets(request)
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

// DescribeFleetsRequest is the request struct for api DescribeFleets
type DescribeFleetsRequest struct {
	*requests.RpcRequest
}

// DescribeFleetsResponse is the response struct for api DescribeFleets
type DescribeFleetsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Fleets     Fleets `json:"Fleets" xml:"Fleets"`
}

// CreateDescribeFleetsRequest creates a request to invoke DescribeFleets API
func CreateDescribeFleetsRequest() (request *DescribeFleetsRequest) {
	request = &DescribeFleetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeFleets", "ecs", "openAPI")
	return
}

// CreateDescribeFleetsResponse creates a response to parse from DescribeFleets response
func CreateDescribeFleetsResponse() (response *DescribeFleetsResponse) {
	response = &DescribeFleetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
