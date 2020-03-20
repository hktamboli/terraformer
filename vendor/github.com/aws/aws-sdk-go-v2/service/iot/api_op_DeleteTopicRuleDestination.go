// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteTopicRuleDestinationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the topic rule destination to delete.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTopicRuleDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTopicRuleDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTopicRuleDestinationInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTopicRuleDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteTopicRuleDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTopicRuleDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTopicRuleDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteTopicRuleDestination = "DeleteTopicRuleDestination"

// DeleteTopicRuleDestinationRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a topic rule destination.
//
//    // Example sending a request using DeleteTopicRuleDestinationRequest.
//    req := client.DeleteTopicRuleDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteTopicRuleDestinationRequest(input *DeleteTopicRuleDestinationInput) DeleteTopicRuleDestinationRequest {
	op := &aws.Operation{
		Name:       opDeleteTopicRuleDestination,
		HTTPMethod: "DELETE",
		HTTPPath:   "/destinations/{arn+}",
	}

	if input == nil {
		input = &DeleteTopicRuleDestinationInput{}
	}

	req := c.newRequest(op, input, &DeleteTopicRuleDestinationOutput{})
	return DeleteTopicRuleDestinationRequest{Request: req, Input: input, Copy: c.DeleteTopicRuleDestinationRequest}
}

// DeleteTopicRuleDestinationRequest is the request type for the
// DeleteTopicRuleDestination API operation.
type DeleteTopicRuleDestinationRequest struct {
	*aws.Request
	Input *DeleteTopicRuleDestinationInput
	Copy  func(*DeleteTopicRuleDestinationInput) DeleteTopicRuleDestinationRequest
}

// Send marshals and sends the DeleteTopicRuleDestination API request.
func (r DeleteTopicRuleDestinationRequest) Send(ctx context.Context) (*DeleteTopicRuleDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTopicRuleDestinationResponse{
		DeleteTopicRuleDestinationOutput: r.Request.Data.(*DeleteTopicRuleDestinationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTopicRuleDestinationResponse is the response type for the
// DeleteTopicRuleDestination API operation.
type DeleteTopicRuleDestinationResponse struct {
	*DeleteTopicRuleDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTopicRuleDestination request.
func (r *DeleteTopicRuleDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
