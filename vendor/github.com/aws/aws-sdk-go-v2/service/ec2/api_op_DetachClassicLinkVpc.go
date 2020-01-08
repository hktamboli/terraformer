// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetachClassicLinkVpcInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance to unlink from the VPC.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`

	// The ID of the VPC to which the instance is linked.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachClassicLinkVpcInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachClassicLinkVpcInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachClassicLinkVpcInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachClassicLinkVpcOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DetachClassicLinkVpcOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachClassicLinkVpc = "DetachClassicLinkVpc"

// DetachClassicLinkVpcRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Unlinks (detaches) a linked EC2-Classic instance from a VPC. After the instance
// has been unlinked, the VPC security groups are no longer associated with
// it. An instance is automatically unlinked from a VPC when it's stopped.
//
//    // Example sending a request using DetachClassicLinkVpcRequest.
//    req := client.DetachClassicLinkVpcRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachClassicLinkVpc
func (c *Client) DetachClassicLinkVpcRequest(input *DetachClassicLinkVpcInput) DetachClassicLinkVpcRequest {
	op := &aws.Operation{
		Name:       opDetachClassicLinkVpc,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachClassicLinkVpcInput{}
	}

	req := c.newRequest(op, input, &DetachClassicLinkVpcOutput{})
	return DetachClassicLinkVpcRequest{Request: req, Input: input, Copy: c.DetachClassicLinkVpcRequest}
}

// DetachClassicLinkVpcRequest is the request type for the
// DetachClassicLinkVpc API operation.
type DetachClassicLinkVpcRequest struct {
	*aws.Request
	Input *DetachClassicLinkVpcInput
	Copy  func(*DetachClassicLinkVpcInput) DetachClassicLinkVpcRequest
}

// Send marshals and sends the DetachClassicLinkVpc API request.
func (r DetachClassicLinkVpcRequest) Send(ctx context.Context) (*DetachClassicLinkVpcResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachClassicLinkVpcResponse{
		DetachClassicLinkVpcOutput: r.Request.Data.(*DetachClassicLinkVpcOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachClassicLinkVpcResponse is the response type for the
// DetachClassicLinkVpc API operation.
type DetachClassicLinkVpcResponse struct {
	*DetachClassicLinkVpcOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachClassicLinkVpc request.
func (r *DetachClassicLinkVpcResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
