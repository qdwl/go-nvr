package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"
	"github.com/qdwl/go-nvr/nvr/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleListLogic {
	return &UserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleListLogic) UserRoleList(req *types.UserRoleListReq) (resp *types.UserRoleListResp, err error) {
	list := make([]*types.UserRole, 0)
	req.Size = utils.If(req.Size == 0, 20, req.Size).(int)
	req.Page = utils.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)

	resp = new(types.UserRoleListResp)

	count, roles, err := repository.GetRoleList(req.Page, req.Size)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
		return resp, err
	}

	for _, role := range roles {
		roleBasic := &types.UserRole{
			Name:   role.Name,
			Type:   role.Type,
			Remark: role.Remark,
		}

		list = append(list, roleBasic)
	}

	resp.Data.Count = count
	resp.Data.List = list

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()

	return
}
