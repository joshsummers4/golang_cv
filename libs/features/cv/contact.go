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
	var contact ContactInfo
	endpoint := "contactinfo"
	req, err := http.NewRequest("GET", config.APIURL+endpoint, nil)
	if err != nil {
		return &contact, err
	}

	req.SetBasicAuth(config.APIUSERNAME, config.APIPASSWORD)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &contact, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &contact, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &contact, err
	}

	err = json.Unmarshal(body, &contact)
	if err != nil {
		return &contact, err
	}
	fmt.Println("contact", contact)
	return &contact, nil
}
