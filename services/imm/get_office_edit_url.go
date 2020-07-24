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

// GetOfficeEditURL invokes the imm.GetOfficeEditURL API synchronously
// api document: https://help.aliyun.com/api/imm/getofficeediturl.html
func (client *Client) GetOfficeEditURL(request *GetOfficeEditURLRequest) (response *GetOfficeEditURLResponse, err error) {
	response = CreateGetOfficeEditURLResponse()
	err = client.DoAction(request, response)
	return
}

// GetOfficeEditURLWithChan invokes the imm.GetOfficeEditURL API asynchronously
// api document: https://help.aliyun.com/api/imm/getofficeediturl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOfficeEditURLWithChan(request *GetOfficeEditURLRequest) (<-chan *GetOfficeEditURLResponse, <-chan error) {
	responseChan := make(chan *GetOfficeEditURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOfficeEditURL(request)
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

// GetOfficeEditURLWithCallback invokes the imm.GetOfficeEditURL API asynchronously
// api document: https://help.aliyun.com/api/imm/getofficeediturl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOfficeEditURLWithCallback(request *GetOfficeEditURLRequest, callback func(response *GetOfficeEditURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOfficeEditURLResponse
		var err error
		defer close(result)
		response, err = client.GetOfficeEditURL(request)
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

// GetOfficeEditURLRequest is the request struct for api GetOfficeEditURL
type GetOfficeEditURLRequest struct {
	*requests.RpcRequest
	SrcType         string `position:"Query" name:"SrcType"`
	Project         string `position:"Query" name:"Project"`
	UserID          string `position:"Query" name:"UserID"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	FileID          string `position:"Query" name:"FileID"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	FileName        string `position:"Query" name:"FileName"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	TgtUri          string `position:"Query" name:"TgtUri"`
	UserName        string `position:"Query" name:"UserName"`
}

// GetOfficeEditURLResponse is the response struct for api GetOfficeEditURL
type GetOfficeEditURLResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	EditURL                 string `json:"EditURL" xml:"EditURL"`
	AccessToken             string `json:"AccessToken" xml:"AccessToken"`
	RefreshToken            string `json:"RefreshToken" xml:"RefreshToken"`
	AccessTokenExpiredTime  string `json:"AccessTokenExpiredTime" xml:"AccessTokenExpiredTime"`
	RefreshTokenExpiredTime string `json:"RefreshTokenExpiredTime" xml:"RefreshTokenExpiredTime"`
}

// CreateGetOfficeEditURLRequest creates a request to invoke GetOfficeEditURL API
func CreateGetOfficeEditURLRequest() (request *GetOfficeEditURLRequest) {
	request = &GetOfficeEditURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetOfficeEditURL", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOfficeEditURLResponse creates a response to parse from GetOfficeEditURL response
func CreateGetOfficeEditURLResponse() (response *GetOfficeEditURLResponse) {
	response = &GetOfficeEditURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
