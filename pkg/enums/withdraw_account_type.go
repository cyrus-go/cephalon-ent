package enums

type BusinessType string

const (
	BusinessTypeYun     BusinessType = "yun"
	BusinessTypeWFT     BusinessType = "wft"
	BusinessTypeCompany BusinessType = "company" // 对公
)

func (obj BusinessType) Values() []string {
	return []string{
		string(BusinessTypeYun),
		string(BusinessTypeWFT),
		string(BusinessTypeCompany),
	}
}

func (obj BusinessType) Ptr() *BusinessType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
