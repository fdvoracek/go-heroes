package model

import "time"

type Request struct {
	Domain string `json:"domain"`
}

type SecurityDefinition struct {
	Id string                `json:"id"`
	ImportRunTimeUtcMs int64 `json:"importRunTimeUtcMs"`
	CreatedTimeUtcMs int64   `json:"createdTimeUtcMs"`
	Resources []Resources    `json:"resources"`
}

type Resources struct {
	Resource string   `json:"resource"`
	Sources []Sources `json:"sources"`
}

type Sources struct {
	LastUpdatedTimeUtcMs int64 `json:"lastUpdatedTimeUtcMs"`
	Source string `json:"source"`
	File string `json:"file"`
	Category string `json:"category"`
	ThreatType string `json:"threatType"`
	Result string `json:"result"`
}

func NewSecurityDefinition(hashId string, resource string) SecurityDefinition {
	var now = time.Now().UnixNano()
	securityDefinition := SecurityDefinition{}
	securityDefinition.Id = hashId
	securityDefinition.ImportRunTimeUtcMs = now
	securityDefinition.CreatedTimeUtcMs = now
	securityDefinition.Resources = NewResources(resource, now)
	return securityDefinition
}

func NewResources(resource string, now int64) []Resources {
	resources := Resources{}
	resources.Resource = resource
	resources.Sources = NewSources(now)

	var result = make([]Resources, 1)
	result[0] = resources
	return result
}

func NewSources(now int64) []Sources {
	sources := Sources{}
	sources.LastUpdatedTimeUtcMs = now
	sources.Result = "CLEAN"

	var result = make([]Sources, 1)
	result[0] = sources
	return result
}