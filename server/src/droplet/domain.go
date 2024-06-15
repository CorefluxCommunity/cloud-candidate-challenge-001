package droplet

type DropletRequest struct {
	Token      string `json:"api_token"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	Size       string `json:"size"`
	Monitoring bool   `json:"monitoring"`
	Ipv6       bool   `json:"ipv6"`
}

func (r DropletRequest) IsValid() bool {
	missingFields := r.Token == "" ||
		r.Image == "" ||
		r.Name == "" ||
		r.Region == "" ||
		r.Size == ""

	return !missingFields
}

type DropletResponse struct {
	Id struct {
		Value string `json:"value"`
	} `json:"droplet_id"`
	Ipv4 struct {
		Value string `json:"value"`
	} `json:"droplet_ip_address"`
}
