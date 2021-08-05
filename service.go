package http_grpc_example

import "context"

type Service struct {

}

var _ BootcampServer = &Service{}

// GetMember only gRPC
func (s Service) GetMember(ctx context.Context, request *GetMemberRequest) (*GetMemberResponse, error) {
	panic("implement me")
}

// GetAllMembers gRPC and HTTP access
func (s Service) GetAllMembers(ctx context.Context, request *GetAllMembersRequest) (*GetAllMembersResponse, error) {
	panic("implement me")
}

// CreateMember gRPC and HTTP access
func (s Service) CreateMember(ctx context.Context, request *CreateMemberRequest) (*CreateMemberResponse, error) {
	panic("implement me")
}