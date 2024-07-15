package enums

type UserType string

const (
	UserTypePersonal      UserType = "personal"
	UserTypeEnterprise    UserType = "enterprise"
	UserTypeAdmin         UserType = "admin"
	UserTypeBoss          UserType = "boss"            // 超级管理员，可管理 admin 用户
	UserTypeAdminReadOnly UserType = "admin_read_only" // 仅可查看后台数据的管理员
)

func (obj UserType) Values() []string {
	return []string{
		string(UserTypePersonal),
		string(UserTypeEnterprise),
		string(UserTypeAdmin),
		string(UserTypeBoss),
		string(UserTypeAdminReadOnly),
	}
}

func (obj UserType) Ptr() *UserType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
