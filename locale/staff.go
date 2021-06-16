package locale

// Constant ...
const (
	StaffKeyRoleInvalidName = "staff_role_invalid_name"
	StaffKeyRoleNameExisted = "staff_role_name_existed"

	StaffKeyInvalidName     = "staff_invalid_name"
	StaffKeyInvalidPassword = "staff_invalid_password"
	StaffKeyInvalidRole     = "staff_invalid_role"
	StaffKeyNotFound        = "staff_not_found"
)

// 300-399
var staff = []Locale{
	{
		Key:     StaffKeyRoleInvalidName,
		Message: "tên nhóm nhân viên không hợp lệ",
		Code:    300,
	},
	{
		Key:     StaffKeyRoleNameExisted,
		Message: "nhóm nhân viên đã tồn tại",
		Code:    301,
	},
	{
		Key:     StaffKeyInvalidName,
		Message: "tên nhân viên không hợp lệ",
		Code:    302,
	},
	{
		Key:     StaffKeyInvalidPassword,
		Message: "mật khẩu phải từ 6-32 ký từ",
		Code:    303,
	},
	{
		Key:     StaffKeyInvalidRole,
		Message: "nhóm quyền không hợp lệ",
		Code:    304,
	},
	{
		Key:     StaffKeyNotFound,
		Message: "nhân viên không tìm thấy",
		Code:    305,
	},
}
