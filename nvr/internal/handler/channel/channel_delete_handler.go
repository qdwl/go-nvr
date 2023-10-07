package channel

import (
	"net/http"

	"github.com/qdwl/go-nvr/nvr/internal/logic/channel"
	"github.com/qdwl/go-nvr/nvr/internal/svc"
	"github.com/qdwl/go-nvr/nvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChannelDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := channel.NewChannelDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ChannelDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
