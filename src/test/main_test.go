package main

import (
	"os"
	"testing"
)

func Test_downloadFile(t *testing.T) {
	type args struct {
		url      string
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				url:      "https://www.baidu.com",
				filepath: "baidu.html",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				url:      "https://www.baidu.com",
				filepath: "baidu.html",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadFile(tt.args.url, tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("downloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unzip(t *testing.T) {
	type args struct {
		uri  string
		src  string
		dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Unzip Valid File",
			args: args{
				uri:  "https://registry.faa.gov/database/ReleasableAircraft.zip",
				src:  "ReleasableAircraft.zip",
				dest: "ReleasableAircraft",
			},
			wantErr: false,
		},
		{
			name: "Test Unzip Invalid File",
			args: args{
				uri:  "https://www.baidu.com",
				src:  "baidu.html",
				dest: "test",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock download file
			downloadFile(tt.args.uri, tt.args.src)
			if err := unzip(tt.args.src, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("unzip() error = %v, wantErr %v", err, tt.wantErr)
			}
			// clean up
			os.RemoveAll(tt.args.dest)
		})
	}
}

func Test_safeSubstring(t *testing.T) {
	type args struct {
		s     string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Numeric String",
			args: args{
				s:     "123456789",
				start: 1,
				end:   4,
			},
			want: "234",
		},
		{
			name: "Numeric String with Letters",
			args: args{
				s:     "123456789xcv",
				start: 4,
				end:   9,
			},
			want: "56789",
		},
		{
			name: "Numeric String with Letters and Comma",
			args: args{
				s:     ",123456789xcv,",
				start: 1,
				end:   15,
			},
			want: "123456789xcv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeSubstring(tt.args.s, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("safeSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertIntoEdgeDB(t *testing.T) {
	type args struct {
		records []NNumberRecord
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertIntoEdgeDB(tt.args.records)
		})
	}
}

func TestEngineType_String(t *testing.T) {
	tests := []struct {
		name string
		e    EngineType
		want string
	}{
		// TODO: Add test cases.
		{
			name: "EngineType Reciprocating",
			e:    EngineType(1),
			want: "Reciprocating",
		},
		{
			name: "EngineType Turbo Prop",
			e:    EngineType(2),
			want: "TurboProp",
		},
		{
			name: "EngineType Turbo Shaft",
			e:    EngineType(3),
			want: "TurboShaft",
		},
		{
			name: "EngineType Invalid",
			e:    EngineType(30),
			want: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("EngineType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
