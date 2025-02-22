package ehpc

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

// SubmitServerlessJob invokes the ehpc.SubmitServerlessJob API synchronously
func (client *Client) SubmitServerlessJob(request *SubmitServerlessJobRequest) (response *SubmitServerlessJobResponse, err error) {
	response = CreateSubmitServerlessJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitServerlessJobWithChan invokes the ehpc.SubmitServerlessJob API asynchronously
func (client *Client) SubmitServerlessJobWithChan(request *SubmitServerlessJobRequest) (<-chan *SubmitServerlessJobResponse, <-chan error) {
	responseChan := make(chan *SubmitServerlessJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitServerlessJob(request)
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

// SubmitServerlessJobWithCallback invokes the ehpc.SubmitServerlessJob API asynchronously
func (client *Client) SubmitServerlessJobWithCallback(request *SubmitServerlessJobRequest, callback func(response *SubmitServerlessJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitServerlessJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitServerlessJob(request)
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

// SubmitServerlessJobRequest is the request struct for api SubmitServerlessJob
type SubmitServerlessJobRequest struct {
	*requests.RpcRequest
	Container        SubmitServerlessJobContainer          `position:"Query" name:"Container"  type:"Struct"`
	Memory           requests.Float                        `position:"Query" name:"Memory"`
	DependsOn        *[]SubmitServerlessJobDependsOn       `position:"Query" name:"DependsOn"  type:"Json"`
	SpotPriceLimit   requests.Float                        `position:"Query" name:"SpotPriceLimit"`
	JobQueue         string                                `position:"Query" name:"JobQueue"`
	Timeout          requests.Integer                      `position:"Query" name:"Timeout"`
	JobUser          string                                `position:"Query" name:"JobUser"`
	InstanceType     *[]string                             `position:"Query" name:"InstanceType"  type:"Json"`
	JobName          string                                `position:"Query" name:"JobName"`
	JobPriority      requests.Integer                      `position:"Query" name:"JobPriority"`
	Cpu              requests.Float                        `position:"Query" name:"Cpu"`
	RamRoleName      string                                `position:"Query" name:"RamRoleName"`
	AcrRegistryInfo  *[]SubmitServerlessJobAcrRegistryInfo `position:"Query" name:"AcrRegistryInfo"  type:"Json"`
	ClusterId        string                                `position:"Query" name:"ClusterId"`
	SpotStrategy     string                                `position:"Query" name:"SpotStrategy"`
	VSwitchId        *[]string                             `position:"Query" name:"VSwitchId"  type:"Json"`
	Volume           *[]SubmitServerlessJobVolume          `position:"Query" name:"Volume"  type:"Json"`
	RetryStrategy    SubmitServerlessJobRetryStrategy      `position:"Query" name:"RetryStrategy"  type:"Struct"`
	EphemeralStorage requests.Integer                      `position:"Query" name:"EphemeralStorage"`
	ArrayProperties  SubmitServerlessJobArrayProperties    `position:"Query" name:"ArrayProperties"  type:"Struct"`
}

// SubmitServerlessJobDependsOn is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobDependsOn struct {
	JobId string `name:"JobId"`
	Type  string `name:"Type"`
}

// SubmitServerlessJobAcrRegistryInfo is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobAcrRegistryInfo struct {
	InstanceName string    `name:"InstanceName"`
	InstanceId   string    `name:"InstanceId"`
	RegionId     string    `name:"RegionId"`
	Domain       *[]string `name:"Domain" type:"Repeated"`
}

// SubmitServerlessJobVolume is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobVolume struct {
	FlexVolumeDriver  string `name:"FlexVolumeDriver"`
	NFSVolumePath     string `name:"NFSVolumePath"`
	FlexVolumeOptions string `name:"FlexVolumeOptions"`
	NFSVolumeReadOnly string `name:"NFSVolumeReadOnly"`
	NFSVolumeServer   string `name:"NFSVolumeServer"`
}

// SubmitServerlessJobContainer is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobContainer struct {
	VolumeMount    *[]SubmitServerlessJobContainerVolumeMountItem    `name:"VolumeMount" type:"Repeated"`
	Image          string                                            `name:"Image"`
	Port           *[]SubmitServerlessJobContainerPortItem           `name:"Port" type:"Repeated"`
	EnvironmentVar *[]SubmitServerlessJobContainerEnvironmentVarItem `name:"EnvironmentVar" type:"Repeated"`
	WorkingDir     string                                            `name:"WorkingDir"`
	Arg            *[]string                                         `name:"Arg" type:"Repeated"`
	Name           string                                            `name:"Name"`
	Gpu            string                                            `name:"Gpu"`
	Command        *[]string                                         `name:"Command" type:"Repeated"`
}

// SubmitServerlessJobRetryStrategy is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobRetryStrategy struct {
	EvaluateOnExit *[]SubmitServerlessJobRetryStrategyEvaluateOnExitItem `name:"EvaluateOnExit" type:"Repeated"`
	Attempts       string                                                `name:"Attempts"`
}

// SubmitServerlessJobArrayProperties is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobArrayProperties struct {
	IndexStart string `name:"IndexStart"`
	IndexStep  string `name:"IndexStep"`
	IndexEnd   string `name:"IndexEnd"`
}

// SubmitServerlessJobContainerVolumeMountItem is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobContainerVolumeMountItem struct {
	FlexVolumeDriver  string `name:"FlexVolumeDriver"`
	MountPath         string `name:"MountPath"`
	ReadOnly          string `name:"ReadOnly"`
	MountPropagation  string `name:"MountPropagation"`
	SubPath           string `name:"SubPath"`
	NFSVolumePath     string `name:"NFSVolumePath"`
	Type              string `name:"Type"`
	FlexVolumeOptions string `name:"FlexVolumeOptions"`
	NFSVolumeReadOnly string `name:"NFSVolumeReadOnly"`
	NFSVolumeServer   string `name:"NFSVolumeServer"`
}

// SubmitServerlessJobContainerPortItem is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobContainerPortItem struct {
	Protocol string `name:"Protocol"`
	Port     string `name:"Port"`
}

// SubmitServerlessJobContainerEnvironmentVarItem is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobContainerEnvironmentVarItem struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// SubmitServerlessJobRetryStrategyEvaluateOnExitItem is a repeated param struct in SubmitServerlessJobRequest
type SubmitServerlessJobRetryStrategyEvaluateOnExitItem struct {
	Action     string `name:"Action"`
	OnExitCode string `name:"OnExitCode"`
}

// SubmitServerlessJobResponse is the response struct for api SubmitServerlessJob
type SubmitServerlessJobResponse struct {
	*responses.BaseResponse
	JobId     string `json:"JobId" xml:"JobId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSubmitServerlessJobRequest creates a request to invoke SubmitServerlessJob API
func CreateSubmitServerlessJobRequest() (request *SubmitServerlessJobRequest) {
	request = &SubmitServerlessJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "SubmitServerlessJob", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitServerlessJobResponse creates a response to parse from SubmitServerlessJob response
func CreateSubmitServerlessJobResponse() (response *SubmitServerlessJobResponse) {
	response = &SubmitServerlessJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
