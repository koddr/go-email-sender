package sender

import (
	"os"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	// Create folder for running tests.
	_ = os.Mkdir("tmp", 0700)

	// Create test file 1.
	file1, _ := os.OpenFile("tmp/test-1.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
	_, _ = file1.Write([]byte("<h1>Hello, World!</h1>"))
	defer file1.Close()

	// Create test file 2.
	file2, _ := os.OpenFile("tmp/test-2.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
	_, _ = file2.Write([]byte("<div>Hello, {{range Wrong}}{{.}}{{end}}!</div>"))
	defer file2.Close()

	// Test cases.
	type args struct {
		file string
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successfully parsing the template file",
			args: args{
				file: "tmp/test-1.html",
				data: nil,
			},
			want:    "<h1>Hello, World!</h1>",
			wantErr: false,
		},
		{
			name: "failed to parse the template file (no file)",
			args: args{
				file: "tmp/file-not-exist.html",
				data: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed to parse the template file (wrong data)",
			args: args{
				file: "tmp/test-2.html",
				data: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTemplate(tt.args.file, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sender.parseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sender.parseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}

	// Remove temp folder.
	_ = os.RemoveAll("tmp")
}
