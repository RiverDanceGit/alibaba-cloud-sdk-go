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

// DetectImageFaces invokes the imm.DetectImageFaces API synchronously
// api document: https://help.aliyun.com/api/imm/detectimagefaces.html
func (client *Client) DetectImageFaces(request *DetectImageFacesRequest) (response *DetectImageFacesResponse, err error) {
	response = CreateDetectImageFacesResponse()
	err = client.DoAction(request, response)
	return
}

// DetectImageFacesWithChan invokes the imm.DetectImageFaces API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageFacesWithChan(request *DetectImageFacesRequest) (<-chan *DetectImageFacesResponse, <-chan error) {
	responseChan := make(chan *DetectImageFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectImageFaces(request)
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

// DetectImageFacesWithCallback invokes the imm.DetectImageFaces API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageFacesWithCallback(request *DetectImageFacesRequest, callback func(response *DetectImageFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectImageFacesResponse
		var err error
		defer close(result)
		response, err = client.DetectImageFaces(request)
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

// DetectImageFacesRequest is the request struct for api DetectImageFaces
type DetectImageFacesRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	RealUid  string `position:"Query" name:"RealUid"`
	ImageUri string `position:"Query" name:"ImageUri"`
}

// DetectImageFacesResponse is the response struct for api DetectImageFaces
type DetectImageFacesResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	ImageUri  string      `json:"ImageUri" xml:"ImageUri"`
	Faces     []FacesItem `json:"Faces" xml:"Faces"`
}

// CreateDetectImageFacesRequest creates a request to invoke DetectImageFaces API
func CreateDetectImageFacesRequest() (request *DetectImageFacesRequest) {
	request = &DetectImageFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DetectImageFaces", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectImageFacesResponse creates a response to parse from DetectImageFaces response
func CreateDetectImageFacesResponse() (response *DetectImageFacesResponse) {
	response = &DetectImageFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
