package code

import "inquora/pkg/xcode"

var (
	UploadFileErr             = xcode.New(30001, "文件上传失败")
	ArtitleTitleEmpty         = xcode.New(30002, "文章标题为空")
	ArticleContentTooFewWords = xcode.New(30003, "文章内容字数太少")
	ArticleCoverEmpty         = xcode.New(30004, "文章封面为空")
)
