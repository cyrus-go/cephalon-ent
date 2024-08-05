package enums

type UserModelRelationStatus string

const (
	UnknownModelRelation UserModelRelationStatus = "unknown"
	StarModel            UserModelRelationStatus = "star"
	UnstartModel         UserModelRelationStatus = "unstar"
	UsedModel            UserModelRelationStatus = "used"
)

func (obj UserModelRelationStatus) Values() []string {
	return []string{
		string(UnknownModelRelation),
		string(StarModel),
		string(UnstartModel),
		string(UsedModel),
	}
}

func (obj UserModelRelationStatus) Ptr() *UserModelRelationStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
