package cv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ContactInfo struct {
	Email    string
	Address  string
	LinkedIn string
	GitHub   string
}

func GetContactInfo() (*ContactInfo, error) {
	endpoint := "contactinfo"
	req, err := http.NewRequest("GET", requestPath+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(user, pass)

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
