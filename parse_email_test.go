package sender

import (
	"os"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	// Create folder and file for running tests.
	_ = os.Mkdir("tmp", 0700)
	file, err := os.OpenFile("tmp/test.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		t.Errorf("Sender.parseTemplate() error = %v", err)
		return
	}
	_, err = file.Write([]byte("<h1>Hello, World!</h1>"))
	if err != nil {
		t.Errorf("Sender.parseTemplate() error = %v", err)
		return
	}
	defer file.Close()

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
				file: "tmp/test.html",
				data: nil,
			},
			want:    "<h1>Hello, World!</h1>",
			wantErr: false,
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
			os.RemoveAll("tmp") // remove temp folder
		})
	}
}
