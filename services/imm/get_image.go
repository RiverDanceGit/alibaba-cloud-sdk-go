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

// GetImage invokes the imm.GetImage API synchronously
func (client *Client) GetImage(request *GetImageRequest) (response *GetImageResponse, err error) {
	response = CreateGetImageResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageWithChan invokes the imm.GetImage API asynchronously
func (client *Client) GetImageWithChan(request *GetImageRequest) (<-chan *GetImageResponse, <-chan error) {
	responseChan := make(chan *GetImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImage(request)
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

// GetImageWithCallback invokes the imm.GetImage API asynchronously
func (client *Client) GetImageWithCallback(request *GetImageRequest, callback func(response *GetImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageResponse
		var err error
		defer close(result)
		response, err = client.GetImage(request)
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

// GetImageRequest is the request struct for api GetImage
type GetImageRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	ImageUri string `position:"Query" name:"ImageUri"`
	SetId    string `position:"Query" name:"SetId"`
}

// GetImageResponse is the response struct for api GetImage
type GetImageResponse struct {
	*responses.BaseResponse
	RequestId                    string                   `json:"RequestId" xml:"RequestId"`
	SetId                        string                   `json:"SetId" xml:"SetId"`
	ImageUri                     string                   `json:"ImageUri" xml:"ImageUri"`
	RemarksA                     string                   `json:"RemarksA" xml:"RemarksA"`
	RemarksB                     string                   `json:"RemarksB" xml:"RemarksB"`
	CreateTime                   string                   `json:"CreateTime" xml:"CreateTime"`
	ModifyTime                   string                   `json:"ModifyTime" xml:"ModifyTime"`
	ImageWidth                   int                      `json:"ImageWidth" xml:"ImageWidth"`
	ImageHeight                  int                      `json:"ImageHeight" xml:"ImageHeight"`
	ImageFormat                  string                   `json:"ImageFormat" xml:"ImageFormat"`
	Exif                         string                   `json:"Exif" xml:"Exif"`
	FileSize                     int                      `json:"FileSize" xml:"FileSize"`
	ImageTime                    string                   `json:"ImageTime" xml:"ImageTime"`
	Orientation                  string                   `json:"Orientation" xml:"Orientation"`
	SourceType                   string                   `json:"SourceType" xml:"SourceType"`
	SourceUri                    string                   `json:"SourceUri" xml:"SourceUri"`
	SourcePosition               string                   `json:"SourcePosition" xml:"SourcePosition"`
	FacesStatus                  string                   `json:"FacesStatus" xml:"FacesStatus"`
	FacesModifyTime              string                   `json:"FacesModifyTime" xml:"FacesModifyTime"`
	Location                     string                   `json:"Location" xml:"Location"`
	OCRStatus                    string                   `json:"OCRStatus" xml:"OCRStatus"`
	OCRModifyTime                string                   `json:"OCRModifyTime" xml:"OCRModifyTime"`
	OCRFailReason                string                   `json:"OCRFailReason" xml:"OCRFailReason"`
	FacesFailReason              string                   `json:"FacesFailReason" xml:"FacesFailReason"`
	TagsFailReason               string                   `json:"TagsFailReason" xml:"TagsFailReason"`
	TagsModifyTime               string                   `json:"TagsModifyTime" xml:"TagsModifyTime"`
	TagsStatus                   string                   `json:"TagsStatus" xml:"TagsStatus"`
	RemarksC                     string                   `json:"RemarksC" xml:"RemarksC"`
	RemarksD                     string                   `json:"RemarksD" xml:"RemarksD"`
	ExternalId                   string                   `json:"ExternalId" xml:"ExternalId"`
	AddressModifyTime            string                   `json:"AddressModifyTime" xml:"AddressModifyTime"`
	AddressStatus                string                   `json:"AddressStatus" xml:"AddressStatus"`
	AddressFailReason            string                   `json:"AddressFailReason" xml:"AddressFailReason"`
	RemarksArrayA                string                   `json:"RemarksArrayA" xml:"RemarksArrayA"`
	RemarksArrayB                string                   `json:"RemarksArrayB" xml:"RemarksArrayB"`
	ImageQualityModifyTime       string                   `json:"ImageQualityModifyTime" xml:"ImageQualityModifyTime"`
	ImageQualityFailReason       string                   `json:"ImageQualityFailReason" xml:"ImageQualityFailReason"`
	ImageQualityStatus           string                   `json:"ImageQualityStatus" xml:"ImageQualityStatus"`
	CroppingSuggestionStatus     string                   `json:"CroppingSuggestionStatus" xml:"CroppingSuggestionStatus"`
	CroppingSuggestionFailReason string                   `json:"CroppingSuggestionFailReason" xml:"CroppingSuggestionFailReason"`
	CroppingSuggestionModifyTime string                   `json:"CroppingSuggestionModifyTime" xml:"CroppingSuggestionModifyTime"`
	ImageQuality                 ImageQuality             `json:"ImageQuality" xml:"ImageQuality"`
	Address                      Address                  `json:"Address" xml:"Address"`
	CroppingSuggestion           []CroppingSuggestionItem `json:"CroppingSuggestion" xml:"CroppingSuggestion"`
	Faces                        []FacesItemInGetImage    `json:"Faces" xml:"Faces"`
	OCR                          []OCRItem                `json:"OCR" xml:"OCR"`
	Tags                         []TagsItem               `json:"Tags" xml:"Tags"`
}

// CreateGetImageRequest creates a request to invoke GetImage API
func CreateGetImageRequest() (request *GetImageRequest) {
	request = &GetImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetImage", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetImageResponse creates a response to parse from GetImage response
func CreateGetImageResponse() (response *GetImageResponse) {
	response = &GetImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
