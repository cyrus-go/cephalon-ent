package enums

type MissionType string

const (
	MissionTypeUnknown            MissionType = "unknown"
	MissionTypeSdTime             MissionType = "sd_time"
	MissionTypeSdProTime          MissionType = "sd_pro_time"
	MissionTypeSdTxt2Img          MissionType = "txt2img"
	MissionTypeSdImg2Img          MissionType = "img2img"
	MissionTypeJpTime             MissionType = "jp_time"
	MissionTypeWtTime             MissionType = "wt_time"
	MissionTypeSdExtraSingleImage MissionType = "extra-single-image"
	MissionTypeSdApi              MissionType = "sd_api"
	MissionTypeKeyPair            MissionType = "key_pair"
	MissionTypeJpDkTime           MissionType = "jp_dk_time"
	MissionTypeSshTime            MissionType = "ssh_time"
	MissionTypeSglangTime         MissionType = "sglang_time"
	MissionTypeSdTomatoTime       MissionType = "sd_tomato_time"
	MissionTypeSdCmdTime          MissionType = "sd_cmd_time"
	MissionTypeSdTianTime         MissionType = "sd_tian_time"
	MissionTypeSdBingoTime        MissionType = "sd_bingo_time"
	MissionTypeFooocusTime        MissionType = "fooocus_time"
	MissionTypeFooocusLanQueTime  MissionType = "fooocus_lan_que_time"
	MissionTypeFooocusShaApiTime  MissionType = "fooocus_sha_api_time"
	MissionTypeTabbyTime          MissionType = "tabby_time"
	MissionTypeOllamaTime         MissionType = "ollama_time"
	MissionTypeJpCondaTime        MissionType = "jp_conda_time"
	MissionTypeJpMlTime           MissionType = "jp_ml_time"
	MissionTypeSdCatTime          MissionType = "sd_cat_time"
	MissionTypeSdFireTime         MissionType = "sd_fire_time"
	MissionTypeComfyuiTime        MissionType = "comfyui_time"
	MissionTypeComfyuiProTime     MissionType = "comfyui_pro_time"
	MissionTypeComfyuiAdvanceTime MissionType = "comfyui_advance_time" // 高级版
	MissionTypeJpDk3Time          MissionType = "jp_dk_3_time"
	MissionTypeSdXlTime           MissionType = "sd_xl_time"
	MissionTypeSdChickTime        MissionType = "sd_chick_time"
	MissionTypeAscendTime         MissionType = "ascend_time"
	MissionTypeSdWuShanTime       MissionType = "sd_wu_shan_time"
	MissionTypeSdLangTime         MissionType = "sd_lang_time"
	MissionTypeSdZhiDaoTime       MissionType = "sd_zhi_dao_time"
	MissionTypeComfyuiKeTime      MissionType = "comfyui_ke_time"
	MissionTypeComfyuiAShuoTime   MissionType = "comfyui_a_shuo_time"
	MissionTypeComfyuiJiaMingTime MissionType = "comfyui_jia_ming_time"
	MissionTypeChatchatTime       MissionType = "chatchat_time"
	MissionTypeChatTTSTime        MissionType = "chat_tts_time"
	MissionTypeLoraTime           MissionType = "lora_time"
	MissionTypeLoraFluxTime       MissionType = "lora_flux_time"
	MissionTypeLoraFluxGymTime    MissionType = "lora_flux_gym_time"
	MissionTypeFooocusWuTime      MissionType = "fooocus_wu_time"
	MissionTypeSvdBackTime        MissionType = "svd_back_time"
	MissionTypeSdJiTime           MissionType = "sd_ji_time"
	MissionTypeSdShangJinTime     MissionType = "sd_shang_jin_time"
	MissionTypeSdXiaoChunTime     MissionType = "sd_xiao_chun_time"
	MissionTypeComfyuiWuTime      MissionType = "comfyui_wu_time"
	MissionTypeComfyuiLiuTime     MissionType = "comfyui_liu_time"
	MissionTypeComfyuiLiTime      MissionType = "comfyui_li_time"
	MissionTypeComfyuiNenlyTime   MissionType = "comfyui_nenly_time"
	MissionTypeComfyuiOuTime      MissionType = "comfyui_ou_time"
	MissionTypeComfyuiLuTime      MissionType = "comfyui_lu_time"
	MissionTypeComfyuiJiangTime   MissionType = "comfyui_jiang_time"
	MissionTypeComfyuiStarTime    MissionType = "comfyui_star_time"
	MissionTypeWaitingTime        MissionType = "waiting_time"     // 等待任务
	MissionTypeWaitingAlTime      MissionType = "waiting_al_time"  // 等待任务
	MissionTypeOpenCLTime         MissionType = "opencl_core_time" // 监控任务 不需要显卡
	MissionTypeIOPaintTime        MissionType = "io_paint_time"

	MissionTypeSdTimePlan             MissionType = "sd_time_plan"
	MissionTypeSdProTimePlan          MissionType = "sd_pro_time_plan"
	MissionTypeWtTimePlan             MissionType = "wt_time_plan"
	MissionTypeJpTimePlan             MissionType = "jp_time_plan"
	MissionTypeJpDkTimePlan           MissionType = "jp_dk_time_plan"
	MissionTypeSshTimePlan            MissionType = "ssh_time_plan"
	MissionTypeSglangTimePlan         MissionType = "sglang_time_plan"
	MissionTypeSdTomatoTimePlan       MissionType = "sd_tomato_time_plan"
	MissionTypeSdCmdTimePlan          MissionType = "sd_cmd_time_plan"
	MissionTypeSdTianTimePlan         MissionType = "sd_tian_time_plan"
	MissionTypeSdBingoTimePlan        MissionType = "sd_bingo_time_plan"
	MissionTypeFooocusTimePlan        MissionType = "fooocus_time_plan"
	MissionTypeFooocusLanQueTimePlan  MissionType = "fooocus_lan_que_time_plan"
	MissionTypeFooocusShaApiTimePlan  MissionType = "fooocus_sha_api_time_plan"
	MissionTypeTabbyTimePlan          MissionType = "tabby_time_plan"
	MissionTypeOllamaTimePlan         MissionType = "ollama_time_plan"
	MissionTypeJpCondaTimePlan        MissionType = "jp_conda_time_plan"
	MissionTypeJpMlTimePlan           MissionType = "jp_ml_time_plan"
	MissionTypeSdCatTimePlan          MissionType = "sd_cat_time_plan"
	MissionTypeSdFireTimePlan         MissionType = "sd_fire_time_plan"
	MissionTypeComfyuiTimePlan        MissionType = "comfyui_time_plan"
	MissionTypeComfyuiProTimePlan     MissionType = "comfyui_pro_time_plan"
	MissionTypeComfyuiAdvanceTimePlan MissionType = "comfyui_advance_time_plan"
	MissionTypeJpDk3TimePlan          MissionType = "jp_dk_3_time_plan"
	MissionTypeSdXlTimePlan           MissionType = "sd_xl_time_plan"
	MissionTypeSdChickTimePlan        MissionType = "sd_chick_time_plan"
	MissionTypeAscendTimePlan         MissionType = "ascend_time_plan"
	MissionTypeSdWuShanTimePlan       MissionType = "sd_wu_shan_time_plan"
	MissionTypeSdLangTimePlan         MissionType = "sd_lang_time_plan"
	MissionTypeSdZhiDaoTimePlan       MissionType = "sd_zhi_dao_time_plan"
	MissionTypeComfyuiKeTimePlan      MissionType = "comfyui_ke_time_plan"
	MissionTypeComfyuiAShuoTimePlan   MissionType = "comfyui_a_shuo_time_plan"
	MissionTypeComfyuiJiaMingTimePlan MissionType = "comfyui_jia_ming_time_plan"
	MissionTypeChatchatTimePlan       MissionType = "chatchat_time_plan"
	MissionTypeChatTTSTimePlan        MissionType = "chat_tts_time_plan"
	MissionTypeLoraTimePlan           MissionType = "lora_time_plan"
	MissionTypeLoraFluxTimePlan       MissionType = "lora_flux_time_plan"
	MissionTypeLoraFluxGymTimePlan    MissionType = "lora_flux_gym_time_plan"
	MissionTypeFooocusWuTimePlan      MissionType = "fooocus_wu_time_plan"
	MissionTypeSvdBackTimePlan        MissionType = "svd_back_time_plan"
	MissionTypeSdJiTimePlan           MissionType = "sd_ji_time_plan"
	MissionTypeSdShangJinTimePlan     MissionType = "sd_shang_jin_time_plan"
	MissionTypeSdXiaoChunTimePlan     MissionType = "sd_xiao_chun_time_plan"
	MissionTypeComfyuiWuTimePlan      MissionType = "comfyui_wu_time_plan"
	MissionTypeComfyuiLiuTimePlan     MissionType = "comfyui_liu_time_plan"
	MissionTypeComfyuiLiTimePlan      MissionType = "comfyui_li_time_plan"
	MissionTypeComfyuiNenlyTimePlan   MissionType = "comfyui_nenly_time_plan"
	MissionTypeComfyuiOuTimePlan      MissionType = "comfyui_ou_time_plan"
	MissionTypeComfyuiLuTimePlan      MissionType = "comfyui_lu_time_plan"
	MissionTypeComfyuiJiangTimePlan   MissionType = "comfyui_jiang_time_plan"
	MissionTypeComfyuiStarTimePlan    MissionType = "comfyui_star_time_plan"
	MissionTypeWaitingTimePlan        MissionType = "waiting_time_plan"     // 等待任务
	MissionTypeWaitingAlTimePlan      MissionType = "waiting_al_time_plan"  // 等待任务
	MissionTypeOpenCLTimePlan         MissionType = "opencl_core_time_plan" // 监控任务 不需要显卡
	MissionTypeIOPaintTimePlan        MissionType = "io_paint_time_plan"
)

