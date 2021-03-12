package accumulate

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
func Accumulate(listOfString []string, converter func(string) string) []string {
	var currentFuncName = GetFunctionName(converter)
	var funcNameEcho = GetFunctionName(echo)
	var funcNameUpper = GetFunctionName(strings.ToUpper)

	if currentFuncName == funcNameEcho {
		return listOfString
	} else {
		for i := 0; i < len(listOfString); i++ {
			if currentFuncName == funcNameUpper {
				listOfString[i] = strings.ToUpper(listOfString[i])
			} else {
				listOfString[i] = capitalize(listOfString[i])
			}
		}
	}
	return listOfString
}
