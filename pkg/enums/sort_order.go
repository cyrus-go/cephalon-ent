package enums

type SortOrder string

const (
	SortOrderDesc SortOrder = "desc"
	SortOrderAsc  SortOrder = "asc"
)

func (obj SortOrder) Values() []string {
	return []string{
		string(SortOrderAsc),
		string(SortOrderDesc),
	}
}

func (obj SortOrder) Ptr() *SortOrder {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
