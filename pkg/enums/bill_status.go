package enums

type BillStatus string

const (
	BillStatusPending  BillStatus = "pending"
	BillStatusCanceled BillStatus = "canceled"
	BillStatusDone     BillStatus = "done"
)

func (BillStatus) Values() []string {
	return []string{
		string(BillStatusPending),
		string(BillStatusCanceled),
		string(BillStatusDone),
	}
}
