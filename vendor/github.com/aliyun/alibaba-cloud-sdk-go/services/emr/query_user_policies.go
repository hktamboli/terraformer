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

// QueryUserPolicies invokes the emr.QueryUserPolicies API synchronously
// api document: https://help.aliyun.com/api/emr/queryuserpolicies.html
func (client *Client) QueryUserPolicies(request *QueryUserPoliciesRequest) (response *QueryUserPoliciesResponse, err error) {
	response = CreateQueryUserPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUserPoliciesWithChan invokes the emr.QueryUserPolicies API asynchronously
// api document: https://help.aliyun.com/api/emr/queryuserpolicies.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUserPoliciesWithChan(request *QueryUserPoliciesRequest) (<-chan *QueryUserPoliciesResponse, <-chan error) {
	responseChan := make(chan *QueryUserPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUserPolicies(request)
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

// QueryUserPoliciesWithCallback invokes the emr.QueryUserPolicies API asynchronously
// api document: https://help.aliyun.com/api/emr/queryuserpolicies.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUserPoliciesWithCallback(request *QueryUserPoliciesRequest, callback func(response *QueryUserPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUserPoliciesResponse
		var err error
		defer close(result)
		response, err = client.QueryUserPolicies(request)
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

// QueryUserPoliciesRequest is the request struct for api QueryUserPolicies
type QueryUserPoliciesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceId      string           `position:"Query" name:"ResourceId"`
	ResourceType    string           `position:"Query" name:"ResourceType"`
}

// QueryUserPoliciesResponse is the response struct for api QueryUserPolicies
type QueryUserPoliciesResponse struct {
	*responses.BaseResponse
	Paging    bool                    `json:"Paging" xml:"Paging"`
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Data      DataInQueryUserPolicies `json:"Data" xml:"Data"`
}

// CreateQueryUserPoliciesRequest creates a request to invoke QueryUserPolicies API
func CreateQueryUserPoliciesRequest() (request *QueryUserPoliciesRequest) {
	request = &QueryUserPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "QueryUserPolicies", "emr", "openAPI")
	return
}

// CreateQueryUserPoliciesResponse creates a response to parse from QueryUserPolicies response
func CreateQueryUserPoliciesResponse() (response *QueryUserPoliciesResponse) {
	response = &QueryUserPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
