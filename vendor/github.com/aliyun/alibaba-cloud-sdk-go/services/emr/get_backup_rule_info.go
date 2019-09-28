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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetBackupRuleInfo invokes the emr.GetBackupRuleInfo API synchronously
// api document: https://help.aliyun.com/api/emr/getbackupruleinfo.html
func (client *Client) GetBackupRuleInfo(request *GetBackupRuleInfoRequest) (response *GetBackupRuleInfoResponse, err error) {
	response = CreateGetBackupRuleInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetBackupRuleInfoWithChan invokes the emr.GetBackupRuleInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/getbackupruleinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBackupRuleInfoWithChan(request *GetBackupRuleInfoRequest) (<-chan *GetBackupRuleInfoResponse, <-chan error) {
	responseChan := make(chan *GetBackupRuleInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBackupRuleInfo(request)
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

// GetBackupRuleInfoWithCallback invokes the emr.GetBackupRuleInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/getbackupruleinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBackupRuleInfoWithCallback(request *GetBackupRuleInfoRequest, callback func(response *GetBackupRuleInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBackupRuleInfoResponse
		var err error
		defer close(result)
		response, err = client.GetBackupRuleInfo(request)
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

// GetBackupRuleInfoRequest is the request struct for api GetBackupRuleInfo
type GetBackupRuleInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BackupRuleId    string           `position:"Query" name:"BackupRuleId"`
}

// GetBackupRuleInfoResponse is the response struct for api GetBackupRuleInfo
type GetBackupRuleInfoResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Id               string `json:"Id" xml:"Id"`
	Name             string `json:"Name" xml:"Name"`
	Description      string `json:"Description" xml:"Description"`
	MetadataType     string `json:"MetadataType" xml:"MetadataType"`
	BackupMethodType string `json:"BackupMethodType" xml:"BackupMethodType"`
	BackupPlanId     string `json:"BackupPlanId" xml:"BackupPlanId"`
}

// CreateGetBackupRuleInfoRequest creates a request to invoke GetBackupRuleInfo API
func CreateGetBackupRuleInfoRequest() (request *GetBackupRuleInfoRequest) {
	request = &GetBackupRuleInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "GetBackupRuleInfo", "emr", "openAPI")
	return
}

// CreateGetBackupRuleInfoResponse creates a response to parse from GetBackupRuleInfo response
func CreateGetBackupRuleInfoResponse() (response *GetBackupRuleInfoResponse) {
	response = &GetBackupRuleInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
