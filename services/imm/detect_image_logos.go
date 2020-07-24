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

// DetectImageLogos invokes the imm.DetectImageLogos API synchronously
// api document: https://help.aliyun.com/api/imm/detectimagelogos.html
func (client *Client) DetectImageLogos(request *DetectImageLogosRequest) (response *DetectImageLogosResponse, err error) {
	response = CreateDetectImageLogosResponse()
	err = client.DoAction(request, response)
	return
}

// DetectImageLogosWithChan invokes the imm.DetectImageLogos API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagelogos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageLogosWithChan(request *DetectImageLogosRequest) (<-chan *DetectImageLogosResponse, <-chan error) {
	responseChan := make(chan *DetectImageLogosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectImageLogos(request)
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

// DetectImageLogosWithCallback invokes the imm.DetectImageLogos API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagelogos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageLogosWithCallback(request *DetectImageLogosRequest, callback func(response *DetectImageLogosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectImageLogosResponse
		var err error
		defer close(result)
		response, err = client.DetectImageLogos(request)
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

// DetectImageLogosRequest is the request struct for api DetectImageLogos
type DetectImageLogosRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	ImageUri string `position:"Query" name:"ImageUri"`
}

// DetectImageLogosResponse is the response struct for api DetectImageLogos
type DetectImageLogosResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	ImageUri  string      `json:"ImageUri" xml:"ImageUri"`
	Logos     []LogosItem `json:"Logos" xml:"Logos"`
}

// CreateDetectImageLogosRequest creates a request to invoke DetectImageLogos API
func CreateDetectImageLogosRequest() (request *DetectImageLogosRequest) {
	request = &DetectImageLogosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DetectImageLogos", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectImageLogosResponse creates a response to parse from DetectImageLogos response
func CreateDetectImageLogosResponse() (response *DetectImageLogosResponse) {
	response = &DetectImageLogosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
