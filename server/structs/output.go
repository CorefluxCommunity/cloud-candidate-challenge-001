package structs

type TerraformOutput struct {
	DropletIP struct {
		Value string `json:"value"`
	} `json:"droplet_ip"`
}
