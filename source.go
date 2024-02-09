// Copyright 2024 LangVM Project
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package purebuild

import (
	"os"
	"strings"
)

func CollectSource(wd string, postfixes ...string) ([]string, error) {
	postfixMap := make(map[string]bool, len(postfixes))
	for _, postfix := range postfixes {
		postfixMap[postfix] = true
	}

	fileEntries, err := os.ReadDir(wd)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, f := range fileEntries {
		sections := strings.Split(f.Name(), ".")
		if len(sections) > 1 {
			if postfixMap[strings.Join(sections[1:], ".")] {
				files = append(files, f.Name())
			}
		}
	}

	return files, nil
}
