package credentialstore

import (
	"github.com/zalando/go-keyring"
)

type Adapter struct {
	credentialStore CredentialStore
}

func NewAdapter() (*Adapter, error) {
	return &Adapter{credentialStore: CredentialStore{}}, nil
}

type CredentialStore struct {
}

// Set stores user and pass in the keyring under the defined service
// name.
func (c *CredentialStore) Set(service, user, pass string) error {
	return keyring.Set(service, user, pass)
}

// Get gets a secret from the keyring given a service name and a user.
func (c *CredentialStore) Get(service, user string) (string, error) {
	return keyring.Get(service, user)
}

// Delete deletes a secret, identified by service & user, from the keyring.
func (c *CredentialStore) Delete(service, user string) error {
	return keyring.Delete(service, user)
}
