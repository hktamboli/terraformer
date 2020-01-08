// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterContainerInstanceInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instance to deregister. If you do not specify a cluster, the
	// default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full ARN of the container instance to deregister.
	// The ARN contains the arn:aws:ecs namespace, followed by the Region of the
	// container instance, the AWS account ID of the container instance owner, the
	// container-instance namespace, and then the container instance ID. For example,
	// arn:aws:ecs:region:aws_account_id:container-instance/container_instance_ID.
	//
	// ContainerInstance is a required field
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`

	// Forces the deregistration of the container instance. If you have tasks running
	// on the container instance when you deregister it with the force option, these
	// tasks remain running until you terminate the instance or the tasks stop through
	// some other means, but they are orphaned (no longer monitored or accounted
	// for by Amazon ECS). If an orphaned task on your container instance is part
	// of an Amazon ECS service, then the service scheduler starts another copy
	// of that task, on a different container instance if possible.
	//
	// Any containers in orphaned service tasks that are registered with a Classic
	// Load Balancer or an Application Load Balancer target group are deregistered.
	// They begin connection draining according to the settings on the load balancer
	// or target group.
	Force *bool `locationName:"force" type:"boolean"`
}

// String returns the string representation
func (s DeregisterContainerInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterContainerInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterContainerInstanceInput"}

	if s.ContainerInstance == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerInstance"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterContainerInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The container instance that was deregistered.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`
}

// String returns the string representation
func (s DeregisterContainerInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterContainerInstance = "DeregisterContainerInstance"

// DeregisterContainerInstanceRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Deregisters an Amazon ECS container instance from the specified cluster.
// This instance is no longer available to run tasks.
//
// If you intend to use the container instance for some other purpose after
// deregistration, you should stop all of the tasks running on the container
// instance before deregistration. That prevents any orphaned tasks from consuming
// resources.
//
// Deregistering a container instance removes the instance from a cluster, but
// it does not terminate the EC2 instance. If you are finished using the instance,
// be sure to terminate it in the Amazon EC2 console to stop billing.
//
// If you terminate a running container instance, Amazon ECS automatically deregisters
// the instance from your cluster (stopped container instances or instances
// with disconnected agents are not automatically deregistered when terminated).
//
//    // Example sending a request using DeregisterContainerInstanceRequest.
//    req := client.DeregisterContainerInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeregisterContainerInstance
func (c *Client) DeregisterContainerInstanceRequest(input *DeregisterContainerInstanceInput) DeregisterContainerInstanceRequest {
	op := &aws.Operation{
		Name:       opDeregisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterContainerInstanceInput{}
	}

	req := c.newRequest(op, input, &DeregisterContainerInstanceOutput{})
	return DeregisterContainerInstanceRequest{Request: req, Input: input, Copy: c.DeregisterContainerInstanceRequest}
}

// DeregisterContainerInstanceRequest is the request type for the
// DeregisterContainerInstance API operation.
type DeregisterContainerInstanceRequest struct {
	*aws.Request
	Input *DeregisterContainerInstanceInput
	Copy  func(*DeregisterContainerInstanceInput) DeregisterContainerInstanceRequest
}

// Send marshals and sends the DeregisterContainerInstance API request.
func (r DeregisterContainerInstanceRequest) Send(ctx context.Context) (*DeregisterContainerInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterContainerInstanceResponse{
		DeregisterContainerInstanceOutput: r.Request.Data.(*DeregisterContainerInstanceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterContainerInstanceResponse is the response type for the
// DeregisterContainerInstance API operation.
type DeregisterContainerInstanceResponse struct {
	*DeregisterContainerInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterContainerInstance request.
func (r *DeregisterContainerInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
