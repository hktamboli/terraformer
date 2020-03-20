// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutImageInput struct {
	_ struct{} `type:"structure"`

	// The image manifest corresponding to the image to be uploaded.
	//
	// ImageManifest is a required field
	ImageManifest *string `locationName:"imageManifest" min:"1" type:"string" required:"true"`

	// The tag to associate with the image. This parameter is required for images
	// that use the Docker Image Manifest V2 Schema 2 or OCI formats.
	ImageTag *string `locationName:"imageTag" min:"1" type:"string"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to put the image. If you do not specify a registry, the default
	// registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository in which to put the image.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutImageInput"}

	if s.ImageManifest == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageManifest"))
	}
	if s.ImageManifest != nil && len(*s.ImageManifest) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageManifest", 1))
	}
	if s.ImageTag != nil && len(*s.ImageTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageTag", 1))
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

type PutImageOutput struct {
	_ struct{} `type:"structure"`

	// Details of the image uploaded.
	Image *Image `locationName:"image" type:"structure"`
}

// String returns the string representation
func (s PutImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutImage = "PutImage"

// PutImageRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Creates or updates the image manifest and tags associated with an image.
//
// This operation is used by the Amazon ECR proxy, and it is not intended for
// general use by customers for pulling and pushing images. In most cases, you
// should use the docker CLI to pull, tag, and push images.
//
//    // Example sending a request using PutImageRequest.
//    req := client.PutImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImage
func (c *Client) PutImageRequest(input *PutImageInput) PutImageRequest {
	op := &aws.Operation{
		Name:       opPutImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutImageInput{}
	}

	req := c.newRequest(op, input, &PutImageOutput{})
	return PutImageRequest{Request: req, Input: input, Copy: c.PutImageRequest}
}

// PutImageRequest is the request type for the
// PutImage API operation.
type PutImageRequest struct {
	*aws.Request
	Input *PutImageInput
	Copy  func(*PutImageInput) PutImageRequest
}

// Send marshals and sends the PutImage API request.
func (r PutImageRequest) Send(ctx context.Context) (*PutImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutImageResponse{
		PutImageOutput: r.Request.Data.(*PutImageOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutImageResponse is the response type for the
// PutImage API operation.
type PutImageResponse struct {
	*PutImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutImage request.
func (r *PutImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
