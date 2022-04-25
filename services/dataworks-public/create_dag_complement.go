package dataworks_public

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

// CreateDagComplement invokes the dataworks_public.CreateDagComplement API synchronously
func (client *Client) CreateDagComplement(request *CreateDagComplementRequest) (response *CreateDagComplementResponse, err error) {
	response = CreateCreateDagComplementResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDagComplementWithChan invokes the dataworks_public.CreateDagComplement API asynchronously
func (client *Client) CreateDagComplementWithChan(request *CreateDagComplementRequest) (<-chan *CreateDagComplementResponse, <-chan error) {
	responseChan := make(chan *CreateDagComplementResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDagComplement(request)
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

// CreateDagComplementWithCallback invokes the dataworks_public.CreateDagComplement API asynchronously
func (client *Client) CreateDagComplementWithCallback(request *CreateDagComplementRequest, callback func(response *CreateDagComplementResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDagComplementResponse
		var err error
		defer close(result)
		response, err = client.CreateDagComplement(request)
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

// CreateDagComplementRequest is the request struct for api CreateDagComplement
type CreateDagComplementRequest struct {
	*requests.RpcRequest
	ProjectEnv     string           `position:"Body" name:"ProjectEnv"`
	StartBizDate   string           `position:"Body" name:"StartBizDate"`
	Parallelism    requests.Boolean `position:"Body" name:"Parallelism"`
	RootNodeId     requests.Integer `position:"Body" name:"RootNodeId"`
	BizBeginTime   string           `position:"Body" name:"BizBeginTime"`
	EndBizDate     string           `position:"Body" name:"EndBizDate"`
	IncludeNodeIds string           `position:"Body" name:"IncludeNodeIds"`
	BizEndTime     string           `position:"Body" name:"BizEndTime"`
	Name           string           `position:"Body" name:"Name"`
	ExcludeNodeIds string           `position:"Body" name:"ExcludeNodeIds"`
	NodeParams     string           `position:"Body" name:"NodeParams"`
}

// CreateDagComplementResponse is the response struct for api CreateDagComplement
type CreateDagComplementResponse struct {
	*responses.BaseResponse
	HttpStatusCode int     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string  `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string  `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool    `json:"Success" xml:"Success"`
	ErrorCode      string  `json:"ErrorCode" xml:"ErrorCode"`
	Data           []int64 `json:"Data" xml:"Data"`
}

// CreateCreateDagComplementRequest creates a request to invoke CreateDagComplement API
func CreateCreateDagComplementRequest() (request *CreateDagComplementRequest) {
	request = &CreateDagComplementRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateDagComplement", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDagComplementResponse creates a response to parse from CreateDagComplement response
func CreateCreateDagComplementResponse() (response *CreateDagComplementResponse) {
	response = &CreateDagComplementResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
