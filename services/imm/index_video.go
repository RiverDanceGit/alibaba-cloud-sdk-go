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

// IndexVideo invokes the imm.IndexVideo API synchronously
// api document: https://help.aliyun.com/api/imm/indexvideo.html
func (client *Client) IndexVideo(request *IndexVideoRequest) (response *IndexVideoResponse, err error) {
	response = CreateIndexVideoResponse()
	err = client.DoAction(request, response)
	return
}

// IndexVideoWithChan invokes the imm.IndexVideo API asynchronously
// api document: https://help.aliyun.com/api/imm/indexvideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IndexVideoWithChan(request *IndexVideoRequest) (<-chan *IndexVideoResponse, <-chan error) {
	responseChan := make(chan *IndexVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.IndexVideo(request)
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

// IndexVideoWithCallback invokes the imm.IndexVideo API asynchronously
// api document: https://help.aliyun.com/api/imm/indexvideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IndexVideoWithCallback(request *IndexVideoRequest, callback func(response *IndexVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *IndexVideoResponse
		var err error
		defer close(result)
		response, err = client.IndexVideo(request)
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

// IndexVideoRequest is the request struct for api IndexVideo
type IndexVideoRequest struct {
	*requests.RpcRequest
	GrabType   string           `position:"Query" name:"GrabType"`
	Project    string           `position:"Query" name:"Project"`
	ExternalId string           `position:"Query" name:"ExternalId"`
	StartTime  string           `position:"Query" name:"StartTime"`
	RemarksB   string           `position:"Query" name:"RemarksB"`
	RemarksA   string           `position:"Query" name:"RemarksA"`
	EndTime    string           `position:"Query" name:"EndTime"`
	VideoUri   string           `position:"Query" name:"VideoUri"`
	SaveType   requests.Boolean `position:"Query" name:"SaveType"`
	RemarksD   string           `position:"Query" name:"RemarksD"`
	RemarksC   string           `position:"Query" name:"RemarksC"`
	SetId      string           `position:"Query" name:"SetId"`
	Interval   string           `position:"Query" name:"Interval"`
	TgtUri     string           `position:"Query" name:"TgtUri"`
}

// IndexVideoResponse is the response struct for api IndexVideo
type IndexVideoResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	SetId      string  `json:"SetId" xml:"SetId"`
	VideoUri   string  `json:"VideoUri" xml:"VideoUri"`
	RemarksA   string  `json:"RemarksA" xml:"RemarksA"`
	RemarksB   string  `json:"RemarksB" xml:"RemarksB"`
	CreateTime string  `json:"CreateTime" xml:"CreateTime"`
	ModifyTime string  `json:"ModifyTime" xml:"ModifyTime"`
	Interval   float64 `json:"Interval" xml:"Interval"`
	GrabType   string  `json:"GrabType" xml:"GrabType"`
	StartTime  string  `json:"StartTime" xml:"StartTime"`
	EndTime    string  `json:"EndTime" xml:"EndTime"`
	SaveType   bool    `json:"SaveType" xml:"SaveType"`
	TgtUri     string  `json:"TgtUri" xml:"TgtUri"`
	RemarksC   string  `json:"RemarksC" xml:"RemarksC"`
	RemarksD   string  `json:"RemarksD" xml:"RemarksD"`
	ExternalId string  `json:"ExternalId" xml:"ExternalId"`
}

// CreateIndexVideoRequest creates a request to invoke IndexVideo API
func CreateIndexVideoRequest() (request *IndexVideoRequest) {
	request = &IndexVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "IndexVideo", "", "")
	request.Method = requests.POST
	return
}

// CreateIndexVideoResponse creates a response to parse from IndexVideo response
func CreateIndexVideoResponse() (response *IndexVideoResponse) {
	response = &IndexVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
