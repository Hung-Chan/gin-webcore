package ipsubnetwhitelistings

// IPSubnetWhitelistingModel .
type IPSubnetWhitelistingModel struct {
	Subnet string `json:"subnet" validate:"required" example:"127.0.0.1/20"`
	Remark string `json:"remark"`
	Enable *int   `json:"enable"`
}
