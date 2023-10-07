package channel

import (
	"context"
	"strings"

	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/qdwl/go-nvr/nvr/internal/repository"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"
	"github.com/qdwl/go-nvr/nvr/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelListLogic {
	return &ChannelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelListLogic) ChannelList(req *types.ChannelListReq) (resp *types.ChannelListResp, err error) {
	list := make([]*types.ChannelInfo, 0)
	req.Size = utils.If(req.Size == 0, 20, req.Size).(int)
	req.Page = utils.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)

	channel := &model.Channel{}
	if req.Id != nil {
		channel.Id = *req.Id
	}
	if len(strings.TrimSpace(req.Name)) > 0 {
		channel.Name = req.Name
	}
	if len(strings.TrimSpace(req.Ip)) > 0 {
		channel.Ip = req.Ip
	}
	// logx.Infof("deviceId:%s, deviceName:%s, status:%d", device.DeviceId, device.Name, device.Status)

	resp = new(types.ChannelListResp)

	count, channels, err := repository.GetChannelList(req.Page, req.Size, channel)
	if err != nil {
		resp.Code = int(types.RESTFUL_ERR_DATABASE_OPERATION_FAILED)
		resp.Msg = types.RESTFUL_ERR_DATABASE_OPERATION_FAILED.String()
		return resp, err
	}

	for _, item := range channels {
		channelInfo := &types.ChannelInfo{}
		utils.Copy(channelInfo, item)

		list = append(list, channelInfo)
	}

	resp.Data.Count = count
	resp.Data.List = list

	resp.Code = int(types.RESTFUL_ERR_OK)
	resp.Msg = types.RESTFUL_ERR_OK.String()

	return
}
