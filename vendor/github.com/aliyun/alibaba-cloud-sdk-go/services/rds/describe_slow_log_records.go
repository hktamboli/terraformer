package rds

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

// DescribeSlowLogRecords invokes the rds.DescribeSlowLogRecords API synchronously
// api document: https://help.aliyun.com/api/rds/describeslowlogrecords.html
func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (response *DescribeSlowLogRecordsResponse, err error) {
	response = CreateDescribeSlowLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlowLogRecordsWithChan invokes the rds.DescribeSlowLogRecords API asynchronously
// api document: https://help.aliyun.com/api/rds/describeslowlogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogRecordsWithChan(request *DescribeSlowLogRecordsRequest) (<-chan *DescribeSlowLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeSlowLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlowLogRecords(request)
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

// DescribeSlowLogRecordsWithCallback invokes the rds.DescribeSlowLogRecords API asynchronously
// api document: https://help.aliyun.com/api/rds/describeslowlogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogRecordsWithCallback(request *DescribeSlowLogRecordsRequest, callback func(response *DescribeSlowLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlowLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlowLogRecords(request)
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

// DescribeSlowLogRecordsRequest is the request struct for api DescribeSlowLogRecords
type DescribeSlowLogRecordsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	SQLHASH              string           `position:"Query" name:"SQLHASH"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	DBName               string           `position:"Query" name:"DBName"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DescribeSlowLogRecordsResponse is the response struct for api DescribeSlowLogRecords
type DescribeSlowLogRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string                        `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string                        `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine           string                        `json:"Engine" xml:"Engine"`
	TotalRecordCount int                           `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                           `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                           `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeSlowLogRecords `json:"Items" xml:"Items"`
}

// CreateDescribeSlowLogRecordsRequest creates a request to invoke DescribeSlowLogRecords API
func CreateDescribeSlowLogRecordsRequest() (request *DescribeSlowLogRecordsRequest) {
	request = &DescribeSlowLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSlowLogRecords", "rds", "openAPI")
	return
}

// CreateDescribeSlowLogRecordsResponse creates a response to parse from DescribeSlowLogRecords response
func CreateDescribeSlowLogRecordsResponse() (response *DescribeSlowLogRecordsResponse) {
	response = &DescribeSlowLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
