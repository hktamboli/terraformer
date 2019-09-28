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

// CreateUserGroup invokes the emr.CreateUserGroup API synchronously
// api document: https://help.aliyun.com/api/emr/createusergroup.html
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
	response = CreateCreateUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserGroupWithChan invokes the emr.CreateUserGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/createusergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserGroupWithChan(request *CreateUserGroupRequest) (<-chan *CreateUserGroupResponse, <-chan error) {
	responseChan := make(chan *CreateUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserGroup(request)
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

// CreateUserGroupWithCallback invokes the emr.CreateUserGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/createusergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserGroupWithCallback(request *CreateUserGroupRequest, callback func(response *CreateUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateUserGroup(request)
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

// CreateUserGroupRequest is the request struct for api CreateUserGroup
type CreateUserGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description     string           `position:"Query" name:"Description"`
	Type            string           `position:"Query" name:"Type"`
	Name            string           `position:"Query" name:"Name"`
	RoleIdList      *[]string        `position:"Query" name:"RoleIdList"  type:"Repeated"`
}

// CreateUserGroupResponse is the response struct for api CreateUserGroup
type CreateUserGroupResponse struct {
	*responses.BaseResponse
	Paging    bool   `json:"Paging" xml:"Paging"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateCreateUserGroupRequest creates a request to invoke CreateUserGroup API
func CreateCreateUserGroupRequest() (request *CreateUserGroupRequest) {
	request = &CreateUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateUserGroup", "emr", "openAPI")
	return
}

// CreateCreateUserGroupResponse creates a response to parse from CreateUserGroup response
func CreateCreateUserGroupResponse() (response *CreateUserGroupResponse) {
	response = &CreateUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
