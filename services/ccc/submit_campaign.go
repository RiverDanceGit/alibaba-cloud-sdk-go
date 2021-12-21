package ccc

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

// SubmitCampaign invokes the ccc.SubmitCampaign API synchronously
func (client *Client) SubmitCampaign(request *SubmitCampaignRequest) (response *SubmitCampaignResponse, err error) {
	response = CreateSubmitCampaignResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitCampaignWithChan invokes the ccc.SubmitCampaign API asynchronously
func (client *Client) SubmitCampaignWithChan(request *SubmitCampaignRequest) (<-chan *SubmitCampaignResponse, <-chan error) {
	responseChan := make(chan *SubmitCampaignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitCampaign(request)
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

// SubmitCampaignWithCallback invokes the ccc.SubmitCampaign API asynchronously
func (client *Client) SubmitCampaignWithCallback(request *SubmitCampaignRequest, callback func(response *SubmitCampaignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitCampaignResponse
		var err error
		defer close(result)
		response, err = client.SubmitCampaign(request)
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

// SubmitCampaignRequest is the request struct for api SubmitCampaign
type SubmitCampaignRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	CampaignId string `position:"Query" name:"CampaignId"`
}

// SubmitCampaignResponse is the response struct for api SubmitCampaign
type SubmitCampaignResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateSubmitCampaignRequest creates a request to invoke SubmitCampaign API
func CreateSubmitCampaignRequest() (request *SubmitCampaignRequest) {
	request = &SubmitCampaignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "SubmitCampaign", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitCampaignResponse creates a response to parse from SubmitCampaign response
func CreateSubmitCampaignResponse() (response *SubmitCampaignResponse) {
	response = &SubmitCampaignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
