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

// UpdateTableTheme invokes the dataworks_public.UpdateTableTheme API synchronously
func (client *Client) UpdateTableTheme(request *UpdateTableThemeRequest) (response *UpdateTableThemeResponse, err error) {
	response = CreateUpdateTableThemeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTableThemeWithChan invokes the dataworks_public.UpdateTableTheme API asynchronously
func (client *Client) UpdateTableThemeWithChan(request *UpdateTableThemeRequest) (<-chan *UpdateTableThemeResponse, <-chan error) {
	responseChan := make(chan *UpdateTableThemeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTableTheme(request)
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

// UpdateTableThemeWithCallback invokes the dataworks_public.UpdateTableTheme API asynchronously
func (client *Client) UpdateTableThemeWithCallback(request *UpdateTableThemeRequest, callback func(response *UpdateTableThemeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTableThemeResponse
		var err error
		defer close(result)
		response, err = client.UpdateTableTheme(request)
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

// UpdateTableThemeRequest is the request struct for api UpdateTableTheme
type UpdateTableThemeRequest struct {
	*requests.RpcRequest
	Name      string           `position:"Query" name:"Name"`
	ThemeId   requests.Integer `position:"Query" name:"ThemeId"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
}

// UpdateTableThemeResponse is the response struct for api UpdateTableTheme
type UpdateTableThemeResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	UpdateResult   bool   `json:"UpdateResult" xml:"UpdateResult"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUpdateTableThemeRequest creates a request to invoke UpdateTableTheme API
func CreateUpdateTableThemeRequest() (request *UpdateTableThemeRequest) {
	request = &UpdateTableThemeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateTableTheme", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTableThemeResponse creates a response to parse from UpdateTableTheme response
func CreateUpdateTableThemeResponse() (response *UpdateTableThemeResponse) {
	response = &UpdateTableThemeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
