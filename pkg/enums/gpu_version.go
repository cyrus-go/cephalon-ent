package enums

type GpuVersion string

const (
	GpuVersion2060 GpuVersion = "RTX 2060"
	GpuVersion3070 GpuVersion = "RTX 3070"
	GpuVersion3080 GpuVersion = "RTX 3080"
	GpuVersion3090 GpuVersion = "RTX 3090"
	GpuVersion4070 GpuVersion = "RTX 4070"
	GpuVersion4080 GpuVersion = "RTX 4080"
	GpuVersion4090 GpuVersion = "RTX 4090"
)

func (g GpuVersion) Values() []string {
	return []string{
		string(GpuVersion2060),
		string(GpuVersion3070),
		string(GpuVersion3080),
		string(GpuVersion3090),
		string(GpuVersion4070),
		string(GpuVersion4080),
		string(GpuVersion4090),
	}
}
