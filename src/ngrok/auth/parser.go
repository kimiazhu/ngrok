// Copyright 2019 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: parser
// Author: ZHU HAIHUA <kimiazhu@gmail.com>
// Since: 2019-06-19 14:31
package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	Path    string
	Content string
	Tokens  map[string]string
}

func NewParser(path string) *Parser {
	return &Parser{
		Path:   path,
		Tokens: make(map[string]string),
	}
}

func (this *Parser) Parse() error {
	file, err := os.Open(this.Path)
	if err != nil {
		err = fmt.Errorf("Failed to read configuration file %s: %v", this.Path, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		toParse := strings.Split(text, "#")[0]
		fields := strings.Fields(toParse)
		if len(fields) != 2 {
			continue
		}
		this.Tokens[fields[0]] = fields[1]
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
