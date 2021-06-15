package locale

// Constant ...
const (
	StaffKeyRoleInvalidName = "staff_role_invalid_name"
	StaffKeyRoleNameExisted = "staff_role_name_existed"
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
}
