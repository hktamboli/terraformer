package emr

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

// JobMigrateInfo is a nested struct in emr response
type JobMigrateInfo struct {
	Id              string `json:"Id" xml:"Id"`
	Name            string `json:"Name" xml:"Name"`
	Type            string `json:"Type" xml:"Type"`
	MaxRetry        int    `json:"MaxRetry" xml:"MaxRetry"`
	FailedAction    string `json:"FailedAction" xml:"FailedAction"`
	Params          string `json:"Params" xml:"Params"`
	NewId           string `json:"NewId" xml:"NewId"`
	PremigratedDate string `json:"PremigratedDate" xml:"PremigratedDate"`
	MigratedDate    string `json:"MigratedDate" xml:"MigratedDate"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
}
