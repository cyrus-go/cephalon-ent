package enums

type UserModelRelation string

const (
	UnknownRelation UserModelRelation = "unknown"
	StarRelation    UserModelRelation = "star"
	UsedRelation    UserModelRelation = "used"
)

func (obj UserModelRelation) Values() []string {
	return []string{
		string(StarRelation),
		string(UsedRelation),
		string(UnknownRelation),
	}
}

func (obj UserModelRelation) Ptr() *UserModelRelation {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
