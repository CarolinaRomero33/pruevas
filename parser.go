package anthon

import (
	"fmt"
	"io"
)

type SelectStatement struct{}
type LOLifstatement struct{}

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*SelectStatement, error) {
	stmt := &SelectStatement{}
	//////////////////////////1////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLIF {
		return nil, fmt.Errorf("encontro %q, esperaba LOLif", lit)
	}
	return stmt, nil
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLELSE {
		return nil, fmt.Errorf("encontro %q, esperaba LOLELSE ", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	/////////////////////////////////////2///////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFOR {
		return nil, fmt.Errorf("encontro %q, esperaba ", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != COMMA {
		return nil, fmt.Errorf("encontro %q, esperaba ,", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != COMMA {
		return nil, fmt.Errorf("encontro %q, esperaba ,", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLINT {
		return nil, fmt.Errorf("encontro %q, esperaba VAR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	/////////////////////////////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba LOLmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	///////////////////////////////////////////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba LOLfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba LOLmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//////////////////LOLIF/////////////////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFOR {
		return nil, fmt.Errorf("encontro %q, esperaba LOLfor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba gofor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	///////////////////////////////////LOLELSEIF//////////////////////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLELSEIF {
		return nil, fmt.Errorf("encontro %q, esperaba LOLelseif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	/////////////////////////LOLSWITH//////////////////////////////////////////////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLSWITCH {
		return nil, fmt.Errorf("encontro %q, esperaba LOLswitch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLBREAK {
		return nil, fmt.Errorf("encontro %q, esperaba LOLbreak", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//////////////////////////////////LOLWHILE//////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLINT {
		return nil, fmt.Errorf("encontro %q, esperaba LOLINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba LOLwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//////////////////LOLDO//////////////////////////////////////////////////////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLDO {
		return nil, fmt.Errorf("encontro %q, esperaba LOLDO", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLDO {
		return nil, fmt.Errorf("encontro %q, esperaba LOLDO", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba LOLwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	///////////////////////LOLTRY/////////////////////////////////////////////////////////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLTRY {
		return nil, fmt.Errorf("encontro %q, esperaba LOLtry", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLCATCH {
		return nil, fmt.Errorf("encontro %q, esperaba LOLcatch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	///////////////////////LOLFINALLY

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFINALLY {
		return nil, fmt.Errorf("encontro %q, esperaba LOLfinally", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	/////////////////////LOLPOINT///////////////////////////////////////////////////////////////

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLPOINT {
		return nil, fmt.Errorf("encontro %q, esperaba LOLPOINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLNEW {
		return nil, fmt.Errorf("encontro %q, esperaba GONEW", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//GOFIXED

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFIXED {
		return nil, fmt.Errorf("encontro %q, esperaba LOLFIXED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	///////////////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLELSEIF {
		return nil, fmt.Errorf("encontro %q, esperaba LOLELSEIF", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppDER {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLMAXCOPE {
		return nil, fmt.Errorf("encontro %q, esperaba  valores", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLASINACION {
		return nil, fmt.Errorf("encontro %q, esperaba LOLASINACION", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}

	//GOUNCHECKED

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba LOLUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLINT {
		return nil, fmt.Errorf("encontro %q, esperaba LOLint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba valores", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba LOLUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	////ARREGLO
	//
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLINT {
		return nil, fmt.Errorf("encontro %q, esperaba LOLINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppISQ {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	///////////////////////////////////////////////////////////////////////////////////////
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLNEW {
		return nil, fmt.Errorf("encontro %q, esperaba LOLnew", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLINT {
		return nil, fmt.Errorf("encontro %q, esperaba LOLint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//
	////goforeach

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFOREACH {
		return nil, fmt.Errorf("encontro %q, esperaba LOLFOREACH", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	/////////////////////////////////////////////////////////////PAQUETES
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLIMPORT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	////////////////////////////////////////////////////////////////////funciones

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba LOLfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba lolend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//gomain

	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba LOLmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLEND {
		return nil, fmt.Errorf("encontro %q, esperaba LOLend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////FUNCION DE TAN
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLTAN {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLCOS {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLSEN {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}

	return stmt, nil

}
func (p *Parser) ParseGoIF() (*LOLifstatement, error) {
	stmt := &LOLifstatement{}
	// First token should be a "SELECT" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != LOLIF {
		return nil, fmt.Errorf("encontro %q, esperaba LOLif", lit)
	}
	return stmt, nil
}

func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}
