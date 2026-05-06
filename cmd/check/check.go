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

func IsValidLicence(filepath string, licence_type string) (bool, error) {
	licence, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	var acceptedLicences []licensecheck.License
	// Special exclude licence
	acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "IgnoreLicence", Text: ZepbenExcludeLicence})

	// TODO: if the number of licences gets large, implement some looping maybe, or array-append
	// Ignore the following licences
	acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "MicrosoftReciprocalLicence", Text: MicrosoftReciprocalLicence})

	switch licence_type {
	case "private":
		acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "ZepbenPrivate", Text: ZepbenPrivateLicence})
	case "public":
		acceptedLicences = append(acceptedLicences, licensecheck.License{Name: "ZepbenPublic", Text: ZepbenPublicLicence})
	default:
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
var version string

func main() {

	var files []string
	var err_files []string
	var err error


	fmt.Printf("Licence-check version: %s\n", version)

	if isFolder(os.Args[1]) {
		files, err = FindFiles(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}

		for _, path := range files {
			valid, err := IsValidLicence(path, os.Args[2])
			if valid {
				continue
			} else {
				fmt.Println(err.Error())
				err_files = append(err_files, path)
			}
		}
		fmt.Printf("Scanned %d files\n", len(files))
		if len(err_files) != 0 {
			os.Exit(-1)
		}
	} else {
		valid, err := IsValidLicence(os.Args[1], os.Args[2])
		if valid {
			os.Exit(0)
		} else {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}
}
