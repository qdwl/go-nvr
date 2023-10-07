package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoListLogic {
	return &UserInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoListLogic) UserInfoList(req *types.UserInfoListReq) (resp *types.UserInfoListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
