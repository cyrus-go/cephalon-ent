package enums

type UserType string

const (
	UserTypePersonal   UserType = "personal"
	UserTypeEnterprise UserType = "enterprise"
)

func (obj UserType) Values() []string {
	return []string{
		string(UserTypePersonal),
		string(UserTypeEnterprise),
	}
}

func (obj UserType) Ptr() *UserType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
