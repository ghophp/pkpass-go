package pkpass

import (
	"io"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	cert, err := os.Open("pass.p12")
	if err != nil {
		t.Fatal(err)
	}
	defer cert.Close()

	r, err := New("DroneEID.pass", "", cert)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("DroneEID.pkpass")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, r)
}
