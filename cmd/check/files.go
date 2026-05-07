package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

/*
 * Copyright 2022 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * Licence, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

func isFolder(path string) bool {
	// Get file info for the given path
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Path '%s' does not exist.\n", path)
		} else {
			fmt.Printf("Error checking path '%s': %v\n", path, err)
		}
		return false
	}

	// Check if the path is a directory
	if info.IsDir() {
		return true
	} else {
		return false
	}
}

func FindFiles(path string) ([]string, error) {

	// Count the files and fill a list
	var files []string
	err := filepath.Walk(path, func(npath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Can't read files from %s", path)
		}

		matched, _ := regexp.MatchString("(.*\\.py|.*\\.java|.*\\.kt|.*\\.cs|.*\\.proto|.*\\.js|.*\\.ts)$", info.Name())
		if matched && !info.IsDir() {
			files = append(files, npath)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
