/*
Copyright 2022 The Kubernetes Authors.

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

package path

import (
	"fmt"
	"os"
	"path/filepath"
)

// Expand expands absolute directory in file paths.
func Expand(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("empty path")
	}

	if path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if len(path) == 1 {
			path = home
		} else if path[1] == '/' || path[1] == '\\' {
			path = Join(home, path[2:])
		}
	}

	return filepath.Abs(path)
}
