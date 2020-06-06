// Code generated with internal/tool/generate/keywordtrie; DO NOT EDIT.

package ruleset

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

func defaultKeywordsRule(s RuneScanner) (token.Type, bool) {
	tok, ok := scanKeyword(s)
	if !ok {
		return token.Unknown, false
	}
	peek, noEof := s.Lookahead()
	if noEof && defaultLiteral.Matches(peek) { // keywords must be terminated with a whitespace
		return token.Unknown, false
	}
	return tok, ok
}

func scanKeyword(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordA(s)

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordB(s)

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordC(s)

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordD(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordE(s)

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordF(s)

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordG(s)

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordH(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordI(s)

	case 'J', 'j':
		s.ConsumeRune()
		return scanKeywordJ(s)

	case 'K', 'k':
		s.ConsumeRune()
		return scanKeywordK(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordL(s)

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordM(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordN(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordO(s)

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordP(s)

	case 'Q', 'q':
		s.ConsumeRune()
		return scanKeywordQ(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordR(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordS(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordT(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordU(s)

	case 'V', 'v':
		s.ConsumeRune()
		return scanKeywordV(s)

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordW(s)
	}
	return token.Unknown, false
}

func scanKeywordA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordAB(s)

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordAC(s)

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordAD(s)

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordAF(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordAL(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordAN(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordAS(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordAT(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordAU(s)
	}
	return token.Unknown, false
}

func scanKeywordAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordABO(s)
	}
	return token.Unknown, false
}

func scanKeywordABO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordABOR(s)
	}
	return token.Unknown, false
}

func scanKeywordABOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordABORT(s)
	}
	return token.Unknown, false
}

func scanKeywordABORT(s RuneScanner) (token.Type, bool) {
	return token.KeywordAbort, true
}

func scanKeywordAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordACT(s)
	}
	return token.Unknown, false
}

func scanKeywordACT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordACTI(s)
	}
	return token.Unknown, false
}

func scanKeywordACTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordACTIO(s)
	}
	return token.Unknown, false
}

func scanKeywordACTIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordACTION(s)
	}
	return token.Unknown, false
}

func scanKeywordACTION(s RuneScanner) (token.Type, bool) {
	return token.KeywordAction, true
}

func scanKeywordAD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordADD(s)
	}
	return token.Unknown, false
}

func scanKeywordADD(s RuneScanner) (token.Type, bool) {
	return token.KeywordAdd, true
}

func scanKeywordAF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordAFT(s)
	}
	return token.Unknown, false
}

func scanKeywordAFT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordAFTE(s)
	}
	return token.Unknown, false
}

func scanKeywordAFTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordAFTER(s)
	}
	return token.Unknown, false
}

func scanKeywordAFTER(s RuneScanner) (token.Type, bool) {
	return token.KeywordAfter, true
}

func scanKeywordAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordALL(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordALT(s)

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordALW(s)
	}
	return token.Unknown, false
}

func scanKeywordALL(s RuneScanner) (token.Type, bool) {
	return token.KeywordAll, true
}

func scanKeywordALT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordALTE(s)
	}
	return token.Unknown, false
}

func scanKeywordALTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordALTER(s)
	}
	return token.Unknown, false
}

func scanKeywordALTER(s RuneScanner) (token.Type, bool) {
	return token.KeywordAlter, true
}

func scanKeywordALW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordALWA(s)
	}
	return token.Unknown, false
}

func scanKeywordALWA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordALWAY(s)
	}
	return token.Unknown, false
}

func scanKeywordALWAY(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordALWAYS(s)
	}
	return token.Unknown, false
}

func scanKeywordALWAYS(s RuneScanner) (token.Type, bool) {
	return token.KeywordAlways, true
}

func scanKeywordAN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordANA(s)

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordAND(s)
	}
	return token.Unknown, false
}

func scanKeywordANA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordANAL(s)
	}
	return token.Unknown, false
}

func scanKeywordANAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordANALY(s)
	}
	return token.Unknown, false
}

func scanKeywordANALY(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Z', 'z':
		s.ConsumeRune()
		return scanKeywordANALYZ(s)
	}
	return token.Unknown, false
}

func scanKeywordANALYZ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordANALYZE(s)
	}
	return token.Unknown, false
}

func scanKeywordANALYZE(s RuneScanner) (token.Type, bool) {
	return token.KeywordAnalyze, true
}

func scanKeywordAND(s RuneScanner) (token.Type, bool) {
	return token.KeywordAnd, true
}

func scanKeywordAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordAs, true
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordASC(s)
	}
	return token.KeywordAs, true
}

func scanKeywordASC(s RuneScanner) (token.Type, bool) {
	return token.KeywordAsc, true
}

func scanKeywordAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordATT(s)
	}
	return token.Unknown, false
}

func scanKeywordATT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordATTA(s)
	}
	return token.Unknown, false
}

func scanKeywordATTA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordATTAC(s)
	}
	return token.Unknown, false
}

func scanKeywordATTAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordATTACH(s)
	}
	return token.Unknown, false
}

func scanKeywordATTACH(s RuneScanner) (token.Type, bool) {
	return token.KeywordAttach, true
}

func scanKeywordAU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordAUT(s)
	}
	return token.Unknown, false
}

func scanKeywordAUT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordAUTO(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordAUTOI(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordAUTOIN(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordAUTOINC(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordAUTOINCR(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINCR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordAUTOINCRE(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINCRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordAUTOINCREM(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINCREM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordAUTOINCREME(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINCREME(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordAUTOINCREMEN(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINCREMEN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordAUTOINCREMENT(s)
	}
	return token.Unknown, false
}

func scanKeywordAUTOINCREMENT(s RuneScanner) (token.Type, bool) {
	return token.KeywordAutoincrement, true
}

func scanKeywordB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordBE(s)

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordBY(s)
	}
	return token.Unknown, false
}

func scanKeywordBE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordBEF(s)

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordBEG(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordBET(s)
	}
	return token.Unknown, false
}

func scanKeywordBEF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordBEFO(s)
	}
	return token.Unknown, false
}

func scanKeywordBEFO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordBEFOR(s)
	}
	return token.Unknown, false
}

func scanKeywordBEFOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordBEFORE(s)
	}
	return token.Unknown, false
}

func scanKeywordBEFORE(s RuneScanner) (token.Type, bool) {
	return token.KeywordBefore, true
}

func scanKeywordBEG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordBEGI(s)
	}
	return token.Unknown, false
}

func scanKeywordBEGI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordBEGIN(s)
	}
	return token.Unknown, false
}

func scanKeywordBEGIN(s RuneScanner) (token.Type, bool) {
	return token.KeywordBegin, true
}

func scanKeywordBET(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordBETW(s)
	}
	return token.Unknown, false
}

func scanKeywordBETW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordBETWE(s)
	}
	return token.Unknown, false
}

func scanKeywordBETWE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordBETWEE(s)
	}
	return token.Unknown, false
}

func scanKeywordBETWEE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordBETWEEN(s)
	}
	return token.Unknown, false
}

func scanKeywordBETWEEN(s RuneScanner) (token.Type, bool) {
	return token.KeywordBetween, true
}

func scanKeywordBY(s RuneScanner) (token.Type, bool) {
	return token.KeywordBy, true
}

func scanKeywordC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCA(s)

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordCH(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordCO(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordCR(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordCU(s)
	}
	return token.Unknown, false
}

func scanKeywordCA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordCAS(s)
	}
	return token.Unknown, false
}

func scanKeywordCAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordCASC(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCASE(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCAST(s)
	}
	return token.Unknown, false
}

