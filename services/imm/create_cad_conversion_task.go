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

// CreateCADConversionTask invokes the imm.CreateCADConversionTask API synchronously
// api document: https://help.aliyun.com/api/imm/createcadconversiontask.html
func (client *Client) CreateCADConversionTask(request *CreateCADConversionTaskRequest) (response *CreateCADConversionTaskResponse, err error) {
	response = CreateCreateCADConversionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCADConversionTaskWithChan invokes the imm.CreateCADConversionTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createcadconversiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCADConversionTaskWithChan(request *CreateCADConversionTaskRequest) (<-chan *CreateCADConversionTaskResponse, <-chan error) {
	responseChan := make(chan *CreateCADConversionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCADConversionTask(request)
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

// CreateCADConversionTaskWithCallback invokes the imm.CreateCADConversionTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createcadconversiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCADConversionTaskWithCallback(request *CreateCADConversionTaskRequest, callback func(response *CreateCADConversionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCADConversionTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateCADConversionTask(request)
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

// CreateCADConversionTaskRequest is the request struct for api CreateCADConversionTask
type CreateCADConversionTaskRequest struct {
	*requests.RpcRequest
	SrcType         string           `position:"Query" name:"SrcType"`
	BaseRow         requests.Integer `position:"Query" name:"BaseRow"`
	Project         string           `position:"Query" name:"Project"`
	ZoomFactor      requests.Integer `position:"Query" name:"ZoomFactor"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	BaseCol         requests.Integer `position:"Query" name:"BaseCol"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	UnitWidth       requests.Integer `position:"Query" name:"UnitWidth"`
	ZoomLevel       requests.Integer `position:"Query" name:"ZoomLevel"`
	ModelId         string           `position:"Query" name:"ModelId"`
	TgtType         string           `position:"Query" name:"TgtType"`
	UnitHeight      requests.Integer `position:"Query" name:"UnitHeight"`
	SrcUri          string           `position:"Query" name:"SrcUri"`
	Thumbnails      requests.Boolean `position:"Query" name:"Thumbnails"`
	TgtUri          string           `position:"Query" name:"TgtUri"`
}

// CreateCADConversionTaskResponse is the response struct for api CreateCADConversionTask
type CreateCADConversionTaskResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
	TgtLoc     string `json:"TgtLoc" xml:"TgtLoc"`
	Status     string `json:"Status" xml:"Status"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	Percent    int    `json:"Percent" xml:"Percent"`
}

// CreateCreateCADConversionTaskRequest creates a request to invoke CreateCADConversionTask API
func CreateCreateCADConversionTaskRequest() (request *CreateCADConversionTaskRequest) {
	request = &CreateCADConversionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateCADConversionTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCADConversionTaskResponse creates a response to parse from CreateCADConversionTask response
func CreateCreateCADConversionTaskResponse() (response *CreateCADConversionTaskResponse) {
	response = &CreateCADConversionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
