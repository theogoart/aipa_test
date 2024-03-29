// Copyright 2018~2022 The aipa Authors
// This file is part of the aipa Chain library.
// Created by  Team of aipa.

// This program is free software: you can distribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with aipa.  If not, see <http://www.gnu.org/licenses/>.

// Copyright 2018 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operators

import (
	"testing"
)

func TestNew(t *testing.T) {
	op1, err := New(Unreachable)
	if err != nil {
		t.Fatalf("unexpected error from New: %v", err)
	}
	if op1.Name != "unreachable" {
		t.Fatalf("0x00: unexpected Op name. got=%s, want=unrechable", op1.Name)
	}
	if !op1.IsValid() {
		t.Fatalf("0x00: operator %v is invalid (should be valid)", op1)
	}

	op2, err := New(0xff)
	if err == nil {
		t.Fatalf("0xff: expected error while getting Op value")
	}
	if op2.IsValid() {
		t.Fatalf("0xff: operator %v is valid (should be invalid)", op2)
	}
}
