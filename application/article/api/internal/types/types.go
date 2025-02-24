// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type PublishRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type PublishResponse struct {
	ArticleId int64 `json:"article_id"`
}

type UploadCoverResponse struct {
	CoverUrl string `json:"cover_url"`
}
