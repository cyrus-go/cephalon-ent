package enums

type MissionCategory string

const (
	MissionCategorySD MissionCategory = "SD"
	MissionCategoryJP MissionCategory = "JP"
	MissionCategoryWT MissionCategory = "WT"
)

func (m MissionCategory) Values() []string {
	return []string{
		string(MissionCategorySD),
		string(MissionCategoryJP),
		string(MissionCategoryWT),
	}
}
