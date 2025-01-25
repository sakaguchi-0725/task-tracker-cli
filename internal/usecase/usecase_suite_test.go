package usecase_test

import (
	"io"
	"os"
)

func captureOutput(fn func()) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	out, _ := io.ReadAll(r)

	return string(out)
}
