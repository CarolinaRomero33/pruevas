package anthon

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwis
	// e read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case ':':
		return DosPunt, string(ch)
	case '(':
		return ParDer, string(ch)
	case ')':
		return ParIsq, string(ch)
	case ';':
		return PuntCom, string(ch)
	case '{':
		return LlaveDer, string(ch)
	case '}':
		return LlaveIsq, string(ch)
	case '\n':
		return SaltoDeL, string(ch)
	case '=':
		return Asignaicon, string(ch)
	case '[':
		return AgruppDER, string(ch)
	case ']':
		return AgruppISQ, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	case "LOLIF":
		return LOLIF, buf.String()
	case "LOLELSE":
		return LOLELSE, buf.String()
	case "LOLEND":
		return LOLEND, buf.String()
	case "LOLSWITCH":
		return LOLSWITCH, buf.String()
	case "LOLMAIN":
		return LOLMAIN, buf.String()
	case "LOLFUNC":
		return LOLFUNC, buf.String()
	case "LOLFOR":
		return LOLFOR, buf.String()
	case "LOLELSEIF":
		return LOLELSEIF, buf.String()
	case "VAR":
		return VAR, buf.String()
	case "LOLBREAK":
		return LOLBREAK, buf.String()
	case "LOLWHILE":
		return LOLWHILE, buf.String()
	case "LOLDO":
		return LOLDO, buf.String()
	case "LOLINT":
		return LOLINT, buf.String()
	case "LOLTRY":
		return LOLTRY, buf.String()
	case "LOLCATCH":
		return LOLCATCH, buf.String()
	case "LOLFINALLY":
		return LOLFINALLY, buf.String()
	case "LOLPOINT":
		return LOLPOINT, buf.String()
	case "LOLNEW":
		return LOLNEW, buf.String()
	case "LOLFIXED":
		return LOLFIXED, buf.String()
	case "LOLUNCHECKED":
		return LOLUNCHECKED, buf.String()
	case "LOLFOREACH":
		return LOLFOREACH, buf.String()
	case "LOLIMPORT":
		return LOLIMPORT, buf.String()
	case "LOLPRINT":
		return LOLPRINT, buf.String()
	case "LOLMAXCOPE":
		return LOLMAXCOPE, buf.String()
	case "LOLASINACION":
		return LOLASINACION, buf.String()
	case "LOLTAN":
		return LOLTAN, buf.String()
	case "LOLCOS":
		return LOLCOS, buf.String()
	case "LOLSEN":
		return LOLSEN, buf.String()
	
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
