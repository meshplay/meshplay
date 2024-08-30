package utils

import (
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

type mockCloser struct {
	closeFunc func() error
}

func (m mockCloser) Close() error {
	return m.closeFunc()

}

func TestSafeClose(t *testing.T) {

	log := logrus.New()
	hook := test.NewGlobal()
	log.AddHook(hook)

	// testcases for SafeClose(co io.Closer)
	t.Run("SafeClose", func(t *testing.T) {
		// define a io.Closer for testing
		expectedErr := errors.New("close error")
		mc := &mockCloser{
			closeFunc: func() error {
				return expectedErr
			},
		}
		SafeClose(mc)

		if len(hook.Entries) != 1 {
			t.Fatal("expected 1 log entry")
		}
	})
}

func TestBackupConfigFile(t *testing.T) {
	cfgFile := "testdata/config.yaml"
	tmpFile, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpFile.Write(data); err != nil {
		t.Fatal(err)
	}
	BackupConfigFile(tmpFile.Name())
	if _, err := os.Stat("/tmp/config.bak.yaml"); os.IsNotExist(err) {
		t.Errorf("BackupConfigFile failed: backup file does not exist")
	}
}
func TestStringWithCharset(t *testing.T) {
	// checking the length, since this function returns random strings everytime
	strLength := 10

	gotString := StringWithCharset(strLength)

	if len(gotString) != strLength {
		t.Errorf("StringWithCharset got = %v want = %v", len(gotString), strLength)
	}
}

func TestPrereq(t *testing.T) {
	_, _, err := prereq()
	if err != nil {
		t.Errorf("prereq error = %v", err)
	}
}

func TestSetFileLocation(t *testing.T) {
	originalMeshplayFolder := MeshplayFolder
	originalDockerComposeFile := DockerComposeFile
	originalAuthConfigFile := AuthConfigFile
	originalDefaultConfigPath := DefaultConfigPath
	defer func() {
		MeshplayFolder = originalMeshplayFolder
		DockerComposeFile = originalDockerComposeFile
		AuthConfigFile = originalAuthConfigFile
		DefaultConfigPath = originalDefaultConfigPath
	}()
	err := SetFileLocation()
	if err != nil {
		t.Errorf("SetFileLocation error = %v", err)
	}
}

func TestNavigateToBrowser(t *testing.T) {
	// opens up a browser window whenever this test runs
	err := NavigateToBrowser("https://khulnasoft.com")
	if err != nil {
		t.Errorf("NavigateToBrowser error: %v", err)
	}
}

func TestContentTypeIsHTML(t *testing.T) {
	tests := []struct {
		name           string
		response       *http.Response
		expectedOutput bool
	}{
		{
			name: "correct content-type",
			response: &http.Response{
				Header: http.Header{
					"Content-Type": []string{"text/html"},
				},
			},
			expectedOutput: true,
		},
		{
			name: "empty content-type",
			response: &http.Response{
				Header: http.Header{
					"Content-Type": []string{},
				},
			},
			expectedOutput: false,
		},
		{
			name: "incorrect content-type",
			response: &http.Response{
				Header: http.Header{
					"Content-Type": []string{"multipart/form-data"},
				},
			},
			expectedOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContentTypeIsHTML(tt.response)
			if got != tt.expectedOutput {
				t.Errorf("ContentTypeIsHTML error = %v want = %v", got, tt.expectedOutput)
			}
		})
	}
}

func TestAskForConfirmation(t *testing.T) {
	tests := []struct {
		name     string
		question string
		input    string
		want     bool
	}{
		{
			name:     "test with input 'yes'",
			question: "question?",
			input:    "yes\n",
			want:     true,
		},
		{
			name:     "test with input 'no'",
			question: "question?",
			input:    "no\n",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mocking stdio
			// see https://stackoverflow.com/a/64518829
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			_, err = w.WriteString(tt.input)
			if err != nil {
				t.Error(err)
			}
			w.Close()
			stdin := os.Stdin
			defer func() { os.Stdin = stdin }()
			os.Stdin = r

			got := AskForConfirmation(tt.question)
			if got != tt.want {
				t.Errorf("AskForConfirmation got = %v want = %v", got, tt.want)
			}
		})
	}
}