func scanKeywordCASC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCASCA(s)
	}
	return token.Unknown, false
}

func scanKeywordCASCA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordCASCAD(s)
	}
	return token.Unknown, false
}

func scanKeywordCASCAD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCASCADE(s)
	}
	return token.Unknown, false
}

func scanKeywordCASCADE(s RuneScanner) (token.Type, bool) {
	return token.KeywordCascade, true
}

func scanKeywordCASE(s RuneScanner) (token.Type, bool) {
	return token.KeywordCase, true
}

func scanKeywordCAST(s RuneScanner) (token.Type, bool) {
	return token.KeywordCast, true
}

func scanKeywordCH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCHE(s)
	}
	return token.Unknown, false
}

func scanKeywordCHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordCHEC(s)
	}
	return token.Unknown, false
}

func scanKeywordCHEC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'K', 'k':
		s.ConsumeRune()
		return scanKeywordCHECK(s)
	}
	return token.Unknown, false
}

func scanKeywordCHECK(s RuneScanner) (token.Type, bool) {
	return token.KeywordCheck, true
}

func scanKeywordCO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordCOL(s)

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordCOM(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordCON(s)
	}
	return token.Unknown, false
}

func scanKeywordCOL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordCOLL(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordCOLU(s)
	}
	return token.Unknown, false
}

func scanKeywordCOLL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCOLLA(s)
	}
	return token.Unknown, false
}

func scanKeywordCOLLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCOLLAT(s)
	}
	return token.Unknown, false
}

func scanKeywordCOLLAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCOLLATE(s)
	}
	return token.Unknown, false
}

func scanKeywordCOLLATE(s RuneScanner) (token.Type, bool) {
	return token.KeywordCollate, true
}

func scanKeywordCOLU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordCOLUM(s)
	}
	return token.Unknown, false
}

func scanKeywordCOLUM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordCOLUMN(s)
	}
	return token.Unknown, false
}

func scanKeywordCOLUMN(s RuneScanner) (token.Type, bool) {
	return token.KeywordColumn, true
}

func scanKeywordCOM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordCOMM(s)
	}
	return token.Unknown, false
}

func scanKeywordCOMM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordCOMMI(s)
	}
	return token.Unknown, false
}

func scanKeywordCOMMI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCOMMIT(s)
	}
	return token.Unknown, false
}

func scanKeywordCOMMIT(s RuneScanner) (token.Type, bool) {
	return token.KeywordCommit, true
}

func scanKeywordCON(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordCONF(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordCONS(s)
	}
	return token.Unknown, false
}

func scanKeywordCONF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordCONFL(s)
	}
	return token.Unknown, false
}

func scanKeywordCONFL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordCONFLI(s)
	}
	return token.Unknown, false
}

func scanKeywordCONFLI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordCONFLIC(s)
	}
	return token.Unknown, false
}

func scanKeywordCONFLIC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCONFLICT(s)
	}
	return token.Unknown, false
}

func scanKeywordCONFLICT(s RuneScanner) (token.Type, bool) {
	return token.KeywordConflict, true
}

func scanKeywordCONS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCONST(s)
	}
	return token.Unknown, false
}

func scanKeywordCONST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordCONSTR(s)
	}
	return token.Unknown, false
}

func scanKeywordCONSTR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCONSTRA(s)
	}
	return token.Unknown, false
}

func scanKeywordCONSTRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordCONSTRAI(s)
	}
	return token.Unknown, false
}

func scanKeywordCONSTRAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordCONSTRAIN(s)
	}
	return token.Unknown, false
}

func scanKeywordCONSTRAIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCONSTRAINT(s)
	}
	return token.Unknown, false
}

func scanKeywordCONSTRAINT(s RuneScanner) (token.Type, bool) {
	return token.KeywordConstraint, true
}

func scanKeywordCR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCRE(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordCRO(s)
	}
	return token.Unknown, false
}

func scanKeywordCRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCREA(s)
	}
	return token.Unknown, false
}

func scanKeywordCREA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCREAT(s)
	}
	return token.Unknown, false
}

func scanKeywordCREAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCREATE(s)
	}
	return token.Unknown, false
}

func scanKeywordCREATE(s RuneScanner) (token.Type, bool) {
	return token.KeywordCreate, true
}

func scanKeywordCRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordCROS(s)
	}
	return token.Unknown, false
}

func scanKeywordCROS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordCROSS(s)
	}
	return token.Unknown, false
}

func scanKeywordCROSS(s RuneScanner) (token.Type, bool) {
	return token.KeywordCross, true
}

func scanKeywordCU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordCUR(s)
	}
	return token.Unknown, false
}

func scanKeywordCUR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordCURR(s)
	}
	return token.Unknown, false
}

func scanKeywordCURR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCURRE(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordCURREN(s)
	}
	return token.Unknown, false
}

func scanKeywordCURREN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCURRENT(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordCurrent, true
	}
	switch next {

	case '_':
		s.ConsumeRune()
		return scanKeywordCURRENTx(s)
	}
	return token.KeywordCurrent, true
}

func scanKeywordCURRENTx(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordCURRENTxD(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCURRENTxT(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCURRENTxDA(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxDA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCURRENTxDAT(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxDAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCURRENTxDATE(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxDATE(s RuneScanner) (token.Type, bool) {
	return token.KeywordCurrentDate, true
}

func scanKeywordCURRENTxT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordCURRENTxTI(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIM(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIME(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTIME(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordCurrentTime, true
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIMES(s)
	}
	return token.KeywordCurrentTime, true
}

func scanKeywordCURRENTxTIMES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIMEST(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTIMEST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIMESTA(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTIMESTA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIMESTAM(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTIMESTAM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordCURRENTxTIMESTAMP(s)
	}
	return token.Unknown, false
}

func scanKeywordCURRENTxTIMESTAMP(s RuneScanner) (token.Type, bool) {
	return token.KeywordCurrentTimestamp, true
}

func scanKeywordD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordDA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDE(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordDI(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordDO(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordDR(s)
	}
	return token.Unknown, false
}

func scanKeywordDA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordDAT(s)
	}
	return token.Unknown, false
}

func scanKeywordDAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordDATA(s)
	}
	return token.Unknown, false
}

func scanKeywordDATA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordDATAB(s)
	}
	return token.Unknown, false
}

func scanKeywordDATAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordDATABA(s)
	}
	return token.Unknown, false
}

func scanKeywordDATABA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordDATABAS(s)
	}
	return token.Unknown, false
}

func scanKeywordDATABAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDATABASE(s)
	}
	return token.Unknown, false
}

func scanKeywordDATABASE(s RuneScanner) (token.Type, bool) {
	return token.KeywordDatabase, true
}

func scanKeywordDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordDEF(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordDEL(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordDES(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordDET(s)
	}
	return token.Unknown, false
}

func scanKeywordDEF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordDEFA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDEFE(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordDEFAU(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFAU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordDEFAUL(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFAUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordDEFAULT(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFAULT(s RuneScanner) (token.Type, bool) {
	return token.KeywordDefault, true
}

func scanKeywordDEFE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordDEFER(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordDEFERR(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFERR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordDEFERRA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDEFERRE(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFERRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordDEFERRAB(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFERRAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordDEFERRABL(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFERRABL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDEFERRABLE(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFERRABLE(s RuneScanner) (token.Type, bool) {
	return token.KeywordDeferrable, true
}

func scanKeywordDEFERRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordDEFERRED(s)
	}
	return token.Unknown, false
}

func scanKeywordDEFERRED(s RuneScanner) (token.Type, bool) {
	return token.KeywordDeferred, true
}

func scanKeywordDEL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDELE(s)
	}
	return token.Unknown, false
}

func scanKeywordDELE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordDELET(s)
	}
	return token.Unknown, false
}

func scanKeywordDELET(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordDELETE(s)
	}
	return token.Unknown, false
}

func scanKeywordDELETE(s RuneScanner) (token.Type, bool) {
	return token.KeywordDelete, true
}

func scanKeywordDES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordDESC(s)
	}
	return token.Unknown, false
}

func scanKeywordDESC(s RuneScanner) (token.Type, bool) {
	return token.KeywordDesc, true
}

func scanKeywordDET(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordDETA(s)
	}
	return token.Unknown, false
}

func scanKeywordDETA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordDETAC(s)
	}
	return token.Unknown, false
}

func scanKeywordDETAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordDETACH(s)
	}
	return token.Unknown, false
}

func scanKeywordDETACH(s RuneScanner) (token.Type, bool) {
	return token.KeywordDetach, true
}

func scanKeywordDI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordDIS(s)
	}
	return token.Unknown, false
}

func scanKeywordDIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordDIST(s)
	}
	return token.Unknown, false
}

func scanKeywordDIST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordDISTI(s)
	}
	return token.Unknown, false
}

func scanKeywordDISTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordDISTIN(s)
	}
	return token.Unknown, false
}

func scanKeywordDISTIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordDISTINC(s)
	}
	return token.Unknown, false
}

func scanKeywordDISTINC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordDISTINCT(s)
	}
	return token.Unknown, false
}

func scanKeywordDISTINCT(s RuneScanner) (token.Type, bool) {
	return token.KeywordDistinct, true
}

func scanKeywordDO(s RuneScanner) (token.Type, bool) {
	return token.KeywordDo, true
}

func scanKeywordDR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordDRO(s)
	}
	return token.Unknown, false
}

func scanKeywordDRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordDROP(s)
	}
	return token.Unknown, false
}

func scanKeywordDROP(s RuneScanner) (token.Type, bool) {
	return token.KeywordDrop, true
}

func scanKeywordE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordEA(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordEL(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordEN(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordES(s)

	case 'X', 'x':
		s.ConsumeRune()
		return scanKeywordEX(s)
	}
	return token.Unknown, false
}

func scanKeywordEA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordEAC(s)
	}
	return token.Unknown, false
}

func scanKeywordEAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordEACH(s)
	}
	return token.Unknown, false
}

func scanKeywordEACH(s RuneScanner) (token.Type, bool) {
	return token.KeywordEach, true
}

func scanKeywordEL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordELS(s)
	}
	return token.Unknown, false
}

func scanKeywordELS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordELSE(s)
	}
	return token.Unknown, false
}

func scanKeywordELSE(s RuneScanner) (token.Type, bool) {
	return token.KeywordElse, true
}

func scanKeywordEN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordEND(s)
	}
	return token.Unknown, false
}

func scanKeywordEND(s RuneScanner) (token.Type, bool) {
	return token.KeywordEnd, true
}

func scanKeywordES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordESC(s)
	}
	return token.Unknown, false
}

func scanKeywordESC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordESCA(s)
	}
	return token.Unknown, false
}

func scanKeywordESCA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordESCAP(s)
	}
	return token.Unknown, false
}

func scanKeywordESCAP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordESCAPE(s)
	}
	return token.Unknown, false
}

func scanKeywordESCAPE(s RuneScanner) (token.Type, bool) {
	return token.KeywordEscape, true
}

func scanKeywordEX(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordEXC(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordEXI(s)

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordEXP(s)
	}
	return token.Unknown, false
}

func scanKeywordEXC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordEXCE(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordEXCL(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordEXCEP(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCEP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordEXCEPT(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCEPT(s RuneScanner) (token.Type, bool) {
	return token.KeywordExcept, true
}

func scanKeywordEXCL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordEXCLU(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCLU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordEXCLUD(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordEXCLUS(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCLUD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordEXCLUDE(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCLUDE(s RuneScanner) (token.Type, bool) {
	return token.KeywordExclude, true
}

func scanKeywordEXCLUS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordEXCLUSI(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCLUSI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'V', 'v':
		s.ConsumeRune()
		return scanKeywordEXCLUSIV(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCLUSIV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordEXCLUSIVE(s)
	}
	return token.Unknown, false
}

func scanKeywordEXCLUSIVE(s RuneScanner) (token.Type, bool) {
	return token.KeywordExclusive, true
}

func scanKeywordEXI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordEXIS(s)
	}
	return token.Unknown, false
}

func scanKeywordEXIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordEXIST(s)
	}
	return token.Unknown, false
}

func scanKeywordEXIST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordEXISTS(s)
	}
	return token.Unknown, false
}

func scanKeywordEXISTS(s RuneScanner) (token.Type, bool) {
	return token.KeywordExists, true
}

func scanKeywordEXP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordEXPL(s)
	}
	return token.Unknown, false
}

func scanKeywordEXPL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordEXPLA(s)
	}
	return token.Unknown, false
}

func scanKeywordEXPLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordEXPLAI(s)
	}
	return token.Unknown, false
}

func scanKeywordEXPLAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordEXPLAIN(s)
	}
	return token.Unknown, false
}

func scanKeywordEXPLAIN(s RuneScanner) (token.Type, bool) {
	return token.KeywordExplain, true
}

func scanKeywordF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordFA(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordFI(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordFO(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordFR(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordFU(s)
	}
	return token.Unknown, false
}

func scanKeywordFA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordFAI(s)
	}
	return token.Unknown, false
}

func scanKeywordFAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordFAIL(s)
	}
	return token.Unknown, false
}

func scanKeywordFAIL(s RuneScanner) (token.Type, bool) {
	return token.KeywordFail, true
}

func scanKeywordFI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordFIL(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordFIR(s)
	}
	return token.Unknown, false
}

func scanKeywordFIL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordFILT(s)
	}
	return token.Unknown, false
}

func scanKeywordFILT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordFILTE(s)
	}
	return token.Unknown, false
}

func scanKeywordFILTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordFILTER(s)
	}
	return token.Unknown, false
}

func scanKeywordFILTER(s RuneScanner) (token.Type, bool) {
	return token.KeywordFilter, true
}

func scanKeywordFIR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordFIRS(s)
	}
	return token.Unknown, false
}

func scanKeywordFIRS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordFIRST(s)
	}
	return token.Unknown, false
}

func scanKeywordFIRST(s RuneScanner) (token.Type, bool) {
	return token.KeywordFirst, true
}

func scanKeywordFO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordFOL(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordFOR(s)
	}
	return token.Unknown, false
}

func scanKeywordFOL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordFOLL(s)
	}
	return token.Unknown, false
}

func scanKeywordFOLL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordFOLLO(s)
	}
	return token.Unknown, false
}

func scanKeywordFOLLO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordFOLLOW(s)
	}
	return token.Unknown, false
}

func scanKeywordFOLLOW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordFOLLOWI(s)
	}
	return token.Unknown, false
}

func scanKeywordFOLLOWI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordFOLLOWIN(s)
	}
	return token.Unknown, false
}

func scanKeywordFOLLOWIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordFOLLOWING(s)
	}
	return token.Unknown, false
}

func scanKeywordFOLLOWING(s RuneScanner) (token.Type, bool) {
	return token.KeywordFollowing, true
}

func scanKeywordFOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordFor, true
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordFORE(s)
	}
	return token.KeywordFor, true
}

func scanKeywordFORE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordFOREI(s)
	}
	return token.Unknown, false
}

func scanKeywordFOREI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordFOREIG(s)
	}
	return token.Unknown, false
}

func scanKeywordFOREIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordFOREIGN(s)
	}
	return token.Unknown, false
}

func scanKeywordFOREIGN(s RuneScanner) (token.Type, bool) {
	return token.KeywordForeign, true
}

func scanKeywordFR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordFRO(s)
	}
	return token.Unknown, false
}

func scanKeywordFRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordFROM(s)
	}
	return token.Unknown, false
}

func scanKeywordFROM(s RuneScanner) (token.Type, bool) {
	return token.KeywordFrom, true
}

func scanKeywordFU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordFUL(s)
	}
	return token.Unknown, false
}

func scanKeywordFUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordFULL(s)
	}
	return token.Unknown, false
}

func scanKeywordFULL(s RuneScanner) (token.Type, bool) {
	return token.KeywordFull, true
}

func scanKeywordG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordGE(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordGL(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordGR(s)
	}
	return token.Unknown, false
}

func scanKeywordGE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordGEN(s)
	}
	return token.Unknown, false
}

func scanKeywordGEN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordGENE(s)
	}
	return token.Unknown, false
}

func scanKeywordGENE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordGENER(s)
	}
	return token.Unknown, false
}

func scanKeywordGENER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordGENERA(s)
	}
	return token.Unknown, false
}

func scanKeywordGENERA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordGENERAT(s)
	}
	return token.Unknown, false
}

func scanKeywordGENERAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordGENERATE(s)
	}
	return token.Unknown, false
}

func scanKeywordGENERATE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordGENERATED(s)
	}
	return token.Unknown, false
}

func scanKeywordGENERATED(s RuneScanner) (token.Type, bool) {
	return token.KeywordGenerated, true
}

func scanKeywordGL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordGLO(s)
	}
	return token.Unknown, false
}

func scanKeywordGLO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordGLOB(s)
	}
	return token.Unknown, false
}

func scanKeywordGLOB(s RuneScanner) (token.Type, bool) {
	return token.KeywordGlob, true
}

func scanKeywordGR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordGRO(s)
	}
	return token.Unknown, false
}

func scanKeywordGRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordGROU(s)
	}
	return token.Unknown, false
}

func scanKeywordGROU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordGROUP(s)
	}
	return token.Unknown, false
}

func scanKeywordGROUP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordGroup, true
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordGROUPS(s)
	}
	return token.KeywordGroup, true
}

func scanKeywordGROUPS(s RuneScanner) (token.Type, bool) {
	return token.KeywordGroups, true
}

func scanKeywordH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordHA(s)
	}
	return token.Unknown, false
}

func scanKeywordHA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'V', 'v':
		s.ConsumeRune()
		return scanKeywordHAV(s)
	}
	return token.Unknown, false
}

func scanKeywordHAV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordHAVI(s)
	}
	return token.Unknown, false
}

func scanKeywordHAVI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordHAVIN(s)
	}
	return token.Unknown, false
}

func scanKeywordHAVIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordHAVING(s)
	}
	return token.Unknown, false
}

func scanKeywordHAVING(s RuneScanner) (token.Type, bool) {
	return token.KeywordHaving, true
}

func scanKeywordI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordIF(s)

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordIG(s)

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordIM(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordIN(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordIS(s)
	}
	return token.Unknown, false
}

func scanKeywordIF(s RuneScanner) (token.Type, bool) {
	return token.KeywordIf, true
}

func scanKeywordIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordIGN(s)
	}
	return token.Unknown, false
}

func scanKeywordIGN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordIGNO(s)
	}
	return token.Unknown, false
}

func scanKeywordIGNO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordIGNOR(s)
	}
	return token.Unknown, false
}

func scanKeywordIGNOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordIGNORE(s)
	}
	return token.Unknown, false
}

func scanKeywordIGNORE(s RuneScanner) (token.Type, bool) {
	return token.KeywordIgnore, true
}

func scanKeywordIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordIMM(s)
	}
	return token.Unknown, false
}

func scanKeywordIMM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordIMME(s)
	}
	return token.Unknown, false
}

func scanKeywordIMME(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordIMMED(s)
	}
	return token.Unknown, false
}

func scanKeywordIMMED(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordIMMEDI(s)
	}
	return token.Unknown, false
}

func scanKeywordIMMEDI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordIMMEDIA(s)
	}
	return token.Unknown, false
}

func scanKeywordIMMEDIA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordIMMEDIAT(s)
	}
	return token.Unknown, false
}

func scanKeywordIMMEDIAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordIMMEDIATE(s)
	}
	return token.Unknown, false
}

func scanKeywordIMMEDIATE(s RuneScanner) (token.Type, bool) {
	return token.KeywordImmediate, true
}

func scanKeywordIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordIn, true
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordIND(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordINI(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordINN(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordINS(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordINT(s)
	}
	return token.KeywordIn, true
}

func scanKeywordIND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINDE(s)
	}
	return token.Unknown, false
}

func scanKeywordINDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'X', 'x':
		s.ConsumeRune()
		return scanKeywordINDEX(s)
	}
	return token.Unknown, false
}

func scanKeywordINDEX(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordIndex, true
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINDEXE(s)
	}
	return token.KeywordIndex, true
}

func scanKeywordINDEXE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordINDEXED(s)
	}
	return token.Unknown, false
}

func scanKeywordINDEXED(s RuneScanner) (token.Type, bool) {
	return token.KeywordIndexed, true
}

func scanKeywordINI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordINIT(s)
	}
	return token.Unknown, false
}

func scanKeywordINIT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordINITI(s)
	}
	return token.Unknown, false
}

func scanKeywordINITI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordINITIA(s)
	}
	return token.Unknown, false
}

func scanKeywordINITIA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordINITIAL(s)
	}
	return token.Unknown, false
}

func scanKeywordINITIAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordINITIALL(s)
	}
	return token.Unknown, false
}

func scanKeywordINITIALL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordINITIALLY(s)
	}
	return token.Unknown, false
}

func scanKeywordINITIALLY(s RuneScanner) (token.Type, bool) {
	return token.KeywordInitially, true
}

func scanKeywordINN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINNE(s)
	}
	return token.Unknown, false
}

func scanKeywordINNE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordINNER(s)
	}
	return token.Unknown, false
}

func scanKeywordINNER(s RuneScanner) (token.Type, bool) {
	return token.KeywordInner, true
}

func scanKeywordINS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINSE(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordINST(s)
	}
	return token.Unknown, false
}

func scanKeywordINSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordINSER(s)
	}
	return token.Unknown, false
}

func scanKeywordINSER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordINSERT(s)
	}
	return token.Unknown, false
}

func scanKeywordINSERT(s RuneScanner) (token.Type, bool) {
	return token.KeywordInsert, true
}

func scanKeywordINST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINSTE(s)
	}
	return token.Unknown, false
}

func scanKeywordINSTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordINSTEA(s)
	}
	return token.Unknown, false
}

func scanKeywordINSTEA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordINSTEAD(s)
	}
	return token.Unknown, false
}

func scanKeywordINSTEAD(s RuneScanner) (token.Type, bool) {
	return token.KeywordInstead, true
}

func scanKeywordINT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINTE(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordINTO(s)
	}
	return token.Unknown, false
}

func scanKeywordINTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordINTER(s)
	}
	return token.Unknown, false
}

func scanKeywordINTER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordINTERS(s)
	}
	return token.Unknown, false
}

func scanKeywordINTERS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordINTERSE(s)
	}
	return token.Unknown, false
}

func scanKeywordINTERSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordINTERSEC(s)
	}
	return token.Unknown, false
}

