/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

const (
	testKey = "KEY"
)

func TestEnvBool(t *testing.T) {
	cases := []struct {
		value string
		def   bool
		exp   bool
	}{
		{"true", true, true},
		{"true", false, true},
		{"", true, true},
		{"", false, false},
		{"false", true, false},
		{"false", false, false},
		{"", true, true},
		{"", false, false},
		{"no true", true, true},
		{"no false", true, true},
	}

	for _, testCase := range cases {
		os.Setenv(testKey, testCase.value)
		val := envBool(testKey, testCase.def)
		if val != testCase.exp {
			t.Fatalf("expected %v but %v returned", testCase.exp, val)
		}
	}
}

func TestEnvString(t *testing.T) {
	cases := []struct {
		value string
		def   string
		exp   string
	}{
		{"true", "true", "true"},
		{"true", "false", "true"},
		{"", "true", "true"},
		{"", "false", "false"},
		{"false", "true", "false"},
		{"false", "false", "false"},
		{"", "true", "true"},
		{"", "false", "false"},
	}

	for _, testCase := range cases {
		os.Setenv(testKey, testCase.value)
		val := envString(testKey, testCase.def)
		if val != testCase.exp {
			t.Fatalf("expected %v but %v returned", testCase.exp, val)
		}
	}
}

func TestEnvInt(t *testing.T) {
	cases := []struct {
		value string
		def   int
		exp   int
	}{
		{"0", 1, 0},
		{"", 0, 0},
		{"-1", 0, -1},
		{"abcd", 0, 0},
		{"abcd", 1, 1},
	}

	for _, testCase := range cases {
		os.Setenv(testKey, testCase.value)
		val := envInt(testKey, testCase.def)
		if val != testCase.exp {
			t.Fatalf("expected %v but %v returned", testCase.exp, val)
		}
	}
}

func TestEnvFloat(t *testing.T) {
	cases := []struct {
		value string
		def   float64
		exp   float64
	}{
		{"0.5", 0, 0.5},
		{"", 0.5, 0.5},
		{"-0.5", 0, -0.5},
		{"abcd", 0.5, 0.5},
		{"abcd", 1.5, 1.5},
	}

	for _, testCase := range cases {
		os.Setenv(testKey, testCase.value)
		val := envFloat(testKey, testCase.def)
		if val != testCase.exp {
			t.Fatalf("expected %v but %v returned", testCase.exp, val)
		}
	}
}

func TestEnvDuration(t *testing.T) {
	cases := []struct {
		value string
		def   time.Duration
		exp   time.Duration
	}{
		{"1s", 0, time.Second},
		{"", time.Minute, time.Minute},
		{"1h", 0, time.Hour},
		{"abcd", time.Second, time.Second},
		{"abcd", time.Minute, time.Minute},
	}

	for _, testCase := range cases {
		os.Setenv(testKey, testCase.value)
		val := envDuration(testKey, testCase.def)
		if val != testCase.exp {
			t.Fatalf("expected %v but %v returned", testCase.exp, val)
		}
	}
}

func TestManualHasNoTabs(t *testing.T) {
	if strings.Contains(manual, "\t") {
		t.Fatal("the manual text contains a tab")
	}
}
