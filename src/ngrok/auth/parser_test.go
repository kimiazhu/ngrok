// Copyright 2019 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: parser_test
// Author: ZHU HAIHUA <kimiazhu@gmail.com>
// Since: 2019-06-19 14:32
package auth

import (
	"testing"
)

func TestParser(t *testing.T) {
	path := "./ngrok-secrets"
	parser := NewParser(path)
	if err := parser.Parse(); err != nil {
		t.Fatal(err)
	}
	t.Log(parser.Tokens)
}
