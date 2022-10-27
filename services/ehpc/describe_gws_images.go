package ehpc

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

// DescribeGWSImages invokes the ehpc.DescribeGWSImages API synchronously
func (client *Client) DescribeGWSImages(request *DescribeGWSImagesRequest) (response *DescribeGWSImagesResponse, err error) {
	response = CreateDescribeGWSImagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGWSImagesWithChan invokes the ehpc.DescribeGWSImages API asynchronously
func (client *Client) DescribeGWSImagesWithChan(request *DescribeGWSImagesRequest) (<-chan *DescribeGWSImagesResponse, <-chan error) {
	responseChan := make(chan *DescribeGWSImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGWSImages(request)
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

// DescribeGWSImagesWithCallback invokes the ehpc.DescribeGWSImages API asynchronously
func (client *Client) DescribeGWSImagesWithCallback(request *DescribeGWSImagesRequest, callback func(response *DescribeGWSImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGWSImagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGWSImages(request)
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

// DescribeGWSImagesRequest is the request struct for api DescribeGWSImages
type DescribeGWSImagesRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeGWSImagesResponse is the response struct for api DescribeGWSImages
type DescribeGWSImagesResponse struct {
	*responses.BaseResponse
	PageSize   int                       `json:"PageSize" xml:"PageSize"`
	RequestId  string                    `json:"RequestId" xml:"RequestId"`
	PageNumber int                       `json:"PageNumber" xml:"PageNumber"`
	TotalCount int                       `json:"TotalCount" xml:"TotalCount"`
	Images     ImagesInDescribeGWSImages `json:"Images" xml:"Images"`
}

// CreateDescribeGWSImagesRequest creates a request to invoke DescribeGWSImages API
func CreateDescribeGWSImagesRequest() (request *DescribeGWSImagesRequest) {
	request = &DescribeGWSImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DescribeGWSImages", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeGWSImagesResponse creates a response to parse from DescribeGWSImages response
func CreateDescribeGWSImagesResponse() (response *DescribeGWSImagesResponse) {
	response = &DescribeGWSImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
