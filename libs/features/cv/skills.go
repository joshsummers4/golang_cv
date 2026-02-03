package cv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/config"
)

type Skills struct {
	Frontend []string `json:"Frontend"`
	Backend  []string `json:"Backend"`
	DevOps   []string `json:"DevOps"`
	Testing  []string `json:"Testing"`
}

func GetSkills(skillsType string) (*Skills, error) {
	var skills Skills
	endpoint := "skills"
	req, err := http.NewRequest("GET", config.APIURL+endpoint, nil)
	if err != nil {
		return &skills, err
	}

	req.Header.Set("type", skillsType)
	req.SetBasicAuth(config.APIUSERNAME, config.APIPASSWORD)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &skills, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &skills, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &skills, err
	}

	err = json.Unmarshal(body, &skills)
	if err != nil {
		return &skills, err
	}

	return &skills, nil
}
