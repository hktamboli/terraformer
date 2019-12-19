// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListFargateProfilesInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster that you would like to listFargate profiles
	// in.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The maximum number of Fargate profile results returned by ListFargateProfiles
	// in paginated output. When you use this parameter, ListFargateProfiles returns
	// only maxResults results in a single page along with a nextToken response
	// element. You can see the remaining results of the initial request by sending
	// another ListFargateProfiles request with the returned nextToken value. This
	// value can be between 1 and 100. If you don't use this parameter, ListFargateProfiles
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListFargateProfiles
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFargateProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFargateProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFargateProfilesInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFargateProfilesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListFargateProfilesOutput struct {
	_ struct{} `type:"structure"`

	// A list of all of the Fargate profiles associated with the specified cluster.
	FargateProfileNames []string `locationName:"fargateProfileNames" type:"list"`

	// The nextToken value to include in a future ListFargateProfiles request. When
	// the results of a ListFargateProfiles request exceed maxResults, you can use
	// this value to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFargateProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFargateProfilesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FargateProfileNames != nil {
		v := s.FargateProfileNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "fargateProfileNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFargateProfiles = "ListFargateProfiles"

// ListFargateProfilesRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Lists the AWS Fargate profiles associated with the specified cluster in your
// AWS account in the specified Region.
//
//    // Example sending a request using ListFargateProfilesRequest.
//    req := client.ListFargateProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/ListFargateProfiles
func (c *Client) ListFargateProfilesRequest(input *ListFargateProfilesInput) ListFargateProfilesRequest {
	op := &aws.Operation{
		Name:       opListFargateProfiles,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/fargate-profiles",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFargateProfilesInput{}
	}

	req := c.newRequest(op, input, &ListFargateProfilesOutput{})
	return ListFargateProfilesRequest{Request: req, Input: input, Copy: c.ListFargateProfilesRequest}
}

// ListFargateProfilesRequest is the request type for the
// ListFargateProfiles API operation.
type ListFargateProfilesRequest struct {
	*aws.Request
	Input *ListFargateProfilesInput
	Copy  func(*ListFargateProfilesInput) ListFargateProfilesRequest
}

// Send marshals and sends the ListFargateProfiles API request.
func (r ListFargateProfilesRequest) Send(ctx context.Context) (*ListFargateProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFargateProfilesResponse{
		ListFargateProfilesOutput: r.Request.Data.(*ListFargateProfilesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFargateProfilesRequestPaginator returns a paginator for ListFargateProfiles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFargateProfilesRequest(input)
//   p := eks.NewListFargateProfilesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFargateProfilesPaginator(req ListFargateProfilesRequest) ListFargateProfilesPaginator {
	return ListFargateProfilesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFargateProfilesInput
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

// ListFargateProfilesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFargateProfilesPaginator struct {
	aws.Pager
}

func (p *ListFargateProfilesPaginator) CurrentPage() *ListFargateProfilesOutput {
	return p.Pager.CurrentPage().(*ListFargateProfilesOutput)
}

// ListFargateProfilesResponse is the response type for the
// ListFargateProfiles API operation.
type ListFargateProfilesResponse struct {
	*ListFargateProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFargateProfiles request.
func (r *ListFargateProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
