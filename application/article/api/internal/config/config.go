package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	ArticleRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	Oss        struct {
		RegionId        string
		BaseUrl         string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
	}
}
