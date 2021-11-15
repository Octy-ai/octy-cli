package credentialstore

import (
	"encoding/base64"
)

func (ca Adapter) SetOctyCredentials(pk string, sk string) error {
	encToken := base64.StdEncoding.EncodeToString([]byte(pk + ":" + sk))
	return ca.credentialStore.Set("Octy Credentials", "Octy", encToken)

}
func (ca Adapter) GetOctyCredentials() (string, error) {
	return ca.credentialStore.Get("Octy Credentials", "Octy")
}
