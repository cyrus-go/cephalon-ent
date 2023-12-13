package enums

type BusinessType string

const (
	BusinessTypeYun BusinessType = "yun"
	BusinessTypeWFT BusinessType = "wft"
)

func (obj BusinessType) Values() []string {
	return []string{
		string(BusinessTypeYun),
		string(BusinessTypeWFT),
	}
}

func (obj BusinessType) Ptr() *BusinessType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
