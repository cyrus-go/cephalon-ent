package enums

type CDKType string

const (
	CDKTypeUnknown   CDKType = "unknown"
	CDKTypeGetCep    CDKType = "get_cep"
	CDKTypeGetGpuUse CDKType = "get_gpu_use"
)

func (CDKType) Values() []string {
	return []string{
		string(CDKTypeUnknown),
		string(CDKTypeGetCep),
		string(CDKTypeGetGpuUse),
	}
}

func (obj CDKType) Ptr() *CDKType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