func (obj MissionType) Values() []string {
	return []string{
		string(MissionTypeUnknown),
		string(MissionTypeSdTime),
		string(MissionTypeSdProTime),
		string(MissionTypeSdTxt2Img),
		string(MissionTypeSdImg2Img),
		string(MissionTypeJpTime),
		string(MissionTypeWtTime),
		string(MissionTypeSdExtraSingleImage),
		string(MissionTypeSdApi),
		string(MissionTypeKeyPair),
		string(MissionTypeJpDkTime),
		string(MissionTypeSshTime),
		string(MissionTypeSglangTime),
		string(MissionTypeSdTimePlan),
		string(MissionTypeSdProTimePlan),
		string(MissionTypeWtTimePlan),
		string(MissionTypeJpTimePlan),
		string(MissionTypeJpDkTimePlan),
		string(MissionTypeSshTimePlan),
		string(MissionTypeSglangTimePlan),
		string(MissionTypeSdTomatoTime),
		string(MissionTypeSdTomatoTimePlan),
		string(MissionTypeSdCmdTime),
		string(MissionTypeSdCmdTimePlan),
		string(MissionTypeSdTianTime),
		string(MissionTypeSdTianTimePlan),
		string(MissionTypeSdBingoTime),
		string(MissionTypeSdBingoTimePlan),
		string(MissionTypeFooocusTime),
		string(MissionTypeFooocusTimePlan),
		string(MissionTypeFooocusLanQueTime),
		string(MissionTypeFooocusLanQueTimePlan),
		string(MissionTypeFooocusShaApiTime),
		string(MissionTypeFooocusShaApiTimePlan),
		string(MissionTypeTabbyTime),
		string(MissionTypeTabbyTimePlan),
		string(MissionTypeOllamaTime),
		string(MissionTypeOllamaTimePlan),
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
		string(MissionTypeComfyuiProTime),
		string(MissionTypeComfyuiProTimePlan),
		string(MissionTypeComfyuiAdvanceTime),
		string(MissionTypeComfyuiAdvanceTimePlan),
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
		string(MissionTypeSdZhiDaoTime),
		string(MissionTypeSdZhiDaoTimePlan),
		string(MissionTypeComfyuiKeTime),
		string(MissionTypeComfyuiKeTimePlan),
		string(MissionTypeComfyuiAShuoTime),
		string(MissionTypeComfyuiAShuoTimePlan),
		string(MissionTypeComfyuiJiaMingTime),
		string(MissionTypeComfyuiJiaMingTimePlan),
		string(MissionTypeChatchatTime),
		string(MissionTypeChatchatTimePlan),
		string(MissionTypeChatTTSTime),
		string(MissionTypeChatTTSTimePlan),
		string(MissionTypeLoraTime),
		string(MissionTypeLoraTimePlan),
		string(MissionTypeLoraFluxTime),
		string(MissionTypeLoraFluxTimePlan),
		string(MissionTypeLoraFluxGymTime),
		string(MissionTypeLoraFluxGymTimePlan),
		string(MissionTypeFooocusWuTime),
		string(MissionTypeFooocusWuTimePlan),
		string(MissionTypeSvdBackTime),
		string(MissionTypeSvdBackTimePlan),
		string(MissionTypeSdJiTime),
		string(MissionTypeSdJiTimePlan),
		string(MissionTypeSdShangJinTime),
		string(MissionTypeSdShangJinTimePlan),
		string(MissionTypeSdXiaoChunTime),
		string(MissionTypeSdXiaoChunTimePlan),
		string(MissionTypeComfyuiWuTime),
		string(MissionTypeComfyuiWuTimePlan),
		string(MissionTypeComfyuiLiuTime),
		string(MissionTypeComfyuiLiuTimePlan),
		string(MissionTypeComfyuiLiTime),
		string(MissionTypeComfyuiLiTimePlan),
		string(MissionTypeComfyuiNenlyTime),
		string(MissionTypeComfyuiNenlyTimePlan),
		string(MissionTypeComfyuiOuTime),
		string(MissionTypeComfyuiOuTimePlan),
		string(MissionTypeComfyuiLuTime),
		string(MissionTypeComfyuiLuTimePlan),
		string(MissionTypeComfyuiJiangTime),
		string(MissionTypeComfyuiJiangTimePlan),
		string(MissionTypeComfyuiStarTime),
		string(MissionTypeComfyuiStarTimePlan),
		string(MissionTypeWaitingTime),
		string(MissionTypeWaitingTimePlan),
		string(MissionTypeWaitingAlTime),
		string(MissionTypeWaitingAlTimePlan),
		string(MissionTypeOpenCLTime),
		string(MissionTypeOpenCLTimePlan),
		string(MissionTypeIOPaintTime),
		string(MissionTypeIOPaintTimePlan),
	}
}

func (obj MissionType) Ptr() *MissionType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
