package ipwhitelisting

// IPWhitelisting .
type IPWhitelisting struct {
	IP     string `json:"ip"`
	Remark string `json:"remark"`
	Enable *int   `json:"enable"`
}
