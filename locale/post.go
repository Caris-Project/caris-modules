package locale

// Constant ...
const (
	PostKeyRequiredContent      = "post_require_content"
	PostKeyRequiredTitle        = "post_require_title"
	PostKeyRequiredBriefContent = "post_require_brief_content"
)

// 500-599
var post = []Locale{
	{
		Key:     PostKeyRequiredContent,
		Message: "Nội dung bài viết không được trống!",
		Code:    500,
	},
	{
		Key:     PostKeyRequiredTitle,
		Message: "Tiêu đề bài viết không được trống!",
		Code:    501,
	},
	{
		Key:     PostKeyRequiredBriefContent,
		Message: "Nội dung bài viết ngắn không được trống!",
		Code:    502,
	},
}
