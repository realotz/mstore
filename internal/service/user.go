package service

import (
	"context"
	v1 "github.com/realotz/mstore/api/core/v1"
	"github.com/realotz/mstore/api/errors"
	userV1 "github.com/realotz/mstore/api/users/v1"
	"github.com/realotz/mstore/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUserService(user *biz.UserUseCase) userV1.UserServiceServer {
	return &UserService{
		user: user,
	}
}

type UserService struct {
	userV1.UnimplementedUserServiceServer
	user *biz.UserUseCase
}

func (u UserService) ListUser(ctx context.Context, req *userV1.ListUserReq) (*userV1.UserListReply, error) {
	user, total, err := u.user.ListUser(ctx, biz.ListUserOption{
		Keyword: req.Keyword,
		Role:    req.Role,
		ListOption: biz.NewListOption(req.Option),
	})
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	var resp = &userV1.UserListReply{
		List:  make([]*v1.User, 0, len(user)),
		Total: total,
	}
	for _, v := range user {
		resp.List = append(resp.List, convertUser(v))
	}
	return resp, nil
}

func (u UserService) GetUser(ctx context.Context, id *v1.Id) (*v1.User, error) {
	user, err := u.user.GetUser(ctx, id.Id)
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return convertUser(user), nil
}

func (u UserService) DelUser(ctx context.Context, ids *v1.Ids) (*v1.Empty, error) {
	for _, v := range ids.Ids {
		err := u.user.DelUser(ctx, v)
		if err != nil {
			return nil, errors.ErrorBusinessError(err.Error())
		}
	}
	return &v1.Empty{}, nil
}

func (u UserService) CreateUser(ctx context.Context, req *userV1.CreateUserReq) (*v1.User, error) {
	user, err := u.user.CreateUser(ctx, &biz.User{
		Name: req.Name,
		Auth: &biz.Auth{
			Account:  req.Account,
			Password: req.Password,
			Role:     req.Role,
		},
	})
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return convertUser(user), nil
}

func (u UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserReq) (*v1.User, error) {
	if req.Id == 0 {
		return nil, errors.ErrorParamsError("id不能为空")
	}
	user, err := u.user.UpdateUser(ctx, req.Id, &biz.User{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	})
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return convertUser(user), nil
}

// user biz dto
func convertUser(user *biz.User) *v1.User {
	if user.Auth == nil {
		return &v1.User{
			Id:        user.Id,
			Name:      user.Name,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		}
	}
	var lastLoginTime *timestamppb.Timestamp
	if user.Auth.LastLoginTime != nil {
		lastLoginTime = timestamppb.New(*user.Auth.LastLoginTime)
	}
	return &v1.User{
		Id:            user.Id,
		Name:          user.Name,
		LastIp:        user.Auth.LastIp,
		LastLoginTime: lastLoginTime,
		Role:          user.Auth.Role,
		CreatedAt:     timestamppb.New(user.CreatedAt),
		UpdatedAt:     timestamppb.New(user.UpdatedAt),
	}
}