func scanKeywordINTERSEC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordINTERSECT(s)
	}
	return token.Unknown, false
}

func scanKeywordINTERSECT(s RuneScanner) (token.Type, bool) {
	return token.KeywordIntersect, true
}

func scanKeywordINTO(s RuneScanner) (token.Type, bool) {
	return token.KeywordInto, true
}

func scanKeywordIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordIs, true
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordISN(s)
	}
	return token.KeywordIs, true
}

func scanKeywordISN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordISNU(s)
	}
	return token.Unknown, false
}

func scanKeywordISNU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordISNUL(s)
	}
	return token.Unknown, false
}

func scanKeywordISNUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordISNULL(s)
	}
	return token.Unknown, false
}

func scanKeywordISNULL(s RuneScanner) (token.Type, bool) {
	return token.KeywordIsnull, true
}

func scanKeywordJ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordJO(s)
	}
	return token.Unknown, false
}

func scanKeywordJO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordJOI(s)
	}
	return token.Unknown, false
}

func scanKeywordJOI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordJOIN(s)
	}
	return token.Unknown, false
}

func scanKeywordJOIN(s RuneScanner) (token.Type, bool) {
	return token.KeywordJoin, true
}

func scanKeywordK(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordKE(s)
	}
	return token.Unknown, false
}

func scanKeywordKE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordKEY(s)
	}
	return token.Unknown, false
}

func scanKeywordKEY(s RuneScanner) (token.Type, bool) {
	return token.KeywordKey, true
}

func scanKeywordL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordLA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordLE(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordLI(s)
	}
	return token.Unknown, false
}

func scanKeywordLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordLAS(s)
	}
	return token.Unknown, false
}

func scanKeywordLAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordLAST(s)
	}
	return token.Unknown, false
}

func scanKeywordLAST(s RuneScanner) (token.Type, bool) {
	return token.KeywordLast, true
}

func scanKeywordLE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordLEF(s)
	}
	return token.Unknown, false
}

func scanKeywordLEF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordLEFT(s)
	}
	return token.Unknown, false
}

func scanKeywordLEFT(s RuneScanner) (token.Type, bool) {
	return token.KeywordLeft, true
}

func scanKeywordLI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'K', 'k':
		s.ConsumeRune()
		return scanKeywordLIK(s)

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordLIM(s)
	}
	return token.Unknown, false
}

func scanKeywordLIK(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordLIKE(s)
	}
	return token.Unknown, false
}

func scanKeywordLIKE(s RuneScanner) (token.Type, bool) {
	return token.KeywordLike, true
}

func scanKeywordLIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordLIMI(s)
	}
	return token.Unknown, false
}

func scanKeywordLIMI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordLIMIT(s)
	}
	return token.Unknown, false
}

func scanKeywordLIMIT(s RuneScanner) (token.Type, bool) {
	return token.KeywordLimit, true
}

func scanKeywordM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordMA(s)
	}
	return token.Unknown, false
}

func scanKeywordMA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordMAT(s)
	}
	return token.Unknown, false
}

func scanKeywordMAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordMATC(s)
	}
	return token.Unknown, false
}

func scanKeywordMATC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordMATCH(s)
	}
	return token.Unknown, false
}

func scanKeywordMATCH(s RuneScanner) (token.Type, bool) {
	return token.KeywordMatch, true
}

func scanKeywordN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordNA(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordNO(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordNU(s)
	}
	return token.Unknown, false
}

func scanKeywordNA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordNAT(s)
	}
	return token.Unknown, false
}

func scanKeywordNAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordNATU(s)
	}
	return token.Unknown, false
}

func scanKeywordNATU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordNATUR(s)
	}
	return token.Unknown, false
}

func scanKeywordNATUR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordNATURA(s)
	}
	return token.Unknown, false
}

func scanKeywordNATURA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordNATURAL(s)
	}
	return token.Unknown, false
}

func scanKeywordNATURAL(s RuneScanner) (token.Type, bool) {
	return token.KeywordNatural, true
}

func scanKeywordNO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordNo, true
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordNOT(s)
	}
	return token.KeywordNo, true
}

func scanKeywordNOT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordNot, true
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordNOTH(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordNOTN(s)
	}
	return token.KeywordNot, true
}

func scanKeywordNOTH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordNOTHI(s)
	}
	return token.Unknown, false
}

func scanKeywordNOTHI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordNOTHIN(s)
	}
	return token.Unknown, false
}

func scanKeywordNOTHIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordNOTHING(s)
	}
	return token.Unknown, false
}

func scanKeywordNOTHING(s RuneScanner) (token.Type, bool) {
	return token.KeywordNothing, true
}

func scanKeywordNOTN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordNOTNU(s)
	}
	return token.Unknown, false
}

func scanKeywordNOTNU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordNOTNUL(s)
	}
	return token.Unknown, false
}

func scanKeywordNOTNUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordNOTNULL(s)
	}
	return token.Unknown, false
}

func scanKeywordNOTNULL(s RuneScanner) (token.Type, bool) {
	return token.KeywordNotnull, true
}

func scanKeywordNU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordNUL(s)
	}
	return token.Unknown, false
}

func scanKeywordNUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordNULL(s)
	}
	return token.Unknown, false
}

func scanKeywordNULL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordNull, true
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordNULLS(s)
	}
	return token.KeywordNull, true
}

func scanKeywordNULLS(s RuneScanner) (token.Type, bool) {
	return token.KeywordNulls, true
}

func scanKeywordO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordOF(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordON(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordOR(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordOT(s)

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordOU(s)

	case 'V', 'v':
		s.ConsumeRune()
		return scanKeywordOV(s)
	}
	return token.Unknown, false
}

func scanKeywordOF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordOf, true
	}
	switch next {

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordOFF(s)
	}
	return token.KeywordOf, true
}

func scanKeywordOFF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordOFFS(s)
	}
	return token.Unknown, false
}

func scanKeywordOFFS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordOFFSE(s)
	}
	return token.Unknown, false
}

func scanKeywordOFFSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordOFFSET(s)
	}
	return token.Unknown, false
}

func scanKeywordOFFSET(s RuneScanner) (token.Type, bool) {
	return token.KeywordOffset, true
}

func scanKeywordON(s RuneScanner) (token.Type, bool) {
	return token.KeywordOn, true
}

func scanKeywordOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordOr, true
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordORD(s)
	}
	return token.KeywordOr, true
}

func scanKeywordORD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordORDE(s)
	}
	return token.Unknown, false
}

func scanKeywordORDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordORDER(s)
	}
	return token.Unknown, false
}

func scanKeywordORDER(s RuneScanner) (token.Type, bool) {
	return token.KeywordOrder, true
}

func scanKeywordOT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordOTH(s)
	}
	return token.Unknown, false
}

func scanKeywordOTH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordOTHE(s)
	}
	return token.Unknown, false
}

func scanKeywordOTHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordOTHER(s)
	}
	return token.Unknown, false
}

func scanKeywordOTHER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordOTHERS(s)
	}
	return token.Unknown, false
}

func scanKeywordOTHERS(s RuneScanner) (token.Type, bool) {
	return token.KeywordOthers, true
}

func scanKeywordOU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordOUT(s)
	}
	return token.Unknown, false
}

func scanKeywordOUT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordOUTE(s)
	}
	return token.Unknown, false
}

func scanKeywordOUTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordOUTER(s)
	}
	return token.Unknown, false
}

func scanKeywordOUTER(s RuneScanner) (token.Type, bool) {
	return token.KeywordOuter, true
}

func scanKeywordOV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordOVE(s)
	}
	return token.Unknown, false
}

