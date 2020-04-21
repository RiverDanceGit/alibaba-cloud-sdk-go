package ivpd

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

// ListPackageDesignModelTypes invokes the ivpd.ListPackageDesignModelTypes API synchronously
// api document: https://help.aliyun.com/api/ivpd/listpackagedesignmodeltypes.html
func (client *Client) ListPackageDesignModelTypes(request *ListPackageDesignModelTypesRequest) (response *ListPackageDesignModelTypesResponse, err error) {
	response = CreateListPackageDesignModelTypesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPackageDesignModelTypesWithChan invokes the ivpd.ListPackageDesignModelTypes API asynchronously
// api document: https://help.aliyun.com/api/ivpd/listpackagedesignmodeltypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPackageDesignModelTypesWithChan(request *ListPackageDesignModelTypesRequest) (<-chan *ListPackageDesignModelTypesResponse, <-chan error) {
	responseChan := make(chan *ListPackageDesignModelTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPackageDesignModelTypes(request)
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

// ListPackageDesignModelTypesWithCallback invokes the ivpd.ListPackageDesignModelTypes API asynchronously
// api document: https://help.aliyun.com/api/ivpd/listpackagedesignmodeltypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPackageDesignModelTypesWithCallback(request *ListPackageDesignModelTypesRequest, callback func(response *ListPackageDesignModelTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPackageDesignModelTypesResponse
		var err error
		defer close(result)
		response, err = client.ListPackageDesignModelTypes(request)
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

// ListPackageDesignModelTypesRequest is the request struct for api ListPackageDesignModelTypes
type ListPackageDesignModelTypesRequest struct {
	*requests.RpcRequest
}

// ListPackageDesignModelTypesResponse is the response struct for api ListPackageDesignModelTypes
type ListPackageDesignModelTypesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListPackageDesignModelTypesRequest creates a request to invoke ListPackageDesignModelTypes API
func CreateListPackageDesignModelTypesRequest() (request *ListPackageDesignModelTypesRequest) {
	request = &ListPackageDesignModelTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "ListPackageDesignModelTypes", "", "")
	return
}

// CreateListPackageDesignModelTypesResponse creates a response to parse from ListPackageDesignModelTypes response
func CreateListPackageDesignModelTypesResponse() (response *ListPackageDesignModelTypesResponse) {
	response = &ListPackageDesignModelTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
