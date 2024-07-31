package enums

type Model string

const (
	UnknownModel  Model = "unknown"
	LanguageModel Model = "language"
)

func (obj Model) Values() []string {
	return []string{
		string(UnknownModel),
		string(LanguageModel),
	}
}

func (obj Model) Ptr() *Model {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
