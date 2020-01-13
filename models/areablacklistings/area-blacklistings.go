package areablacklistings

// AreaBlacklistingModel .
type AreaBlacklistingModel struct {
	Country string `json:"country" validate:"required,max=20" example:"TW"`
	Remark  string `json:"remark" validate:"max=20"`
	Enable  *int   `json:"enable"`
}
