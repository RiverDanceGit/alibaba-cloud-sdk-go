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

// ListIntervalSkillGroupReport invokes the ccc.ListIntervalSkillGroupReport API synchronously
func (client *Client) ListIntervalSkillGroupReport(request *ListIntervalSkillGroupReportRequest) (response *ListIntervalSkillGroupReportResponse, err error) {
	response = CreateListIntervalSkillGroupReportResponse()
	err = client.DoAction(request, response)
	return
}

// ListIntervalSkillGroupReportWithChan invokes the ccc.ListIntervalSkillGroupReport API asynchronously
func (client *Client) ListIntervalSkillGroupReportWithChan(request *ListIntervalSkillGroupReportRequest) (<-chan *ListIntervalSkillGroupReportResponse, <-chan error) {
	responseChan := make(chan *ListIntervalSkillGroupReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIntervalSkillGroupReport(request)
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

// ListIntervalSkillGroupReportWithCallback invokes the ccc.ListIntervalSkillGroupReport API asynchronously
func (client *Client) ListIntervalSkillGroupReportWithCallback(request *ListIntervalSkillGroupReportRequest, callback func(response *ListIntervalSkillGroupReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIntervalSkillGroupReportResponse
		var err error
		defer close(result)
		response, err = client.ListIntervalSkillGroupReport(request)
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

// ListIntervalSkillGroupReportRequest is the request struct for api ListIntervalSkillGroupReport
type ListIntervalSkillGroupReportRequest struct {
	*requests.RpcRequest
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	Interval     string           `position:"Query" name:"Interval"`
}

// ListIntervalSkillGroupReportResponse is the response struct for api ListIntervalSkillGroupReport
type ListIntervalSkillGroupReportResponse struct {
	*responses.BaseResponse
	Code           string     `json:"Code" xml:"Code"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateListIntervalSkillGroupReportRequest creates a request to invoke ListIntervalSkillGroupReport API
func CreateListIntervalSkillGroupReportRequest() (request *ListIntervalSkillGroupReportRequest) {
	request = &ListIntervalSkillGroupReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListIntervalSkillGroupReport", "", "")
	request.Method = requests.POST
	return
}

// CreateListIntervalSkillGroupReportResponse creates a response to parse from ListIntervalSkillGroupReport response
func CreateListIntervalSkillGroupReportResponse() (response *ListIntervalSkillGroupReportResponse) {
	response = &ListIntervalSkillGroupReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
