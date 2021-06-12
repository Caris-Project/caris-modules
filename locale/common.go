package locale

// Constant
const (
	CommonKeySuccess         = "common_success"
	CommonKeyBadRequest      = "common_bad_request"
	CommonKeyUnauthorized    = "common_unauthorized"
	CommonKeyNotFound        = "common_not_found"
	CommonKeyInvalidChecksum = "common_invalid_checksum"
)

// 1-99
var common = []Locale{
	{
		Key:     CommonKeySuccess,
		Message: "thành công",
		Code:    1,
	},
	{
		Key:     CommonKeyBadRequest,
		Message: "dữ liệu không hợp lệ",
		Code:    2,
	},
	{
		Key:     CommonKeyUnauthorized,
		Message: "bạn không có quyền thực hiện hành động này",
		Code:    3,
	},
	{
		Key:     CommonKeyNotFound,
		Message: "dữ liệu không tìm thấy",
		Code:    4,
	},
	{
		Key:     CommonKeyInvalidChecksum,
		Message: "xác thực dữ liệu thất bại",
		Code:    5,
	},
}
