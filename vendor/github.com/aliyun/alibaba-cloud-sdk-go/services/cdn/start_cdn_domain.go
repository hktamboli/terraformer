package cdn

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

// StartCdnDomain invokes the cdn.StartCdnDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/startcdndomain.html
func (client *Client) StartCdnDomain(request *StartCdnDomainRequest) (response *StartCdnDomainResponse, err error) {
	response = CreateStartCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StartCdnDomainWithChan invokes the cdn.StartCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/startcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartCdnDomainWithChan(request *StartCdnDomainRequest) (<-chan *StartCdnDomainResponse, <-chan error) {
	responseChan := make(chan *StartCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartCdnDomain(request)
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

// StartCdnDomainWithCallback invokes the cdn.StartCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/startcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartCdnDomainWithCallback(request *StartCdnDomainRequest, callback func(response *StartCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.StartCdnDomain(request)
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

// StartCdnDomainRequest is the request struct for api StartCdnDomain
type StartCdnDomainRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// StartCdnDomainResponse is the response struct for api StartCdnDomain
type StartCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartCdnDomainRequest creates a request to invoke StartCdnDomain API
func CreateStartCdnDomainRequest() (request *StartCdnDomainRequest) {
	request = &StartCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "StartCdnDomain", "cdn", "openAPI")
	return
}

// CreateStartCdnDomainResponse creates a response to parse from StartCdnDomain response
func CreateStartCdnDomainResponse() (response *StartCdnDomainResponse) {
	response = &StartCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
