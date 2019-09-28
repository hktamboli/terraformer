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

// ModifyJobExecutionPlanFolder invokes the emr.ModifyJobExecutionPlanFolder API synchronously
// api document: https://help.aliyun.com/api/emr/modifyjobexecutionplanfolder.html
func (client *Client) ModifyJobExecutionPlanFolder(request *ModifyJobExecutionPlanFolderRequest) (response *ModifyJobExecutionPlanFolderResponse, err error) {
	response = CreateModifyJobExecutionPlanFolderResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyJobExecutionPlanFolderWithChan invokes the emr.ModifyJobExecutionPlanFolder API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyjobexecutionplanfolder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyJobExecutionPlanFolderWithChan(request *ModifyJobExecutionPlanFolderRequest) (<-chan *ModifyJobExecutionPlanFolderResponse, <-chan error) {
	responseChan := make(chan *ModifyJobExecutionPlanFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyJobExecutionPlanFolder(request)
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

// ModifyJobExecutionPlanFolderWithCallback invokes the emr.ModifyJobExecutionPlanFolder API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyjobexecutionplanfolder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyJobExecutionPlanFolderWithCallback(request *ModifyJobExecutionPlanFolderRequest, callback func(response *ModifyJobExecutionPlanFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyJobExecutionPlanFolderResponse
		var err error
		defer close(result)
		response, err = client.ModifyJobExecutionPlanFolder(request)
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

// ModifyJobExecutionPlanFolderRequest is the request struct for api ModifyJobExecutionPlanFolder
type ModifyJobExecutionPlanFolderRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Name            string           `position:"Query" name:"Name"`
	Id              requests.Integer `position:"Query" name:"Id"`
	ParentId        requests.Integer `position:"Query" name:"ParentId"`
}

// ModifyJobExecutionPlanFolderResponse is the response struct for api ModifyJobExecutionPlanFolder
type ModifyJobExecutionPlanFolderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	FolderId  string `json:"FolderId" xml:"FolderId"`
}

// CreateModifyJobExecutionPlanFolderRequest creates a request to invoke ModifyJobExecutionPlanFolder API
func CreateModifyJobExecutionPlanFolderRequest() (request *ModifyJobExecutionPlanFolderRequest) {
	request = &ModifyJobExecutionPlanFolderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyJobExecutionPlanFolder", "emr", "openAPI")
	return
}

// CreateModifyJobExecutionPlanFolderResponse creates a response to parse from ModifyJobExecutionPlanFolder response
func CreateModifyJobExecutionPlanFolderResponse() (response *ModifyJobExecutionPlanFolderResponse) {
	response = &ModifyJobExecutionPlanFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
