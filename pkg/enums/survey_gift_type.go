package enums

type SurveyGiftType string

const (
	SurveyGiftTypeUnknown SurveyGiftType = "unknown"
	SurveyGiftTypeSubmit  SurveyGiftType = "submit"  // 提交就赠送
	SurveyGiftTypeApprove SurveyGiftType = "approve" // 审批通过赠送
)

func (obj SurveyGiftType) Values() []string {
	return []string{
		string(SurveyGiftTypeUnknown),
		string(SurveyGiftTypeSubmit),
		string(SurveyGiftTypeApprove),
	}
}

func (obj SurveyGiftType) Ptr() *SurveyGiftType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
