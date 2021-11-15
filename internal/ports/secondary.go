package ports

type RestPort interface {

	// auth
	Authenticate(pk string, sk string) error
}

type CredentialStorePort interface {
	SetOctyCredentials(pk string, sk string) error
	GetOctyCredentials() (string, error)
}
