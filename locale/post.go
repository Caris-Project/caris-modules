package locale

// Constant ...
const (
	PostKeyRequiredContent      = "post_require_content"
	PostKeyRequiredTitle        = "post_require_title"
	PostKeyRequiredBriefContent = "post_require_brief_content"
	PostKeyStatusNotPending     = "post_status_not_pending"
	PostKeyIsPublishTimer       = "post_is_publish_timer"
	PostKeyChangePublishError  = "post_change_publish_err"
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
	{
		Key:     PostKeyStatusNotPending,
		Message: "Yêu cầu bài viết trạng thái chờ duyệt!",
		Code:    503,
	},
	{
		Key:     PostKeyIsPublishTimer,
		Message: "Yêu cầu bài viết loại hẹn giờ!",
		Code:    504,
	},
	{
		Key:     PostKeyChangePublishError,
		Message: "Đã xảy ra lỗi! Vui lòng thử lại!",
		Code:    505,
	},
}
