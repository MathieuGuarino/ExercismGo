package hamming

import "errors"


func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of two strings is different")
	}
	var distance int = 0
	for index:=0; index<len(a); index++ {               
        if a[index]!=b[index] { distance++ }
    }                   
	return distance, nil
}

 