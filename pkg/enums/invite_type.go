package enums

type InviteType string

const (
	InviteTypeUnknown       InviteType = "unknown"
	InviteTypeShareRegister InviteType = "share_register"
	InviteTypeAppInvite     InviteType = "app_invite"
)

func (InviteType) Values() []string {
	return []string{
		string(InviteTypeUnknown),
		string(InviteTypeShareRegister),
		string(InviteTypeAppInvite),
	}
}

func (obj InviteType) Ptr() *InviteType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
