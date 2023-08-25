package enums

type PlatformWalletType string

const (
	PlatformWalletTypeProfit PlatformWalletType = "profit"
)

func (obj PlatformWalletType) Values() []string {
	return []string{
		string(PlatformWalletTypeProfit),
	}
}

func (obj PlatformWalletType) Ptr() *PlatformWalletType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
