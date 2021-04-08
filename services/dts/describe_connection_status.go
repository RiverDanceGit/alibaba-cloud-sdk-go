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

// DescribeConnectionStatus invokes the dts.DescribeConnectionStatus API synchronously
func (client *Client) DescribeConnectionStatus(request *DescribeConnectionStatusRequest) (response *DescribeConnectionStatusResponse, err error) {
	response = CreateDescribeConnectionStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConnectionStatusWithChan invokes the dts.DescribeConnectionStatus API asynchronously
func (client *Client) DescribeConnectionStatusWithChan(request *DescribeConnectionStatusRequest) (<-chan *DescribeConnectionStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeConnectionStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConnectionStatus(request)
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

// DescribeConnectionStatusWithCallback invokes the dts.DescribeConnectionStatus API asynchronously
func (client *Client) DescribeConnectionStatusWithCallback(request *DescribeConnectionStatusRequest, callback func(response *DescribeConnectionStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConnectionStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeConnectionStatus(request)
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

// DescribeConnectionStatusRequest is the request struct for api DescribeConnectionStatus
type DescribeConnectionStatusRequest struct {
	*requests.RpcRequest
	SourceEndpointRegion            string `position:"Query" name:"SourceEndpointRegion"`
	SourceEndpointArchitecture      string `position:"Query" name:"SourceEndpointArchitecture"`
	DestinationEndpointInstanceType string `position:"Query" name:"DestinationEndpointInstanceType"`
	SourceEndpointInstanceID        string `position:"Query" name:"SourceEndpointInstanceID"`
	SourceEndpointUserName          string `position:"Query" name:"SourceEndpointUserName"`
	SourceEndpointDatabaseName      string `position:"Query" name:"SourceEndpointDatabaseName"`
	DestinationEndpointRegion       string `position:"Query" name:"DestinationEndpointRegion"`
	SourceEndpointIP                string `position:"Query" name:"SourceEndpointIP"`
	DestinationEndpointUserName     string `position:"Query" name:"DestinationEndpointUserName"`
	DestinationEndpointArchitecture string `position:"Query" name:"DestinationEndpointArchitecture"`
	DestinationEndpointOracleSID    string `position:"Query" name:"DestinationEndpointOracleSID"`
	DestinationEndpointEngineName   string `position:"Query" name:"DestinationEndpointEngineName"`
	DestinationEndpointInstanceID   string `position:"Query" name:"DestinationEndpointInstanceID"`
	DestinationEndpointPort         string `position:"Query" name:"DestinationEndpointPort"`
	SourceEndpointPassword          string `position:"Query" name:"SourceEndpointPassword"`
	SourceEndpointPort              string `position:"Query" name:"SourceEndpointPort"`
	DestinationEndpointIP           string `position:"Query" name:"DestinationEndpointIP"`
	SourceEndpointInstanceType      string `position:"Query" name:"SourceEndpointInstanceType"`
	SourceEndpointOracleSID         string `position:"Query" name:"SourceEndpointOracleSID"`
	DestinationEndpointDatabaseName string `position:"Query" name:"DestinationEndpointDatabaseName"`
	DestinationEndpointPassword     string `position:"Query" name:"DestinationEndpointPassword"`
	SourceEndpointEngineName        string `position:"Query" name:"SourceEndpointEngineName"`
}

// DescribeConnectionStatusResponse is the response struct for api DescribeConnectionStatus
type DescribeConnectionStatusResponse struct {
	*responses.BaseResponse
	ErrCode                     string                 `json:"ErrCode" xml:"ErrCode"`
	ErrMessage                  string                 `json:"ErrMessage" xml:"ErrMessage"`
	RequestId                   string                 `json:"RequestId" xml:"RequestId"`
	Success                     string                 `json:"Success" xml:"Success"`
	SourceConnectionStatus      map[string]interface{} `json:"SourceConnectionStatus" xml:"SourceConnectionStatus"`
	DestinationConnectionStatus map[string]interface{} `json:"DestinationConnectionStatus" xml:"DestinationConnectionStatus"`
}

// CreateDescribeConnectionStatusRequest creates a request to invoke DescribeConnectionStatus API
func CreateDescribeConnectionStatusRequest() (request *DescribeConnectionStatusRequest) {
	request = &DescribeConnectionStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeConnectionStatus", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeConnectionStatusResponse creates a response to parse from DescribeConnectionStatus response
func CreateDescribeConnectionStatusResponse() (response *DescribeConnectionStatusResponse) {
	response = &DescribeConnectionStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
