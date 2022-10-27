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

// ListContainerImages invokes the ehpc.ListContainerImages API synchronously
func (client *Client) ListContainerImages(request *ListContainerImagesRequest) (response *ListContainerImagesResponse, err error) {
	response = CreateListContainerImagesResponse()
	err = client.DoAction(request, response)
	return
}

// ListContainerImagesWithChan invokes the ehpc.ListContainerImages API asynchronously
func (client *Client) ListContainerImagesWithChan(request *ListContainerImagesRequest) (<-chan *ListContainerImagesResponse, <-chan error) {
	responseChan := make(chan *ListContainerImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListContainerImages(request)
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

// ListContainerImagesWithCallback invokes the ehpc.ListContainerImages API asynchronously
func (client *Client) ListContainerImagesWithCallback(request *ListContainerImagesRequest, callback func(response *ListContainerImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListContainerImagesResponse
		var err error
		defer close(result)
		response, err = client.ListContainerImages(request)
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

// ListContainerImagesRequest is the request struct for api ListContainerImages
type ListContainerImagesRequest struct {
	*requests.RpcRequest
	ClusterId     string           `position:"Query" name:"ClusterId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	ContainerType string           `position:"Query" name:"ContainerType"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListContainerImagesResponse is the response struct for api ListContainerImages
type ListContainerImagesResponse struct {
	*responses.BaseResponse
	PageSize   int                         `json:"PageSize" xml:"PageSize"`
	PageNumber int                         `json:"PageNumber" xml:"PageNumber"`
	RequestId  string                      `json:"RequestId" xml:"RequestId"`
	TotalCount int                         `json:"TotalCount" xml:"TotalCount"`
	DBInfo     string                      `json:"DBInfo" xml:"DBInfo"`
	Images     ImagesInListContainerImages `json:"Images" xml:"Images"`
}

// CreateListContainerImagesRequest creates a request to invoke ListContainerImages API
func CreateListContainerImagesRequest() (request *ListContainerImagesRequest) {
	request = &ListContainerImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListContainerImages", "", "")
	request.Method = requests.GET
	return
}

// CreateListContainerImagesResponse creates a response to parse from ListContainerImages response
func CreateListContainerImagesResponse() (response *ListContainerImagesResponse) {
	response = &ListContainerImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
