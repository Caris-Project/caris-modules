package locale

// Constant
const (
	CommonKeySuccess                 = "common_success"
	CommonKeyBadRequest              = "common_bad_request"
	CommonKeyUnauthorized            = "common_unauthorized"
	CommonKeyNotFound                = "common_not_found"
	CommonKeyInvalidChecksum         = "common_invalid_checksum"
	CommonKeyInvalidOTP              = "common_invalid_otp"
	CommonKeyInvalidPhoneCountryCode = "common_invalid_phone_country_code"
	CommonKeyInvalidPhoneNumber      = "common_invalid_phone_number"
	CommonKeyInvalidFacebookToken    = "common_invalid_facebook_token"
	CommonKeyInvalidGoogleToken      = "common_invalid_google_token"
	CommonKeyInvalidAppleToken       = "common_invalid_apple_token"
)

// 1-199
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
	{
		Key:     CommonKeyInvalidOTP,
		Message: "mã xác nhận không chính xác hoặc đã hết hạn",
		Code:    6,
	},
	{
		Key:     CommonKeyInvalidPhoneCountryCode,
		Message: "mã quốc gia không hợp lệ",
		Code:    7,
	},
	{
		Key:     CommonKeyInvalidPhoneNumber,
		Message: "số điện thoại không hợp lệ",
		Code:    8,
	},
	{
		Key:     CommonKeyInvalidFacebookToken,
		Message: "thông tin Facebook không hợp lệ",
		Code:    9,
	},
	{
		Key:     CommonKeyInvalidGoogleToken,
		Message: "thông tin Google không hợp lệ",
		Code:    10,
	},
	{
		Key:     CommonKeyInvalidAppleToken,
		Message: "thông tin Apple không hợp lệ",
		Code:    11,
	},
}
