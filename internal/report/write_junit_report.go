package report

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/ci-on-dev/ci.on-cli/internal/assert"
)

type Testsuites struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Testsuites []Testsuite `xml:"testsuite"`
}

type Testsuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Failures  int        `xml:"failures,attr"`
	Testcases []Testcase `xml:"testcase"`
}

type Testcase struct {
	XMLName xml.Name `xml:"testcase"`
	Class   string   `xml:"classname,attr"`
	Name    string   `xml:"name,attr"`
	Failure *Failure `xml:"failure,omitempty"`
}

type Failure struct {
	Message string `xml:"message,attr"`
	Detail  string `xml:",chardata"`
}

func WriteJUnitFile(results assert.Results, filePath string) error {
	var suites []Testsuite
	suiteMap := make(map[string]*Testsuite)

	for _, r := range results.Results {
		if _, ok := suiteMap[r.Suite]; !ok {
			suiteMap[r.Suite] = &Testsuite{
				Name:      r.Suite,
				Testcases: []Testcase{},
			}
		}

		tc := Testcase{
			Class: r.Suite,
			Name:  r.Test,
		}

		if !r.Passed {
			tc.Failure = &Failure{
				Message: "assertion failed",
				Detail:  collectAssertionMessages(r.Assertions),
			}
			suiteMap[r.Suite].Failures++
		}

		suiteMap[r.Suite].Tests++
		suiteMap[r.Suite].Testcases = append(suiteMap[r.Suite].Testcases, tc)
	}

	for _, s := range suiteMap {
		suites = append(suites, *s)
	}

	data, err := xml.MarshalIndent(Testsuites{Testsuites: suites}, "", "  ")
	if err != nil {
		return err
	}
	data = []byte(xml.Header + string(data))

	// se o filePath for só um diretório, adiciona junit.xml
	info, err := os.Stat(filePath)
	if err == nil && info.IsDir() {
		filePath = filepath.Join(filePath, "junit.xml")
	}

	fileFolder := filepath.Dir(filePath)
	os.MkdirAll(fileFolder, 0755)

	return os.WriteFile(filePath, data, 0644)
}

func collectAssertionMessages(assertions []assert.AssertionResult) string {
	msg := ""
	for _, a := range assertions {
		if !a.Passed {
			msg += a.Message + "\n"
		}
	}
	return msg
}
