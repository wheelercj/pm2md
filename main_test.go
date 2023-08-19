// Copyright 2023 Chris Wheeler

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"reflect"
	"testing"
)

func TestParseStatusRanges(t *testing.T) {
	tests := []struct {
		str  string
		want [][]int
	}{
		{"", nil},
		{"200-299", [][]int{{200, 299}}},
		{"200-299,400-499", [][]int{{200, 299}, {400, 499}}},
		{"200-200", [][]int{{200, 200}}},
	}
	for _, test := range tests {
		ans := parseStatusRanges(test.str)
		if !reflect.DeepEqual(ans, test.want) {
			t.Errorf("parseStatusRanges(%q) = %v, want %v", test.str, ans, test.want)
		}
	}
}