package dds

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

// ModifyAuditPolicy invokes the dds.ModifyAuditPolicy API synchronously
// api document: https://help.aliyun.com/api/dds/modifyauditpolicy.html
func (client *Client) ModifyAuditPolicy(request *ModifyAuditPolicyRequest) (response *ModifyAuditPolicyResponse, err error) {
	response = CreateModifyAuditPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAuditPolicyWithChan invokes the dds.ModifyAuditPolicy API asynchronously
// api document: https://help.aliyun.com/api/dds/modifyauditpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAuditPolicyWithChan(request *ModifyAuditPolicyRequest) (<-chan *ModifyAuditPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyAuditPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAuditPolicy(request)
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

// ModifyAuditPolicyWithCallback invokes the dds.ModifyAuditPolicy API asynchronously
// api document: https://help.aliyun.com/api/dds/modifyauditpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAuditPolicyWithCallback(request *ModifyAuditPolicyRequest, callback func(response *ModifyAuditPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAuditPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyAuditPolicy(request)
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

// ModifyAuditPolicyRequest is the request struct for api ModifyAuditPolicy
type ModifyAuditPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AuditStatus          string           `position:"Query" name:"AuditStatus"`
	StoragePeriod        requests.Integer `position:"Query" name:"StoragePeriod"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyAuditPolicyResponse is the response struct for api ModifyAuditPolicy
type ModifyAuditPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAuditPolicyRequest creates a request to invoke ModifyAuditPolicy API
func CreateModifyAuditPolicyRequest() (request *ModifyAuditPolicyRequest) {
	request = &ModifyAuditPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyAuditPolicy", "dds", "openAPI")
	return
}

// CreateModifyAuditPolicyResponse creates a response to parse from ModifyAuditPolicy response
func CreateModifyAuditPolicyResponse() (response *ModifyAuditPolicyResponse) {
	response = &ModifyAuditPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}