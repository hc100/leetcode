package main

import (
	"reflect"
	"testing"
)

func TestExample1Success(t *testing.T) {
	tickets := [][]string{[]string{"MUC", "LHR"}, []string{"JFK", "MUC"}, []string{"SFO", "SJC"}, []string{"LHR", "SFO"}}
	output := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	result := findItinerary(tickets)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample2Success(t *testing.T) {
	tickets := [][]string{[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"}}
	output := []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	result := findItinerary(tickets)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample3Success(t *testing.T) {
	tickets := [][]string{[]string{"JFK", "ATL"}, []string{"ATL", "JFK"}}
	output := []string{"JFK", "ATL", "JFK"}
	result := findItinerary(tickets)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample4Success(t *testing.T) {
	tickets := [][]string{[]string{"JFK", "KUL"}, []string{"JFK", "NRT"}, []string{"NRT", "JFK"}}
	output := []string{"JFK", "NRT", "JFK", "KUL"}
	result := findItinerary(tickets)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample5Success(t *testing.T) {
	tickets := [][]string{[]string{"JFK", "SFO"}}
	output := []string{"JFK", "SFO"}
	result := findItinerary(tickets)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample6Success(t *testing.T) {
	tickets := [][]string{[]string{"JFK", "AAA"}, []string{"AAA", "JFK"}, []string{"JFK", "BBB"}, []string{"JFK", "CCC"}, []string{"CCC", "JFK"}}
	output := []string{"JFK", "AAA", "JFK", "CCC", "JFK", "BBB"}
	result := findItinerary(tickets)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}
