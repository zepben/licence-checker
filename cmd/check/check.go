/*
 * Copyright 2022 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * Licence, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"os"

	"github.com/google/licensecheck"
)

const RequiredMatchPercentage = 88

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsValidLicence(filepath string, licence_type string) (bool, error) {
	licence, err := os.ReadFile(filepath)
    
	check(err)

	var acceptedLicences []licensecheck.License
    // Special exclude licence
    acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "IgnoreLicence", Text: ZepbenExcludeLicence})

	if licence_type == "private" {
		acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "ZepbenPrivate", Text: ZepbenPrivateLicence})
	} else if licence_type == "public" {
		acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "ZepbenPublic", Text: ZepbenPublicLicence})
	} else {
		return false, fmt.Errorf("Wrong licence type '%s'! Use 'private' or 'public'", licence_type)
	}

	checker := licensecheck.New(acceptedLicences)

	_, matches := checker.Cover(licence, licensecheck.Options{MinLength: 2, Threshold: RequiredMatchPercentage, Slop: 8})
	if matches {
		return true, nil
	} else {
		return false, fmt.Errorf("Licence check failed with a <%d%% match for %s. Ensure the %s licence is present and correct in the file.", RequiredMatchPercentage, filepath, licence_type)
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

	valid, err := IsValidLicence(os.Args[1], os.Args[2])
	if valid {
		os.Exit(0)
	} else {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
