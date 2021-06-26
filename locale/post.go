package locale

// Constant ...
const (
	PostKeyRequiredContent = "post_require_content"
	PostKeyRequiredTitle   = "post_require_title"
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
}
