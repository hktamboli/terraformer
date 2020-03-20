// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListTopicRuleDestinationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListTopicRuleDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTopicRuleDestinationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTopicRuleDestinationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTopicRuleDestinationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListTopicRuleDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about a topic rule destination.
	DestinationSummaries []TopicRuleDestinationSummary `locationName:"destinationSummaries" type:"list"`

	// The token to retrieve the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListTopicRuleDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTopicRuleDestinationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DestinationSummaries != nil {
		v := s.DestinationSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "destinationSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListTopicRuleDestinations = "ListTopicRuleDestinations"

// ListTopicRuleDestinationsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists all the topic rule destinations in your AWS account.
//
//    // Example sending a request using ListTopicRuleDestinationsRequest.
//    req := client.ListTopicRuleDestinationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListTopicRuleDestinationsRequest(input *ListTopicRuleDestinationsInput) ListTopicRuleDestinationsRequest {
	op := &aws.Operation{
		Name:       opListTopicRuleDestinations,
		HTTPMethod: "GET",
		HTTPPath:   "/destinations",
	}

	if input == nil {
		input = &ListTopicRuleDestinationsInput{}
	}

	req := c.newRequest(op, input, &ListTopicRuleDestinationsOutput{})
	return ListTopicRuleDestinationsRequest{Request: req, Input: input, Copy: c.ListTopicRuleDestinationsRequest}
}

// ListTopicRuleDestinationsRequest is the request type for the
// ListTopicRuleDestinations API operation.
type ListTopicRuleDestinationsRequest struct {
	*aws.Request
	Input *ListTopicRuleDestinationsInput
	Copy  func(*ListTopicRuleDestinationsInput) ListTopicRuleDestinationsRequest
}

// Send marshals and sends the ListTopicRuleDestinations API request.
func (r ListTopicRuleDestinationsRequest) Send(ctx context.Context) (*ListTopicRuleDestinationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTopicRuleDestinationsResponse{
		ListTopicRuleDestinationsOutput: r.Request.Data.(*ListTopicRuleDestinationsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTopicRuleDestinationsResponse is the response type for the
// ListTopicRuleDestinations API operation.
type ListTopicRuleDestinationsResponse struct {
	*ListTopicRuleDestinationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTopicRuleDestinations request.
func (r *ListTopicRuleDestinationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
