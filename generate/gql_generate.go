package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

type BSONMapping struct {
	Model      string `json:"model"`
	Field      string `json:"field"`
	TagPostfix string `json:"tagPostfix"`
}

var mappings []BSONMapping
var mappingFile string

func main() {

	flag.StringVar(&mappingFile, "mappingFile", "", "Path to the file that contains the mappings to apply")
	flag.Parse()

	loadMappingsFile()

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Fatal("Failed to load config", err.Error())
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		log.Fatal("Could not execute gqlgen based code generation correctly", err.Error())
	}
}

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			for _, mapping := range mappings {
				if (mapping.Field == "*" || mapping.Field == field.Name) && (mapping.Model == "*" || mapping.Model == model.Name) {
					field.Tag += mapping.TagPostfix
				}
			}
		}
	}
	return b
}

func loadMappingsFile() {
	jsonData, err := ioutil.ReadFile(mappingFile)
	if os.IsNotExist(err) {
		log.Fatal("Error, could not find mappings file")
	} else if err != nil {
		log.Fatal("Error reading mappings file: ", err)
	}
	err = json.Unmarshal(jsonData, &mappings)
	if err != nil {
		log.Fatal("Error parsing mappings: ", err)
	}
}
