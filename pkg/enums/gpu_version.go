package enums

type GpuVersion string

const (
	GpuVersion2060   GpuVersion = "RTX2060"
	GpuVersion2060Ti GpuVersion = "RTX2060Ti"
	GpuVersion2070   GpuVersion = "RTX2070"
	GpuVersion2070Ti GpuVersion = "RTX2070Ti"
	GpuVersion2080   GpuVersion = "RTX2080"
	GpuVersion2080Ti GpuVersion = "RTX2080Ti"
	GpuVersion3060   GpuVersion = "RTX3060"
	GpuVersion3060Ti GpuVersion = "RTX3060Ti"
	GpuVersion3070   GpuVersion = "RTX3070"
	GpuVersion3070Ti GpuVersion = "RTX3070Ti"
	GpuVersion3080   GpuVersion = "RTX3080"
	GpuVersion3080Ti GpuVersion = "RTX3080Ti"
	GpuVersion3090   GpuVersion = "RTX3090"
	GpuVersion3090Ti GpuVersion = "RTX3090Ti"
	GpuVersion4060   GpuVersion = "RTX4060"
	GpuVersion4060Ti GpuVersion = "RTX4060Ti"
	GpuVersion4070   GpuVersion = "RTX4070"
	GpuVersion4070Ti GpuVersion = "RTX4070Ti"
	GpuVersion4080   GpuVersion = "RTX4080"
	GpuVersion4090   GpuVersion = "RTX4090"
	GpuVersionA800   GpuVersion = "A800"
	GpuVersionA100   GpuVersion = "A100"
	GpuVersionV100   GpuVersion = "V100"
)

func (obj GpuVersion) Values() []string {
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
		string(GpuVersionA800),
		string(GpuVersionA100),
		string(GpuVersionV100),
	}
}

func (obj GpuVersion) Ptr() *GpuVersion {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
