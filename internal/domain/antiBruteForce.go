package domain

type IncomingRequest struct {
	Login    string `json:"login" validate:"required,max=50"`
	Password string `json:"password" validate:"required,max=50"`
	IP       string `json:"ip" validate:"required,max=50"`
}

type ResponseIsAccess struct {
	IsAccess    bool `json:"isAccess"`
}