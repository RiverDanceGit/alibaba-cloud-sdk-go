package kms

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

// GenerateDataKey invokes the kms.GenerateDataKey API synchronously
func (client *Client) GenerateDataKey(request *GenerateDataKeyRequest) (response *GenerateDataKeyResponse, err error) {
	response = CreateGenerateDataKeyResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateDataKeyWithChan invokes the kms.GenerateDataKey API asynchronously
func (client *Client) GenerateDataKeyWithChan(request *GenerateDataKeyRequest) (<-chan *GenerateDataKeyResponse, <-chan error) {
	responseChan := make(chan *GenerateDataKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateDataKey(request)
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

// GenerateDataKeyWithCallback invokes the kms.GenerateDataKey API asynchronously
func (client *Client) GenerateDataKeyWithCallback(request *GenerateDataKeyRequest, callback func(response *GenerateDataKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateDataKeyResponse
		var err error
		defer close(result)
		response, err = client.GenerateDataKey(request)
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

// GenerateDataKeyRequest is the request struct for api GenerateDataKey
type GenerateDataKeyRequest struct {
	*requests.RpcRequest
	KeyId             string           `position:"Query" name:"KeyId"`
	KeySpec           string           `position:"Query" name:"KeySpec"`
	NumberOfBytes     requests.Integer `position:"Query" name:"NumberOfBytes"`
	EncryptionContext string           `position:"Query" name:"EncryptionContext"`
}

// GenerateDataKeyResponse is the response struct for api GenerateDataKey
type GenerateDataKeyResponse struct {
	*responses.BaseResponse
	KeyVersionId   string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId          string `json:"KeyId" xml:"KeyId"`
	CiphertextBlob string `json:"CiphertextBlob" xml:"CiphertextBlob"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Plaintext      string `json:"Plaintext" xml:"Plaintext"`
}

// CreateGenerateDataKeyRequest creates a request to invoke GenerateDataKey API
func CreateGenerateDataKeyRequest() (request *GenerateDataKeyRequest) {
	request = &GenerateDataKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "GenerateDataKey", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGenerateDataKeyResponse creates a response to parse from GenerateDataKey response
func CreateGenerateDataKeyResponse() (response *GenerateDataKeyResponse) {
	response = &GenerateDataKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
