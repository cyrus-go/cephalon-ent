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
	MissionTypeSdCmdTime          MissionType = "sd_cmd_time"
	MissionTypeSdBingoTime        MissionType = "sd_bingo_time"
	MissionTypeFooocusTime        MissionType = "fooocus_time"
	MissionTypeTabbyTime          MissionType = "tabby_time"
	MissionTypeJpCondaTime        MissionType = "jp_conda_time"
	MissionTypeJpMlTime           MissionType = "jp_ml_time"
	MissionTypeSdCatTime          MissionType = "sd_cat_time"
	MissionTypeSdFireTime         MissionType = "sd_fire_time"
	MissionTypeComfyuiTime        MissionType = "comfyui_time"
	MissionTypeJpDk3Time          MissionType = "jp_dk_3_time"
	MissionTypeSdXlTime           MissionType = "sd_xl_time"
	MissionTypeSdChickTime        MissionType = "sd_chick_time"
	MissionTypeAscendTime         MissionType = "ascend_time"
	MissionTypeSdWuShanTime       MissionType = "sd_wu_shan_time"
	MissionTypeSdLangTime         MissionType = "sd_lang_time"
	MissionTypeComfyuiKeTime      MissionType = "comfyui_ke_time"
	MissionTypeChatchatTime       MissionType = "chatchat_time"
	MissionTypeLoraTime           MissionType = "lora_time"
	MissionTypeFooocusWuTime      MissionType = "fooocus_wu_time"
	MissionTypeSvdBackTime        MissionType = "svd_back_time"
	MissionTypeSdJiTime           MissionType = "sd_ji_time"
	MissionTypeSdShangJinTime     MissionType = "sd_shang_jin_time"

	MissionTypeSdTimePlan         MissionType = "sd_time_plan"
	MissionTypeWtTimePlan         MissionType = "wt_time_plan"
	MissionTypeJpTimePlan         MissionType = "jp_time_plan"
	MissionTypeJpDkTimePlan       MissionType = "jp_dk_time_plan"
	MissionTypeSshTimePlan        MissionType = "ssh_time_plan"
	MissionTypeSdTomatoTimePlan   MissionType = "sd_tomato_time_plan"
	MissionTypeSdCmdTimePlan      MissionType = "sd_cmd_time_plan"
	MissionTypeSdBingoTimePlan    MissionType = "sd_bingo_time_plan"
	MissionTypeFooocusTimePlan    MissionType = "fooocus_time_plan"
	MissionTypeTabbyTimePlan      MissionType = "tabby_time_plan"
	MissionTypeJpCondaTimePlan    MissionType = "jp_conda_time_plan"
	MissionTypeJpMlTimePlan       MissionType = "jp_ml_time_plan"
	MissionTypeSdCatTimePlan      MissionType = "sd_cat_time_plan"
	MissionTypeSdFireTimePlan     MissionType = "sd_fire_time_plan"
	MissionTypeComfyuiTimePlan    MissionType = "comfyui_time_plan"
	MissionTypeJpDk3TimePlan      MissionType = "jp_dk_3_time_plan"
	MissionTypeSdXlTimePlan       MissionType = "sd_xl_time_plan"
	MissionTypeSdChickTimePlan    MissionType = "sd_chick_time_plan"
	MissionTypeAscendTimePlan     MissionType = "ascend_time_plan"
	MissionTypeSdWuShanTimePlan   MissionType = "sd_wu_shan_time_plan"
	MissionTypeSdLangTimePlan     MissionType = "sd_lang_time_plan"
	MissionTypeComfyuiKeTimePlan  MissionType = "comfyui_ke_time_plan"
	MissionTypeChatchatTimePlan   MissionType = "chatchat_time_plan"
	MissionTypeLoraTimePlan       MissionType = "lora_time_plan"
	MissionTypeFooocusWuTimePlan  MissionType = "fooocus_wu_time_plan"
	MissionTypeSvdBackTimePlan    MissionType = "svd_back_time_plan"
	MissionTypeSdJiTimePlan       MissionType = "sd_ji_time_plan"
	MissionTypeSdShangJinTimePlan MissionType = "sd_shang_jin_time_plan"
	MissionTypeWaitingTime        MissionType = "waiting_time"      //等待任务
	MissionTypeWaitingTimePlan    MissionType = "waiting_time_plan" //等待任务
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
		string(MissionTypeSdCmdTime),
		string(MissionTypeSdCmdTimePlan),
		string(MissionTypeSdBingoTime),
		string(MissionTypeSdBingoTimePlan),
		string(MissionTypeFooocusTime),
		string(MissionTypeFooocusTimePlan),
		string(MissionTypeTabbyTime),
		string(MissionTypeTabbyTimePlan),
		string(MissionTypeJpCondaTime),
		string(MissionTypeJpCondaTimePlan),
		string(MissionTypeJpMlTime),
		string(MissionTypeJpMlTimePlan),
		string(MissionTypeSdCatTime),
		string(MissionTypeSdCatTimePlan),
		string(MissionTypeSdFireTime),
		string(MissionTypeSdFireTimePlan),
		string(MissionTypeComfyuiTime),
		string(MissionTypeComfyuiTimePlan),
		string(MissionTypeJpDk3Time),
		string(MissionTypeJpDk3TimePlan),
		string(MissionTypeSdXlTime),
		string(MissionTypeSdXlTimePlan),
		string(MissionTypeSdChickTime),
		string(MissionTypeSdChickTimePlan),
		string(MissionTypeAscendTime),
		string(MissionTypeAscendTimePlan),
		string(MissionTypeSdWuShanTime),
		string(MissionTypeSdWuShanTimePlan),
		string(MissionTypeSdLangTime),
		string(MissionTypeSdLangTimePlan),
		string(MissionTypeComfyuiKeTime),
		string(MissionTypeComfyuiKeTimePlan),
		string(MissionTypeChatchatTime),
		string(MissionTypeChatchatTimePlan),
		string(MissionTypeLoraTime),
		string(MissionTypeLoraTimePlan),
		string(MissionTypeFooocusWuTime),
		string(MissionTypeFooocusWuTimePlan),
		string(MissionTypeSvdBackTime),
		string(MissionTypeSvdBackTimePlan),
		string(MissionTypeSdJiTime),
		string(MissionTypeSdJiTimePlan),
		string(MissionTypeSdShangJinTime),
		string(MissionTypeSdShangJinTimePlan),
		string(MissionTypeWaitingTime),
		string(MissionTypeWaitingTimePlan),
	}
}

func (obj MissionType) Ptr() *MissionType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
