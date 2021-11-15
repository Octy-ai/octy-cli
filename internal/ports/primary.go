package ports

type APIPort interface {

	// auth
	SetOctyCredentials(pk string, sk string) error
}
