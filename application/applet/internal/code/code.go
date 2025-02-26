package code

import "inquora/pkg/xcode"

var (
	RegisterMobileEmpty     = xcode.New(10001, "注册手机号不能为空")
	VerificationCodeEmpty   = xcode.New(10002, "验证码不能为空")
	MobileHasRegistered     = xcode.New(10003, "手机号已经注册")
	LoginMobileEmpty        = xcode.New(10004, "手机号不能为空")
	VerificationCodeExpire  = xcode.New(10005, "验证码已过期")
	VerificationCodeInvalid = xcode.New(10006, "验证码无效")

	UploadFileErr             = xcode.New(30001, "文件上传失败")
	ArtitleTitleEmpty         = xcode.New(30002, "文章标题为空")
	ArticleContentTooFewWords = xcode.New(30003, "文章内容字数太少")
	ArticleCoverEmpty         = xcode.New(30004, "文章封面为空")
)
