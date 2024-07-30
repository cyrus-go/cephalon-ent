package enums

type StarStatus string

const (
	UnknownStartStatus StarStatus = "unknown"
	Star               Model      = "star"
	UnStar             Model      = "unstar"
)

func (obj StarStatus) Values() []string {
	return []string{
		string(UnknownStartStatus),
		string(Star),
		string(UnStar),
	}
}

func (obj StarStatus) Ptr() *StarStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