func scanKeywordOVE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordOVER(s)
	}
	return token.Unknown, false
}

func scanKeywordOVER(s RuneScanner) (token.Type, bool) {
	return token.KeywordOver, true
}

func scanKeywordP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordPA(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordPL(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordPR(s)
	}
	return token.Unknown, false
}

func scanKeywordPA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordPAR(s)
	}
	return token.Unknown, false
}

func scanKeywordPAR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordPART(s)
	}
	return token.Unknown, false
}

func scanKeywordPART(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordPARTI(s)
	}
	return token.Unknown, false
}

func scanKeywordPARTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordPARTIT(s)
	}
	return token.Unknown, false
}

func scanKeywordPARTIT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordPARTITI(s)
	}
	return token.Unknown, false
}

func scanKeywordPARTITI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordPARTITIO(s)
	}
	return token.Unknown, false
}

func scanKeywordPARTITIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordPARTITION(s)
	}
	return token.Unknown, false
}

func scanKeywordPARTITION(s RuneScanner) (token.Type, bool) {
	return token.KeywordPartition, true
}

func scanKeywordPL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordPLA(s)
	}
	return token.Unknown, false
}

func scanKeywordPLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordPLAN(s)
	}
	return token.Unknown, false
}

func scanKeywordPLAN(s RuneScanner) (token.Type, bool) {
	return token.KeywordPlan, true
}

func scanKeywordPR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordPRA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordPRE(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordPRI(s)
	}
	return token.Unknown, false
}

func scanKeywordPRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordPRAG(s)
	}
	return token.Unknown, false
}

func scanKeywordPRAG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordPRAGM(s)
	}
	return token.Unknown, false
}

func scanKeywordPRAGM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordPRAGMA(s)
	}
	return token.Unknown, false
}

func scanKeywordPRAGMA(s RuneScanner) (token.Type, bool) {
	return token.KeywordPragma, true
}

func scanKeywordPRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordPREC(s)
	}
	return token.Unknown, false
}

func scanKeywordPREC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordPRECE(s)
	}
	return token.Unknown, false
}

func scanKeywordPRECE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordPRECED(s)
	}
	return token.Unknown, false
}

func scanKeywordPRECED(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordPRECEDI(s)
	}
	return token.Unknown, false
}

func scanKeywordPRECEDI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordPRECEDIN(s)
	}
	return token.Unknown, false
}

func scanKeywordPRECEDIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordPRECEDING(s)
	}
	return token.Unknown, false
}

func scanKeywordPRECEDING(s RuneScanner) (token.Type, bool) {
	return token.KeywordPreceding, true
}

func scanKeywordPRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordPRIM(s)
	}
	return token.Unknown, false
}

func scanKeywordPRIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordPRIMA(s)
	}
	return token.Unknown, false
}

func scanKeywordPRIMA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordPRIMAR(s)
	}
	return token.Unknown, false
}

func scanKeywordPRIMAR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordPRIMARY(s)
	}
	return token.Unknown, false
}

func scanKeywordPRIMARY(s RuneScanner) (token.Type, bool) {
	return token.KeywordPrimary, true
}

func scanKeywordQ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordQU(s)
	}
	return token.Unknown, false
}

func scanKeywordQU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordQUE(s)
	}
	return token.Unknown, false
}

func scanKeywordQUE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordQUER(s)
	}
	return token.Unknown, false
}

func scanKeywordQUER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordQUERY(s)
	}
	return token.Unknown, false
}

func scanKeywordQUERY(s RuneScanner) (token.Type, bool) {
	return token.KeywordQuery, true
}

func scanKeywordR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordRA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRE(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordRI(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordRO(s)
	}
	return token.Unknown, false
}

func scanKeywordRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordRAI(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordRAN(s)
	}
	return token.Unknown, false
}

func scanKeywordRAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordRAIS(s)
	}
	return token.Unknown, false
}

func scanKeywordRAIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRAISE(s)
	}
	return token.Unknown, false
}

func scanKeywordRAISE(s RuneScanner) (token.Type, bool) {
	return token.KeywordRaise, true
}

func scanKeywordRAN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordRANG(s)
	}
	return token.Unknown, false
}

func scanKeywordRANG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRANGE(s)
	}
	return token.Unknown, false
}

func scanKeywordRANGE(s RuneScanner) (token.Type, bool) {
	return token.KeywordRange, true
}

func scanKeywordRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordREC(s)

	case 'F', 'f':
		s.ConsumeRune()
		return scanKeywordREF(s)

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordREG(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordREI(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordREL(s)

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordREN(s)

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordREP(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordRES(s)
	}
	return token.Unknown, false
}

func scanKeywordREC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordRECU(s)
	}
	return token.Unknown, false
}

func scanKeywordRECU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordRECUR(s)
	}
	return token.Unknown, false
}

func scanKeywordRECUR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordRECURS(s)
	}
	return token.Unknown, false
}

func scanKeywordRECURS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordRECURSI(s)
	}
	return token.Unknown, false
}

func scanKeywordRECURSI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'V', 'v':
		s.ConsumeRune()
		return scanKeywordRECURSIV(s)
	}
	return token.Unknown, false
}

func scanKeywordRECURSIV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRECURSIVE(s)
	}
	return token.Unknown, false
}

func scanKeywordRECURSIVE(s RuneScanner) (token.Type, bool) {
	return token.KeywordRecursive, true
}

func scanKeywordREF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordREFE(s)
	}
	return token.Unknown, false
}

func scanKeywordREFE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordREFER(s)
	}
	return token.Unknown, false
}

func scanKeywordREFER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordREFERE(s)
	}
	return token.Unknown, false
}

func scanKeywordREFERE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordREFEREN(s)
	}
	return token.Unknown, false
}

func scanKeywordREFEREN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordREFERENC(s)
	}
	return token.Unknown, false
}

func scanKeywordREFERENC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordREFERENCE(s)
	}
	return token.Unknown, false
}

func scanKeywordREFERENCE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordREFERENCES(s)
	}
	return token.Unknown, false
}

func scanKeywordREFERENCES(s RuneScanner) (token.Type, bool) {
	return token.KeywordReferences, true
}

func scanKeywordREG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordREGE(s)
	}
	return token.Unknown, false
}

func scanKeywordREGE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'X', 'x':
		s.ConsumeRune()
		return scanKeywordREGEX(s)
	}
	return token.Unknown, false
}

func scanKeywordREGEX(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordREGEXP(s)
	}
	return token.Unknown, false
}

func scanKeywordREGEXP(s RuneScanner) (token.Type, bool) {
	return token.KeywordRegexp, true
}

func scanKeywordREI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordREIN(s)
	}
	return token.Unknown, false
}

func scanKeywordREIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordREIND(s)
	}
	return token.Unknown, false
}

func scanKeywordREIND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordREINDE(s)
	}
	return token.Unknown, false
}

func scanKeywordREINDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'X', 'x':
		s.ConsumeRune()
		return scanKeywordREINDEX(s)
	}
	return token.Unknown, false
}

func scanKeywordREINDEX(s RuneScanner) (token.Type, bool) {
	return token.KeywordReIndex, true
}

func scanKeywordREL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRELE(s)
	}
	return token.Unknown, false
}

func scanKeywordRELE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordRELEA(s)
	}
	return token.Unknown, false
}

func scanKeywordRELEA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordRELEAS(s)
	}
	return token.Unknown, false
}

func scanKeywordRELEAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRELEASE(s)
	}
	return token.Unknown, false
}

func scanKeywordRELEASE(s RuneScanner) (token.Type, bool) {
	return token.KeywordRelease, true
}

