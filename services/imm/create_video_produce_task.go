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

// CreateVideoProduceTask invokes the imm.CreateVideoProduceTask API synchronously
// api document: https://help.aliyun.com/api/imm/createvideoproducetask.html
func (client *Client) CreateVideoProduceTask(request *CreateVideoProduceTaskRequest) (response *CreateVideoProduceTaskResponse, err error) {
	response = CreateCreateVideoProduceTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVideoProduceTaskWithChan invokes the imm.CreateVideoProduceTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createvideoproducetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVideoProduceTaskWithChan(request *CreateVideoProduceTaskRequest) (<-chan *CreateVideoProduceTaskResponse, <-chan error) {
	responseChan := make(chan *CreateVideoProduceTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVideoProduceTask(request)
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

// CreateVideoProduceTaskWithCallback invokes the imm.CreateVideoProduceTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createvideoproducetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVideoProduceTaskWithCallback(request *CreateVideoProduceTaskRequest, callback func(response *CreateVideoProduceTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVideoProduceTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateVideoProduceTask(request)
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

// CreateVideoProduceTaskRequest is the request struct for api CreateVideoProduceTask
type CreateVideoProduceTaskRequest struct {
	*requests.RpcRequest
	Project         string           `position:"Query" name:"Project"`
	Music           string           `position:"Query" name:"Music"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	TargetUri       string           `position:"Query" name:"TargetUri"`
	TemplateName    string           `position:"Query" name:"TemplateName"`
	Height          requests.Integer `position:"Query" name:"Height"`
	CustomMessage   string           `position:"Query" name:"CustomMessage"`
	Images          string           `position:"Query" name:"Images"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	Width           requests.Integer `position:"Query" name:"Width"`
}

// CreateVideoProduceTaskResponse is the response struct for api CreateVideoProduceTask
type CreateVideoProduceTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateVideoProduceTaskRequest creates a request to invoke CreateVideoProduceTask API
func CreateCreateVideoProduceTaskRequest() (request *CreateVideoProduceTaskRequest) {
	request = &CreateVideoProduceTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateVideoProduceTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateVideoProduceTaskResponse creates a response to parse from CreateVideoProduceTask response
func CreateCreateVideoProduceTaskResponse() (response *CreateVideoProduceTaskResponse) {
	response = &CreateVideoProduceTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
