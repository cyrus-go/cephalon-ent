package enums

type ModelStatus string

const (
	UnknownModelStatus ModelStatus = "unknown"
	InitModelStatus    ModelStatus = "init"
)

func (obj ModelStatus) Values() []string {
	return []string{
		string(UnknownModel),
		string(InitModelStatus),
	}
}

func (obj ModelStatus) Ptr() *ModelStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
