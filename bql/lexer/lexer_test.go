// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lexer

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	table := []struct {
		input  string
		tokens []Token
	}{
		{"", []Token{
			{Type: ItemEOF}}},
		{"{}()", []Token{
			{Type: ItemLBracket, Text: "{"},
			{Type: ItemRBracket, Text: "}"},
			{Type: ItemLPar, Text: "("},
			{Type: ItemRPar, Text: ")"},
			{Type: ItemEOF}}},
	}

	for _, test := range table {
		_, c := lex(test.input)
		idx := 0
		for got := range c {
			fmt.Printf("%v\n", got)
			if want := test.tokens[idx]; got != want {
				t.Errorf("lex(%q) failed to provide %v, got %v instead", test.input, want, got)
			}
			idx++
			if idx > len(test.tokens) {
				t.Fatalf("lex(%q) has not finished producing tokens when it should have.", test.input)
			}
		}
	}
}
