package enums

type UserStatus string

const (
	UserStatusNormal UserStatus = "normal"
	UserStatusFrozen UserStatus = "frozen"
)

func (obj UserStatus) Values() []string {
	return []string{
		string(UserStatusNormal),
		string(UserStatusFrozen),
	}
}

func (obj UserStatus) Ptr() *UserStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
