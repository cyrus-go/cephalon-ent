package enums

type CreateWay string

const (
	CreateWayUnknown     CreateWay = "unknown"
	CreateWayUser        CreateWay = "user"
	CreateWayLoadBalance CreateWay = "load_balance"
)

func (CreateWay) Values() []string {
	return []string{
		string(CreateWayUnknown),
		string(CreateWayUser),
		string(CreateWayLoadBalance),
	}
}

func (obj CreateWay) Ptr() *CreateWay {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
