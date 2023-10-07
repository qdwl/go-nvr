package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"
	"github.com/qdwl/go-nvr/nvr/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("login params invalid")
	}

	user, err := repository.GetUserByName(req.Username)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, errors.New("username or password error")
	}

	iat := time.Now().Unix()
	seconds := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, iat, seconds, user.Id)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		Base: types.Base{
			Code: int(types.RESTFUL_ERR_OK),
			Msg:  types.RESTFUL_ERR_OK.String(),
		},
		Data: types.UserLoginToken{
			AccessToken:  jwtToken,
			AccessExpire: iat + seconds,
			RefreshAfter: iat + seconds/2,
		},
	}, nil
}
