package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	contestID = "abc146"
	overwrite = false
)

func TestGenCodeforcesContestTemplates(t *testing.T) {
	rootPath := fmt.Sprintf("../../dash/%s/", contestID)
	if err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := filepath.Join(path, fileName)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	tips := fmt.Sprintf("cd %[1]s\ncf submit %[1]s a a/main.go\n", contestID)
	if err := ioutil.WriteFile(rootPath+"tips.txt", []byte(tips), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGenCodeforcesNormalTemplates(t *testing.T) {
	const problemURL = "https://codeforces.ml/contest/1265/problem/D"
	problemID := parseProblemIDFromURL(problemURL)
	mainStr := fmt.Sprintf(`package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol%[1]s(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() {
	Sol%[1]s(os.Stdin, os.Stdout)
}
`, problemID)
	mainTestStr := fmt.Sprintf(`package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol%[1]s(t *testing.T) {
	// just copy from website
	rawText := `+"`\n`"+`
	testutil.AssertEqualCase(t, rawText, 0, Sol%[1]s)
}
`, problemID)
	rootPath := "../../main/"
	if err := ioutil.WriteFile(rootPath+problemID+".go", []byte(mainStr), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(rootPath+problemID+"_test.go", []byte(mainTestStr), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGenNormalTemplates(t *testing.T) {
	const rootPath = "../../nowcoder/2720/"
	for i := 'a'; i <= 'h'; i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := dir + fileName
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				t.Fatal(err)
			}
		}
	}
}
