package enums

type SurveyResponseStatus string

const (
	SurveyResponseStatusPending SurveyResponseStatus = "pending" // 等待
	SurveyResponseStatusPass    SurveyResponseStatus = "pass"    // 通过
	SurveyResponseStatusReject  SurveyResponseStatus = "reject"  // 拒绝
)

func (obj SurveyResponseStatus) Values() []string {
	return []string{
		string(SurveyResponseStatusPending),
		string(SurveyResponseStatusPass),
		string(SurveyResponseStatusReject),
	}
}

func (obj SurveyResponseStatus) Ptr() *SurveyResponseStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
