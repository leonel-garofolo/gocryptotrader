package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"gopkg.in/yaml.v3"
)

// An example showing how to unmarshal embedded
// structs from YAML.

type DataTest struct {
	Apikey              string `yaml:"apikey"`
	ApiAccountPlanLevel string `yaml:"apiAccountPlanLevel"`
}

func TestYaml(t *testing.T) {
	var b DataTest

	buf, err := ioutil.ReadFile("enviroment_test.yml")
	if err != nil {
		fmt.Println(err)
	}

	err2 := yaml.Unmarshal(buf, &b)
	if err2 != nil {
		log.Fatalf("cannot unmarshal data: %v", err2)
	}
	fmt.Println(b.Apikey)
	fmt.Println(b.ApiAccountPlanLevel)
}
