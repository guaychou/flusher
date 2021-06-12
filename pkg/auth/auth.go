package auth

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/vault/api"
)

// GetAuthCredential  for auth
func GetAuthCredential() (string, string) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	//secret:= &Data{}
	data, err := client.Logical().Read("secret/config/flusher")
	if err != nil {
		log.Fatal(err)
	}
	if data == nil {
		log.Fatal(err)
	}
	username := fmt.Sprintf("%s", data.Data["username"])
	password := fmt.Sprintf("%s", data.Data["password"])
	return username, password
}
