package enums

type InnerMethod string

const (
	InnerMethodPost InnerMethod = "POST"
	InnerMethodGet  InnerMethod = "GET"
	InnerMethodHead InnerMethod = "HEAD"
)

func (obj InnerMethod) Values() []string {
	return []string{
		string(InnerMethodPost),
		string(InnerMethodGet),
		string(InnerMethodHead),
	}
}

func (obj InnerMethod) Ptr() *InnerMethod {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
