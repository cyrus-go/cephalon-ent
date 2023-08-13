package enums

type GPU string

const (
	GPU3070   GPU = "3070"
	GPU3070Ti GPU = "3070Ti"
	GPU3080   GPU = "3080"
	GPU3080Ti GPU = "3080Ti"
	GPU3090   GPU = "3090"
	GPU3090Ti GPU = "3090Ti"
	GPU4070   GPU = "4070"
	GPU4070Ti GPU = "4070Ti"
	GPU4080   GPU = "4080"
	GPU4080Ti GPU = "4080Ti"
	GPU4090   GPU = "4090"
	GPU4090Ti GPU = "4090Ti"

	A100 GPU = "A100"
	V100 GPU = "V100"
)

func (g GPU) Values() []string {
	return []string{
		string(GPU3070),
		string(GPU3070Ti),
		string(GPU3080),
		string(GPU3080Ti),
		string(GPU3090),
		string(GPU3090Ti),
		string(GPU4070),
		string(GPU4070Ti),
		string(GPU4080),
		string(GPU4080Ti),
		string(GPU4090),
		string(GPU4090Ti),
		string(A100),
		string(V100),
	}
}
