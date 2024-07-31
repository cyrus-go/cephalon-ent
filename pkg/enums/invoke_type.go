package enums

type Invoke_type string

const (
	UnKnownInvokeType Model = "unknown"
	ApiInvokeType     Model = "api"
	ChatInvokeType    Model = "web"
)

func (obj Invoke_type) Values() []string {
	return []string{
		string(UnknownModel),
		string(ApiInvokeType),
		string(ChatInvokeType),
	}
}

func (obj Invoke_type) Ptr() *Invoke_type {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
