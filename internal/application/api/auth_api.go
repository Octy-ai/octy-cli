package api

//
// Public methods
//

func (api Application) SetOctyCredentials(pk string, sk string) error {

	// verify credentials
	err := api.rest.Authenticate(pk, sk)
	if err != nil {
		return err
	}

	// set credentials
	err = api.cs.SetOctyCredentials(pk, sk)
	if err != nil {
		return err
	}

	return nil
}
