package enums

type MissionCategory string

const (
	MissionCategorySD   MissionCategory = "SD"
	MissionCategoryJP   MissionCategory = "JP"
	MissionCategoryWT   MissionCategory = "WT"
	MissionCategoryJPDK MissionCategory = "JP_DK"
)

func (obj MissionCategory) Values() []string {
	return []string{
		string(MissionCategorySD),
		string(MissionCategoryJP),
		string(MissionCategoryWT),
		string(MissionCategoryJPDK),
	}
}

func (obj MissionCategory) Ptr() *MissionCategory {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
