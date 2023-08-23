package enums

type MissionType string

const (
	MissionTypeSdTime             MissionType = "sd_time"
	MissionTypeSdTxt2Img          MissionType = "txt2img"
	MissionTypeSdImg2Img          MissionType = "img2img"
	MissionTypeJpTime             MissionType = "jp_time"
	MissionTypeWtTime             MissionType = "wt_time"
	MissionTypeSdExtraSingleImage MissionType = "extra-single-image"
)

func (m MissionType) Values() []string {
	return []string{
		string(MissionTypeSdTime),
		string(MissionTypeSdTxt2Img),
		string(MissionTypeSdImg2Img),
		string(MissionTypeJpTime),
		string(MissionTypeWtTime),
		string(MissionTypeSdExtraSingleImage),
	}
}
