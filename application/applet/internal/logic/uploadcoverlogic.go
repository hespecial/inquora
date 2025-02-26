package logic

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"inquora/application/applet/internal/code"
	"net/http"
	"time"

	"inquora/application/applet/internal/svc"
	"inquora/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	maxFileSize    = 10 << 20 // 10MB
	coverUploadKey = "cover"
)

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover(req *http.Request) (resp *types.UploadCoverResponse, err error) {
	if err = req.ParseMultipartForm(maxFileSize); err != nil {
		return nil, err
	}
	file, header, err := req.FormFile(coverUploadKey)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	key := l.genFilename(header.Filename)
	err = l.svcCtx.OssUploader.UploadReader(l.ctx, file, &uploader.ObjectOptions{
		BucketName: l.svcCtx.Config.Oss.BucketName,
		ObjectName: &key,
		FileName:   key,
	}, nil)
	if err != nil {
		l.Logger.Errorf("UploadCover err: %v", err)
		return nil, code.UploadFileErr
	}

	return &types.UploadCoverResponse{
		CoverUrl: l.genFileURL(key),
	}, nil
}

func (l *UploadCoverLogic) genFilename(filename string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixMilli(), filename)
}

func (l *UploadCoverLogic) genFileURL(key string) string {
	return fmt.Sprintf("%s/%s", l.svcCtx.Config.Oss.BaseUrl, key)
}
