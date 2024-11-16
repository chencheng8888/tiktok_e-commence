package service

import (
	"context"

	pb "github.com/chencheng8888/tiktok_e-commence/authService/api/auth/v1"
)

type AuthServiceService struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceService() *AuthServiceService {
	return &AuthServiceService{}
}

func (s *AuthServiceService) DeliverTokenByRPC(ctx context.Context, req *pb.DeliverTokenReq) (*pb.DeliveryResp, error) {
	return &pb.DeliveryResp{}, nil
}
func (s *AuthServiceService) VerifyTokenByRPC(ctx context.Context, req *pb.VerifyTokenReq) (*pb.VerifyResp, error) {
	return &pb.VerifyResp{}, nil
}
func (s *AuthServiceService) AddUserToBlackList(ctx context.Context, req *pb.AddUserToBlackListReq) (*pb.AddUserToBlackListResp, error) {
	return &pb.AddUserToBlackListResp{}, nil
}
func (s *AuthServiceService) RemoveUserFromBlackList(ctx context.Context, req *pb.RemoveUserFromBlackListReq) (*pb.RemoveUserFromBlackListResp, error) {
	return &pb.RemoveUserFromBlackListResp{}, nil
}
