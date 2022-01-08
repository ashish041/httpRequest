package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func TestHttpRequest_1(t *testing.T) {
	domains := []string{"https://google.com", "https://yahoo.com", "http://facebook.com"}
	for i := 0; i < len(domains); i++ {
		got, err := httpRequest(domains[i])
		if err != nil {
			t.Error(
				"For", domains[i],
				"expected", "Non empty string",
				"got", err,
			)
		}
		if got == "" {
			t.Error(
				"For", domains[i],
				"expected", "Non empty string",
				"got", got,
			)
		}
	}
}

func TestHttpRequest_2(t *testing.T) {
	domains := []string{"https://google_com", "https://yahoo_com", "http://facebook_com"}
	for i := 0; i < len(domains); i++ {
		got, err := httpRequest(domains[i])
		if err == nil {
			t.Error(
				"For", domains[i],
				"expected", "Error",
				"got", err,
			)
		}
		if got != "" {
			t.Error(
				"For", domains[i],
				"expected", "Empty string",
				"got", got,
			)
		}
	}
}

func TestDistribute_1(t *testing.T) {
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	_for := 1
	expected_1 := 8
	expected_2 := [][]string{[]string{"a"},
		[]string{"b"},
		[]string{"c"},
		[]string{"d"},
		[]string{"e"},
		[]string{"f"},
		[]string{"g"},
		[]string{"h"},
	}
	_, got := distribute(str, _for)
	if !reflect.DeepEqual(len(got), expected_1) {
		t.Error(
			"For", _for,
			"expected", expected_1,
			"got", got,
		)
	}
	if !reflect.DeepEqual(got, expected_2) {
		t.Error(
			"For", _for,
			"expected", expected_2,
			"got", got,
		)
	}
}

func TestDistribute_2(t *testing.T) {
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	_for := 2
	expected_1 := 4
	expected_2 := [][]string{
		[]string{"a", "b"},
		[]string{"c", "d"},
		[]string{"e", "f"},
		[]string{"g", "h"},
	}

	_, got := distribute(str, _for)
	if !reflect.DeepEqual(len(got), expected_1) {
		t.Error(
			"For", _for,
			"expected", expected_1,
			"got", got,
		)
	}
	if !reflect.DeepEqual(got, expected_2) {
		t.Error(
			"For", _for,
			"expected", expected_2,
			"got", got,
		)
	}
}

func TestDistribute_3(t *testing.T) {
	str := []string{"a", "b", "c"}
	_for := 4
	expected_1 := 3
	expected_2 := [][]string{
		[]string{"a"},
		[]string{"b"},
		[]string{"c"},
	}

	_, got := distribute(str, _for)
	if !reflect.DeepEqual(len(got), expected_1) {
		t.Error(
			"For", _for,
			"expected", expected_1,
			"got", got,
		)
	}
	if !reflect.DeepEqual(got, expected_2) {
		t.Error(
			"For", _for,
			"expected", expected_2,
			"got", got,
		)
	}
}

func TestGetMD5Hash(t *testing.T) {
	_for := []string{
		"BpLnfgDsc2",
		"WD8F2qNfHK",
		"a84jjJkwzD",
		"9h2fhfUVuS",
		"8uVbhV3vC5",
	}
	expected := []string{
		"7334055ceb42aeb0357a003a15d3ab45",
		"e8e23b19657d03cd8fcf8e4d6451c315",
		"64fda72830d13960bc13a989ae5d3d0c",
		"9bf2a5e17c8e3e85c4ed69ee3fdcd5a0",
		"d01511379a5ab059f1d9b931acbbe7cf",
	}
	for i := 0; i < 5; i++ {
		got := getMD5Hash(_for[i])
		if !reflect.DeepEqual(got, expected[i]) {
			t.Error(
				"For", _for[i],
				"expected", expected[i],
				"got", got,
			)
		}
	}
}

func TestPrefixUrl(t *testing.T) {
	_for := []string{
		"google.com",
		"https://google.com",
		"http://www.google.com",
		"twitter.com",
		"facebook.com",
	}
	expected := []string{
		"http://google.com",
		"https://google.com",
		"http://www.google.com",
		"http://twitter.com",
		"http://facebook.com",
	}
	for i := 0; i < 5; i++ {
		got := prefixUrl(_for[i])
		if !reflect.DeepEqual(got, expected[i]) {
			t.Error(
				"For", _for[i],
				"expected", expected[i],
				"got", got,
			)
		}
	}
}
