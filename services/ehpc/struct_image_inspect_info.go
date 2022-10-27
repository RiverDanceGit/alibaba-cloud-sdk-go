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

// ImageInspectInfo is a nested struct in ehpc response
type ImageInspectInfo struct {
	BuildArch        string `json:"BuildArch" xml:"BuildArch"`
	BuildDate        string `json:"BuildDate" xml:"BuildDate"`
	BootStrap        string `json:"BootStrap" xml:"BootStrap"`
	DefFrom          string `json:"DefFrom" xml:"DefFrom"`
	ContainerVersion string `json:"ContainerVersion" xml:"ContainerVersion"`
	SchemaVersion    string `json:"SchemaVersion" xml:"SchemaVersion"`
}
