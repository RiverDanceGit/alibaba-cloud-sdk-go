package imm

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

// DetectImageTags invokes the imm.DetectImageTags API synchronously
// api document: https://help.aliyun.com/api/imm/detectimagetags.html
func (client *Client) DetectImageTags(request *DetectImageTagsRequest) (response *DetectImageTagsResponse, err error) {
	response = CreateDetectImageTagsResponse()
	err = client.DoAction(request, response)
	return
}

// DetectImageTagsWithChan invokes the imm.DetectImageTags API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageTagsWithChan(request *DetectImageTagsRequest) (<-chan *DetectImageTagsResponse, <-chan error) {
	responseChan := make(chan *DetectImageTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectImageTags(request)
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

// DetectImageTagsWithCallback invokes the imm.DetectImageTags API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageTagsWithCallback(request *DetectImageTagsRequest, callback func(response *DetectImageTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectImageTagsResponse
		var err error
		defer close(result)
		response, err = client.DetectImageTags(request)
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

// DetectImageTagsRequest is the request struct for api DetectImageTags
type DetectImageTagsRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	RealUid  string `position:"Query" name:"RealUid"`
	ImageUri string `position:"Query" name:"ImageUri"`
}

// DetectImageTagsResponse is the response struct for api DetectImageTags
type DetectImageTagsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	ImageUri  string     `json:"ImageUri" xml:"ImageUri"`
	Tags      []TagsItem `json:"Tags" xml:"Tags"`
}

// CreateDetectImageTagsRequest creates a request to invoke DetectImageTags API
func CreateDetectImageTagsRequest() (request *DetectImageTagsRequest) {
	request = &DetectImageTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DetectImageTags", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectImageTagsResponse creates a response to parse from DetectImageTags response
func CreateDetectImageTagsResponse() (response *DetectImageTagsResponse) {
	response = &DetectImageTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
