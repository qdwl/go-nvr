package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleUpdateLogic {
	return &UserRoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleUpdateLogic) UserRoleUpdate(req *types.UserRoleUpdateReq) (resp *types.UserRoleUpdateResp, err error) {
	role := &model.Role{
		Id:     req.Id,
		Name:   req.Name,
		Type:   req.Type,
		Remark: req.Remark,
	}
	resp = new(types.UserRoleUpdateResp)

	err = repository.UpdateRole(role)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
	}

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()

	return
}
