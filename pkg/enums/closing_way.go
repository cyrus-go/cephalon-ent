package enums

type CloseWay string

const (
	CloseWayUnknown          CloseWay = "unknown"
	CloseWayUser             CloseWay = "user"
	CloseWayBalanceNotEnough CloseWay = "balance_not_enough"
	CloseWayExpired          CloseWay = "expired"
	CloseWayAdmin            CloseWay = "admin"
)

func (CloseWay) Values() []string {
	return []string{
		string(CloseWayUnknown),
		string(CloseWayUser),
		string(CloseWayBalanceNotEnough),
		string(CloseWayExpired),
		string(CloseWayAdmin),
	}
}

func (obj CloseWay) Ptr() *CloseWay {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