func scanKeywordREN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordRENA(s)
	}
	return token.Unknown, false
}

func scanKeywordRENA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordRENAM(s)
	}
	return token.Unknown, false
}

func scanKeywordRENAM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordRENAME(s)
	}
	return token.Unknown, false
}

func scanKeywordRENAME(s RuneScanner) (token.Type, bool) {
	return token.KeywordRename, true
}

func scanKeywordREP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordREPL(s)
	}
	return token.Unknown, false
}

func scanKeywordREPL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordREPLA(s)
	}
	return token.Unknown, false
}

func scanKeywordREPLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordREPLAC(s)
	}
	return token.Unknown, false
}

func scanKeywordREPLAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordREPLACE(s)
	}
	return token.Unknown, false
}

func scanKeywordREPLACE(s RuneScanner) (token.Type, bool) {
	return token.KeywordReplace, true
}

func scanKeywordRES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordREST(s)
	}
	return token.Unknown, false
}

func scanKeywordREST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordRESTR(s)
	}
	return token.Unknown, false
}

func scanKeywordRESTR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordRESTRI(s)
	}
	return token.Unknown, false
}

func scanKeywordRESTRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordRESTRIC(s)
	}
	return token.Unknown, false
}

func scanKeywordRESTRIC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordRESTRICT(s)
	}
	return token.Unknown, false
}

func scanKeywordRESTRICT(s RuneScanner) (token.Type, bool) {
	return token.KeywordRestrict, true
}

func scanKeywordRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordRIG(s)
	}
	return token.Unknown, false
}

func scanKeywordRIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordRIGH(s)
	}
	return token.Unknown, false
}

func scanKeywordRIGH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordRIGHT(s)
	}
	return token.Unknown, false
}

func scanKeywordRIGHT(s RuneScanner) (token.Type, bool) {
	return token.KeywordRight, true
}

func scanKeywordRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordROL(s)

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordROW(s)
	}
	return token.Unknown, false
}

func scanKeywordROL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordROLL(s)
	}
	return token.Unknown, false
}

func scanKeywordROLL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordROLLB(s)
	}
	return token.Unknown, false
}

func scanKeywordROLLB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordROLLBA(s)
	}
	return token.Unknown, false
}

func scanKeywordROLLBA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordROLLBAC(s)
	}
	return token.Unknown, false
}

func scanKeywordROLLBAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'K', 'k':
		s.ConsumeRune()
		return scanKeywordROLLBACK(s)
	}
	return token.Unknown, false
}

func scanKeywordROLLBACK(s RuneScanner) (token.Type, bool) {
	return token.KeywordRollback, true
}

func scanKeywordROW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordRow, true
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordROWS(s)
	}
	return token.KeywordRow, true
}

func scanKeywordROWS(s RuneScanner) (token.Type, bool) {
	return token.KeywordRows, true
}

func scanKeywordS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordSA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordSE(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordST(s)
	}
	return token.Unknown, false
}

func scanKeywordSA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'V', 'v':
		s.ConsumeRune()
		return scanKeywordSAV(s)
	}
	return token.Unknown, false
}

func scanKeywordSAV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordSAVE(s)
	}
	return token.Unknown, false
}

func scanKeywordSAVE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordSAVEP(s)
	}
	return token.Unknown, false
}

func scanKeywordSAVEP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordSAVEPO(s)
	}
	return token.Unknown, false
}

func scanKeywordSAVEPO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordSAVEPOI(s)
	}
	return token.Unknown, false
}

func scanKeywordSAVEPOI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordSAVEPOIN(s)
	}
	return token.Unknown, false
}

func scanKeywordSAVEPOIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordSAVEPOINT(s)
	}
	return token.Unknown, false
}

func scanKeywordSAVEPOINT(s RuneScanner) (token.Type, bool) {
	return token.KeywordSavepoint, true
}

func scanKeywordSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordSEL(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordSET(s)
	}
	return token.Unknown, false
}

func scanKeywordSEL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordSELE(s)
	}
	return token.Unknown, false
}

func scanKeywordSELE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordSELEC(s)
	}
	return token.Unknown, false
}

func scanKeywordSELEC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordSELECT(s)
	}
	return token.Unknown, false
}

func scanKeywordSELECT(s RuneScanner) (token.Type, bool) {
	return token.KeywordSelect, true
}

func scanKeywordSET(s RuneScanner) (token.Type, bool) {
	return token.KeywordSet, true
}

func scanKeywordST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordSTO(s)
	}
	return token.Unknown, false
}

func scanKeywordSTO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordSTOR(s)
	}
	return token.Unknown, false
}

func scanKeywordSTOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordSTORE(s)
	}
	return token.Unknown, false
}

func scanKeywordSTORE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordSTORED(s)
	}
	return token.Unknown, false
}

func scanKeywordSTORED(s RuneScanner) (token.Type, bool) {
	return token.KeywordStored, true
}

func scanKeywordT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordTA(s)

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordTE(s)

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordTH(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordTI(s)

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordTO(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordTR(s)
	}
	return token.Unknown, false
}

func scanKeywordTA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordTAB(s)
	}
	return token.Unknown, false
}

func scanKeywordTAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordTABL(s)
	}
	return token.Unknown, false
}

func scanKeywordTABL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordTABLE(s)
	}
	return token.Unknown, false
}

func scanKeywordTABLE(s RuneScanner) (token.Type, bool) {
	return token.KeywordTable, true
}

func scanKeywordTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordTEM(s)
	}
	return token.Unknown, false
}

func scanKeywordTEM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordTEMP(s)
	}
	return token.Unknown, false
}

func scanKeywordTEMP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordTemp, true
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordTEMPO(s)
	}
	return token.KeywordTemp, true
}

func scanKeywordTEMPO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordTEMPOR(s)
	}
	return token.Unknown, false
}

func scanKeywordTEMPOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordTEMPORA(s)
	}
	return token.Unknown, false
}

func scanKeywordTEMPORA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordTEMPORAR(s)
	}
	return token.Unknown, false
}

func scanKeywordTEMPORAR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'Y', 'y':
		s.ConsumeRune()
		return scanKeywordTEMPORARY(s)
	}
	return token.Unknown, false
}

func scanKeywordTEMPORARY(s RuneScanner) (token.Type, bool) {
	return token.KeywordTemporary, true
}

func scanKeywordTH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordTHE(s)
	}
	return token.Unknown, false
}

func scanKeywordTHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordTHEN(s)
	}
	return token.Unknown, false
}

func scanKeywordTHEN(s RuneScanner) (token.Type, bool) {
	return token.KeywordThen, true
}

func scanKeywordTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordTIE(s)
	}
	return token.Unknown, false
}

func scanKeywordTIE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordTIES(s)
	}
	return token.Unknown, false
}

func scanKeywordTIES(s RuneScanner) (token.Type, bool) {
	return token.KeywordTies, true
}

func scanKeywordTO(s RuneScanner) (token.Type, bool) {
	return token.KeywordTo, true
}

func scanKeywordTR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordTRA(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordTRI(s)
	}
	return token.Unknown, false
}

func scanKeywordTRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordTRAN(s)
	}
	return token.Unknown, false
}

func scanKeywordTRAN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordTRANS(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordTRANSA(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANSA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordTRANSAC(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANSAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordTRANSACT(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANSACT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordTRANSACTI(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANSACTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordTRANSACTIO(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANSACTIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordTRANSACTION(s)
	}
	return token.Unknown, false
}

func scanKeywordTRANSACTION(s RuneScanner) (token.Type, bool) {
	return token.KeywordTransaction, true
}

func scanKeywordTRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordTRIG(s)
	}
	return token.Unknown, false
}

func scanKeywordTRIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordTRIGG(s)
	}
	return token.Unknown, false
}

