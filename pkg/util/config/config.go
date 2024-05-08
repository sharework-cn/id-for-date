/*
Copyright 2021 the Velero contributors.

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
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	ConfigKeyFeatures  = "features"
	ConfigKeyColorized = "colorized"
)

// IdpConfig is a map of strings to interface{} for deserializing Velero client config options.
// The alias is a way to attach type-asserting convenience methods.
type IdpConfig map[string]interface{}

// LoadConfig loads the Velero client configuration file and returns it as a VeleroConfig. If the
// file does not exist, an empty map is returned.
func LoadConfig() (IdpConfig, error) {
	fileName := ConfigFileName()

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		// If the file isn't there, just return an empty map
		return IdpConfig{}, nil
	}
	if err != nil {
		// For any other Stat() error, return it
		return nil, errors.WithStack(err)
	}

	configFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer configFile.Close()

	var config IdpConfig
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}

// SaveConfig saves the passed in config map to the Velero client configuration file.
func SaveConfig(config IdpConfig) error {
	fileName := ConfigFileName()

	// Try to make the directory in case it doesn't exist
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return errors.WithStack(err)
	}

	configFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return errors.WithStack(err)
	}
	defer configFile.Close()

	return json.NewEncoder(configFile).Encode(&config)
}

func (c IdpConfig) Features() []string {
	val, ok := c[ConfigKeyFeatures]
	if !ok {
		return []string{}
	}

	features, ok := val.(string)
	if !ok {
		return []string{}
	}

	return strings.Split(features, ",")
}

func (c IdpConfig) Colorized() bool {
	val, ok := c[ConfigKeyColorized]
	if !ok {
		return true
	}

	valString, ok := val.(string)
	if !ok {
		return true
	}

	colorized, err := strconv.ParseBool(valString)
	if err != nil {
		return true
	}

	return colorized
}

func ConfigFileName() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "idp", "idp.yaml")
}
