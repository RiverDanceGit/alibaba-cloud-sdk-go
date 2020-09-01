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

// UpdateCustomLine invokes the alidns.UpdateCustomLine API synchronously
// api document: https://help.aliyun.com/api/alidns/updatecustomline.html
func (client *Client) UpdateCustomLine(request *UpdateCustomLineRequest) (response *UpdateCustomLineResponse, err error) {
	response = CreateUpdateCustomLineResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCustomLineWithChan invokes the alidns.UpdateCustomLine API asynchronously
// api document: https://help.aliyun.com/api/alidns/updatecustomline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCustomLineWithChan(request *UpdateCustomLineRequest) (<-chan *UpdateCustomLineResponse, <-chan error) {
	responseChan := make(chan *UpdateCustomLineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCustomLine(request)
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

// UpdateCustomLineWithCallback invokes the alidns.UpdateCustomLine API asynchronously
// api document: https://help.aliyun.com/api/alidns/updatecustomline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCustomLineWithCallback(request *UpdateCustomLineRequest, callback func(response *UpdateCustomLineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCustomLineResponse
		var err error
		defer close(result)
		response, err = client.UpdateCustomLine(request)
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

// UpdateCustomLineRequest is the request struct for api UpdateCustomLine
type UpdateCustomLineRequest struct {
	*requests.RpcRequest
	LineId       requests.Integer             `position:"Query" name:"LineId"`
	IpSegment    *[]UpdateCustomLineIpSegment `position:"Query" name:"IpSegment"  type:"Repeated"`
	UserClientIp string                       `position:"Query" name:"UserClientIp"`
	LineName     string                       `position:"Query" name:"LineName"`
	Lang         string                       `position:"Query" name:"Lang"`
}

// UpdateCustomLineIpSegment is a repeated param struct in UpdateCustomLineRequest
type UpdateCustomLineIpSegment struct {
	EndIp   string `name:"EndIp"`
	StartIp string `name:"StartIp"`
}

// UpdateCustomLineResponse is the response struct for api UpdateCustomLine
type UpdateCustomLineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCustomLineRequest creates a request to invoke UpdateCustomLine API
func CreateUpdateCustomLineRequest() (request *UpdateCustomLineRequest) {
	request = &UpdateCustomLineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateCustomLine", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateCustomLineResponse creates a response to parse from UpdateCustomLine response
func CreateUpdateCustomLineResponse() (response *UpdateCustomLineResponse) {
	response = &UpdateCustomLineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
