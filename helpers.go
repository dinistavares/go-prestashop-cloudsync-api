package prestashop

import "fmt"

func ParseBooleanToPointer(value bool) *bool {
	return &value
}

func ParseIntToPointer(value int) *int {
	return &value
}

func (client *Client) getResourceTypeRaw() string {
	return fmt.Sprintf("%s/%s", ResourceTypeRaw, client.config.RestEndpointVersion)
}

func (client *Client) getResourceTypeReporting() string {
	return fmt.Sprintf("%s/%s", ResourceTypeReporting, client.config.RestEndpointVersion)
}

func (client *Client) getResourceTypeSync() string {
	return fmt.Sprintf("%s/%s", ResourceTypeSync, client.config.RestEndpointVersion)
}
