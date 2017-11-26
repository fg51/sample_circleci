package main


import ("testing")


func TestExampleSuccess(t *testing.T) {
	result, err := is_apple("apple")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != true {
		t.Fatal("failed test")
	}
}


func TestExampleFailed(t *testing.T) {
	result, err := is_apple("banana")
	if err == nil {
		t.Fatal("failed test")
	}
	if result != false {
		t.Fatal("failed test")
	}
}
