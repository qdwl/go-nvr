package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoAddLogic {
	return &UserInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoAddLogic) UserInfoAdd(req *types.UserInfoAddReq) (resp *types.UserInfoAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
