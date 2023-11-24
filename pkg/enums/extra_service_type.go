package enums

type ExtraServiceType string

const (
	ExtraServiceTypeUnknown ExtraServiceType = "unknown"
	ExtraServiceTypeVPN     ExtraServiceType = "vpn"
)

func (obj ExtraServiceType) Values() []string {
	return []string{
		string(ExtraServiceTypeUnknown),
		string(ExtraServiceTypeVPN),
	}
}
