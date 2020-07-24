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

// FindSimilarFaces invokes the imm.FindSimilarFaces API synchronously
// api document: https://help.aliyun.com/api/imm/findsimilarfaces.html
func (client *Client) FindSimilarFaces(request *FindSimilarFacesRequest) (response *FindSimilarFacesResponse, err error) {
	response = CreateFindSimilarFacesResponse()
	err = client.DoAction(request, response)
	return
}

// FindSimilarFacesWithChan invokes the imm.FindSimilarFaces API asynchronously
// api document: https://help.aliyun.com/api/imm/findsimilarfaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindSimilarFacesWithChan(request *FindSimilarFacesRequest) (<-chan *FindSimilarFacesResponse, <-chan error) {
	responseChan := make(chan *FindSimilarFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindSimilarFaces(request)
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

// FindSimilarFacesWithCallback invokes the imm.FindSimilarFaces API asynchronously
// api document: https://help.aliyun.com/api/imm/findsimilarfaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindSimilarFacesWithCallback(request *FindSimilarFacesRequest, callback func(response *FindSimilarFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindSimilarFacesResponse
		var err error
		defer close(result)
		response, err = client.FindSimilarFaces(request)
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

// FindSimilarFacesRequest is the request struct for api FindSimilarFaces
type FindSimilarFacesRequest struct {
	*requests.RpcRequest
	Project        string           `position:"Query" name:"Project"`
	MinSimilarity  requests.Float   `position:"Query" name:"MinSimilarity"`
	ResponseFormat string           `position:"Query" name:"ResponseFormat"`
	Limit          requests.Integer `position:"Query" name:"Limit"`
	FaceId         string           `position:"Query" name:"FaceId"`
	ImageUri       string           `position:"Query" name:"ImageUri"`
	SetId          string           `position:"Query" name:"SetId"`
}

// FindSimilarFacesResponse is the response struct for api FindSimilarFaces
type FindSimilarFacesResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Faces     []FacesItem `json:"Faces" xml:"Faces"`
}

// CreateFindSimilarFacesRequest creates a request to invoke FindSimilarFaces API
func CreateFindSimilarFacesRequest() (request *FindSimilarFacesRequest) {
	request = &FindSimilarFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "FindSimilarFaces", "", "")
	request.Method = requests.POST
	return
}

// CreateFindSimilarFacesResponse creates a response to parse from FindSimilarFaces response
func CreateFindSimilarFacesResponse() (response *FindSimilarFacesResponse) {
	response = &FindSimilarFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
