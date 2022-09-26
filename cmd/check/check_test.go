/*
 * Copyright 2022 Zeppelin Bend Pty Ltd
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"testing"
)

func TestIsValidLicense(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "Zepben 2023 C-style", args: args{"Zepben-2023-C-style.txt"}, want: true, wantErr: false},
		{name: "Zepben 2023 shell-style", args: args{"Zepben-2023-shell-style.txt"}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidLicense(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidLicense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidLicense() got = %v, want %v", got, tt.want)
			}
		})
	}
}
