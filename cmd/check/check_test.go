/*
 * Copyright 2022 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * Licence, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"testing"
)

func TestIsValidLicence(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := map[string][]struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		"private": {
			{name: "Zepben Private C-style", args: args{"ZepbenPrivate-C-style.txt"}, want: true, wantErr: false},
			{name: "Zepben Private shell-style", args: args{"ZepbenPrivate-shell-style.txt"}, want: true, wantErr: false},
		},
		"public": {
			{name: "Zepben Public C-style", args: args{"ZepbenMPL-C-style.txt"}, want: true, wantErr: false},
			{name: "Zepben Public shell-style", args: args{"ZepbenMPL-shell-style.txt"}, want: true, wantErr: false},
		}}

	// Private tests
	for _, tt := range tests["private"] {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidLicence(tt.args.filepath, "private")
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidLicence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidLicence() got = %v, want %v", got, tt.want)
			}
		})
	}

	// Public tests
	for _, tt := range tests["public"] {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidLicence(tt.args.filepath, "public")
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidLicence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidLicence() got = %v, want %v", got, tt.want)
			}
		})
	}
}
