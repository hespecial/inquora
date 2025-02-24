package svc

import (
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/region"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/zeromicro/go-zero/zrpc"
	"inquora/application/article/api/internal/config"
	"inquora/application/article/rpc/article"
)

type ServiceContext struct {
	Config      config.Config
	OssUploader *uploader.UploadManager
	ArticleRpc  article.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OssUploader: uploader.NewUploadManager(
			&uploader.UploadManagerOptions{
				Options: http_client.Options{
					Regions:     region.GetRegionByID(c.Oss.RegionId, true),
					Credentials: credentials.NewCredentials(c.Oss.AccessKeyId, c.Oss.AccessKeySecret),
				},
			},
		),
		ArticleRpc: article.NewArticle(zrpc.MustNewClient(c.ArticleRpc)),
	}
}
