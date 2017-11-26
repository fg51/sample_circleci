package main


import ("errors")


func is_apple(fruit string) (bool, error) {
	if fruit == "apple" {
		return true, nil
	}
	return false, errors.New("fruit must be apple")
}
