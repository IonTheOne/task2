package env

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppHost        string        `envconfig:"APP-HOST" required:"true"`
	WriteTimeout   time.Duration `envconfig:"WRITE-TIMEOUT" default:"15s"`
	ReadTimeout    time.Duration `envconfig:"READ-TIMEOUT" default:"15s"`
	ImmuDBAdress   string        `envconfig:"IMMUDB-ADRESS" required:"true"`
	ImmuBDPort     int           `envconfig:"IMMUDB-PORT" required:"true"`
	ImmuDBName     string        `envconfig:"IMMUDB-NAME" required:"true"`
	ImmuDBUser     string        `envconfig:"IMMUDB-USER" required:"true"`
	ImmuDBPassword string        `envconfig:"IMMUDB-PASSWORD" required:"true"`
}

func LoadConfig(cfg interface{}) {
	err := ProcessConfig(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ProcessConfig(cfg interface{}) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}

	t := reflect.ValueOf(cfg).Elem()
	typeOfSpec := t.Type()

	for i := 0; i < t.NumField(); i++ {
		ftype := typeOfSpec.Field(i)
		if ftype.Tag.Get("required") == "true" && t.Field(i).IsZero() {
			return fmt.Errorf("required key %s empty value", ftype.Tag.Get("envconfig"))
		}
	}

	return nil
}
