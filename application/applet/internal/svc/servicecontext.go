package svc

import (
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/region"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"inquora/application/applet/internal/config"
	"inquora/application/article/rpc/article"
	"inquora/application/user/rpc/user"
	"inquora/pkg/interceptors"
)

type ServiceContext struct {
	Config      config.Config
	OssUploader *uploader.UploadManager
	BizRedis    *redis.Redis
	UserRpc     user.User
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
		BizRedis:   redis.MustNewRedis(c.BizRedis),
		ArticleRpc: article.NewArticle(zrpc.MustNewClient(c.ArticleRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
	}
}
