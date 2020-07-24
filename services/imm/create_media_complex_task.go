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

// CreateMediaComplexTask invokes the imm.CreateMediaComplexTask API synchronously
// api document: https://help.aliyun.com/api/imm/createmediacomplextask.html
func (client *Client) CreateMediaComplexTask(request *CreateMediaComplexTaskRequest) (response *CreateMediaComplexTaskResponse, err error) {
	response = CreateCreateMediaComplexTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMediaComplexTaskWithChan invokes the imm.CreateMediaComplexTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createmediacomplextask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMediaComplexTaskWithChan(request *CreateMediaComplexTaskRequest) (<-chan *CreateMediaComplexTaskResponse, <-chan error) {
	responseChan := make(chan *CreateMediaComplexTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMediaComplexTask(request)
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

// CreateMediaComplexTaskWithCallback invokes the imm.CreateMediaComplexTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createmediacomplextask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMediaComplexTaskWithCallback(request *CreateMediaComplexTaskRequest, callback func(response *CreateMediaComplexTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMediaComplexTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateMediaComplexTask(request)
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

// CreateMediaComplexTaskRequest is the request struct for api CreateMediaComplexTask
type CreateMediaComplexTaskRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	Parameters      string `position:"Query" name:"Parameters"`
}

// CreateMediaComplexTaskResponse is the response struct for api CreateMediaComplexTask
type CreateMediaComplexTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateMediaComplexTaskRequest creates a request to invoke CreateMediaComplexTask API
func CreateCreateMediaComplexTaskRequest() (request *CreateMediaComplexTaskRequest) {
	request = &CreateMediaComplexTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateMediaComplexTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMediaComplexTaskResponse creates a response to parse from CreateMediaComplexTask response
func CreateCreateMediaComplexTaskResponse() (response *CreateMediaComplexTaskResponse) {
	response = &CreateMediaComplexTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
