package cloudapi

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

// DescribeLogConfig invokes the cloudapi.DescribeLogConfig API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describelogconfig.html
func (client *Client) DescribeLogConfig(request *DescribeLogConfigRequest) (response *DescribeLogConfigResponse, err error) {
	response = CreateDescribeLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogConfigWithChan invokes the cloudapi.DescribeLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describelogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLogConfigWithChan(request *DescribeLogConfigRequest) (<-chan *DescribeLogConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogConfig(request)
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

// DescribeLogConfigWithCallback invokes the cloudapi.DescribeLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describelogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLogConfigWithCallback(request *DescribeLogConfigRequest, callback func(response *DescribeLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogConfig(request)
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

// DescribeLogConfigRequest is the request struct for api DescribeLogConfig
type DescribeLogConfigRequest struct {
	*requests.RpcRequest
	LogType       string `position:"Query" name:"LogType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeLogConfigResponse is the response struct for api DescribeLogConfig
type DescribeLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	LogInfos  LogInfos `json:"LogInfos" xml:"LogInfos"`
}

// CreateDescribeLogConfigRequest creates a request to invoke DescribeLogConfig API
func CreateDescribeLogConfigRequest() (request *DescribeLogConfigRequest) {
	request = &DescribeLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeLogConfig", "apigateway", "openAPI")
	return
}

// CreateDescribeLogConfigResponse creates a response to parse from DescribeLogConfig response
func CreateDescribeLogConfigResponse() (response *DescribeLogConfigResponse) {
	response = &DescribeLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
