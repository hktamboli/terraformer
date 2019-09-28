package drds

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

// DescribeDrdsDbRelationInfo invokes the drds.DescribeDrdsDbRelationInfo API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbrelationinfo.html
func (client *Client) DescribeDrdsDbRelationInfo(request *DescribeDrdsDbRelationInfoRequest) (response *DescribeDrdsDbRelationInfoResponse, err error) {
	response = CreateDescribeDrdsDbRelationInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDbRelationInfoWithChan invokes the drds.DescribeDrdsDbRelationInfo API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbrelationinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbRelationInfoWithChan(request *DescribeDrdsDbRelationInfoRequest) (<-chan *DescribeDrdsDbRelationInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDbRelationInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDbRelationInfo(request)
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

// DescribeDrdsDbRelationInfoWithCallback invokes the drds.DescribeDrdsDbRelationInfo API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbrelationinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbRelationInfoWithCallback(request *DescribeDrdsDbRelationInfoRequest, callback func(response *DescribeDrdsDbRelationInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDbRelationInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDbRelationInfo(request)
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

// DescribeDrdsDbRelationInfoRequest is the request struct for api DescribeDrdsDbRelationInfo
type DescribeDrdsDbRelationInfoRequest struct {
	*requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeDrdsDbRelationInfoResponse is the response struct for api DescribeDrdsDbRelationInfo
type DescribeDrdsDbRelationInfoResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Success   bool                             `json:"Success" xml:"Success"`
	Data      DataInDescribeDrdsDbRelationInfo `json:"Data" xml:"Data"`
}

// CreateDescribeDrdsDbRelationInfoRequest creates a request to invoke DescribeDrdsDbRelationInfo API
func CreateDescribeDrdsDbRelationInfoRequest() (request *DescribeDrdsDbRelationInfoRequest) {
	request = &DescribeDrdsDbRelationInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDbRelationInfo", "drds", "openAPI")
	return
}

// CreateDescribeDrdsDbRelationInfoResponse creates a response to parse from DescribeDrdsDbRelationInfo response
func CreateDescribeDrdsDbRelationInfoResponse() (response *DescribeDrdsDbRelationInfoResponse) {
	response = &DescribeDrdsDbRelationInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
