// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type DeleteBucketAnalyticsConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket from which an analytics configuration is deleted.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The ID that identifies the analytics configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketAnalyticsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketAnalyticsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketAnalyticsConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketAnalyticsConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketAnalyticsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *DeleteBucketAnalyticsConfigurationInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeleteBucketAnalyticsConfigurationInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeleteBucketAnalyticsConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketAnalyticsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketAnalyticsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketAnalyticsConfiguration = "DeleteBucketAnalyticsConfiguration"

// DeleteBucketAnalyticsConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes an analytics configuration for the bucket (specified by the analytics
// configuration ID).
//
// To use this operation, you must have permissions to perform the s3:PutAnalyticsConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev//using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// For information about the Amazon S3 analytics feature, see Amazon S3 Analytics
// – Storage Class Analysis (https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html).
//
// The following operations are related to DeleteBucketAnalyticsConfiguration:
//
//    *
//
//    *
//
//    *
//
//    // Example sending a request using DeleteBucketAnalyticsConfigurationRequest.
//    req := client.DeleteBucketAnalyticsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketAnalyticsConfiguration
func (c *Client) DeleteBucketAnalyticsConfigurationRequest(input *DeleteBucketAnalyticsConfigurationInput) DeleteBucketAnalyticsConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketAnalyticsConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?analytics",
	}

	if input == nil {
		input = &DeleteBucketAnalyticsConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketAnalyticsConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: c.DeleteBucketAnalyticsConfigurationRequest}
}

// DeleteBucketAnalyticsConfigurationRequest is the request type for the
// DeleteBucketAnalyticsConfiguration API operation.
type DeleteBucketAnalyticsConfigurationRequest struct {
	*aws.Request
	Input *DeleteBucketAnalyticsConfigurationInput
	Copy  func(*DeleteBucketAnalyticsConfigurationInput) DeleteBucketAnalyticsConfigurationRequest
}

// Send marshals and sends the DeleteBucketAnalyticsConfiguration API request.
func (r DeleteBucketAnalyticsConfigurationRequest) Send(ctx context.Context) (*DeleteBucketAnalyticsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketAnalyticsConfigurationResponse{
		DeleteBucketAnalyticsConfigurationOutput: r.Request.Data.(*DeleteBucketAnalyticsConfigurationOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketAnalyticsConfigurationResponse is the response type for the
// DeleteBucketAnalyticsConfiguration API operation.
type DeleteBucketAnalyticsConfigurationResponse struct {
	*DeleteBucketAnalyticsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketAnalyticsConfiguration request.
func (r *DeleteBucketAnalyticsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
