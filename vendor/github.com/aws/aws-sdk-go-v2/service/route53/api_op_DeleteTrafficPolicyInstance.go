// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to delete a specified traffic policy instance.
type DeleteTrafficPolicyInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the traffic policy instance that you want to delete.
	//
	// When you delete a traffic policy instance, Amazon Route 53 also deletes all
	// of the resource record sets that were created when you created the traffic
	// policy instance.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrafficPolicyInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrafficPolicyInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrafficPolicyInstanceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTrafficPolicyInstanceInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// An empty element.
type DeleteTrafficPolicyInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTrafficPolicyInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTrafficPolicyInstanceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteTrafficPolicyInstance = "DeleteTrafficPolicyInstance"

// DeleteTrafficPolicyInstanceRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Deletes a traffic policy instance and all of the resource record sets that
// Amazon Route 53 created when you created the instance.
//
// In the Route 53 console, traffic policy instances are known as policy records.
//
//    // Example sending a request using DeleteTrafficPolicyInstanceRequest.
//    req := client.DeleteTrafficPolicyInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteTrafficPolicyInstance
func (c *Client) DeleteTrafficPolicyInstanceRequest(input *DeleteTrafficPolicyInstanceInput) DeleteTrafficPolicyInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficPolicyInstance,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2013-04-01/trafficpolicyinstance/{Id}",
	}

	if input == nil {
		input = &DeleteTrafficPolicyInstanceInput{}
	}

	req := c.newRequest(op, input, &DeleteTrafficPolicyInstanceOutput{})
	return DeleteTrafficPolicyInstanceRequest{Request: req, Input: input, Copy: c.DeleteTrafficPolicyInstanceRequest}
}

// DeleteTrafficPolicyInstanceRequest is the request type for the
// DeleteTrafficPolicyInstance API operation.
type DeleteTrafficPolicyInstanceRequest struct {
	*aws.Request
	Input *DeleteTrafficPolicyInstanceInput
	Copy  func(*DeleteTrafficPolicyInstanceInput) DeleteTrafficPolicyInstanceRequest
}

// Send marshals and sends the DeleteTrafficPolicyInstance API request.
func (r DeleteTrafficPolicyInstanceRequest) Send(ctx context.Context) (*DeleteTrafficPolicyInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficPolicyInstanceResponse{
		DeleteTrafficPolicyInstanceOutput: r.Request.Data.(*DeleteTrafficPolicyInstanceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficPolicyInstanceResponse is the response type for the
// DeleteTrafficPolicyInstance API operation.
type DeleteTrafficPolicyInstanceResponse struct {
	*DeleteTrafficPolicyInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficPolicyInstance request.
func (r *DeleteTrafficPolicyInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
