package mygotest

import (
	"fmt"
	"go/token"
	scanner "std/go/scanner"
	"testing"
)

func TestMyScanner1(t *testing.T) {
	var src = "var s = 0"
	var s = scanner.Scanner{}
	var tf = token.NewFileSet()
	file := tf.AddFile("", tf.Base(), len(src)) // register input "file"
	var mode = scanner.ScanComments
	s.Init(file, []byte(src), nil, mode)
	for {
		pos, t, tStr := s.Scan()
		if t == token.EOF {
			break
		}
		fmt.Println(pos, t, tStr)
	}
}
