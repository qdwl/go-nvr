package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleDeleteLogic {
	return &UserRoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleDeleteLogic) UserRoleDelete(req *types.UserRoleDeleteReq) (resp *types.UserRoleDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
