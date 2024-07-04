package enums

type UserCloseType string

const (
	UserCloseTypeUnknown UserCloseType = "unknown"
	UserCloseTypeSelf    UserCloseType = "self"  // 自己注销
	UserCloseTypeAdmin   UserCloseType = "admin" // 管理员注销
)

func (obj UserCloseType) Values() []string {
	return []string{
		string(UserCloseTypeUnknown),
		string(UserCloseTypeSelf),
		string(UserCloseTypeAdmin),
	}
}

func (obj UserCloseType) Ptr() *UserCloseType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
