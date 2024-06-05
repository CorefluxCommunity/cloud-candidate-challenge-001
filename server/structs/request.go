package structs

type TerraformRequest struct {
	Token       string `json:"do_token"`
	Image       string `json:"image"`
	Region      string `json:"region"`
	Size        string `json:"size"`
	DropletName string `json:"droplet_name"`
}
