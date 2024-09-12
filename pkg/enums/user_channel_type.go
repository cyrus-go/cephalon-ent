package enums

type UserChannelType string

const (
	UserChannelTypeNoChannel     UserChannelType = "no_channel"
	UserChannelTypeNormalChannel UserChannelType = "normal_channel"
)

func (obj UserChannelType) Values() []string {
	return []string{
		string(UserChannelTypeNoChannel),
		string(UserChannelTypeNormalChannel),
	}
}

func (obj UserChannelType) Ptr() *UserChannelType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
