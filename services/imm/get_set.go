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

// GetSet invokes the imm.GetSet API synchronously
// api document: https://help.aliyun.com/api/imm/getset.html
func (client *Client) GetSet(request *GetSetRequest) (response *GetSetResponse, err error) {
	response = CreateGetSetResponse()
	err = client.DoAction(request, response)
	return
}

// GetSetWithChan invokes the imm.GetSet API asynchronously
// api document: https://help.aliyun.com/api/imm/getset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSetWithChan(request *GetSetRequest) (<-chan *GetSetResponse, <-chan error) {
	responseChan := make(chan *GetSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSet(request)
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

// GetSetWithCallback invokes the imm.GetSet API asynchronously
// api document: https://help.aliyun.com/api/imm/getset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSetWithCallback(request *GetSetRequest, callback func(response *GetSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSetResponse
		var err error
		defer close(result)
		response, err = client.GetSet(request)
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

// GetSetRequest is the request struct for api GetSet
type GetSetRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

// GetSetResponse is the response struct for api GetSet
type GetSetResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	SetId       string `json:"SetId" xml:"SetId"`
	SetName     string `json:"SetName" xml:"SetName"`
	CreateTime  string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime  string `json:"ModifyTime" xml:"ModifyTime"`
	FaceCount   int    `json:"FaceCount" xml:"FaceCount"`
	ImageCount  int    `json:"ImageCount" xml:"ImageCount"`
	VideoCount  int    `json:"VideoCount" xml:"VideoCount"`
	VideoLength int    `json:"VideoLength" xml:"VideoLength"`
}

// CreateGetSetRequest creates a request to invoke GetSet API
func CreateGetSetRequest() (request *GetSetRequest) {
	request = &GetSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetSet", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSetResponse creates a response to parse from GetSet response
func CreateGetSetResponse() (response *GetSetResponse) {
	response = &GetSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
