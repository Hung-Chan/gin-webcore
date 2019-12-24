package areablacklistings

// AreaBlacklistingModel .
type AreaBlacklistingModel struct {
	Country string `json:"country" validate:"required" example:"TW"`
	Remark  string `json:"remark"`
	Enable  *int   `json:"enable"`
}
