package utils_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/CoopHive/hive/utils"
)

type EnsureDirTest struct {
	name string
	dir  string
	// want    string
	err error
}

func TestEnsureDir(t *testing.T) {

	tests := []EnsureDirTest{
		{"no directory", ".", nil},
		{"tempDir", t.TempDir(), nil},
	}

	for i := 0; i < 10; i++ {
		test := EnsureDirTest{fmt.Sprintf("tempDir%d", i), t.TempDir(), nil}
		tests = append(tests, test)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.EnsureDir(tt.dir)
			if !errors.Is(err, tt.err) {
				t.Errorf("EnsureDir() error = %v, wantErr %v", err, tt.err)
				return
			}
			if tt.dir != got {
				t.Errorf("EnsureDir() got = %v, want %v", got, tt.dir)
			}
		})
	}
}

func genTableTest(t *testing.T) EnsureDirTest {
	tempDir := t.TempDir()

	return EnsureDirTest{
		tempDir,
		tempDir,
		nil,
	}
}
