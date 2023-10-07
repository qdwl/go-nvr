package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoDeleteLogic {
	return &UserInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoDeleteLogic) UserInfoDelete(req *types.UserInfoDeleteReq) (resp *types.UserInfoDeleteResp, err error) {
	resp = new(types.UserInfoDeleteResp)
	err = repository.DeleteUser(req.Id)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
	}

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()

	return
}
