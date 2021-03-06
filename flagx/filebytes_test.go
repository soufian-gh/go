package flagx_test

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/m-lab/go/rtx"

	"github.com/m-lab/go/flagx"
)

func TestFileBytes(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "okay",
			content: "1234567890abcdef",
		},
		{
			name:    "error-bad-filename",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fname string
			var f *os.File
			var err error

			if !tt.wantErr {
				f, err = ioutil.TempFile("", "-filebytes")
				rtx.Must(err, "Failed to create tempfile")
				defer os.Remove(f.Name())
				f.WriteString(tt.content)
				fname = f.Name()
				f.Close()
			} else {
				fname = "this-is-not-a-file"
			}

			fb := &flagx.FileBytes{}
			if err := fb.Set(fname); (err != nil) != tt.wantErr {
				t.Errorf("FileBytes.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && tt.content != (string)(fb.Get().(flagx.FileBytes)) {
				t.Errorf("FileBytes.Get() want = %q, got %q", tt.content, string(*fb))
			}
			if !tt.wantErr && tt.content != fb.String() {
				t.Errorf("FileBytes.String() want = %q, got %q", tt.content, fb.String())
			}
		})
	}
}

// Successful compilation of this function means that FileBytes implements the
// flag.Getter interface. The function need not be called.
func assertFlagGetter(b flagx.FileBytes) {
	func(in flag.Getter) {}(&b)
}
