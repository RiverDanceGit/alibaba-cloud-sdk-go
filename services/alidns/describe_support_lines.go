package alidns

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

// DescribeSupportLines invokes the alidns.DescribeSupportLines API synchronously
// api document: https://help.aliyun.com/api/alidns/describesupportlines.html
func (client *Client) DescribeSupportLines(request *DescribeSupportLinesRequest) (response *DescribeSupportLinesResponse, err error) {
	response = CreateDescribeSupportLinesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSupportLinesWithChan invokes the alidns.DescribeSupportLines API asynchronously
// api document: https://help.aliyun.com/api/alidns/describesupportlines.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSupportLinesWithChan(request *DescribeSupportLinesRequest) (<-chan *DescribeSupportLinesResponse, <-chan error) {
	responseChan := make(chan *DescribeSupportLinesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSupportLines(request)
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

// DescribeSupportLinesWithCallback invokes the alidns.DescribeSupportLines API asynchronously
// api document: https://help.aliyun.com/api/alidns/describesupportlines.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSupportLinesWithCallback(request *DescribeSupportLinesRequest, callback func(response *DescribeSupportLinesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSupportLinesResponse
		var err error
		defer close(result)
		response, err = client.DescribeSupportLines(request)
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

// DescribeSupportLinesRequest is the request struct for api DescribeSupportLines
type DescribeSupportLinesRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeSupportLinesResponse is the response struct for api DescribeSupportLines
type DescribeSupportLinesResponse struct {
	*responses.BaseResponse
	RequestId   string                            `json:"RequestId" xml:"RequestId"`
	RecordLines RecordLinesInDescribeSupportLines `json:"RecordLines" xml:"RecordLines"`
}

// CreateDescribeSupportLinesRequest creates a request to invoke DescribeSupportLines API
func CreateDescribeSupportLinesRequest() (request *DescribeSupportLinesRequest) {
	request = &DescribeSupportLinesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeSupportLines", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSupportLinesResponse creates a response to parse from DescribeSupportLines response
func CreateDescribeSupportLinesResponse() (response *DescribeSupportLinesResponse) {
	response = &DescribeSupportLinesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
