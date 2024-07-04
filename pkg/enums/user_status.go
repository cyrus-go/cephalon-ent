package enums

type UserStatus string

const (
	UserStatusNormal UserStatus = "normal"
	UserStatusFrozen UserStatus = "frozen"
	UserStatusClosed UserStatus = "closed" // 已注销
)

func (obj UserStatus) Values() []string {
	return []string{
		string(UserStatusNormal),
		string(UserStatusFrozen),
		string(UserStatusClosed),
	}
}

func (obj UserStatus) Ptr() *UserStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
