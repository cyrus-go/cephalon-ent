package enums

type SurveyQuestionType string

const (
	SurveyQuestionTypeUnknown  SurveyQuestionType = "unknown"
	SurveyQuestionTypeSingle   SurveyQuestionType = "single"   // 单选问题
	SurveyQuestionTypeMultiple SurveyQuestionType = "multiple" // 多选问题
)

func (obj SurveyQuestionType) Values() []string {
	return []string{
		string(SurveyQuestionTypeUnknown),
		string(SurveyQuestionTypeSingle),
		string(SurveyQuestionTypeMultiple),
	}
}

func (obj SurveyQuestionType) Ptr() *SurveyQuestionType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
