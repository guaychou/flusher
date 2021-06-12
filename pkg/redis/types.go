package rds

// Type that GdnRedis have
type Rds struct {
	// Redis Type we use , either it sentinel / standalone
	Data struct {
		Address          string `json:"address,omitempty"`
		Port             int    `json:"port,omitempty"`
		Master           string `json:"master,omitempty"`
		Database         int    `json:"database,omitempty"`
		Key              string `json:"key,omitempty"`
		Value            string `json:"value,omitempty"`
		Password         string `json:"password,omitempty"`
		SentinelPassword string `json:"sentinel_password,omitempty"`
	} `json:"data,omitempty"`
}
