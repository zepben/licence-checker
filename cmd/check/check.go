/*
 * Copyright 2022 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * Licence, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"github.com/google/licensecheck"
)
import (
	"fmt"
	"os"
)

const RequiredMatchPercentage = 88

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsValidLicence(filepath string, private bool) (bool, error) {
	licence, err := os.ReadFile(filepath)
	check(err)

	var acceptedLicences []licensecheck.License

	if private {
		acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "Zepben", Text: ZepbenLicence})
	} else {
		for _, l := range licensecheck.BuiltinLicenses() {
			if (l.URL == "") && (l.Name == "AGPL-Header" || l.Name == "AGPL-v3.0" || l.Name == "MPL-2.0" || l.Name == "MPL-2.0-Header" || l.Name == "MIT") {
				acceptedLicences = append(acceptedLicences, l)
			}
		}
	}

	checker := licensecheck.New(acceptedLicences)

	_, succ := checker.Cover(licence, licensecheck.Options{MinLength: 10, Threshold: RequiredMatchPercentage, Slop: 8})
	if succ {
		return true, nil
	} else {
		return false, fmt.Errorf("Licence check failed with a <%d%% match for %s. Ensure the AGPL/MPL/MIT/Zepben licence is present and correct in the file.", RequiredMatchPercentage, filepath)
	}
}

// Simple program for checking for accepted licences. The accepted licences are:
//   - AGPL v3
//   - MPL v2
//   - MIT
//   - Zepben (closed source)
//
// Takes a single argument: The path of the file to check
// Returns 0 on success and -1 if either the licence or header snippet did not meet an 80% match.
// Should be used on either source files with licence headers or COPYING files.
func main() {

	var private bool
	if os.Args[2] == "private" {
		private = true
	} else {
		private = false
	}

	valid, err := IsValidLicence(os.Args[1], private)
	if valid {
		os.Exit(0)
	} else {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
