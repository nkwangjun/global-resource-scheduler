/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type controllerTypes struct {
	Types []controllerType `json:"controllers"`
}

type controllerType struct {
	Type    string `json:"type"`
	Workers int32  `json:"workers"`
}

// LocalConfig is the config to load all the local configurations
type LocalConfig interface {
	Get() (map[string]interface{}, error)
	GetByKey(key string) (interface{}, error)
}

// ControllerConfig is the config to load controller configurations
type ControllerConfig struct {
	typemap map[string]int32
}

// NewControllerConfig to load configuration from a local file
func NewControllerConfig(filePath string) (ControllerConfig, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return ControllerConfig{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return ControllerConfig{}, err
	}

	var types controllerTypes

	json.Unmarshal([]byte(byteValue), &types)

	var controllerMap map[string]int32
	controllerMap = make(map[string]int32)
	for _, controllerType := range types.Types {
		controllerMap[controllerType.Type] = controllerType.Workers
	}
	return ControllerConfig{typemap: controllerMap}, nil
}