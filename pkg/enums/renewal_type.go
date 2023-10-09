package enums

type RenewalType string

const (
	RenewalTypeUnknow RenewalType = "unknow"
	RenewalTypeHour   RenewalType = "hour"
	RenewalTypeDay    RenewalType = "day"
	RenewalTypeMonth  RenewalType = "month"
)

func (obj RenewalType) Values() []string {
	return []string{
		string(RenewalTypeUnknow),
		string(RenewalTypeHour),
		string(RenewalTypeDay),
		string(RenewalTypeMonth),
	}
}

func (obj RenewalType) Ptr() *RenewalType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
