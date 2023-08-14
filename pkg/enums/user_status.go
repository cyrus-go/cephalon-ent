package enums

type UserStatus string

const (
	UserStatusNormal UserStatus = "normal"
	UserStatusFrozen UserStatus = "frozen"
)

func (m UserStatus) Values() []string {
	return []string{
		string(UserStatusNormal),
		string(UserStatusFrozen),
	}
}
