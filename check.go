/*
 * Copyright 2020 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"github.com/google/licensecheck"
)
import (
	"fmt"
	"io/ioutil"
	"os"
)

var REQUIRED_MATCH_PERCENTAGE = 88

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Simple program for checking for AGPL licences.
// Takes a single argument: The path of the file to check
// Returns 0 on success and -1 if neither the AGPL licence or header snippet did not meet an 80% match.
// Should be used on either source files with licence headers or COPYING files.
func main() {
	var licences = licensecheck.BuiltinLicenses()
	var gpls []licensecheck.License
	filepath := os.Args[1]

	for _, l := range licences {
		if l.Name == "AGPL-Header" || l.Name == "AGPL-v3.0" || l.Name == "MPL-2.0" || l.Name == "MPL-2.0-Header" || l.Name == "MIT" {
			gpls = append(gpls, l)
		}
	}
	checker := licensecheck.New(gpls)
		file, err := ioutil.ReadFile(filepath)
		check(err)
		_, succ := checker.Cover(file, licensecheck.Options{10, REQUIRED_MATCH_PERCENTAGE, 8})
		if (succ) {
			os.Exit(0)
		} else {
			fmt.Println("Licence check failed with a <", REQUIRED_MATCH_PERCENTAGE, "% match for", filepath, "Ensure the AGPL or MPL licence is present and correct in the file.")
			os.Exit(-1)
		}
}
