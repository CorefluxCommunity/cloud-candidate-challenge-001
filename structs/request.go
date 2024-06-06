package structs

type CreateDropletRequest struct {
	Name   string `json:"droplet_name"`
	Region string `json:"region"`
	Size   string `json:"size"`
	Image  string `json:"image"`
}
