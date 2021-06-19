package database

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	envs = [...]string{"DB_HOST", "DB_USER", "DB_NAME", "DB_PASSWORD", "DB_PORT", "DB_SSL_MODE"}
)

func getEnv(env string) (string, error) {
	value := os.Getenv(env)
	if value == "" {
		return "", errors.New(fmt.Sprintf("$%s", env))
	}
	return value, nil
}

func dsn() (string, error) {

	vars := make([]string, 0)
	errs := make([]string, 0)

	for _, e := range envs {
		v, err := getEnv(e)
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			vars = append(vars, v)
		}
	}

	if len(errs) > 0 {
		return "", errors.New(fmt.Sprintf("Missing %s variables", strings.Join(errs, ",")))
	}

	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", vars[0], vars[1], vars[2], vars[3], vars[4], vars[5]), nil
}

func New(name string) (*gorm.DB, error) {

	s, err := dsn()

	if err != nil {
		return nil, err
	}

	var database *gorm.DB

	for i := 0; i < 10; i++ {
		database, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: name,
			DSN:        s,
		}))
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, err
	}

	return database, nil
}
