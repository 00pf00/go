// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go2go

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
	"unicode"
)

// We use Arabic digit zero as a separator.
// Do not use this character in your own identifiers.
const nameSep = '٠'

// We use Vai digit one to introduce a special character code.
// Do not use this character in your own identifiers.
const nameIntro = '꘡'

var nameCodes = map[rune]int{
	' ': 0,
	'*': 1,
	';': 2,
	',': 3,
	'{': 4,
	'}': 5,
	'[': 6,
	']': 7,
	'(': 8,
	')': 9,
}

// instantiatedName returns the name of a newly instantiated function.
func (t *translator) instantiatedName(fnident *ast.Ident, types []types.Type) (string, error) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "_instantiate%c%s", nameSep, fnident.Name)
	for _, typ := range types {
		sb.WriteRune(nameSep)
		s := typ.String()

		// We have to uniquely translate s into a valid Go identifier.
		// This is not possible in general but we assume that
		// identifiers will not contain
		for _, r := range s {
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				sb.WriteRune(r)
			} else {
				code, ok := nameCodes[r]
				if !ok {
					panic(fmt.Sprintf("unexpected type string character %q", r))
				}
				fmt.Fprintf(&sb, "%c%d", nameIntro, code)
			}
		}
	}
	return sb.String(), nil
}
