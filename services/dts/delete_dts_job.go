package dts

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

// DeleteDtsJob invokes the dts.DeleteDtsJob API synchronously
func (client *Client) DeleteDtsJob(request *DeleteDtsJobRequest) (response *DeleteDtsJobResponse, err error) {
	response = CreateDeleteDtsJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDtsJobWithChan invokes the dts.DeleteDtsJob API asynchronously
func (client *Client) DeleteDtsJobWithChan(request *DeleteDtsJobRequest) (<-chan *DeleteDtsJobResponse, <-chan error) {
	responseChan := make(chan *DeleteDtsJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDtsJob(request)
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

// DeleteDtsJobWithCallback invokes the dts.DeleteDtsJob API asynchronously
func (client *Client) DeleteDtsJobWithCallback(request *DeleteDtsJobRequest, callback func(response *DeleteDtsJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDtsJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteDtsJob(request)
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

// DeleteDtsJobRequest is the request struct for api DeleteDtsJob
type DeleteDtsJobRequest struct {
	*requests.RpcRequest
	DtsJobId                 string `position:"Query" name:"DtsJobId"`
	DtsInstanceId            string `position:"Query" name:"DtsInstanceId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
}

// DeleteDtsJobResponse is the response struct for api DeleteDtsJob
type DeleteDtsJobResponse struct {
	*responses.BaseResponse
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDtsJobRequest creates a request to invoke DeleteDtsJob API
func CreateDeleteDtsJobRequest() (request *DeleteDtsJobRequest) {
	request = &DeleteDtsJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DeleteDtsJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDtsJobResponse creates a response to parse from DeleteDtsJob response
func CreateDeleteDtsJobResponse() (response *DeleteDtsJobResponse) {
	response = &DeleteDtsJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
