// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListHandshakesForOrganizationInput struct {
	_ struct{} `type:"structure"`

	// A filter of the handshakes that you want included in the response. The default
	// is all types. Use the ActionType element to limit the output to only a specified
	// type, such as INVITE, ENABLE-ALL-FEATURES, or APPROVE-ALL-FEATURES. Alternatively,
	// for the ENABLE-ALL-FEATURES handshake that generates a separate child handshake
	// for each member account, you can specify the ParentHandshakeId to see only
	// the handshakes that were generated by that parent request.
	Filter *HandshakeFilter `type:"structure"`

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific
	// to the operation. If additional items exist beyond the maximum you specify,
	// the NextToken response element is present and has a value (is not null).
	// Include that value as the NextToken request parameter in the next call to
	// the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that
	// you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more
	// output is available. Set this parameter to the value of the previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHandshakesForOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHandshakesForOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHandshakesForOrganizationInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListHandshakesForOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// A list of Handshake objects with details about each of the handshakes that
	// are associated with an organization.
	Handshakes []Handshake `type:"list"`

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You
	// should repeat this until the NextToken response element comes back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHandshakesForOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opListHandshakesForOrganization = "ListHandshakesForOrganization"

// ListHandshakesForOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the handshakes that are associated with the organization that the requesting
// user is part of. The ListHandshakesForOrganization operation returns a list
// of handshake structures. Each structure contains details and status about
// a handshake.
//
// Handshakes that are ACCEPTED, DECLINED, or CANCELED appear in the results
// of this API for only 30 days after changing to that state. After that, they're
// deleted and no longer accessible.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account
// or by a member account that is a delegated administrator for an AWS service.
//
//    // Example sending a request using ListHandshakesForOrganizationRequest.
//    req := client.ListHandshakesForOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListHandshakesForOrganization
func (c *Client) ListHandshakesForOrganizationRequest(input *ListHandshakesForOrganizationInput) ListHandshakesForOrganizationRequest {
	op := &aws.Operation{
		Name:       opListHandshakesForOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListHandshakesForOrganizationInput{}
	}

	req := c.newRequest(op, input, &ListHandshakesForOrganizationOutput{})
	return ListHandshakesForOrganizationRequest{Request: req, Input: input, Copy: c.ListHandshakesForOrganizationRequest}
}

// ListHandshakesForOrganizationRequest is the request type for the
// ListHandshakesForOrganization API operation.
type ListHandshakesForOrganizationRequest struct {
	*aws.Request
	Input *ListHandshakesForOrganizationInput
	Copy  func(*ListHandshakesForOrganizationInput) ListHandshakesForOrganizationRequest
}

// Send marshals and sends the ListHandshakesForOrganization API request.
func (r ListHandshakesForOrganizationRequest) Send(ctx context.Context) (*ListHandshakesForOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHandshakesForOrganizationResponse{
		ListHandshakesForOrganizationOutput: r.Request.Data.(*ListHandshakesForOrganizationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHandshakesForOrganizationRequestPaginator returns a paginator for ListHandshakesForOrganization.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHandshakesForOrganizationRequest(input)
//   p := organizations.NewListHandshakesForOrganizationRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHandshakesForOrganizationPaginator(req ListHandshakesForOrganizationRequest) ListHandshakesForOrganizationPaginator {
	return ListHandshakesForOrganizationPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHandshakesForOrganizationInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListHandshakesForOrganizationPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHandshakesForOrganizationPaginator struct {
	aws.Pager
}

func (p *ListHandshakesForOrganizationPaginator) CurrentPage() *ListHandshakesForOrganizationOutput {
	return p.Pager.CurrentPage().(*ListHandshakesForOrganizationOutput)
}

// ListHandshakesForOrganizationResponse is the response type for the
// ListHandshakesForOrganization API operation.
type ListHandshakesForOrganizationResponse struct {
	*ListHandshakesForOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHandshakesForOrganization request.
func (r *ListHandshakesForOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
