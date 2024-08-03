package enums

type StarStatus string

const (
	UnknownModelStart StarStatus = "unknown"
	ModelStar         StarStatus = "star"
	ModelUnStar       StarStatus = "unstar"
)

func (obj StarStatus) Values() []string {
	return []string{
		string(UnknownModelStart),
		string(ModelStar),
		string(ModelUnStar),
	}
}

func (obj StarStatus) Ptr() *StarStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
