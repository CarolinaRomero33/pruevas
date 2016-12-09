package anthon_test

import (
	"github.com/CarolinaRomero33/pruevas"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok anthon.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: anthon.EOF},
		{s: `#`, tok: anthon.ILLEGAL, lit: `#`},
		{s: ` `, tok: anthon.WS, lit: " "},
		{s: "\t", tok: anthon.WS, lit: "\t"},
		{s: "\r", tok: anthon.WS, lit: "\r"},

		// Misc characters
		{s: `*`, tok: anthon.ASTERISK, lit: "*"},
		{s: `:`, tok: anthon.DosPunt, lit: ":"},
		{s: `(`, tok: anthon.ParDer, lit: "("},
		{s: `)`, tok: anthon.ParIsq, lit: ")"},
		{s: `;`, tok: anthon.PuntCom, lit: ";"},
		{s: `=`, tok: anthon.Asignaicon, lit: "="},

		// Identifiers
		{s: `foo`, tok: anthon.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: anthon.IDENT, lit: `Zx12_3U_`},

		//Keywords
		{s: `LOLIF`, tok: anthon.LOLIF, lit: "LOLIF"},
		{s: `LOLFOR`, tok: anthon.LOLFOR, lit: "LOLFOR"},
		{s: `LOLELSEIF`, tok: anthon.LOLELSEIF, lit: "LOLELSEIF"},
		{s: `LOLMAIN`, tok: anthon.LOLMAIN, lit: "LOLMAIN"},
		{s: `LOLELSE`, tok: anthon.LOLELSE, lit: "LOLELSE"},
		{s: `LOLSWITCH`, tok: anthon.LOLSWITCH, lit: "LOLSWITCH"},
		{s: `LOLFUNC`, tok: anthon.LOLFUNC, lit: "LOLFUNC"},
		{s: `LOLINT`, tok: anthon.LOLINT, lit: "LOLINT"},
		{s: `LOLEND`, tok: anthon.LOLEND, lit: "LOLEND"},
		{s: `VAR`, tok: anthon.VAR, lit: "VAR"},
		{s: `LOLWHILE`, tok: anthon.LOLWHILE, lit: "LOLWHILE"},
		{s: `LOLDO`, tok: anthon.LOLDO, lit: "LOLDO"},
		{s: `LOLTRY`, tok: anthon.LOLTRY, lit: "LOLTRY"},
		{s: `LOLCATCH`, tok: anthon.LOLCATCH, lit: "LOLCATCH"},
		{s: `LOLFINALLY`, tok: anthon.LOLFINALLY, lit: "LOLFINALLY"},
		{s: `LOLPOINT`, tok: anthon.LOLPOINT, lit: "LOLPOINT"},
		{s: `LOLFIXED`, tok: anthon.LOLFIXED, lit: "LOLFIXED"},
		{s: `LOLUNCHECKED`, tok: anthon.LOLUNCHECKED, lit: "LOLUNCHECKED"},
		{s: `LOLNEW`, tok: anthon.LOLNEW, lit: "LOLNEW"},
		{s: `LOLFOREACH`, tok: anthon.LOLFOREACH, lit: "LOLFOREACH"},
		{s: `LOLIMPORT`, tok: anthon.LOLIMPORT, lit: "LOLIMPORT"},
		{s: `LOLMAXCOPE`, tok: anthon.LOLMAXCOPE, lit: "LOLMAXCOPE"},
		{s: `LOLASINACION`, tok: anthon.LOLASINACION, lit: "LOLASINACION"},
		{s: `LOLTAN`, tok: anthon.LOLTAN, lit: "LOLTAN"},
		{s: `LOLCOS`, tok: anthon.LOLCOS, lit: "LOLCOS"},
		{s: `LOLSEN`, tok: anthon.LOLSEN, lit: "LOLSEN"},
	}

	for i, tt := range tests {
		s := anthon.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
