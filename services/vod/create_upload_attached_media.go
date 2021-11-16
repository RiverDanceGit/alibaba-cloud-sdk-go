package vod

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

// CreateUploadAttachedMedia invokes the vod.CreateUploadAttachedMedia API synchronously
func (client *Client) CreateUploadAttachedMedia(request *CreateUploadAttachedMediaRequest) (response *CreateUploadAttachedMediaResponse, err error) {
	response = CreateCreateUploadAttachedMediaResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUploadAttachedMediaWithChan invokes the vod.CreateUploadAttachedMedia API asynchronously
func (client *Client) CreateUploadAttachedMediaWithChan(request *CreateUploadAttachedMediaRequest) (<-chan *CreateUploadAttachedMediaResponse, <-chan error) {
	responseChan := make(chan *CreateUploadAttachedMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUploadAttachedMedia(request)
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

// CreateUploadAttachedMediaWithCallback invokes the vod.CreateUploadAttachedMedia API asynchronously
func (client *Client) CreateUploadAttachedMediaWithCallback(request *CreateUploadAttachedMediaRequest, callback func(response *CreateUploadAttachedMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUploadAttachedMediaResponse
		var err error
		defer close(result)
		response, err = client.CreateUploadAttachedMedia(request)
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

// CreateUploadAttachedMediaRequest is the request struct for api CreateUploadAttachedMedia
type CreateUploadAttachedMediaRequest struct {
	*requests.RpcRequest
	Icon            string           `position:"Query" name:"Icon"`
	Description     string           `position:"Query" name:"Description"`
	FileSize        string           `position:"Query" name:"FileSize"`
	Title           string           `position:"Query" name:"Title"`
	BusinessType    string           `position:"Query" name:"BusinessType"`
	StorageLocation string           `position:"Query" name:"StorageLocation"`
	UserData        string           `position:"Query" name:"UserData"`
	CateId          requests.Integer `position:"Query" name:"CateId"`
	CateIds         string           `position:"Query" name:"CateIds"`
	Tags            string           `position:"Query" name:"Tags"`
	MediaExt        string           `position:"Query" name:"MediaExt"`
	FileName        string           `position:"Query" name:"FileName"`
	AppId           string           `position:"Query" name:"AppId"`
}

// CreateUploadAttachedMediaResponse is the response struct for api CreateUploadAttachedMedia
type CreateUploadAttachedMediaResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	MediaId       string `json:"MediaId" xml:"MediaId"`
	MediaURL      string `json:"MediaURL" xml:"MediaURL"`
	UploadAddress string `json:"UploadAddress" xml:"UploadAddress"`
	UploadAuth    string `json:"UploadAuth" xml:"UploadAuth"`
	FileURL       string `json:"FileURL" xml:"FileURL"`
}

// CreateCreateUploadAttachedMediaRequest creates a request to invoke CreateUploadAttachedMedia API
func CreateCreateUploadAttachedMediaRequest() (request *CreateUploadAttachedMediaRequest) {
	request = &CreateUploadAttachedMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "CreateUploadAttachedMedia", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateUploadAttachedMediaResponse creates a response to parse from CreateUploadAttachedMedia response
func CreateCreateUploadAttachedMediaResponse() (response *CreateUploadAttachedMediaResponse) {
	response = &CreateUploadAttachedMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
