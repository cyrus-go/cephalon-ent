package enums

type TransferOrderGiftType string

const (
	TransferOrderGiftTypeNo       TransferOrderGiftType = "no"
	TransferOrderGiftTypeRegister TransferOrderGiftType = "register"
	TransferOrderGiftTypeSurvey   TransferOrderGiftType = "survey"
)

func (TransferOrderGiftType) Values() []string {
	return []string{
		string(TransferOrderGiftTypeNo),
		string(TransferOrderGiftTypeRegister),
		string(TransferOrderGiftTypeSurvey),
	}
}

func (obj TransferOrderGiftType) Ptr() *TransferOrderGiftType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
