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

// CreateGroupFacesJob invokes the imm.CreateGroupFacesJob API synchronously
// api document: https://help.aliyun.com/api/imm/creategroupfacesjob.html
func (client *Client) CreateGroupFacesJob(request *CreateGroupFacesJobRequest) (response *CreateGroupFacesJobResponse, err error) {
	response = CreateCreateGroupFacesJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGroupFacesJobWithChan invokes the imm.CreateGroupFacesJob API asynchronously
// api document: https://help.aliyun.com/api/imm/creategroupfacesjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupFacesJobWithChan(request *CreateGroupFacesJobRequest) (<-chan *CreateGroupFacesJobResponse, <-chan error) {
	responseChan := make(chan *CreateGroupFacesJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGroupFacesJob(request)
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

// CreateGroupFacesJobWithCallback invokes the imm.CreateGroupFacesJob API asynchronously
// api document: https://help.aliyun.com/api/imm/creategroupfacesjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupFacesJobWithCallback(request *CreateGroupFacesJobRequest, callback func(response *CreateGroupFacesJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGroupFacesJobResponse
		var err error
		defer close(result)
		response, err = client.CreateGroupFacesJob(request)
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

// CreateGroupFacesJobRequest is the request struct for api CreateGroupFacesJob
type CreateGroupFacesJobRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	SetId           string `position:"Query" name:"SetId"`
}

// CreateGroupFacesJobResponse is the response struct for api CreateGroupFacesJob
type CreateGroupFacesJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
	SetId     string `json:"SetId" xml:"SetId"`
	JobType   string `json:"JobType" xml:"JobType"`
}

// CreateCreateGroupFacesJobRequest creates a request to invoke CreateGroupFacesJob API
func CreateCreateGroupFacesJobRequest() (request *CreateGroupFacesJobRequest) {
	request = &CreateGroupFacesJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateGroupFacesJob", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateGroupFacesJobResponse creates a response to parse from CreateGroupFacesJob response
func CreateCreateGroupFacesJobResponse() (response *CreateGroupFacesJobResponse) {
	response = &CreateGroupFacesJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
