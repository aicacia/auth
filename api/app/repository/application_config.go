package repository

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type ApplicationConfigRowST struct {
	ApplicationId int32     `db:"application_id"`
	Key           string    `db:"key"`
	Value         string    `db:"value"`
	UpdatedAt     time.Time `db:"updated_at"`
	CreatedAt     time.Time `db:"created_at"`
}

func GetApplicationConfigs(applicationId int32) ([]ApplicationConfigRowST, error) {
	return All[ApplicationConfigRowST]("SELECT * FROM application_configs WHERE application_id = $1;", applicationId)
}

func ApplicationConfigRowsToMap(rows []ApplicationConfigRowST) map[string]interface{} {
	result := make(map[string]interface{})
	for _, row := range rows {
		entry := result
		parts := strings.Split(row.Key, ".")
		for _, part := range parts[:len(parts)-1] {
			if _, ok := entry[part]; !ok {
				entry[part] = make(map[string]interface{})
			}
			entry = entry[part].(map[string]interface{})
		}
		var value interface{}
		err := json.Unmarshal([]byte(row.Value), &value)
		if err != nil {
			log.Printf("invalid json %d-%s: %v", row.ApplicationId, row.Key, err)
			continue
		}
		entry[parts[len(parts)-1]] = value
	}
	return result
}
