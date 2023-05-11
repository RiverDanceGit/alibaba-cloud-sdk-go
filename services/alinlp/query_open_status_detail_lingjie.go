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

// QueryOpenStatusDetailLingjie invokes the alinlp.QueryOpenStatusDetailLingjie API synchronously
func (client *Client) QueryOpenStatusDetailLingjie(request *QueryOpenStatusDetailLingjieRequest) (response *QueryOpenStatusDetailLingjieResponse, err error) {
	response = CreateQueryOpenStatusDetailLingjieResponse()
	err = client.DoAction(request, response)
	return
}

// QueryOpenStatusDetailLingjieWithChan invokes the alinlp.QueryOpenStatusDetailLingjie API asynchronously
func (client *Client) QueryOpenStatusDetailLingjieWithChan(request *QueryOpenStatusDetailLingjieRequest) (<-chan *QueryOpenStatusDetailLingjieResponse, <-chan error) {
	responseChan := make(chan *QueryOpenStatusDetailLingjieResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryOpenStatusDetailLingjie(request)
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

// QueryOpenStatusDetailLingjieWithCallback invokes the alinlp.QueryOpenStatusDetailLingjie API asynchronously
func (client *Client) QueryOpenStatusDetailLingjieWithCallback(request *QueryOpenStatusDetailLingjieRequest, callback func(response *QueryOpenStatusDetailLingjieResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryOpenStatusDetailLingjieResponse
		var err error
		defer close(result)
		response, err = client.QueryOpenStatusDetailLingjie(request)
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

// QueryOpenStatusDetailLingjieRequest is the request struct for api QueryOpenStatusDetailLingjie
type QueryOpenStatusDetailLingjieRequest struct {
	*requests.RpcRequest
	Commodity   string `position:"Query" name:"Commodity"`
	TabName     string `position:"Query" name:"TabName"`
	ServiceCode string `position:"Query" name:"ServiceCode"`
	Region      string `position:"Query" name:"Region"`
}

// QueryOpenStatusDetailLingjieResponse is the response struct for api QueryOpenStatusDetailLingjie
type QueryOpenStatusDetailLingjieResponse struct {
	*responses.BaseResponse
	HttpCode     string `json:"HttpCode" xml:"HttpCode"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         string `json:"Data" xml:"Data"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateQueryOpenStatusDetailLingjieRequest creates a request to invoke QueryOpenStatusDetailLingjie API
func CreateQueryOpenStatusDetailLingjieRequest() (request *QueryOpenStatusDetailLingjieRequest) {
	request = &QueryOpenStatusDetailLingjieRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "QueryOpenStatusDetailLingjie", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryOpenStatusDetailLingjieResponse creates a response to parse from QueryOpenStatusDetailLingjie response
func CreateQueryOpenStatusDetailLingjieResponse() (response *QueryOpenStatusDetailLingjieResponse) {
	response = &QueryOpenStatusDetailLingjieResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
