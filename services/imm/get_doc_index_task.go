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

// GetDocIndexTask invokes the imm.GetDocIndexTask API synchronously
// api document: https://help.aliyun.com/api/imm/getdocindextask.html
func (client *Client) GetDocIndexTask(request *GetDocIndexTaskRequest) (response *GetDocIndexTaskResponse, err error) {
	response = CreateGetDocIndexTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetDocIndexTaskWithChan invokes the imm.GetDocIndexTask API asynchronously
// api document: https://help.aliyun.com/api/imm/getdocindextask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocIndexTaskWithChan(request *GetDocIndexTaskRequest) (<-chan *GetDocIndexTaskResponse, <-chan error) {
	responseChan := make(chan *GetDocIndexTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDocIndexTask(request)
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

// GetDocIndexTaskWithCallback invokes the imm.GetDocIndexTask API asynchronously
// api document: https://help.aliyun.com/api/imm/getdocindextask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocIndexTaskWithCallback(request *GetDocIndexTaskRequest, callback func(response *GetDocIndexTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDocIndexTaskResponse
		var err error
		defer close(result)
		response, err = client.GetDocIndexTask(request)
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

// GetDocIndexTaskRequest is the request struct for api GetDocIndexTask
type GetDocIndexTaskRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

// GetDocIndexTaskResponse is the response struct for api GetDocIndexTask
type GetDocIndexTaskResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Status     string     `json:"Status" xml:"Status"`
	TaskId     string     `json:"TaskId" xml:"TaskId"`
	CreateTime string     `json:"CreateTime" xml:"CreateTime"`
	FinishTime string     `json:"FinishTime" xml:"FinishTime"`
	FailDetail FailDetail `json:"FailDetail" xml:"FailDetail"`
}

// CreateGetDocIndexTaskRequest creates a request to invoke GetDocIndexTask API
func CreateGetDocIndexTaskRequest() (request *GetDocIndexTaskRequest) {
	request = &GetDocIndexTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetDocIndexTask", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDocIndexTaskResponse creates a response to parse from GetDocIndexTask response
func CreateGetDocIndexTaskResponse() (response *GetDocIndexTaskResponse) {
	response = &GetDocIndexTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
