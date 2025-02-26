package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Oss struct {
		RegionId        string
		BaseUrl         string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
	}
	BizRedis   redis.RedisConf
	UserRpc    zrpc.RpcClientConf
	ArticleRpc zrpc.RpcClientConf
}
