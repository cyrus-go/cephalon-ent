package enums

type PlatformWalletType string

const (
	PlatformWalletTypeProfit PlatformWalletType = "profit"
)

func (m PlatformWalletType) Values() []string {
	return []string{
		string(PlatformWalletTypeProfit),
	}
}
