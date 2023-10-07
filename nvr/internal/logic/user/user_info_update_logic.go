package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateReq) (resp *types.UserInfoUpdateResp, err error) {
	user := &model.User{
		Id:       req.Id,
		UserId:   req.UserId,
		UserName: req.UserName,
		Password: req.Password,
		RoleId:   req.RoleId,
		Remark:   req.Remark,
	}
	resp = new(types.UserInfoUpdateResp)

	err = repository.UpdateUser(user)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
	}

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()

	return
}
