package config

import (
	"fmt"
	"os"
)

func GetDatabasePath() string {
	projectId := os.Getenv("PROJECT_ID")
	instancesName := os.Getenv("INSTANCE_NAME")
	databaseName := os.Getenv("DATABASE_NAME")

	return fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instancesName, databaseName)
}

func GetTableName() string {
	return os.Getenv("TABLE_NAME")
}
