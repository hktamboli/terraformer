// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the get device request.
type GetDeviceInput struct {
	_ struct{} `type:"structure"`

	// The device type's ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeviceInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a get device request.
type GetDeviceOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about the requested device.
	Device *Device `locationName:"device" type:"structure"`
}

// String returns the string representation
func (s GetDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDevice = "GetDevice"

// GetDeviceRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about a unique device type.
//
//    // Example sending a request using GetDeviceRequest.
//    req := client.GetDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetDevice
func (c *Client) GetDeviceRequest(input *GetDeviceInput) GetDeviceRequest {
	op := &aws.Operation{
		Name:       opGetDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeviceInput{}
	}

	req := c.newRequest(op, input, &GetDeviceOutput{})
	return GetDeviceRequest{Request: req, Input: input, Copy: c.GetDeviceRequest}
}

// GetDeviceRequest is the request type for the
// GetDevice API operation.
type GetDeviceRequest struct {
	*aws.Request
	Input *GetDeviceInput
	Copy  func(*GetDeviceInput) GetDeviceRequest
}

// Send marshals and sends the GetDevice API request.
func (r GetDeviceRequest) Send(ctx context.Context) (*GetDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceResponse{
		GetDeviceOutput: r.Request.Data.(*GetDeviceOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceResponse is the response type for the
// GetDevice API operation.
type GetDeviceResponse struct {
	*GetDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevice request.
func (r *GetDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
