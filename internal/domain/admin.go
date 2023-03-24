package domain

const (
	MaxLimitCommon = "MAX_LIMIT_COMMON"
	MaxLimitIp = "MAX_LIMIT_IP"
	MaxLimitLogin = "MAX_LIMIT_LOGIN"
	MaxLimitPassword = "MAX_LIMIT_PASSWORD"
	ExpairBacket = "EXPAIR_BACKET"
)

type ConfigSetting struct {
	MaxLimitCommon int
	MaxLimitIp int
	MaxLimitLogin int
	MaxLimitPassword int
	ExpairBacket int
}

type ResetBucket struct {
	Login    string `json:"login" validate:"required,max=50"`
	IP       string `json:"ip" validate:"required,max=50"`
}

type ListType string

const (
    ListBlack ListType = "black"
    ListWhite ListType = "white"
)

type ListsActions struct {
	TitleList    ListType `json:"titleList" validate:"required,max=50,oneof=black white"`
	IP       string `json:"ip" validate:"required,max=50"`
}

