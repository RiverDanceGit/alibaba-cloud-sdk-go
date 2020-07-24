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

// UpdateFaceGroup invokes the imm.UpdateFaceGroup API synchronously
// api document: https://help.aliyun.com/api/imm/updatefacegroup.html
func (client *Client) UpdateFaceGroup(request *UpdateFaceGroupRequest) (response *UpdateFaceGroupResponse, err error) {
	response = CreateUpdateFaceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFaceGroupWithChan invokes the imm.UpdateFaceGroup API asynchronously
// api document: https://help.aliyun.com/api/imm/updatefacegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFaceGroupWithChan(request *UpdateFaceGroupRequest) (<-chan *UpdateFaceGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateFaceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFaceGroup(request)
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

// UpdateFaceGroupWithCallback invokes the imm.UpdateFaceGroup API asynchronously
// api document: https://help.aliyun.com/api/imm/updatefacegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFaceGroupWithCallback(request *UpdateFaceGroupRequest, callback func(response *UpdateFaceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFaceGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateFaceGroup(request)
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

// UpdateFaceGroupRequest is the request struct for api UpdateFaceGroup
type UpdateFaceGroupRequest struct {
	*requests.RpcRequest
	Project          string `position:"Query" name:"Project"`
	ExternalId       string `position:"Query" name:"ExternalId"`
	GroupId          string `position:"Query" name:"GroupId"`
	RemarksB         string `position:"Query" name:"RemarksB"`
	RemarksA         string `position:"Query" name:"RemarksA"`
	GroupName        string `position:"Query" name:"GroupName"`
	RemarksArrayA    string `position:"Query" name:"RemarksArrayA"`
	RemarksArrayB    string `position:"Query" name:"RemarksArrayB"`
	RemarksD         string `position:"Query" name:"RemarksD"`
	RemarksC         string `position:"Query" name:"RemarksC"`
	SetId            string `position:"Query" name:"SetId"`
	GroupCoverFaceId string `position:"Query" name:"GroupCoverFaceId"`
}

// UpdateFaceGroupResponse is the response struct for api UpdateFaceGroup
type UpdateFaceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SetId     string `json:"SetId" xml:"SetId"`
	GroupId   string `json:"GroupId" xml:"GroupId"`
}

// CreateUpdateFaceGroupRequest creates a request to invoke UpdateFaceGroup API
func CreateUpdateFaceGroupRequest() (request *UpdateFaceGroupRequest) {
	request = &UpdateFaceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "UpdateFaceGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateFaceGroupResponse creates a response to parse from UpdateFaceGroup response
func CreateUpdateFaceGroupResponse() (response *UpdateFaceGroupResponse) {
	response = &UpdateFaceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
