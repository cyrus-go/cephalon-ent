package enums

type MissionType string

const (
	MissionTypeUnknown            MissionType = "unknown"
	MissionTypeSdTime             MissionType = "sd_time"
	MissionTypeSdTxt2Img          MissionType = "txt2img"
	MissionTypeSdImg2Img          MissionType = "img2img"
	MissionTypeJpTime             MissionType = "jp_time"
	MissionTypeWtTime             MissionType = "wt_time"
	MissionTypeSdExtraSingleImage MissionType = "extra-single-image"
	MissionTypeSdApi              MissionType = "sd_api"
	MissionTypeKeyPair            MissionType = "key_pair"
	MissionTypeJpDkTime           MissionType = "jp_dk_time"
	MissionTypeSshTime            MissionType = "ssh_time"
	MissionTypeSdTomatoTime       MissionType = "sd_tomato_time"
	MissionTypeSdTimePlan         MissionType = "sd_time_plan"
	MissionTypeWtTimePlan         MissionType = "wt_time_plan"
	MissionTypeJpTimePlan         MissionType = "jp_time_plan"
	MissionTypeJpDkTimePlan       MissionType = "jp_dk_time_plan"
	MissionTypeSshTimePlan        MissionType = "ssh_time_plan"
	MissionTypeSdTomatoTimePlan   MissionType = "sd_tomato_time_plan"
)

func (obj MissionType) Values() []string {
	return []string{
		string(MissionTypeUnknown),
		string(MissionTypeSdTime),
		string(MissionTypeSdTxt2Img),
		string(MissionTypeSdImg2Img),
		string(MissionTypeJpTime),
		string(MissionTypeWtTime),
		string(MissionTypeSdExtraSingleImage),
		string(MissionTypeSdApi),
		string(MissionTypeKeyPair),
		string(MissionTypeJpDkTime),
		string(MissionTypeSshTime),
		string(MissionTypeSdTimePlan),
		string(MissionTypeWtTimePlan),
		string(MissionTypeJpTimePlan),
		string(MissionTypeJpDkTimePlan),
		string(MissionTypeSshTimePlan),
		string(MissionTypeSdTomatoTime),
		string(MissionTypeSdTomatoTimePlan),
	}
}

func (obj MissionType) Ptr() *MissionType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
