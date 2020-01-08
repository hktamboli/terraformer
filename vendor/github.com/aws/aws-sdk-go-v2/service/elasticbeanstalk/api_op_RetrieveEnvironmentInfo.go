// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to download logs retrieved with RequestEnvironmentInfo.
type RetrieveEnvironmentInfoInput struct {
	_ struct{} `type:"structure"`

	// The ID of the data's environment.
	//
	// If no such environment is found, returns an InvalidParameterValue error.
	//
	// Condition: You must specify either this or an EnvironmentName, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentId *string `type:"string"`

	// The name of the data's environment.
	//
	// If no such environment is found, returns an InvalidParameterValue error.
	//
	// Condition: You must specify either this or an EnvironmentId, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentName *string `min:"4" type:"string"`

	// The type of information to retrieve.
	//
	// InfoType is a required field
	InfoType EnvironmentInfoType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s RetrieveEnvironmentInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RetrieveEnvironmentInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RetrieveEnvironmentInfoInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}
	if len(s.InfoType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("InfoType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message containing a description of the requested environment info.
type RetrieveEnvironmentInfoOutput struct {
	_ struct{} `type:"structure"`

	// The EnvironmentInfoDescription of the environment.
	EnvironmentInfo []EnvironmentInfoDescription `type:"list"`
}

// String returns the string representation
func (s RetrieveEnvironmentInfoOutput) String() string {
	return awsutil.Prettify(s)
}

const opRetrieveEnvironmentInfo = "RetrieveEnvironmentInfo"

// RetrieveEnvironmentInfoRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Retrieves the compiled information from a RequestEnvironmentInfo request.
//
// Related Topics
//
//    * RequestEnvironmentInfo
//
//    // Example sending a request using RetrieveEnvironmentInfoRequest.
//    req := client.RetrieveEnvironmentInfoRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/RetrieveEnvironmentInfo
func (c *Client) RetrieveEnvironmentInfoRequest(input *RetrieveEnvironmentInfoInput) RetrieveEnvironmentInfoRequest {
	op := &aws.Operation{
		Name:       opRetrieveEnvironmentInfo,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RetrieveEnvironmentInfoInput{}
	}

	req := c.newRequest(op, input, &RetrieveEnvironmentInfoOutput{})
	return RetrieveEnvironmentInfoRequest{Request: req, Input: input, Copy: c.RetrieveEnvironmentInfoRequest}
}

// RetrieveEnvironmentInfoRequest is the request type for the
// RetrieveEnvironmentInfo API operation.
type RetrieveEnvironmentInfoRequest struct {
	*aws.Request
	Input *RetrieveEnvironmentInfoInput
	Copy  func(*RetrieveEnvironmentInfoInput) RetrieveEnvironmentInfoRequest
}

// Send marshals and sends the RetrieveEnvironmentInfo API request.
func (r RetrieveEnvironmentInfoRequest) Send(ctx context.Context) (*RetrieveEnvironmentInfoResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RetrieveEnvironmentInfoResponse{
		RetrieveEnvironmentInfoOutput: r.Request.Data.(*RetrieveEnvironmentInfoOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RetrieveEnvironmentInfoResponse is the response type for the
// RetrieveEnvironmentInfo API operation.
type RetrieveEnvironmentInfoResponse struct {
	*RetrieveEnvironmentInfoOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RetrieveEnvironmentInfo request.
func (r *RetrieveEnvironmentInfoResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
