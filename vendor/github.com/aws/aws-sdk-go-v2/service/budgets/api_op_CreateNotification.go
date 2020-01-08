// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of CreateNotification
type CreateNotificationInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget that you want to create
	// a notification for.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The name of the budget that you want AWS to notify you about. Budget names
	// must be unique within an account.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`

	// The notification that you want to create.
	//
	// Notification is a required field
	Notification *Notification `type:"structure" required:"true"`

	// A list of subscribers that you want to associate with the notification. Each
	// notification can have one SNS subscriber and up to 10 email subscribers.
	//
	// Subscribers is a required field
	Subscribers []Subscriber `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNotificationInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}

	if s.Notification == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notification"))
	}

	if s.Subscribers == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subscribers"))
	}
	if s.Subscribers != nil && len(s.Subscribers) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Subscribers", 1))
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			invalidParams.AddNested("Notification", err.(aws.ErrInvalidParams))
		}
	}
	if s.Subscribers != nil {
		for i, v := range s.Subscribers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Subscribers", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of CreateNotification
type CreateNotificationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNotification = "CreateNotification"

// CreateNotificationRequest returns a request value for making API operation for
// AWS Budgets.
//
// Creates a notification. You must create the budget before you create the
// associated notification.
//
//    // Example sending a request using CreateNotificationRequest.
//    req := client.CreateNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateNotificationRequest(input *CreateNotificationInput) CreateNotificationRequest {
	op := &aws.Operation{
		Name:       opCreateNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNotificationInput{}
	}

	req := c.newRequest(op, input, &CreateNotificationOutput{})
	return CreateNotificationRequest{Request: req, Input: input, Copy: c.CreateNotificationRequest}
}

// CreateNotificationRequest is the request type for the
// CreateNotification API operation.
type CreateNotificationRequest struct {
	*aws.Request
	Input *CreateNotificationInput
	Copy  func(*CreateNotificationInput) CreateNotificationRequest
}

// Send marshals and sends the CreateNotification API request.
func (r CreateNotificationRequest) Send(ctx context.Context) (*CreateNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNotificationResponse{
		CreateNotificationOutput: r.Request.Data.(*CreateNotificationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNotificationResponse is the response type for the
// CreateNotification API operation.
type CreateNotificationResponse struct {
	*CreateNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNotification request.
func (r *CreateNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
