package cms

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

// DescribeMonitorGroupInstanceAttribute invokes the cms.DescribeMonitorGroupInstanceAttribute API synchronously
func (client *Client) DescribeMonitorGroupInstanceAttribute(request *DescribeMonitorGroupInstanceAttributeRequest) (response *DescribeMonitorGroupInstanceAttributeResponse, err error) {
	response = CreateDescribeMonitorGroupInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorGroupInstanceAttributeWithChan invokes the cms.DescribeMonitorGroupInstanceAttribute API asynchronously
func (client *Client) DescribeMonitorGroupInstanceAttributeWithChan(request *DescribeMonitorGroupInstanceAttributeRequest) (<-chan *DescribeMonitorGroupInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorGroupInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorGroupInstanceAttribute(request)
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

// DescribeMonitorGroupInstanceAttributeWithCallback invokes the cms.DescribeMonitorGroupInstanceAttribute API asynchronously
func (client *Client) DescribeMonitorGroupInstanceAttributeWithCallback(request *DescribeMonitorGroupInstanceAttributeRequest, callback func(response *DescribeMonitorGroupInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorGroupInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorGroupInstanceAttribute(request)
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

// DescribeMonitorGroupInstanceAttributeRequest is the request struct for api DescribeMonitorGroupInstanceAttribute
type DescribeMonitorGroupInstanceAttributeRequest struct {
	*requests.RpcRequest
	GroupId     requests.Integer `position:"Query" name:"GroupId"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	Total       requests.Boolean `position:"Query" name:"Total"`
	InstanceIds string           `position:"Query" name:"InstanceIds"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Category    string           `position:"Query" name:"Category"`
	Keyword     string           `position:"Query" name:"Keyword"`
}

// DescribeMonitorGroupInstanceAttributeResponse is the response struct for api DescribeMonitorGroupInstanceAttribute
type DescribeMonitorGroupInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId  string                                           `json:"RequestId" xml:"RequestId"`
	Success    bool                                             `json:"Success" xml:"Success"`
	Code       int                                              `json:"Code" xml:"Code"`
	Message    string                                           `json:"Message" xml:"Message"`
	PageNumber int                                              `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                                              `json:"PageSize" xml:"PageSize"`
	Total      int                                              `json:"Total" xml:"Total"`
	Resources  ResourcesInDescribeMonitorGroupInstanceAttribute `json:"Resources" xml:"Resources"`
}

// CreateDescribeMonitorGroupInstanceAttributeRequest creates a request to invoke DescribeMonitorGroupInstanceAttribute API
func CreateDescribeMonitorGroupInstanceAttributeRequest() (request *DescribeMonitorGroupInstanceAttributeRequest) {
	request = &DescribeMonitorGroupInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitorGroupInstanceAttribute", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMonitorGroupInstanceAttributeResponse creates a response to parse from DescribeMonitorGroupInstanceAttribute response
func CreateDescribeMonitorGroupInstanceAttributeResponse() (response *DescribeMonitorGroupInstanceAttributeResponse) {
	response = &DescribeMonitorGroupInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
