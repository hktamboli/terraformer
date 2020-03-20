// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to change information about the current Account resource.
type UpdateAccountInput struct {
	_ struct{} `type:"structure"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`
}

// String returns the string representation
func (s UpdateAccountInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PatchOperations != nil {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Represents an AWS account that is associated with API Gateway.
//
// To view the account info, call GET on this resource.
//
// Error Codes
//
// The following exception may be thrown when the request fails.
//
//    * UnauthorizedException
//
//    * NotFoundException
//
//    * TooManyRequestsException
//
// For detailed error code information, including the corresponding HTTP Status
// Codes, see API Gateway Error Codes (https://docs.aws.amazon.com/apigateway/api-reference/handling-errors/#api-error-codes)
//
// Example: Get the information about an account.
//
// Request
//  GET /account HTTP/1.1 Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
//  X-Amz-Date: 20160531T184618Z Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/us-east-1/apigateway/aws4_request,
//  SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
// Response
//
// The successful response returns a 200 OK status code and a payload similar
// to the following:
//  { "_links": { "curies": { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/account-apigateway-{rel}.html",
//  "name": "account", "templated": true }, "self": { "href": "/account" },
//  "account:update": { "href": "/account" } }, "cloudwatchRoleArn": "arn:aws:iam::123456789012:role/apigAwsProxyRole",
//  "throttleSettings": { "rateLimit": 500, "burstLimit": 1000 } }
// In addition to making the REST API call directly, you can use the AWS CLI
// and an AWS SDK to access this resource.
//
// API Gateway Limits (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-limits.html)
// Developer Guide (https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html),
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-account.html)
type UpdateAccountOutput struct {
	_ struct{} `type:"structure"`

	// The version of the API keys used for the account.
	ApiKeyVersion *string `locationName:"apiKeyVersion" type:"string"`

	// The ARN of an Amazon CloudWatch role for the current Account.
	CloudwatchRoleArn *string `locationName:"cloudwatchRoleArn" type:"string"`

	// A list of features supported for the account. When usage plans are enabled,
	// the features list will include an entry of "UsagePlans".
	Features []string `locationName:"features" type:"list"`

	// Specifies the API request limits configured for the current Account.
	ThrottleSettings *ThrottleSettings `locationName:"throttleSettings" type:"structure"`
}

// String returns the string representation
func (s UpdateAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiKeyVersion != nil {
		v := *s.ApiKeyVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeyVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CloudwatchRoleArn != nil {
		v := *s.CloudwatchRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cloudwatchRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Features != nil {
		v := s.Features

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "features", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ThrottleSettings != nil {
		v := s.ThrottleSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "throttleSettings", v, metadata)
	}
	return nil
}

const opUpdateAccount = "UpdateAccount"

// UpdateAccountRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about the current Account resource.
//
//    // Example sending a request using UpdateAccountRequest.
//    req := client.UpdateAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateAccountRequest(input *UpdateAccountInput) UpdateAccountRequest {
	op := &aws.Operation{
		Name:       opUpdateAccount,
		HTTPMethod: "PATCH",
		HTTPPath:   "/account",
	}

	if input == nil {
		input = &UpdateAccountInput{}
	}

	req := c.newRequest(op, input, &UpdateAccountOutput{})
	return UpdateAccountRequest{Request: req, Input: input, Copy: c.UpdateAccountRequest}
}

// UpdateAccountRequest is the request type for the
// UpdateAccount API operation.
type UpdateAccountRequest struct {
	*aws.Request
	Input *UpdateAccountInput
	Copy  func(*UpdateAccountInput) UpdateAccountRequest
}

// Send marshals and sends the UpdateAccount API request.
func (r UpdateAccountRequest) Send(ctx context.Context) (*UpdateAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccountResponse{
		UpdateAccountOutput: r.Request.Data.(*UpdateAccountOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccountResponse is the response type for the
// UpdateAccount API operation.
type UpdateAccountResponse struct {
	*UpdateAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccount request.
func (r *UpdateAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