func TestCreateConfigFile(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Error(err, "failed to get users home directory")
	}
	originalDefaultConfigPath := DefaultConfigPath
	defer func() { DefaultConfigPath = originalDefaultConfigPath }()
	DefaultConfigPath = filepath.Join(home, "config.yaml")
	err = CreateConfigFile()
	if err != nil {
		t.Errorf("CreateConfigFile error = %v", err)
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr string
	}{
		{
			name: "Correct URL",
			url:  "https://www.khulnasoft.com",
		},
		{
			name:    "Unsupported scheme",
			url:     "mqtt://www.khulnasoft.com",
			wantErr: "mqtt is not a supported protocol",
		},
		{
			name:    "invalid URL",
			url:     "khulnasoft.com",
			wantErr: "parse \"khulnasoft.com\": invalid URI for request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.url)
			if err != nil {
				// check error message
				if err.Error() != tt.wantErr {
					t.Errorf("ValidateURL error = %v want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

// func TestReadToken(t *testing.T) {
// }

func TestTruncateID(t *testing.T) {
	id := "1234567890"
	want := "12345678"
	got := TruncateID(id)
	if got != want {
		t.Errorf("TruncateID got = %v want = %v", got, want)
	}
}

func TestStringContainedInSlice(t *testing.T) {
	tests := []struct {
		name  string
		str   string
		slice []string
		want  int
	}{
		{
			name:  "test with present string",
			str:   "data2",
			slice: []string{"data1", "data2"},
			want:  1,
		},
		{
			name:  "test with absent string",
			str:   "data3",
			slice: []string{"data1", "data2"},
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringContainedInSlice(tt.str, tt.slice)
			if got != tt.want {
				t.Errorf("StringContainedInSlice got = %v want = %v", got, tt.want)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	tests := []struct {
		name  string
		str   string
		slice []string
		want  bool
	}{
		{
			name:  "test with string present",
			str:   "data1",
			slice: []string{"data1", "data2"},
			want:  true,
		},
		{
			name:  "test with string absent",
			str:   "data3",
			slice: []string{"data1", "data2"},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringInSlice(tt.str, tt.slice)
			if got != tt.want {
				t.Errorf("StringInSlice got = %v want = %v", got, tt.want)
			}
		})
	}
}

func TestAskForInput(t *testing.T) {
	input := "data1"

	// mocking stdio
	// see https://stackoverflow.com/a/64518829
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	_, err = w.WriteString(input + "\n")
	if err != nil {
		t.Error(err)
	}
	w.Close()
	stdin := os.Stdin
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	got := AskForInput("Prompt", []string{"data1", "data2"})
	if got != input {
		t.Errorf("AskForInput got = %v want = %v", got, input)
	}
}

func TestParseURLGithub(t *testing.T) {
	tests := []struct {
		name          string
		url           string
		rawRepoOutput string
		pathOutput    string
		expectedError string
	}{
		{
			name:          "test with non-github url",
			url:           "https://khulnasoft.com",
			rawRepoOutput: "https://khulnasoft.com",
			pathOutput:    "",
			expectedError: "only github urls are supported",
		},
		{
			name:          "test with github.com",
			url:           "https://github.com/khulnasoft/meshplay/blob/master/.goreleaser.yml",
			rawRepoOutput: "https://github.com/khulnasoft/meshplay/master",
			pathOutput:    ".goreleaser.yml",
		},
		{
			name:          "test with raw.githubusercontent.com",
			url:           "https://raw.githubusercontent.com/khulnasoft/meshplay/master/.goreleaser.yml",
			rawRepoOutput: "https://raw.githubusercontent.com/khulnasoft/meshplay/master/.goreleaser.yml",
			pathOutput:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRawRepoOutput, gotPathOutput, gotError := ParseURLGithub(tt.url)
			// gotRawOutput != tt.rawRepoOutput || gotPathOutput != tt.pathOutput ||
			if gotRawRepoOutput != tt.rawRepoOutput {
				t.Errorf("ParseURLGithub got = %v, want = %v", gotRawRepoOutput, tt.rawRepoOutput)
			}
			if gotPathOutput != tt.pathOutput {
				t.Errorf("ParseURLGithub got = %v, want = %v", gotPathOutput, tt.pathOutput)
			}
			if gotError != nil {
				if gotError.Error() != tt.expectedError {
					t.Errorf("ParseURLGithub error = %v, want = %v", gotError, tt.expectedError)
				}
			}
		})
	}
}

// func TestPrintToTableInStringFormat(t *testing.T) {
// 	want := NewGoldenFile(t, "PrintToTableInStringFormat.golden", fixturesDir).Load()
// 	got := PrintToTableInStringFormat([]string{"firstheader", "secondheader"}, [][]string{{"data1", "data2"}, {"data3", "data4"}})
// 	if got != want {
// 		t.Errorf("PrintToTableInStringFormat got = %v want = %v", got, want)
// 	}
// }

func TestCreateDefaultSpinner(t *testing.T) {
	// only checking for Suffix and FinalMSG
	got := CreateDefaultSpinner("suffix", "message")
	want := struct {
		Suffix   string
		FinalMsg string
	}{
		Suffix:   " suffix", // The leading space is intentional
		FinalMsg: "message\n",
	}
	if want.Suffix != got.Suffix {
		t.Errorf("CreateDefaultSpinner got = %v want = %v", got.Suffix, want.Suffix)
	}
	if want.FinalMsg != got.FinalMSG {
		t.Errorf("CreateDefaultSpinner got = %v want = %v", got.FinalMSG, want.FinalMsg)
	}
}

func TestContainsStringPrefix(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		str   string
		want  bool
	}{
		{
			name:  "str is present in the slice",
			slice: []string{"data1", "data2"},
			str:   "data2",
			want:  true,
		},
		{
			name:  "str is not present in the slice",
			slice: []string{"data1", "data2"},
			str:   "data3",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsStringPrefix(tt.slice, tt.str)
			if got != tt.want {
				t.Errorf("ContainsStringPrefix got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetOverrideValues(t *testing.T) {
	testChannel := "testChannel"

	tests := []struct {
		name                string
		ctx                 *config.Context
		meshplayImageVersion string
		want                map[string]interface{}
	}{
		{
			name: "Context contains no components and no meshplay image version",
			ctx: &config.Context{
				Components: nil,
				Channel:    testChannel,
			},
			meshplayImageVersion: "",
			want: map[string]interface{}{
				"meshplay-app-mesh": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-istio": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-cilium": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-linkerd": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-consul": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-kuma": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-nsm": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-nginx-sm": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-traefik-mesh": map[string]interface{}{
					"enabled": false,
				},
				"image": map[string]interface{}{
					"tag": testChannel + "-",
				},
			},
		},
		{
			name: "Context contains part of all available components and meshplay image version",
			ctx: &config.Context{
				Components: []string{"meshplay-istio", "meshplay-nsm"},
				Channel:    testChannel,
			},
			meshplayImageVersion: "testImageVersion",
			want: map[string]interface{}{
				"meshplay-app-mesh": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-istio": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-cilium": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-linkerd": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-consul": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-kuma": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-nsm": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-nginx-sm": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-traefik-mesh": map[string]interface{}{
					"enabled": false,
				},
				"image": map[string]interface{}{
					"tag": testChannel + "-testImageVersion",
				},
			},
		},
		{
			name: "Context contains all available components and meshplay image version",
			ctx: &config.Context{
				Components: []string{"meshplay-istio", "meshplay-linkerd", "meshplay-consul", "meshplay-kuma",
					"meshplay-nsm", "meshplay-nginx-sm", "meshplay-traefik-mesh", "meshplay-cilium"},
				Channel: testChannel,
			},
			meshplayImageVersion: "testImageVersion",
			want: map[string]interface{}{
				"meshplay-app-mesh": map[string]interface{}{
					"enabled": false,
				},
				"meshplay-istio": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-cilium": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-linkerd": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-consul": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-kuma": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-nsm": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-nginx-sm": map[string]interface{}{
					"enabled": true,
				},
				"meshplay-traefik-mesh": map[string]interface{}{
					"enabled": true,
				},
				"image": map[string]interface{}{
					"tag": testChannel + "-testImageVersion",
				},
			},
		},
	}

	for _, tt := range tests {
		got := SetOverrideValues(tt.ctx, tt.meshplayImageVersion, "", "")
		eq := reflect.DeepEqual(got, tt.want)
		if !eq {
			t.Errorf("SetOverrideValues %s got = %v want = %v", tt.name, got, tt.want)
		}
	}
}
