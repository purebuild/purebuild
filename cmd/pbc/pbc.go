// Copyright 2024 LangVM Project
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"purebuild"
)

func main() {
	files, err := purebuild.CollectSource(".", "c", "cc", "cxx", "cpp")
	if err != nil {
		fmt.Println(err)
		return
	}
}
