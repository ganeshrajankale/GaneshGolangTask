package main_test

import (
	p "tech_m"
	"testing"
)

func TestJsonfilecreation(t *testing.T) {
	apiUri:="https://http.dog/200.json"
	outputFilepath:="./JsonOutput"
	err := p.Jsonfilecreation(apiUri,outputFilepath)
	if err != nil {
		t.Fail()
	}
}

func TestNegativeJsonfilecreation(t *testing.T) {
	apiUri:="https://http.dog/200.json"
	outputFilepath:="./JsonOutput1"
	err := p.Jsonfilecreation(apiUri,outputFilepath)
	if err == nil {
		t.Fail()
	}
}

func TestNegativeuriJsonfilecreation(t *testing.T) {
	apiUri:="https://http.dog/2001.json"
	outputFilepath:="./JsonOutput1"
	err := p.Jsonfilecreation(apiUri,outputFilepath)
	if err == nil {
		t.Fail()
	}
}
