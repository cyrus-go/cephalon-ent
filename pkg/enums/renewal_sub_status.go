package enums

type RenewalSubStatus string

const (
	RenewalSubStatusUnknow      RenewalSubStatus = "unknow"
	RenewalSubStatusSubscribing RenewalSubStatus = "subscribing"
	RenewalSubStatusFinished    RenewalSubStatus = "finished"
)

func (obj RenewalSubStatus) Values() []string {
	return []string{
		string(RenewalSubStatusUnknow),
		string(RenewalSubStatusSubscribing),
		string(RenewalSubStatusFinished),
	}
}

func (obj RenewalSubStatus) Ptr() *RenewalSubStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
