package service

import (
	"context"

	pb "api/Seckill-Core/v1"

	"Seckill-Core/internal/biz"
)

// CreateUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.CreateUserReply
//	@return error
func (s *Seckill-CoreService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, err := s.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{Message: "trytest"}, nil
}
func (s *Seckill-CoreService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}

// DeleteUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.DeleteUserReply
//	@return error
func (s *Seckill-CoreService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}

// GetUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.GetUserReply
//	@return error
func (s *Seckill-CoreService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	userInfo, err := s.uc.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Code:    0,
		Message: "success",
		Data: &pb.GetUserReplyData{
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}

// GetUserByName
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.GetUserByNameReply
//	@return error
func (s *Seckill-CoreService) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetUserByNameReply, error) {
	userInfo, err := s.uc.GetUserInfoByName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByNameReply{
		Code:    0,
		Message: "success",
		Data: &pb.GetUserReplyData{
			Id:       int64(userInfo.UserID),
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}

// ListUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.ListUserReply
//	@return error
func (s *Seckill-CoreService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

// 执行定时任务
func (u *Seckill-CoreService) Cronjob(ctx context.Context, req *pb.CreateUserRequest) {
	u.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
}
