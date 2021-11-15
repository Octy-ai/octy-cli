package rest

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Octy-ai/octy-cli/pkg/globals"
)

func (ha Adapter) Authenticate(pk string, sk string) error {

	req, err := http.NewRequest("GET", globals.AuthRoute, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(pk+":"+sk)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 200 {
		return errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}
