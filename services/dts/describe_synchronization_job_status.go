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

// DescribeSynchronizationJobStatus invokes the dts.DescribeSynchronizationJobStatus API synchronously
func (client *Client) DescribeSynchronizationJobStatus(request *DescribeSynchronizationJobStatusRequest) (response *DescribeSynchronizationJobStatusResponse, err error) {
	response = CreateDescribeSynchronizationJobStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSynchronizationJobStatusWithChan invokes the dts.DescribeSynchronizationJobStatus API asynchronously
func (client *Client) DescribeSynchronizationJobStatusWithChan(request *DescribeSynchronizationJobStatusRequest) (<-chan *DescribeSynchronizationJobStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSynchronizationJobStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSynchronizationJobStatus(request)
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

// DescribeSynchronizationJobStatusWithCallback invokes the dts.DescribeSynchronizationJobStatus API asynchronously
func (client *Client) DescribeSynchronizationJobStatusWithCallback(request *DescribeSynchronizationJobStatusRequest, callback func(response *DescribeSynchronizationJobStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSynchronizationJobStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSynchronizationJobStatus(request)
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

// DescribeSynchronizationJobStatusRequest is the request struct for api DescribeSynchronizationJobStatus
type DescribeSynchronizationJobStatusRequest struct {
	*requests.RpcRequest
	ClientToken              string `position:"Query" name:"ClientToken"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
	SynchronizationJobId     string `position:"Query" name:"SynchronizationJobId"`
	AccountId                string `position:"Query" name:"AccountId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
}

// DescribeSynchronizationJobStatusResponse is the response struct for api DescribeSynchronizationJobStatus
type DescribeSynchronizationJobStatusResponse struct {
	*responses.BaseResponse
	Checkpoint                    string                                           `json:"Checkpoint" xml:"Checkpoint"`
	DataInitialization            string                                           `json:"DataInitialization" xml:"DataInitialization"`
	Delay                         string                                           `json:"Delay" xml:"Delay"`
	DelayMillis                   int64                                            `json:"DelayMillis" xml:"DelayMillis"`
	ErrCode                       string                                           `json:"ErrCode" xml:"ErrCode"`
	ErrMessage                    string                                           `json:"ErrMessage" xml:"ErrMessage"`
	ErrorMessage                  string                                           `json:"ErrorMessage" xml:"ErrorMessage"`
	ExpireTime                    string                                           `json:"ExpireTime" xml:"ExpireTime"`
	PayType                       string                                           `json:"PayType" xml:"PayType"`
	RequestId                     string                                           `json:"RequestId" xml:"RequestId"`
	Status                        string                                           `json:"Status" xml:"Status"`
	StructureInitialization       string                                           `json:"StructureInitialization" xml:"StructureInitialization"`
	Success                       string                                           `json:"Success" xml:"Success"`
	SynchronizationDirection      string                                           `json:"SynchronizationDirection" xml:"SynchronizationDirection"`
	SynchronizationJobClass       string                                           `json:"SynchronizationJobClass" xml:"SynchronizationJobClass"`
	SynchronizationJobId          string                                           `json:"SynchronizationJobId" xml:"SynchronizationJobId"`
	SynchronizationJobName        string                                           `json:"SynchronizationJobName" xml:"SynchronizationJobName"`
	TaskId                        string                                           `json:"TaskId" xml:"TaskId"`
	DataInitializationStatus      DataInitializationStatus                         `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	DataSynchronizationStatus     DataSynchronizationStatus                        `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	DestinationEndpoint           DestinationEndpoint                              `json:"DestinationEndpoint" xml:"DestinationEndpoint"`
	Performance                   Performance                                      `json:"Performance" xml:"Performance"`
	PrecheckStatus                PrecheckStatusInDescribeSynchronizationJobStatus `json:"PrecheckStatus" xml:"PrecheckStatus"`
	SourceEndpoint                SourceEndpoint                                   `json:"SourceEndpoint" xml:"SourceEndpoint"`
	StructureInitializationStatus StructureInitializationStatus                    `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
	SynchronizationObjects        []SynchronizationObject                          `json:"SynchronizationObjects" xml:"SynchronizationObjects"`
}

// CreateDescribeSynchronizationJobStatusRequest creates a request to invoke DescribeSynchronizationJobStatus API
func CreateDescribeSynchronizationJobStatusRequest() (request *DescribeSynchronizationJobStatusRequest) {
	request = &DescribeSynchronizationJobStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeSynchronizationJobStatus", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSynchronizationJobStatusResponse creates a response to parse from DescribeSynchronizationJobStatus response
func CreateDescribeSynchronizationJobStatusResponse() (response *DescribeSynchronizationJobStatusResponse) {
	response = &DescribeSynchronizationJobStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
