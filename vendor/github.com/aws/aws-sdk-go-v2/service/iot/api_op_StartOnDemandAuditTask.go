// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StartOnDemandAuditTaskInput struct {
	_ struct{} `type:"structure"`

	// Which checks are performed during the audit. The checks you specify must
	// be enabled for your account or an exception occurs. Use DescribeAccountAuditConfiguration
	// to see the list of all checks, including those that are enabled or UpdateAccountAuditConfiguration
	// to select which checks are enabled.
	//
	// TargetCheckNames is a required field
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list" required:"true"`
}

// String returns the string representation
func (s StartOnDemandAuditTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartOnDemandAuditTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartOnDemandAuditTaskInput"}

	if s.TargetCheckNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetCheckNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartOnDemandAuditTaskInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TargetCheckNames != nil {
		v := s.TargetCheckNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetCheckNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type StartOnDemandAuditTaskOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the on-demand audit you started.
	TaskId *string `locationName:"taskId" min:"1" type:"string"`
}

// String returns the string representation
func (s StartOnDemandAuditTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartOnDemandAuditTaskOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "taskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartOnDemandAuditTask = "StartOnDemandAuditTask"

// StartOnDemandAuditTaskRequest returns a request value for making API operation for
// AWS IoT.
//
// Starts an on-demand Device Defender audit.
//
//    // Example sending a request using StartOnDemandAuditTaskRequest.
//    req := client.StartOnDemandAuditTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartOnDemandAuditTaskRequest(input *StartOnDemandAuditTaskInput) StartOnDemandAuditTaskRequest {
	op := &aws.Operation{
		Name:       opStartOnDemandAuditTask,
		HTTPMethod: "POST",
		HTTPPath:   "/audit/tasks",
	}

	if input == nil {
		input = &StartOnDemandAuditTaskInput{}
	}

	req := c.newRequest(op, input, &StartOnDemandAuditTaskOutput{})
	return StartOnDemandAuditTaskRequest{Request: req, Input: input, Copy: c.StartOnDemandAuditTaskRequest}
}

// StartOnDemandAuditTaskRequest is the request type for the
// StartOnDemandAuditTask API operation.
type StartOnDemandAuditTaskRequest struct {
	*aws.Request
	Input *StartOnDemandAuditTaskInput
	Copy  func(*StartOnDemandAuditTaskInput) StartOnDemandAuditTaskRequest
}

// Send marshals and sends the StartOnDemandAuditTask API request.
func (r StartOnDemandAuditTaskRequest) Send(ctx context.Context) (*StartOnDemandAuditTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartOnDemandAuditTaskResponse{
		StartOnDemandAuditTaskOutput: r.Request.Data.(*StartOnDemandAuditTaskOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartOnDemandAuditTaskResponse is the response type for the
// StartOnDemandAuditTask API operation.
type StartOnDemandAuditTaskResponse struct {
	*StartOnDemandAuditTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartOnDemandAuditTask request.
func (r *StartOnDemandAuditTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
