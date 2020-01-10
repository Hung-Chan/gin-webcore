package ipwhitelistings

// IPWhitelistingModel .
type IPWhitelistingModel struct {
	IP     string `json:"ip" validate:"required,max=20" example:"127.0.0.1"`
	Remark string `json:"remark" validate:"max=20"`
	Enable *int   `json:"enable"`
}
