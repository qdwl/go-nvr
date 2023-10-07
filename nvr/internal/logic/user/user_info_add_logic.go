package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/qdwl/go-nvr/nvr/internal/repository"
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
	user := &model.User{
		UserId:   req.UserId,
		UserName: req.UserName,
		Password: req.Password,
		RoleId:   req.RoleId,
		Remark:   req.Remark,
	}
	resp = new(types.UserInfoAddResp)

	err = repository.AddUser(user)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
	}

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()
	return
}
