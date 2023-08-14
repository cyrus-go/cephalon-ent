package enums

type MissionCategory string

const (
	MissionCategorySD      MissionCategory = "SD"
	MissionCategoryJupyter MissionCategory = "Jupyter"
	MissionCategoryWeTTY   MissionCategory = "WeTTy"
)

func (m MissionCategory) Values() []string {
	return []string{
		string(MissionCategorySD),
		string(MissionCategoryJupyter),
		string(MissionCategoryWeTTY),
	}
}
