package ipwhitelistings

// IPWhitelistingModel .
type IPWhitelistingModel struct {
	IP     string `json:"ip" validate:"required" example:"127.0.0.1"`
	Remark string `json:"remark"`
	Enable *int   `json:"enable"`
}
