// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutImageScanningConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The image scanning configuration for the repository. This setting determines
	// whether images are scanned for known vulnerabilities after being pushed to
	// the repository.
	//
	// ImageScanningConfiguration is a required field
	ImageScanningConfiguration *ImageScanningConfiguration `locationName:"imageScanningConfiguration" type:"structure" required:"true"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to update the image scanning configuration setting. If you do not
	// specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository in which to update the image scanning configuration
	// setting.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutImageScanningConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutImageScanningConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutImageScanningConfigurationInput"}

	if s.ImageScanningConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageScanningConfiguration"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutImageScanningConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The image scanning configuration setting for the repository.
	ImageScanningConfiguration *ImageScanningConfiguration `locationName:"imageScanningConfiguration" type:"structure"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s PutImageScanningConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutImageScanningConfiguration = "PutImageScanningConfiguration"

// PutImageScanningConfigurationRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Updates the image scanning configuration for the specified repository.
//
//    // Example sending a request using PutImageScanningConfigurationRequest.
//    req := client.PutImageScanningConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImageScanningConfiguration
func (c *Client) PutImageScanningConfigurationRequest(input *PutImageScanningConfigurationInput) PutImageScanningConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutImageScanningConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutImageScanningConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutImageScanningConfigurationOutput{})
	return PutImageScanningConfigurationRequest{Request: req, Input: input, Copy: c.PutImageScanningConfigurationRequest}
}

// PutImageScanningConfigurationRequest is the request type for the
// PutImageScanningConfiguration API operation.
type PutImageScanningConfigurationRequest struct {
	*aws.Request
	Input *PutImageScanningConfigurationInput
	Copy  func(*PutImageScanningConfigurationInput) PutImageScanningConfigurationRequest
}

// Send marshals and sends the PutImageScanningConfiguration API request.
func (r PutImageScanningConfigurationRequest) Send(ctx context.Context) (*PutImageScanningConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutImageScanningConfigurationResponse{
		PutImageScanningConfigurationOutput: r.Request.Data.(*PutImageScanningConfigurationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutImageScanningConfigurationResponse is the response type for the
// PutImageScanningConfiguration API operation.
type PutImageScanningConfigurationResponse struct {
	*PutImageScanningConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutImageScanningConfiguration request.
func (r *PutImageScanningConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
