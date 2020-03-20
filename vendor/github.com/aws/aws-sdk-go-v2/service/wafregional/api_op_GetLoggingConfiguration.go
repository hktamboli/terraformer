// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLoggingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the web ACL for which you want to get the
	// LoggingConfiguration.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLoggingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLoggingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLoggingConfigurationInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLoggingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The LoggingConfiguration for the specified web ACL.
	LoggingConfiguration *LoggingConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetLoggingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLoggingConfiguration = "GetLoggingConfiguration"

// GetLoggingConfigurationRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the LoggingConfiguration for the specified web ACL.
//
//    // Example sending a request using GetLoggingConfigurationRequest.
//    req := client.GetLoggingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetLoggingConfiguration
func (c *Client) GetLoggingConfigurationRequest(input *GetLoggingConfigurationInput) GetLoggingConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetLoggingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLoggingConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetLoggingConfigurationOutput{})
	return GetLoggingConfigurationRequest{Request: req, Input: input, Copy: c.GetLoggingConfigurationRequest}
}

// GetLoggingConfigurationRequest is the request type for the
// GetLoggingConfiguration API operation.
type GetLoggingConfigurationRequest struct {
	*aws.Request
	Input *GetLoggingConfigurationInput
	Copy  func(*GetLoggingConfigurationInput) GetLoggingConfigurationRequest
}

// Send marshals and sends the GetLoggingConfiguration API request.
func (r GetLoggingConfigurationRequest) Send(ctx context.Context) (*GetLoggingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoggingConfigurationResponse{
		GetLoggingConfigurationOutput: r.Request.Data.(*GetLoggingConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoggingConfigurationResponse is the response type for the
// GetLoggingConfiguration API operation.
type GetLoggingConfigurationResponse struct {
	*GetLoggingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoggingConfiguration request.
func (r *GetLoggingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
