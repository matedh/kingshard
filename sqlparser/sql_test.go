// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package sqlparser

import (
	"strings"
	"testing"
)

func testParse(t *testing.T, sql string) {
	_, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSet(t *testing.T) {
	sql := "set names gbk"
	testParse(t, sql)
}

func TestSimpleSelect(t *testing.T) {
	sql := "select last_insert_id() as a"
	testParse(t, sql)
}

func TestMixer(t *testing.T) {
	sql := `admin upnode("node1", "master", "127.0.0.1")`
	testParse(t, sql)

	sql = "show databases"
	testParse(t, sql)

	sql = "show tables from abc"
	testParse(t, sql)

	sql = "show tables from abc like a"
	testParse(t, sql)

	sql = "show tables from abc where a = 1"
	testParse(t, sql)

	sql = "show proxy abc"
	testParse(t, sql)
}

func TestDefaultValueExpressionParsing(t *testing.T) {
	cases := []string{
		"insert into t (a, b) values (default, ?)",
		"update t set col = default where id = 1",
		"set names default",
	}

	for _, sql := range cases {
		testParse(t, sql)
	}
}

func TestInsertDefaultRoundTrip(t *testing.T) {
	sql := "insert into t (a, b) values (default, ?)"
	stmt, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

	normalized := String(stmt)
	if !strings.Contains(strings.ToLower(normalized), "default") {
		t.Fatalf("expected formatted SQL to contain default, got %s", normalized)
	}
}
