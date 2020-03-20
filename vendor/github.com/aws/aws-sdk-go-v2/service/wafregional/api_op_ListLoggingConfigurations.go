// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListLoggingConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of LoggingConfigurations that you want AWS WAF to return
	// for this request. If you have more LoggingConfigurations than the number
	// that you specify for Limit, the response includes a NextMarker value that
	// you can use to get another batch of LoggingConfigurations.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more LoggingConfigurations
	// than the value of Limit, AWS WAF returns a NextMarker value in the response
	// that allows you to list another group of LoggingConfigurations. For the second
	// and subsequent ListLoggingConfigurations requests, specify the value of NextMarker
	// from the previous response to get information about another batch of ListLoggingConfigurations.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListLoggingConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListLoggingConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListLoggingConfigurationsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListLoggingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// An array of LoggingConfiguration objects.
	LoggingConfigurations []LoggingConfiguration `type:"list"`

	// If you have more LoggingConfigurations than the number that you specified
	// for Limit in the request, the response includes a NextMarker value. To list
	// more LoggingConfigurations, submit another ListLoggingConfigurations request,
	// and specify the NextMarker value from the response in the NextMarker value
	// in the next request.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListLoggingConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListLoggingConfigurations = "ListLoggingConfigurations"

// ListLoggingConfigurationsRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns an array of LoggingConfiguration objects.
//
//    // Example sending a request using ListLoggingConfigurationsRequest.
//    req := client.ListLoggingConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListLoggingConfigurations
func (c *Client) ListLoggingConfigurationsRequest(input *ListLoggingConfigurationsInput) ListLoggingConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListLoggingConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListLoggingConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListLoggingConfigurationsOutput{})
	return ListLoggingConfigurationsRequest{Request: req, Input: input, Copy: c.ListLoggingConfigurationsRequest}
}

// ListLoggingConfigurationsRequest is the request type for the
// ListLoggingConfigurations API operation.
type ListLoggingConfigurationsRequest struct {
	*aws.Request
	Input *ListLoggingConfigurationsInput
	Copy  func(*ListLoggingConfigurationsInput) ListLoggingConfigurationsRequest
}

// Send marshals and sends the ListLoggingConfigurations API request.
func (r ListLoggingConfigurationsRequest) Send(ctx context.Context) (*ListLoggingConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLoggingConfigurationsResponse{
		ListLoggingConfigurationsOutput: r.Request.Data.(*ListLoggingConfigurationsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListLoggingConfigurationsResponse is the response type for the
// ListLoggingConfigurations API operation.
type ListLoggingConfigurationsResponse struct {
	*ListLoggingConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLoggingConfigurations request.
func (r *ListLoggingConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
