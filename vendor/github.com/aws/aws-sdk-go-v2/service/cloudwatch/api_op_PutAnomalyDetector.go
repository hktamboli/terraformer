// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutAnomalyDetectorInput struct {
	_ struct{} `type:"structure"`

	// The configuration specifies details about how the anomaly detection model
	// is to be trained, including time ranges to exclude when training and updating
	// the model. You can specify as many as 10 time ranges.
	//
	// The configuration can also include the time zone to use for the metric.
	//
	// You can in
	Configuration *AnomalyDetectorConfiguration `type:"structure"`

	// The metric dimensions to create the anomaly detection model for.
	Dimensions []Dimension `type:"list"`

	// The name of the metric to create the anomaly detection model for.
	//
	// MetricName is a required field
	MetricName *string `min:"1" type:"string" required:"true"`

	// The namespace of the metric to create the anomaly detection model for.
	//
	// Namespace is a required field
	Namespace *string `min:"1" type:"string" required:"true"`

	// The statistic to use for the metric and the anomaly detection model.
	//
	// Stat is a required field
	Stat *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutAnomalyDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutAnomalyDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutAnomalyDetectorInput"}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}

	if s.Stat == nil {
		invalidParams.Add(aws.NewErrParamRequired("Stat"))
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			invalidParams.AddNested("Configuration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutAnomalyDetectorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAnomalyDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutAnomalyDetector = "PutAnomalyDetector"

// PutAnomalyDetectorRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Creates an anomaly detection model for a CloudWatch metric. You can use the
// model to display a band of expected normal values when the metric is graphed.
//
// For more information, see CloudWatch Anomaly Detection (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html).
//
//    // Example sending a request using PutAnomalyDetectorRequest.
//    req := client.PutAnomalyDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutAnomalyDetector
func (c *Client) PutAnomalyDetectorRequest(input *PutAnomalyDetectorInput) PutAnomalyDetectorRequest {
	op := &aws.Operation{
		Name:       opPutAnomalyDetector,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAnomalyDetectorInput{}
	}

	req := c.newRequest(op, input, &PutAnomalyDetectorOutput{})
	return PutAnomalyDetectorRequest{Request: req, Input: input, Copy: c.PutAnomalyDetectorRequest}
}

// PutAnomalyDetectorRequest is the request type for the
// PutAnomalyDetector API operation.
type PutAnomalyDetectorRequest struct {
	*aws.Request
	Input *PutAnomalyDetectorInput
	Copy  func(*PutAnomalyDetectorInput) PutAnomalyDetectorRequest
}

// Send marshals and sends the PutAnomalyDetector API request.
func (r PutAnomalyDetectorRequest) Send(ctx context.Context) (*PutAnomalyDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAnomalyDetectorResponse{
		PutAnomalyDetectorOutput: r.Request.Data.(*PutAnomalyDetectorOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAnomalyDetectorResponse is the response type for the
// PutAnomalyDetector API operation.
type PutAnomalyDetectorResponse struct {
	*PutAnomalyDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAnomalyDetector request.
func (r *PutAnomalyDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
