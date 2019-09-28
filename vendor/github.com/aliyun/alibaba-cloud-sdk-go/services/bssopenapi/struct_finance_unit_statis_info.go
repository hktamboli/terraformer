package bssopenapi

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

// FinanceUnitStatisInfo is a nested struct in bssopenapi response
type FinanceUnitStatisInfo struct {
	ResourceCount           int64 `json:"ResourceCount" xml:"ResourceCount"`
	ResourceGroupCount      int64 `json:"ResourceGroupCount" xml:"ResourceGroupCount"`
	SubUnitCount            int64 `json:"SubUnitCount" xml:"SubUnitCount"`
	UserCount               int64 `json:"UserCount" xml:"UserCount"`
	TotalResourceCount      int64 `json:"TotalResourceCount" xml:"TotalResourceCount"`
	TotalUserCount          int64 `json:"TotalUserCount" xml:"TotalUserCount"`
	TotalResourceGroupCount int64 `json:"TotalResourceGroupCount" xml:"TotalResourceGroupCount"`
}
