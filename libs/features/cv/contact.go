package cv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/config"
)

type ContactInfo struct {
	Email    string
	Address  string
	LinkedIn string
	GitHub   string
}

func GetContactInfo() (*ContactInfo, error) {
	endpoint := "contactinfo"
	req, err := http.NewRequest("GET", config.APIURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.APIUSERNAME, config.APIPASSWORD)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var contact ContactInfo
	err = json.Unmarshal(body, &contact)
	if err != nil {
		return nil, err
	}
	fmt.Println("contact", contact)
	return &contact, nil
}
