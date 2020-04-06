// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteResourceServerInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the resource server.
	//
	// Identifier is a required field
	Identifier *string `min:"1" type:"string" required:"true"`

	// The user pool ID for the user pool that hosts the resource server.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourceServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourceServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourceServerInput"}

	if s.Identifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identifier"))
	}
	if s.Identifier != nil && len(*s.Identifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Identifier", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteResourceServerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourceServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteResourceServer = "DeleteResourceServer"

// DeleteResourceServerRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Deletes a resource server.
//
//    // Example sending a request using DeleteResourceServerRequest.
//    req := client.DeleteResourceServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteResourceServer
func (c *Client) DeleteResourceServerRequest(input *DeleteResourceServerInput) DeleteResourceServerRequest {
	op := &aws.Operation{
		Name:       opDeleteResourceServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResourceServerInput{}
	}

	req := c.newRequest(op, input, &DeleteResourceServerOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteResourceServerRequest{Request: req, Input: input, Copy: c.DeleteResourceServerRequest}
}

// DeleteResourceServerRequest is the request type for the
// DeleteResourceServer API operation.
type DeleteResourceServerRequest struct {
	*aws.Request
	Input *DeleteResourceServerInput
	Copy  func(*DeleteResourceServerInput) DeleteResourceServerRequest
}

// Send marshals and sends the DeleteResourceServer API request.
func (r DeleteResourceServerRequest) Send(ctx context.Context) (*DeleteResourceServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourceServerResponse{
		DeleteResourceServerOutput: r.Request.Data.(*DeleteResourceServerOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourceServerResponse is the response type for the
// DeleteResourceServer API operation.
type DeleteResourceServerResponse struct {
	*DeleteResourceServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourceServer request.
func (r *DeleteResourceServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
