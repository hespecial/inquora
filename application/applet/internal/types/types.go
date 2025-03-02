// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type ArticleDetailRequest struct {
	ArticleId int64 `form:"article_id"`
}

type ArticleDetailResponse struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
}

type ArticleInfo struct {
	ArticleId   int64  `json:"article_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type ArticleListRequest struct {
	AuthorId  int64 `form:"author_id"`
	Cursor    int64 `form:"cursor"`
	PageSize  int64 `form:"page_size"`
	SortType  int32 `form:"sort_type"`
	ArticleId int64 `form:"article_id"`
}

type ArticleListResponse struct {
	Articles []ArticleInfo `json:"articles"`
}

type LoginRequest struct {
	Mobile           string `json:"mobile"`
	VerificationCode string `json:"verification_code"`
}

type LoginResponse struct {
	UserId int64 `json:"user_id"`
	Token  Token `json:"token"`
}

type PublishRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type PublishResponse struct {
	ArticleId int64 `json:"article_id"`
}

type RegisterRequest struct {
	Name             string `json:"name"`
	Mobile           string `json:"mobile"`
	VerificationCode string `json:"verification_code"`
}

type RegisterResponse struct {
	UserId int64 `json:"user_id"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type UploadCoverResponse struct {
	CoverUrl string `json:"cover_url"`
}

type UserInfoResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type VerificationRequest struct {
	Mobile string `json:"mobile"`
}

type VerificationResponse struct {
}
