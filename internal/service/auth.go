package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	v1 "github.com/realotz/mstore/api/core/v1"
	"github.com/realotz/mstore/api/errors"
	userV1 "github.com/realotz/mstore/api/users/v1"
	"github.com/realotz/mstore/internal/biz"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/pkg/myctx"
	"github.com/realotz/mstore/pkg/utils"
	"time"
)

// 认证授权服务
func NewAuthService(uc *biz.AuthUseCase) userV1.AuthServiceServer {
	return &AuthService{uc: uc}
}

type AuthService struct {
	userV1.UnimplementedAuthServiceServer
	uc *biz.AuthUseCase
}

func (s AuthService) Captcha(ctx context.Context, req *userV1.CaptchaReq) (*userV1.CaptchaReply, error) {
	captchaImgData, err := s.uc.Captcha(ctx, req.Uuid)
	if err != nil {
		return nil, errors.ErrorBusinessError("获取图形验证码失败 %s", err.Error())
	}
	return &userV1.CaptchaReply{ImgBase64: captchaImgData}, nil
}

func (s AuthService) Login(ctx context.Context, req *userV1.LoginReq) (*userV1.LoginToken, error) {
	if req.Account == "" {
		return nil, errors.ErrorParamsError("请输入账号")
	}
	if req.Passwd == "" {
		return nil, errors.ErrorParamsError("请输入密码")
	}
	var user *biz.User
	var err error
	// 风险判定 输错密码达到4次
	if s.uc.CheckRisk(ctx, req.Account) {
		if req.Code == "" {
			return nil, errors.ErrorPreconditionRequired("请输入图形验证码")
		}
		if req.Uuid == "" {
			return nil, errors.ErrorPreconditionRequired("uuid不能为空")
		}
		if !s.uc.ValidateCaptcha(ctx, req.Uuid, req.Code) {
			return nil, errors.ErrorPreconditionRequired("图形验证码错误")
		}
	}
	user, err = s.uc.LoginForPassword(ctx, req.Account, req.Passwd)
	if err != nil {
		if s.uc.CheckRisk(ctx, req.Account) {
			return nil, errors.ErrorPreconditionRequired(err.Error())
		}
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return s.sendToken(ctx, user)
}

func (s AuthService) RefreshToken(ctx context.Context, empty *v1.Empty) (*userV1.LoginToken, error) {
	uid, err := myctx.FormUserId(ctx)
	if err != nil {
		return nil, errors.ErrorNotLogin(err.Error())
	}
	account, err := s.uc.GetUser(ctx, uid)
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return s.sendToken(ctx, account)
}

func (s AuthService) NewPasswd(ctx context.Context, req *userV1.NewPasswdReq) (*v1.Empty, error) {
	if req.Passwd == "" {
		return nil, errors.ErrorParamsError("请输入密码")
	}
	if req.Passwd != "" {
		if req.Passwd != req.PasswdConfirm {
			return nil, errors.ErrorParamsError("两次输入的密码不相同")
		}
		if !utils.VerifyIsPasswd(req.Passwd) {
			return nil, errors.ErrorParamsError("新密码必须包含数字、英文字符、字母大小写")
		}
	}
	err := s.uc.UpdatePassword(ctx, req.Account, req.Passwd, req.OldPasswd)
	if req.OldPasswd != "" {
		_, err := s.uc.LoginForPassword(ctx, req.Account, req.Passwd)
		if err != nil {
			return nil, errors.ErrorBusinessError(err.Error())
		}
	}
	if err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}

// 发送登录token
func (s *AuthService) sendToken(ctx context.Context, user *biz.User) (*userV1.LoginToken, error) {
	token, err := s.uc.CreateUserToken(myctx.JwtProfile{
		ID:   user.Id,
		Role: user.Auth.Role,
	}, "web")
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	tokenWrap := token
	claims, _ := s.uc.ParseUserToken(ctx, token)
	expires := time.Unix(claims.ExpiresAt, 0).Format(conf.TimeFormat)
	// header写入
	if info, ok := transport.FromServerContext(ctx); ok {
		info.ReplyHeader().Set("Authorization", tokenWrap)
		info.ReplyHeader().Set("Expires", expires)
	}
	return &userV1.LoginToken{
		Token:        token,
		TokenExpires: expires,
		UserInfo:     convertUser(user),
	}, nil
}
