// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateCommitInput struct {
	_ struct{} `type:"structure"`

	// The name of the author who created the commit. This information is used as
	// both the author and committer for the commit.
	AuthorName *string `locationName:"authorName" type:"string"`

	// The name of the branch where you create the commit.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// The commit message you want to include in the commit. Commit messages are
	// limited to 256 KB. If no message is specified, a default message is used.
	CommitMessage *string `locationName:"commitMessage" type:"string"`

	// The files to delete in this commit. These files still exist in earlier commits.
	DeleteFiles []DeleteFileEntry `locationName:"deleteFiles" type:"list"`

	// The email address of the person who created the commit.
	Email *string `locationName:"email" type:"string"`

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If true, a ..gitkeep file is created
	// for empty folders. The default is false.
	KeepEmptyFolders *bool `locationName:"keepEmptyFolders" type:"boolean"`

	// The ID of the commit that is the parent of the commit you create. Not required
	// if this is an empty repository.
	ParentCommitId *string `locationName:"parentCommitId" type:"string"`

	// The files to add or update in this commit.
	PutFiles []PutFileEntry `locationName:"putFiles" type:"list"`

	// The name of the repository where you create the commit.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The file modes to update for files in this commit.
	SetFileModes []SetFileModeEntry `locationName:"setFileModes" type:"list"`
}

// String returns the string representation
func (s CreateCommitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCommitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCommitInput"}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}
	if s.DeleteFiles != nil {
		for i, v := range s.DeleteFiles {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DeleteFiles", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.PutFiles != nil {
		for i, v := range s.PutFiles {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PutFiles", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SetFileModes != nil {
		for i, v := range s.SetFileModes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SetFileModes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCommitOutput struct {
	_ struct{} `type:"structure"`

	// The full commit ID of the commit that contains your committed file changes.
	CommitId *string `locationName:"commitId" type:"string"`

	// The files added as part of the committed file changes.
	FilesAdded []FileMetadata `locationName:"filesAdded" type:"list"`

	// The files deleted as part of the committed file changes.
	FilesDeleted []FileMetadata `locationName:"filesDeleted" type:"list"`

	// The files updated as part of the commited file changes.
	FilesUpdated []FileMetadata `locationName:"filesUpdated" type:"list"`

	// The full SHA-1 pointer of the tree information for the commit that contains
	// the commited file changes.
	TreeId *string `locationName:"treeId" type:"string"`
}

// String returns the string representation
func (s CreateCommitOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCommit = "CreateCommit"

// CreateCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Creates a commit for a repository on the tip of a specified branch.
//
//    // Example sending a request using CreateCommitRequest.
//    req := client.CreateCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/CreateCommit
func (c *Client) CreateCommitRequest(input *CreateCommitInput) CreateCommitRequest {
	op := &aws.Operation{
		Name:       opCreateCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCommitInput{}
	}

	req := c.newRequest(op, input, &CreateCommitOutput{})
	return CreateCommitRequest{Request: req, Input: input, Copy: c.CreateCommitRequest}
}

// CreateCommitRequest is the request type for the
// CreateCommit API operation.
type CreateCommitRequest struct {
	*aws.Request
	Input *CreateCommitInput
	Copy  func(*CreateCommitInput) CreateCommitRequest
}

// Send marshals and sends the CreateCommit API request.
func (r CreateCommitRequest) Send(ctx context.Context) (*CreateCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCommitResponse{
		CreateCommitOutput: r.Request.Data.(*CreateCommitOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCommitResponse is the response type for the
// CreateCommit API operation.
type CreateCommitResponse struct {
	*CreateCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCommit request.
func (r *CreateCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
