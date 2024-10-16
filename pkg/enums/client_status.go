package enums

type ClientStatus string

const (
	ClientStatusDisable ClientStatus = "disable"
	ClientStatusEnable  ClientStatus = "enable"
)

func (ClientStatus) Values() []string {
	return []string{
		string(ClientStatusDisable),
		string(ClientStatusEnable),
	}
}

func (obj ClientStatus) Ptr() *ClientStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
