package enums

type UserType string

const (
	UserTypePersonal   UserType = "personal"
	UserTypeEnterprise UserType = "enterprise"
)

func (m UserType) Values() []string {
	return []string{
		string(UserTypePersonal),
		string(UserTypeEnterprise),
	}
}
