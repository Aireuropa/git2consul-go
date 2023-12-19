/*
Copyright 2019 Kohl's Department Stores, Inc.

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
	"fmt"
	"path/filepath"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetHandler(discard.New())
}

func TestLoad(t *testing.T) {
	file := filepath.Join("test-fixtures", "local.json")

	_, err := Load(file)
	assert.NoError(t, err)
}

func TestLoadInvalidConfig(t *testing.T) {
	file := filepath.Join("test-fixtures", "invalid_config.json")

	_, err := Load(file)
	assert.Error(t, err)
}

func TestIsValidGitURL(t *testing.T) {
	// Casos de prueba positivos
	validURLs := []string{
		"https://github.com/usuario/repo.git",
		"git://github.com/usuario/repo.git",
		"https://gitlab.com/usuario/repo.git",
		"git://gitlab.com/usuario/repo.git",
		"git@bitbucket.org:workspace/repo.git",
	}

	for _, url := range validURLs {
		t.Run(fmt.Sprintf("ValidURL: %s", url), func(t *testing.T) {
			if !isValidGitURL(url) {
				t.Errorf("Se esperaba que la URL fuera válida, pero no lo es.")
			}
		})
	}

	// Casos de prueba negativos
	invalidURLs := []string{
		"invalid-url",
		"https://example.com",
		"ftp://github.com/usuario/repo.git",
		"https://bitbucket.org/usuario/repo",
	}

	for _, url := range invalidURLs {
		t.Run(fmt.Sprintf("InvalidURL: %s", url), func(t *testing.T) {
			if isValidGitURL(url) {
				t.Errorf("Se esperaba que la URL fuera inválida, pero es válida.")
			}
		})
	}
}