func scanKeywordTRIGG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordTRIGGE(s)
	}
	return token.Unknown, false
}

func scanKeywordTRIGGE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordTRIGGER(s)
	}
	return token.Unknown, false
}

func scanKeywordTRIGGER(s RuneScanner) (token.Type, bool) {
	return token.KeywordTrigger, true
}

func scanKeywordU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordUN(s)

	case 'P', 'p':
		s.ConsumeRune()
		return scanKeywordUP(s)

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordUS(s)
	}
	return token.Unknown, false
}

func scanKeywordUN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'B', 'b':
		s.ConsumeRune()
		return scanKeywordUNB(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordUNI(s)
	}
	return token.Unknown, false
}

func scanKeywordUNB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordUNBO(s)
	}
	return token.Unknown, false
}

func scanKeywordUNBO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordUNBOU(s)
	}
	return token.Unknown, false
}

func scanKeywordUNBOU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordUNBOUN(s)
	}
	return token.Unknown, false
}

func scanKeywordUNBOUN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordUNBOUND(s)
	}
	return token.Unknown, false
}

func scanKeywordUNBOUND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordUNBOUNDE(s)
	}
	return token.Unknown, false
}

func scanKeywordUNBOUNDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordUNBOUNDED(s)
	}
	return token.Unknown, false
}

func scanKeywordUNBOUNDED(s RuneScanner) (token.Type, bool) {
	return token.KeywordUnbounded, true
}

func scanKeywordUNI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordUNIO(s)

	case 'Q', 'q':
		s.ConsumeRune()
		return scanKeywordUNIQ(s)
	}
	return token.Unknown, false
}

func scanKeywordUNIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordUNION(s)
	}
	return token.Unknown, false
}

func scanKeywordUNION(s RuneScanner) (token.Type, bool) {
	return token.KeywordUnion, true
}

func scanKeywordUNIQ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordUNIQU(s)
	}
	return token.Unknown, false
}

func scanKeywordUNIQU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordUNIQUE(s)
	}
	return token.Unknown, false
}

func scanKeywordUNIQUE(s RuneScanner) (token.Type, bool) {
	return token.KeywordUnique, true
}

func scanKeywordUP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordUPD(s)
	}
	return token.Unknown, false
}

func scanKeywordUPD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordUPDA(s)
	}
	return token.Unknown, false
}

func scanKeywordUPDA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordUPDAT(s)
	}
	return token.Unknown, false
}

func scanKeywordUPDAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordUPDATE(s)
	}
	return token.Unknown, false
}

func scanKeywordUPDATE(s RuneScanner) (token.Type, bool) {
	return token.KeywordUpdate, true
}

func scanKeywordUS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordUSI(s)
	}
	return token.Unknown, false
}

func scanKeywordUSI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordUSIN(s)
	}
	return token.Unknown, false
}

func scanKeywordUSIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'G', 'g':
		s.ConsumeRune()
		return scanKeywordUSING(s)
	}
	return token.Unknown, false
}

func scanKeywordUSING(s RuneScanner) (token.Type, bool) {
	return token.KeywordUsing, true
}

func scanKeywordV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordVA(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordVI(s)
	}
	return token.Unknown, false
}

func scanKeywordVA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'C', 'c':
		s.ConsumeRune()
		return scanKeywordVAC(s)

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordVAL(s)
	}
	return token.Unknown, false
}

func scanKeywordVAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordVACU(s)
	}
	return token.Unknown, false
}

func scanKeywordVACU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordVACUU(s)
	}
	return token.Unknown, false
}

func scanKeywordVACUU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'M', 'm':
		s.ConsumeRune()
		return scanKeywordVACUUM(s)
	}
	return token.Unknown, false
}

func scanKeywordVACUUM(s RuneScanner) (token.Type, bool) {
	return token.KeywordVacuum, true
}

func scanKeywordVAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordVALU(s)
	}
	return token.Unknown, false
}

func scanKeywordVALU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordVALUE(s)
	}
	return token.Unknown, false
}

func scanKeywordVALUE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'S', 's':
		s.ConsumeRune()
		return scanKeywordVALUES(s)
	}
	return token.Unknown, false
}

func scanKeywordVALUES(s RuneScanner) (token.Type, bool) {
	return token.KeywordValues, true
}

func scanKeywordVI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordVIE(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordVIR(s)
	}
	return token.Unknown, false
}

func scanKeywordVIE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordVIEW(s)
	}
	return token.Unknown, false
}

func scanKeywordVIEW(s RuneScanner) (token.Type, bool) {
	return token.KeywordView, true
}

func scanKeywordVIR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordVIRT(s)
	}
	return token.Unknown, false
}

func scanKeywordVIRT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordVIRTU(s)
	}
	return token.Unknown, false
}

func scanKeywordVIRTU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'A', 'a':
		s.ConsumeRune()
		return scanKeywordVIRTUA(s)
	}
	return token.Unknown, false
}

func scanKeywordVIRTUA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'L', 'l':
		s.ConsumeRune()
		return scanKeywordVIRTUAL(s)
	}
	return token.Unknown, false
}

func scanKeywordVIRTUAL(s RuneScanner) (token.Type, bool) {
	return token.KeywordVirtual, true
}

func scanKeywordW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordWH(s)

	case 'I', 'i':
		s.ConsumeRune()
		return scanKeywordWI(s)
	}
	return token.Unknown, false
}

func scanKeywordWH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordWHE(s)
	}
	return token.Unknown, false
}

func scanKeywordWHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordWHEN(s)

	case 'R', 'r':
		s.ConsumeRune()
		return scanKeywordWHER(s)
	}
	return token.Unknown, false
}

func scanKeywordWHEN(s RuneScanner) (token.Type, bool) {
	return token.KeywordWhen, true
}

func scanKeywordWHER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'E', 'e':
		s.ConsumeRune()
		return scanKeywordWHERE(s)
	}
	return token.Unknown, false
}

func scanKeywordWHERE(s RuneScanner) (token.Type, bool) {
	return token.KeywordWhere, true
}

func scanKeywordWI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'N', 'n':
		s.ConsumeRune()
		return scanKeywordWIN(s)

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordWIT(s)
	}
	return token.Unknown, false
}

func scanKeywordWIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'D', 'd':
		s.ConsumeRune()
		return scanKeywordWIND(s)
	}
	return token.Unknown, false
}

func scanKeywordWIND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordWINDO(s)
	}
	return token.Unknown, false
}

func scanKeywordWINDO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'W', 'w':
		s.ConsumeRune()
		return scanKeywordWINDOW(s)
	}
	return token.Unknown, false
}

func scanKeywordWINDOW(s RuneScanner) (token.Type, bool) {
	return token.KeywordWindow, true
}

func scanKeywordWIT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'H', 'h':
		s.ConsumeRune()
		return scanKeywordWITH(s)
	}
	return token.Unknown, false
}

func scanKeywordWITH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordWith, true
	}
	switch next {

	case 'O', 'o':
		s.ConsumeRune()
		return scanKeywordWITHO(s)
	}
	return token.KeywordWith, true
}

func scanKeywordWITHO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'U', 'u':
		s.ConsumeRune()
		return scanKeywordWITHOU(s)
	}
	return token.Unknown, false
}

func scanKeywordWITHOU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {

	case 'T', 't':
		s.ConsumeRune()
		return scanKeywordWITHOUT(s)
	}
	return token.Unknown, false
}

func scanKeywordWITHOUT(s RuneScanner) (token.Type, bool) {
	return token.KeywordWithout, true
}
