package eventbridge

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MauricioAntonioMartinez/aws/model"
)

type svc struct {
	Alias   string   `json:"alias"`
	Name    string   `json:"name"`
	Methods []string `json:"methods"`
}

func (s *EventBridgeService) Seed() error {
	data, err := os.ReadFile("eventbridge/services.json")

	if err != nil {
		return err
	}
	var services []svc

	if err = json.Unmarshal(data, &services); err != nil {
		return err
	}

	for _, service := range services {
		for _, method := range service.Methods {
			s.db.Create(&model.ServiceEvent{
				Path:   fmt.Sprintf("%s/%s", service.Name, method),
				Name:   service.Alias,
				Method: method,
			})
		}
	}

	return nil
}
