package domain

const (
	MaxLimitCommon = "MAX_LIMIT_COMMON"
	MaxLimitIp = "MAX_LIMIT_IP"
	MaxLimitLogin = "MAX_LIMIT_LOGIN"
	MaxLimitPassword = "MAX_LIMIT_PASSWORD"
	ExpairBacket = "EXPAIR_BACKET"
)

type ConfigSetting struct {
	MaxLimitCommon 		int `json:"maxLimitCommon" validate:"max=10"`
	MaxLimitIp 			int `json:"maxLimitIp" validate:"max=10"`
	MaxLimitLogin 		int	`json:"maxLimitLogin" validate:"max=10"`
	MaxLimitPassword 	int	`json:"maxLimitPassword" validate:"max=10"`
	ExpairBacket 		int	`json:"expairBacket" validate:"max=10"`
}

type ResetBucket struct {
	Login    string `json:"login" validate:"max=50"`
	IP       string `json:"ip" validate:"max=50"`
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

