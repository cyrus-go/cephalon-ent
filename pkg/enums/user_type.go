package enums

type UserType string

const (
	UserTypePersonal   UserType = "personal"
	UserTypeEnterprise UserType = "enterprise"
	UserTypeAdmin      UserType = "admin"
)

func (obj UserType) Values() []string {
	return []string{
		string(UserTypePersonal),
		string(UserTypeEnterprise),
		string(UserTypeAdmin),
	}
}

func (obj UserType) Ptr() *UserType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
