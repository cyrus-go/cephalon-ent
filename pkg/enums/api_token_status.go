package enums

type ApiTokenStatus string

const (
	UnknownApiTokenStatus   ApiTokenStatus = "unknown"
	InitApiTokenStatus      ApiTokenStatus = "init"
	DelApiTokenStatus       ApiTokenStatus = "del"
	ForbiddenApiTokenStatus ApiTokenStatus = "forbidden"
)

func (obj ApiTokenStatus) Values() []string {
	return []string{
		string(UnknownApiTokenStatus),
		string(InitApiTokenStatus),
		string(ForbiddenApiTokenStatus),
	}
}

func (obj ApiTokenStatus) Ptr() *ApiTokenStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
