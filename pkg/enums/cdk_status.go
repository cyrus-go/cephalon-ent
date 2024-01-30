package enums

type CDKStatus string

const (
	CDKStatusUnknown  CDKStatus = "unknown"
	CDKStatusNormal   CDKStatus = "normal"
	CDKStatusFreeze   CDKStatus = "freeze"
	CDKStatusUsed     CDKStatus = "used"
	CDKStatusCanceled CDKStatus = "canceled"
)

func (CDKStatus) Values() []string {
	return []string{
		string(CDKStatusUnknown),
		string(CDKStatusNormal),
		string(CDKStatusFreeze),
		string(CDKStatusUsed),
		string(CDKStatusCanceled),
	}
}

func (obj CDKStatus) Ptr() *CDKStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
