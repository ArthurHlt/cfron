package models

type CFContext struct {
	Platform         string `json:"platform"`
	OrganizationGUID string `json:"organization_guid"`
	OrganizationName string `json:"organization_name"`
	SpaceGUID        string `json:"space_guid"`
	SpaceName        string `json:"space_name"`
	InstanceName     string `json:"instance_name"`
}

type CFPlatform struct {
	UserId string `json:"user_id"`
}
