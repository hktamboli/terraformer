// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateTypeInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The type definition, in GraphQL Schema Definition Language (SDL) format.
	//
	// For more information, see the GraphQL SDL documentation (http://graphql.org/learn/schema/).
	//
	// Definition is a required field
	Definition *string `locationName:"definition" type:"string" required:"true"`

	// The type format: SDL or JSON.
	//
	// Format is a required field
	Format TypeDefinitionFormat `locationName:"format" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTypeInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTypeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Definition != nil {
		v := *s.Definition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "definition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateTypeOutput struct {
	_ struct{} `type:"structure"`

	// The Type object.
	Type *Type `locationName:"type" type:"structure"`
}

// String returns the string representation
func (s CreateTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Type != nil {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "type", v, metadata)
	}
	return nil
}

const opCreateType = "CreateType"

// CreateTypeRequest returns a request value for making API operation for
// AWS AppSync.
//
// Creates a Type object.
//
//    // Example sending a request using CreateTypeRequest.
//    req := client.CreateTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateType
func (c *Client) CreateTypeRequest(input *CreateTypeInput) CreateTypeRequest {
	op := &aws.Operation{
		Name:       opCreateType,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/types",
	}

	if input == nil {
		input = &CreateTypeInput{}
	}

	req := c.newRequest(op, input, &CreateTypeOutput{})
	return CreateTypeRequest{Request: req, Input: input, Copy: c.CreateTypeRequest}
}

// CreateTypeRequest is the request type for the
// CreateType API operation.
type CreateTypeRequest struct {
	*aws.Request
	Input *CreateTypeInput
	Copy  func(*CreateTypeInput) CreateTypeRequest
}

// Send marshals and sends the CreateType API request.
func (r CreateTypeRequest) Send(ctx context.Context) (*CreateTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTypeResponse{
		CreateTypeOutput: r.Request.Data.(*CreateTypeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTypeResponse is the response type for the
// CreateType API operation.
type CreateTypeResponse struct {
	*CreateTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateType request.
func (r *CreateTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
