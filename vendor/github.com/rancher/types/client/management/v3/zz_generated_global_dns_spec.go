package client

const (
	GlobalDNSSpecType                   = "globalDnsSpec"
	GlobalDNSSpecFieldFQDN              = "fqdn"
	GlobalDNSSpecFieldMultiClusterAppID = "multiClusterAppId"
	GlobalDNSSpecFieldProjectIDs        = "projectIds"
	GlobalDNSSpecFieldProviderID        = "providerId"
)

type GlobalDNSSpec struct {
	FQDN              string   `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	MultiClusterAppID string   `json:"multiClusterAppId,omitempty" yaml:"multiClusterAppId,omitempty"`
	ProjectIDs        []string `json:"projectIds,omitempty" yaml:"projectIds,omitempty"`
	ProviderID        string   `json:"providerId,omitempty" yaml:"providerId,omitempty"`
}
