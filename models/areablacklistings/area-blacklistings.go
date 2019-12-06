package areablacklistings

// AreaBlacklisting .
type AreaBlacklisting struct {
	Country string `json:"country"`
	Remark  string `json:"remark"`
	Enable  *int   `json:"enable"`
}
