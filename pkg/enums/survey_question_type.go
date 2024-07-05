package enums

type SurveyQuestionType string

const (
	SurveyQuestionTypeUnknown  SurveyQuestionType = "unknown"
	SurveyQuestionTypeText     SurveyQuestionType = "text"     // 文本问题
	SurveyQuestionTypeSingle   SurveyQuestionType = "single"   // 单选问题
	SurveyQuestionTypeMultiple SurveyQuestionType = "multiple" // 多选问题
)

func (obj SurveyQuestionType) Values() []string {
	return []string{
		string(SurveyQuestionTypeUnknown),
		string(SurveyQuestionTypeText),
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
