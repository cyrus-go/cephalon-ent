package enums

type UserMissionTag string

const (
	UserMissionTagNo     UserMissionTag = "no"     // 该标签不能跳过任何类型任务的验证
	UserMissionTagMining UserMissionTag = "mining" // 该标签表示可以跳过接可能会被用户用于挖矿类型任务的验证逻辑
)

func (UserMissionTag) Values() []string {
	return []string{
		string(UserMissionTagNo),
		string(UserMissionTagMining),
	}
}

func (obj UserMissionTag) Ptr() *UserMissionTag {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
