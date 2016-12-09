package anthon

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK      // *
	COMMA         // ,
	DosPunt       // :
	ParDer        // (
	ParIsq        // )
	PuntCom       // ;
	LlaveDer      // {
	LlaveIsq      // }
	SaltoDeL      //n
	Asignaicon    // =
	AgruppDER     // [
	AgruppISQ     // ]
	comillasdoble //""

	// Keywords
	SELECT
	FROM
	LOLIF
	LOLSWITCH
	LOLEND
	LOLINT
	LOLFUNC
	LOLMAIN
	LOLELSE
	LOLFOR
	LOLELSEIF
	VAR
	LOLBREAK
	LOLWHILE
	LOLDO
	LOLTRY
	LOLCATCH
	LOLFINALLY
	LOLPOINT
	LOLNEW
	LOLFIXED
	LOLFOREACH
	LOLUNCHECKED
	LOLIMPORT
	LOLPRINT
	LOLMAXCOPE
	LOLASINACION
	LOLCOS
	LOLTAN
	LOLSEN
)
