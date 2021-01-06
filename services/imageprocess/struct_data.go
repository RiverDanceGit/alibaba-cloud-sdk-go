package imageprocess

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

// Data is a nested struct in imageprocess response
type Data struct {
	Text              string                 `json:"Text" xml:"Text"`
	ErrorMessage      string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	NormalProbability string                 `json:"NormalProbability" xml:"NormalProbability"`
	NewProbability    string                 `json:"NewProbability" xml:"NewProbability"`
	Score             string                 `json:"Score" xml:"Score"`
	DUrl              string                 `json:"DUrl" xml:"DUrl"`
	ImageUrl          string                 `json:"ImageUrl" xml:"ImageUrl"`
	Mask              string                 `json:"Mask" xml:"Mask"`
	ErrorCode         string                 `json:"ErrorCode" xml:"ErrorCode"`
	JobId             string                 `json:"JobId" xml:"JobId"`
	OrgId             string                 `json:"OrgId" xml:"OrgId"`
	Result            string                 `json:"Result" xml:"Result"`
	OtherProbability  string                 `json:"OtherProbability" xml:"OtherProbability"`
	Words             int64                  `json:"Words" xml:"Words"`
	LesionRatio       string                 `json:"LesionRatio" xml:"LesionRatio"`
	NUrl              string                 `json:"NUrl" xml:"NUrl"`
	Status            string                 `json:"Status" xml:"Status"`
	Answer            string                 `json:"Answer" xml:"Answer"`
	Results           map[string]interface{} `json:"Results" xml:"Results"`
	OrgName           string                 `json:"OrgName" xml:"OrgName"`
	ResultURL         string                 `json:"ResultURL" xml:"ResultURL"`
	ResultUrl         string                 `json:"ResultUrl" xml:"ResultUrl"`
	Spacing           []float64              `json:"Spacing" xml:"Spacing"`
	SimilarQuestion   []string               `json:"SimilarQuestion" xml:"SimilarQuestion"`
	Origin            []float64              `json:"Origin" xml:"Origin"`
	LungNodule        LungNodule             `json:"LungNodule" xml:"LungNodule"`
	CACS              CACS                   `json:"CACS" xml:"CACS"`
	Covid             Covid                  `json:"Covid" xml:"Covid"`
	Fractures         []FracturesItem        `json:"Fractures" xml:"Fractures"`
	Discs             []Disc                 `json:"Discs" xml:"Discs"`
	Detections        []DetectionsItem       `json:"Detections" xml:"Detections"`
	Series            []Serie                `json:"Series" xml:"Series"`
	Vertebras         []Vertebra             `json:"Vertebras" xml:"Vertebras"`
	KLDetections      []KLDetectionsItem     `json:"KLDetections" xml:"KLDetections"`
	KeyPoints         []KeyPointsItem        `json:"KeyPoints" xml:"KeyPoints"`
}
