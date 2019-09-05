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

// DescribeVirtualBorderRoutersForPhysicalConnection invokes the ecs.DescribeVirtualBorderRoutersForPhysicalConnection API synchronously
// api document: https://help.aliyun.com/api/ecs/describevirtualborderroutersforphysicalconnection.html
func (client *Client) DescribeVirtualBorderRoutersForPhysicalConnection(request *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) (response *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, err error) {
	response = CreateDescribeVirtualBorderRoutersForPhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVirtualBorderRoutersForPhysicalConnectionWithChan invokes the ecs.DescribeVirtualBorderRoutersForPhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/describevirtualborderroutersforphysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVirtualBorderRoutersForPhysicalConnectionWithChan(request *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) (<-chan *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVirtualBorderRoutersForPhysicalConnection(request)
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

// DescribeVirtualBorderRoutersForPhysicalConnectionWithCallback invokes the ecs.DescribeVirtualBorderRoutersForPhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/describevirtualborderroutersforphysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVirtualBorderRoutersForPhysicalConnectionWithCallback(request *DescribeVirtualBorderRoutersForPhysicalConnectionRequest, callback func(response *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVirtualBorderRoutersForPhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.DescribeVirtualBorderRoutersForPhysicalConnection(request)
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

// DescribeVirtualBorderRoutersForPhysicalConnectionRequest is the request struct for api DescribeVirtualBorderRoutersForPhysicalConnection
type DescribeVirtualBorderRoutersForPhysicalConnectionRequest struct {
	*requests.RpcRequest
	Filter               *[]DescribeVirtualBorderRoutersForPhysicalConnectionFilter `position:"Query" name:"Filter"  type:"Repeated"`
	OwnerId              requests.Integer                                           `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                                                     `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer                                           `position:"Query" name:"ResourceOwnerId"`
	PhysicalConnectionId string                                                     `position:"Query" name:"PhysicalConnectionId"`
	PageNumber           requests.Integer                                           `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer                                           `position:"Query" name:"PageSize"`
}

// DescribeVirtualBorderRoutersForPhysicalConnectionFilter is a repeated param struct in DescribeVirtualBorderRoutersForPhysicalConnectionRequest
type DescribeVirtualBorderRoutersForPhysicalConnectionFilter struct {
	Key   string    `name:"Key"`
	Value *[]string `name:"Value" type:"Repeated"`
}

// DescribeVirtualBorderRoutersForPhysicalConnectionResponse is the response struct for api DescribeVirtualBorderRoutersForPhysicalConnection
type DescribeVirtualBorderRoutersForPhysicalConnectionResponse struct {
	*responses.BaseResponse
	RequestId                                   string                                      `json:"RequestId" xml:"RequestId"`
	PageNumber                                  int                                         `json:"PageNumber" xml:"PageNumber"`
	PageSize                                    int                                         `json:"PageSize" xml:"PageSize"`
	TotalCount                                  int                                         `json:"TotalCount" xml:"TotalCount"`
	VirtualBorderRouterForPhysicalConnectionSet VirtualBorderRouterForPhysicalConnectionSet `json:"VirtualBorderRouterForPhysicalConnectionSet" xml:"VirtualBorderRouterForPhysicalConnectionSet"`
}

// CreateDescribeVirtualBorderRoutersForPhysicalConnectionRequest creates a request to invoke DescribeVirtualBorderRoutersForPhysicalConnection API
func CreateDescribeVirtualBorderRoutersForPhysicalConnectionRequest() (request *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) {
	request = &DescribeVirtualBorderRoutersForPhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVirtualBorderRoutersForPhysicalConnection", "ecs", "openAPI")
	return
}

// CreateDescribeVirtualBorderRoutersForPhysicalConnectionResponse creates a response to parse from DescribeVirtualBorderRoutersForPhysicalConnection response
func CreateDescribeVirtualBorderRoutersForPhysicalConnectionResponse() (response *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) {
	response = &DescribeVirtualBorderRoutersForPhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
