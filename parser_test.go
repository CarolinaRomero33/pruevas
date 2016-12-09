package anthon_test

import (
	"github.com/CarolinaRomero33/pruevas"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

// Ensure the parser can parse strings into Statement ASTs.

func TestParser_ParseStatement(t *testing.T) {
	f, _ := ioutil.ReadFile("C:\\porno\\Gramatico.txt")
	z := string(f[:])

	var tests = []struct {
		s    string
		stmt *anthon.SelectStatement
		err  string
	}{
		{
			s:    z,
			stmt: &anthon.SelectStatement{},
		},
	}
	for i, tt := range tests {
		stmt, err := anthon.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}

}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
