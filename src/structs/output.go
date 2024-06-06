package structs

type Droplet struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	PublicIPv4 string `json:"ipv4"`
	Region     string `json:"region"`
	Size       string `json:"size"`
	Image      string `json:"image"`
}

type DropletListOutput struct {
	Value []Droplet `json:"value"`
}

type CreateDropletOutput struct {
	DropletIP struct {
		Value string `json:"value"`
	} `json:"droplet_ip"`
}
