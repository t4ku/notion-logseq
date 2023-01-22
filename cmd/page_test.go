package cmd

import (
	"bytes"
	"os"
	"testing"
)

var originalArgs = os.Args

func setArgs(args []string) {
	os.Args = append(originalArgs, args...)
}

func Test_pageCmd(t *testing.T) {
	setArgs([]string{"page", "64a270aa-031e-4c52-8e8f-01b588dbd413"})

	got := PickStdout(t, func() { pageCmd.Execute() })
	// want := "sub called"
	// if got != want {
	// 	t.Errorf("subCmd.Execute() = %v, want = %v", got, want)
	// }
	if len(got) < 1 {
		t.Errorf("pageCmd.Execute() does not return anything")
	}
}

func PickStdout(t *testing.T, fnc func()) string {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("fail pipe: %v", err)
	}
	os.Stdout = w
	fnc()
	w.Close()
	var buffer bytes.Buffer
	if n, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v - number: %v", err, n)
	}
	s := buffer.String()
	return s
	// return s[:len(s)-1]
}
