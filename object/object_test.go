package object

import "testing"

func TestStringHashKey(t *testing.T) {
	helloWorld1 := &String{Value: "Hello World"}
	helloWorld2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is sagnikc"}
	diff2 := &String{Value: "My name is sangikc"}
	if helloWorld1.HashKey() != helloWorld2.HashKey() ||
		diff1.HashKey() != diff2.HashKey() ||
		helloWorld1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
}
