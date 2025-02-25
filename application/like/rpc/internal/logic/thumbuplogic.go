package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"inquora/application/like/rpc/internal/svc"
	"inquora/application/like/rpc/internal/types"
	"inquora/application/like/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpLogic {
	return &ThumbUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbUpLogic) ThumbUp(in *pb.ThumbUpRequest) (*pb.ThumbUpResponse, error) {
	// todo: add your logic here and delete this line
	msg := types.ThumbUpMsg{
		BizId:    in.BizId,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
		LikeType: in.LikeType,
	}
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("ThumbUp msg marshal error: %v", err)
			return
		}
		if err = l.svcCtx.KqPusherClient.Push(l.ctx, string(data)); err != nil {
			l.Logger.Errorf("ThumbUp msg push error: %v", err)
		}
	})
	return &pb.ThumbUpResponse{
		BizId:      in.BizId,
		ObjId:      in.ObjId,
		LikeNum:    666,
		DislikeNum: 666,
	}, nil
}
