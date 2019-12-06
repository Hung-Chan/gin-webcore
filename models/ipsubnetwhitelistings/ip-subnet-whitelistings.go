package ipsubnetwhitelistings

// IPSubnetWhitelisting .
type IPSubnetWhitelisting struct {
	Subnet string `json:"subnet"`
	Remark string `json:"remark"`
	Enable *int   `json:"enable"`
}
