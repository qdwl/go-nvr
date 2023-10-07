package user

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"
	"github.com/qdwl/go-nvr/nvr/internal/utils"

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
	list := make([]*types.UserInfo, 0)
	req.Size = utils.If(req.Size == 0, 20, req.Size).(int)
	req.Page = utils.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)

	resp = new(types.UserInfoListResp)

	count, users, err := repository.GetUserList(req.Page, req.Size)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
		return resp, err
	}

	for _, user := range users {
		userInfo := &types.UserInfo{
			Id:       user.Id,
			UserId:   user.UserId,
			UserName: user.UserName,
			Password: user.Password,
			RoleId:   user.RoleId,
			Remark:   user.Remark,
		}

		list = append(list, userInfo)
	}

	resp.Data.Count = count
	resp.Data.List = list

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()

	return
}
