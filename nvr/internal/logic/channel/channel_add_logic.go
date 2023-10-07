package channel

import (
	"context"

	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"
	"github.com/qdwl/go-nvr/nvr/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelAddLogic {
	return &ChannelAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelAddLogic) ChannelAdd(req *types.ChannelAddReq) (resp *types.ChannelAddResp, err error) {
	channel := new(model.Channel)
	utils.Copy(channel, req.ChannelInfo)

	resp = new(types.ChannelAddResp)

	err = repository.AddChannel(channel)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
		return
	}

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()
	return
}
