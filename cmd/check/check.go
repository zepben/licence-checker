/*
 * Copyright 2022 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * Licence, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"github.com/google/licencecheck"
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

func IsValidLicence(filepath string) (bool, error) {
	licence, err := os.ReadFile(filepath)
	check(err)

	var acceptedLicences []licencecheck.Licence

	for _, l := range licencecheck.BuiltinLicences() {
		if (l.URL == "") && (l.Name == "AGPL-Header" || l.Name == "AGPL-v3.0" || l.Name == "MPL-2.0" || l.Name == "MPL-2.0-Header" || l.Name == "MIT") {
			acceptedLicences = append(acceptedLicences, l)
		}
	}
	acceptedLicences = append(acceptedLicences, licencecheck.Licence{Name: "Zepben", Text: ZepbenLicence})
	checker := licencecheck.New(acceptedLicences)

	_, succ := checker.Cover(licence, licencecheck.Options{MinLength: 10, Threshold: RequiredMatchPercentage, Slop: 8})
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
	valid, err := IsValidLicence(os.Args[1])
	if valid {
		os.Exit(0)
	} else {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
