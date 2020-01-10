package ipsubnetwhitelistings

// IPSubnetWhitelistingModel .
type IPSubnetWhitelistingModel struct {
	Subnet string `json:"subnet" validate:"required,max=20" example:"127.0.0.1/20"`
	Remark string `json:"remark" validate:"max=20"`
	Enable *int   `json:"enable"`
}
