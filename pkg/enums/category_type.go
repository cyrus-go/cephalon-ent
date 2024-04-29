package enums

type CategoryType string

const (
	CategoryTypeUnknown CategoryType = "unknown"
	CategoryTypeHot     CategoryType = "hot"
)

func (CategoryType) Values() []string {
	return []string{
		string(CategoryTypeUnknown),
		string(CategoryTypeHot),
	}
}

func (obj CategoryType) Ptr() *CategoryType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
