package enums

type GpuVersion string

const (
	GpuVersion2060   GpuVersion = "RTX 2060"
	GpuVersion2060Ti GpuVersion = "RTX 2060 Ti"
	GpuVersion2070   GpuVersion = "RTX 2070"
	GpuVersion2070Ti GpuVersion = "RTX 2070 Ti"
	GpuVersion2080   GpuVersion = "RTX 2080"
	GpuVersion2080Ti GpuVersion = "RTX 2080 Ti"
	GpuVersion3060   GpuVersion = "RTX 3060"
	GpuVersion3060Ti GpuVersion = "RTX 3060 Ti"
	GpuVersion3070   GpuVersion = "RTX 3070"
	GpuVersion3070Ti GpuVersion = "RTX 3070 Ti"
	GpuVersion3080   GpuVersion = "RTX 3080"
	GpuVersion3080Ti GpuVersion = "RTX 3080 Ti"
	GpuVersion3090   GpuVersion = "RTX 3090"
	GpuVersion3090Ti GpuVersion = "RTX 3090 Ti"
	GpuVersion4060   GpuVersion = "RTX 4060"
	GpuVersion4060Ti GpuVersion = "RTX 4060 Ti"
	GpuVersion4070   GpuVersion = "RTX 4070"
	GpuVersion4070Ti GpuVersion = "RTX 4070 Ti"
	GpuVersion4080   GpuVersion = "RTX 4080"
	GpuVersion4090   GpuVersion = "RTX 4090"
)

func (g GpuVersion) Values() []string {
	return []string{
		string(GpuVersion2060),
		string(GpuVersion2060Ti),
		string(GpuVersion2070),
		string(GpuVersion2070Ti),
		string(GpuVersion2080),
		string(GpuVersion2080Ti),
		string(GpuVersion3060),
		string(GpuVersion3060Ti),
		string(GpuVersion3070),
		string(GpuVersion3070Ti),
		string(GpuVersion3080),
		string(GpuVersion3080Ti),
		string(GpuVersion3090),
		string(GpuVersion3090Ti),
		string(GpuVersion4060),
		string(GpuVersion4060Ti),
		string(GpuVersion4070),
		string(GpuVersion4070Ti),
		string(GpuVersion4080),
		string(GpuVersion4090),
	}
}
