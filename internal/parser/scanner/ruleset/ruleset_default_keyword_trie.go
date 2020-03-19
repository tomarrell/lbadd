// Code generated; DO NOT EDIT.

package ruleset

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

func defaultKeywordsRule(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordH(s)
	case 'M':
		return _scanKeywordM(s)
	case 'Q':
		return _scanKeywordQ(s)
	case 'V':
		return _scanKeywordV(s)
	case 'N':
		return _scanKeywordN(s)
	case 'O':
		return _scanKeywordO(s)
	case 'U':
		return _scanKeywordU(s)
	case 'G':
		return _scanKeywordG(s)
	case 'C':
		return _scanKeywordC(s)
	case 'I':
		return _scanKeywordI(s)
	case 'R':
		return _scanKeywordR(s)
	case 'A':
		return _scanKeywordA(s)
	case 'D':
		return _scanKeywordD(s)
	case 'E':
		return _scanKeywordE(s)
	case 'W':
		return _scanKeywordW(s)
	case 'S':
		return _scanKeywordS(s)
	case 'P':
		return _scanKeywordP(s)
	case 'L':
		return _scanKeywordL(s)
	case 'K':
		return _scanKeywordK(s)
	case 'J':
		return _scanKeywordJ(s)
	case 'B':
		return _scanKeywordB(s)
	case 'F':
		return _scanKeywordF(s)
	case 'T':
		return _scanKeywordT(s)
	}
	return token.Unknown, false
}

func _scanKeywordB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordBY(s)
	case 'E':
		return _scanKeywordBE(s)
	}
	return token.Unknown, false
}

func _scanKeywordBY(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordBy, true
	}
	return token.Unknown, false
}

func _scanKeywordBE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'F':
		return _scanKeywordBEF(s)
	case 'G':
		return _scanKeywordBEG(s)
	case 'T':
		return _scanKeywordBET(s)
	}
	return token.Unknown, false
}

func _scanKeywordBEF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordBEFO(s)
	}
	return token.Unknown, false
}

func _scanKeywordBEFO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordBEFOR(s)
	}
	return token.Unknown, false
}

func _scanKeywordBEFOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordBEFORE(s)
	}
	return token.Unknown, false
}

func _scanKeywordBEFORE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordBefore, true
	}
	return token.Unknown, false
}

func _scanKeywordBEG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordBEGI(s)
	}
	return token.Unknown, false
}

func _scanKeywordBEGI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordBEGIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordBEGIN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordBegin, true
	}
	return token.Unknown, false
}

func _scanKeywordBET(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'W':
		return _scanKeywordBETW(s)
	}
	return token.Unknown, false
}

func _scanKeywordBETW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordBETWE(s)
	}
	return token.Unknown, false
}

func _scanKeywordBETWE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordBETWEE(s)
	}
	return token.Unknown, false
}

func _scanKeywordBETWEE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordBETWEEN(s)
	}
	return token.Unknown, false
}

func _scanKeywordBETWEEN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordBetween, true
	}
	return token.Unknown, false
}

func _scanKeywordF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordFI(s)
	case 'R':
		return _scanKeywordFR(s)
	case 'O':
		return _scanKeywordFO(s)
	case 'U':
		return _scanKeywordFU(s)
	case 'A':
		return _scanKeywordFA(s)
	}
	return token.Unknown, false
}

func _scanKeywordFA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordFAI(s)
	}
	return token.Unknown, false
}

func _scanKeywordFAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordFAIL(s)
	}
	return token.Unknown, false
}

func _scanKeywordFAIL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordFail, true
	}
	return token.Unknown, false
}

func _scanKeywordFI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordFIL(s)
	case 'R':
		return _scanKeywordFIR(s)
	}
	return token.Unknown, false
}

func _scanKeywordFIR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordFIRS(s)
	}
	return token.Unknown, false
}

func _scanKeywordFIRS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordFIRST(s)
	}
	return token.Unknown, false
}

func _scanKeywordFIRST(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordFirst, true
	}
	return token.Unknown, false
}

func _scanKeywordFIL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordFILT(s)
	}
	return token.Unknown, false
}

func _scanKeywordFILT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordFILTE(s)
	}
	return token.Unknown, false
}

func _scanKeywordFILTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordFILTER(s)
	}
	return token.Unknown, false
}

func _scanKeywordFILTER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordFilter, true
	}
	return token.Unknown, false
}

func _scanKeywordFR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordFRO(s)
	}
	return token.Unknown, false
}

func _scanKeywordFRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordFROM(s)
	}
	return token.Unknown, false
}

func _scanKeywordFROM(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordFrom, true
	}
	return token.Unknown, false
}

func _scanKeywordFO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordFOR(s)
	case 'L':
		return _scanKeywordFOL(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordFor, true
	}
	switch next {
	case 'E':
		return _scanKeywordFORE(s)
	}
	return token.Unknown, false
}

func _scanKeywordFORE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordFOREI(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOREI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordFOREIG(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOREIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordFOREIGN(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOREIGN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordForeign, true
	}
	return token.Unknown, false
}

func _scanKeywordFOL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordFOLL(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOLL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordFOLLO(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOLLO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'W':
		return _scanKeywordFOLLOW(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOLLOW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordFOLLOWI(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOLLOWI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordFOLLOWIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOLLOWIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordFOLLOWING(s)
	}
	return token.Unknown, false
}

func _scanKeywordFOLLOWING(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordFollowing, true
	}
	return token.Unknown, false
}

func _scanKeywordFU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordFUL(s)
	}
	return token.Unknown, false
}

func _scanKeywordFUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordFULL(s)
	}
	return token.Unknown, false
}

func _scanKeywordFULL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordFull, true
	}
	return token.Unknown, false
}

func _scanKeywordT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordTE(s)
	case 'R':
		return _scanKeywordTR(s)
	case 'A':
		return _scanKeywordTA(s)
	case 'H':
		return _scanKeywordTH(s)
	case 'O':
		return _scanKeywordTO(s)
	case 'I':
		return _scanKeywordTI(s)
	}
	return token.Unknown, false
}

func _scanKeywordTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordTIE(s)
	}
	return token.Unknown, false
}

func _scanKeywordTIE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordTIES(s)
	}
	return token.Unknown, false
}

func _scanKeywordTIES(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordTies, true
	}
	return token.Unknown, false
}

func _scanKeywordTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordTEM(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordTEMP(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEMP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordTemp, true
	}
	switch next {
	case 'O':
		return _scanKeywordTEMPO(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEMPO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordTEMPOR(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEMPOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordTEMPORA(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEMPORA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordTEMPORAR(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEMPORAR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordTEMPORARY(s)
	}
	return token.Unknown, false
}

func _scanKeywordTEMPORARY(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordTemporary, true
	}
	return token.Unknown, false
}

func _scanKeywordTR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordTRI(s)
	case 'A':
		return _scanKeywordTRA(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordTRIG(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordTRIGG(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRIGG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordTRIGGE(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRIGGE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordTRIGGER(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRIGGER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordTrigger, true
	}
	return token.Unknown, false
}

func _scanKeywordTRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordTRAN(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRAN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordTRANS(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordTRANSA(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANSA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordTRANSAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANSAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordTRANSACT(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANSACT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordTRANSACTI(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANSACTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordTRANSACTIO(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANSACTIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordTRANSACTION(s)
	}
	return token.Unknown, false
}

func _scanKeywordTRANSACTION(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordTransaction, true
	}
	return token.Unknown, false
}

func _scanKeywordTA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'B':
		return _scanKeywordTAB(s)
	}
	return token.Unknown, false
}

func _scanKeywordTAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordTABL(s)
	}
	return token.Unknown, false
}

func _scanKeywordTABL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordTABLE(s)
	}
	return token.Unknown, false
}

func _scanKeywordTABLE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordTable, true
	}
	return token.Unknown, false
}

func _scanKeywordTH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordTHE(s)
	}
	return token.Unknown, false
}

func _scanKeywordTHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordTHEN(s)
	}
	return token.Unknown, false
}

func _scanKeywordTHEN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordThen, true
	}
	return token.Unknown, false
}

func _scanKeywordTO(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordTo, true
	}
	return token.Unknown, false
}

func _scanKeywordJ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordJO(s)
	}
	return token.Unknown, false
}

func _scanKeywordJO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordJOI(s)
	}
	return token.Unknown, false
}

func _scanKeywordJOI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordJOIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordJOIN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordJoin, true
	}
	return token.Unknown, false
}

func _scanKeywordM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordMA(s)
	}
	return token.Unknown, false
}

func _scanKeywordMA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordMAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordMAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordMATC(s)
	}
	return token.Unknown, false
}

func _scanKeywordMATC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordMATCH(s)
	}
	return token.Unknown, false
}

func _scanKeywordMATCH(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordMatch, true
	}
	return token.Unknown, false
}

func _scanKeywordQ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordQU(s)
	}
	return token.Unknown, false
}

func _scanKeywordQU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordQUE(s)
	}
	return token.Unknown, false
}

func _scanKeywordQUE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordQUER(s)
	}
	return token.Unknown, false
}

func _scanKeywordQUER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordQUERY(s)
	}
	return token.Unknown, false
}

func _scanKeywordQUERY(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordQuery, true
	}
	return token.Unknown, false
}

func _scanKeywordV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordVA(s)
	case 'I':
		return _scanKeywordVI(s)
	}
	return token.Unknown, false
}

func _scanKeywordVI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordVIE(s)
	case 'R':
		return _scanKeywordVIR(s)
	}
	return token.Unknown, false
}

func _scanKeywordVIE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'W':
		return _scanKeywordVIEW(s)
	}
	return token.Unknown, false
}

func _scanKeywordVIEW(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordView, true
	}
	return token.Unknown, false
}

func _scanKeywordVIR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordVIRT(s)
	}
	return token.Unknown, false
}

func _scanKeywordVIRT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordVIRTU(s)
	}
	return token.Unknown, false
}

func _scanKeywordVIRTU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordVIRTUA(s)
	}
	return token.Unknown, false
}

func _scanKeywordVIRTUA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordVIRTUAL(s)
	}
	return token.Unknown, false
}

func _scanKeywordVIRTUAL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordVirtual, true
	}
	return token.Unknown, false
}

func _scanKeywordVA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordVAL(s)
	case 'C':
		return _scanKeywordVAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordVAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordVALU(s)
	}
	return token.Unknown, false
}

func _scanKeywordVALU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordVALUE(s)
	}
	return token.Unknown, false
}

func _scanKeywordVALUE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordVALUES(s)
	}
	return token.Unknown, false
}

func _scanKeywordVALUES(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordValues, true
	}
	return token.Unknown, false
}

func _scanKeywordVAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordVACU(s)
	}
	return token.Unknown, false
}

func _scanKeywordVACU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordVACUU(s)
	}
	return token.Unknown, false
}

func _scanKeywordVACUU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordVACUUM(s)
	}
	return token.Unknown, false
}

func _scanKeywordVACUUM(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordVacuum, true
	}
	return token.Unknown, false
}

func _scanKeywordN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordNO(s)
	case 'A':
		return _scanKeywordNA(s)
	case 'U':
		return _scanKeywordNU(s)
	}
	return token.Unknown, false
}

func _scanKeywordNU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordNUL(s)
	}
	return token.Unknown, false
}

func _scanKeywordNUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordNULL(s)
	}
	return token.Unknown, false
}

func _scanKeywordNULL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordNull, true
	}
	return token.Unknown, false
}

func _scanKeywordNO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordNo, true
	}
	switch next {
	case 'T':
		return _scanKeywordNOT(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordNot, true
	}
	switch next {
	case 'N':
		return _scanKeywordNOTN(s)
	case 'H':
		return _scanKeywordNOTH(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordNOTNU(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTNU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordNOTNUL(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTNUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordNOTNULL(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTNULL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordNotnull, true
	}
	return token.Unknown, false
}

func _scanKeywordNOTH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordNOTHI(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTHI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordNOTHIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTHIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordNOTHING(s)
	}
	return token.Unknown, false
}

func _scanKeywordNOTHING(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordNothing, true
	}
	return token.Unknown, false
}

func _scanKeywordNA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordNAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordNAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordNATU(s)
	}
	return token.Unknown, false
}

func _scanKeywordNATU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordNATUR(s)
	}
	return token.Unknown, false
}

func _scanKeywordNATUR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordNATURA(s)
	}
	return token.Unknown, false
}

func _scanKeywordNATURA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordNATURAL(s)
	}
	return token.Unknown, false
}

func _scanKeywordNATURAL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordNatural, true
	}
	return token.Unknown, false
}

func _scanKeywordH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordHA(s)
	}
	return token.Unknown, false
}

func _scanKeywordHA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'V':
		return _scanKeywordHAV(s)
	}
	return token.Unknown, false
}

func _scanKeywordHAV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordHAVI(s)
	}
	return token.Unknown, false
}

func _scanKeywordHAVI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordHAVIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordHAVIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordHAVING(s)
	}
	return token.Unknown, false
}

func _scanKeywordHAVING(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordHaving, true
	}
	return token.Unknown, false
}

func _scanKeywordU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordUS(s)
	case 'N':
		return _scanKeywordUN(s)
	case 'P':
		return _scanKeywordUP(s)
	}
	return token.Unknown, false
}

func _scanKeywordUP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordUPD(s)
	}
	return token.Unknown, false
}

func _scanKeywordUPD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordUPDA(s)
	}
	return token.Unknown, false
}

func _scanKeywordUPDA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordUPDAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordUPDAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordUPDATE(s)
	}
	return token.Unknown, false
}

func _scanKeywordUPDATE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordUpdate, true
	}
	return token.Unknown, false
}

func _scanKeywordUS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordUSI(s)
	}
	return token.Unknown, false
}

func _scanKeywordUSI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordUSIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordUSIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordUSING(s)
	}
	return token.Unknown, false
}

func _scanKeywordUSING(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordUsing, true
	}
	return token.Unknown, false
}

func _scanKeywordUN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'B':
		return _scanKeywordUNB(s)
	case 'I':
		return _scanKeywordUNI(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordUNBO(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNBO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordUNBOU(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNBOU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordUNBOUN(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNBOUN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordUNBOUND(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNBOUND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordUNBOUNDE(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNBOUNDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordUNBOUNDED(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNBOUNDED(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordUnbounded, true
	}
	return token.Unknown, false
}

func _scanKeywordUNI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordUNIO(s)
	case 'Q':
		return _scanKeywordUNIQ(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNIQ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordUNIQU(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNIQU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordUNIQUE(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNIQUE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordUnique, true
	}
	return token.Unknown, false
}

func _scanKeywordUNIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordUNION(s)
	}
	return token.Unknown, false
}

func _scanKeywordUNION(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordUnion, true
	}
	return token.Unknown, false
}

func _scanKeywordG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordGR(s)
	case 'L':
		return _scanKeywordGL(s)
	}
	return token.Unknown, false
}

func _scanKeywordGL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordGLO(s)
	}
	return token.Unknown, false
}

func _scanKeywordGLO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'B':
		return _scanKeywordGLOB(s)
	}
	return token.Unknown, false
}

func _scanKeywordGLOB(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordGlob, true
	}
	return token.Unknown, false
}

func _scanKeywordGR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordGRO(s)
	}
	return token.Unknown, false
}

func _scanKeywordGRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordGROU(s)
	}
	return token.Unknown, false
}

func _scanKeywordGROU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordGROUP(s)
	}
	return token.Unknown, false
}

func _scanKeywordGROUP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordGroup, true
	}
	switch next {
	case 'S':
		return _scanKeywordGROUPS(s)
	}
	return token.Unknown, false
}

func _scanKeywordGROUPS(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordGroups, true
	}
	return token.Unknown, false
}

func _scanKeywordC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordCU(s)
	case 'R':
		return _scanKeywordCR(s)
	case 'H':
		return _scanKeywordCH(s)
	case 'A':
		return _scanKeywordCA(s)
	case 'O':
		return _scanKeywordCO(s)
	}
	return token.Unknown, false
}

func _scanKeywordCU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordCUR(s)
	}
	return token.Unknown, false
}

func _scanKeywordCUR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordCURR(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCURRE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordCURREN(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURREN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCURRENT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordCurrent, true
	}
	switch next {
	case '_':
		return _scanKeywordCURRENT_(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordCURRENT_D(s)
	case 'T':
		return _scanKeywordCURRENT_T(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_D(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordCURRENT_DA(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_DA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCURRENT_DAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_DAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCURRENT_DATE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_DATE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCurrentDate, true
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_T(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordCURRENT_TI(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordCURRENT_TIM(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCURRENT_TIME(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIME(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordCurrentTime, true
	}
	switch next {
	case 'S':
		return _scanKeywordCURRENT_TIMES(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIMES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCURRENT_TIMEST(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIMEST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordCURRENT_TIMESTA(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIMESTA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordCURRENT_TIMESTAM(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIMESTAM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordCURRENT_TIMESTAMP(s)
	}
	return token.Unknown, false
}

func _scanKeywordCURRENT_TIMESTAMP(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCurrentTimestamp, true
	}
	return token.Unknown, false
}

func _scanKeywordCR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCRE(s)
	case 'O':
		return _scanKeywordCRO(s)
	}
	return token.Unknown, false
}

func _scanKeywordCRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordCREA(s)
	}
	return token.Unknown, false
}

func _scanKeywordCREA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCREAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCREAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCREATE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCREATE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCreate, true
	}
	return token.Unknown, false
}

func _scanKeywordCRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordCROS(s)
	}
	return token.Unknown, false
}

func _scanKeywordCROS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordCROSS(s)
	}
	return token.Unknown, false
}

func _scanKeywordCROSS(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCross, true
	}
	return token.Unknown, false
}

func _scanKeywordCH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCHE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordCHEC(s)
	}
	return token.Unknown, false
}

func _scanKeywordCHEC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'K':
		return _scanKeywordCHECK(s)
	}
	return token.Unknown, false
}

func _scanKeywordCHECK(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCheck, true
	}
	return token.Unknown, false
}

func _scanKeywordCA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordCAS(s)
	}
	return token.Unknown, false
}

func _scanKeywordCAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCAST(s)
	case 'C':
		return _scanKeywordCASC(s)
	case 'E':
		return _scanKeywordCASE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCASE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCase, true
	}
	return token.Unknown, false
}

func _scanKeywordCAST(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCast, true
	}
	return token.Unknown, false
}

func _scanKeywordCASC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordCASCA(s)
	}
	return token.Unknown, false
}

func _scanKeywordCASCA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordCASCAD(s)
	}
	return token.Unknown, false
}

func _scanKeywordCASCAD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCASCADE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCASCADE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCascade, true
	}
	return token.Unknown, false
}

func _scanKeywordCO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordCOM(s)
	case 'N':
		return _scanKeywordCON(s)
	case 'L':
		return _scanKeywordCOL(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordCOLL(s)
	case 'U':
		return _scanKeywordCOLU(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOLL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordCOLLA(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOLLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCOLLAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOLLAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordCOLLATE(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOLLATE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCollate, true
	}
	return token.Unknown, false
}

func _scanKeywordCOLU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordCOLUM(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOLUM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordCOLUMN(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOLUMN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordColumn, true
	}
	return token.Unknown, false
}

func _scanKeywordCOM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordCOMM(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOMM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordCOMMI(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOMMI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCOMMIT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCOMMIT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordCommit, true
	}
	return token.Unknown, false
}

func _scanKeywordCON(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordCONS(s)
	case 'F':
		return _scanKeywordCONF(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCONST(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordCONSTR(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONSTR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordCONSTRA(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONSTRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordCONSTRAI(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONSTRAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordCONSTRAIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONSTRAIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCONSTRAINT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONSTRAINT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordConstraint, true
	}
	return token.Unknown, false
}

func _scanKeywordCONF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordCONFL(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONFL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordCONFLI(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONFLI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordCONFLIC(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONFLIC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordCONFLICT(s)
	}
	return token.Unknown, false
}

func _scanKeywordCONFLICT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordConflict, true
	}
	return token.Unknown, false
}

func _scanKeywordI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'F':
		return _scanKeywordIF(s)
	case 'M':
		return _scanKeywordIM(s)
	case 'G':
		return _scanKeywordIG(s)
	case 'N':
		return _scanKeywordIN(s)
	case 'S':
		return _scanKeywordIS(s)
	}
	return token.Unknown, false
}

func _scanKeywordIF(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordIf, true
	}
	return token.Unknown, false
}

func _scanKeywordIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordIMM(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordIMME(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMME(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordIMMED(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMMED(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordIMMEDI(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMMEDI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordIMMEDIA(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMMEDIA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordIMMEDIAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMMEDIAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordIMMEDIATE(s)
	}
	return token.Unknown, false
}

func _scanKeywordIMMEDIATE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordImmediate, true
	}
	return token.Unknown, false
}

func _scanKeywordIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordIGN(s)
	}
	return token.Unknown, false
}

func _scanKeywordIGN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordIGNO(s)
	}
	return token.Unknown, false
}

func _scanKeywordIGNO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordIGNOR(s)
	}
	return token.Unknown, false
}

func _scanKeywordIGNOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordIGNORE(s)
	}
	return token.Unknown, false
}

func _scanKeywordIGNORE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordIgnore, true
	}
	return token.Unknown, false
}

func _scanKeywordIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordIn, true
	}
	switch next {
	case 'N':
		return _scanKeywordINN(s)
	case 'I':
		return _scanKeywordINI(s)
	case 'T':
		return _scanKeywordINT(s)
	case 'D':
		return _scanKeywordIND(s)
	case 'S':
		return _scanKeywordINS(s)
	}
	return token.Unknown, false
}

func _scanKeywordINI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordINIT(s)
	}
	return token.Unknown, false
}

func _scanKeywordINIT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordINITI(s)
	}
	return token.Unknown, false
}

func _scanKeywordINITI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordINITIA(s)
	}
	return token.Unknown, false
}

func _scanKeywordINITIA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordINITIAL(s)
	}
	return token.Unknown, false
}

func _scanKeywordINITIAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordINITIALL(s)
	}
	return token.Unknown, false
}

func _scanKeywordINITIALL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordINITIALLY(s)
	}
	return token.Unknown, false
}

func _scanKeywordINITIALLY(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordInitially, true
	}
	return token.Unknown, false
}

func _scanKeywordINT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordINTE(s)
	case 'O':
		return _scanKeywordINTO(s)
	}
	return token.Unknown, false
}

func _scanKeywordINTO(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordInto, true
	}
	return token.Unknown, false
}

func _scanKeywordINTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordINTER(s)
	}
	return token.Unknown, false
}

func _scanKeywordINTER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordINTERS(s)
	}
	return token.Unknown, false
}

func _scanKeywordINTERS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordINTERSE(s)
	}
	return token.Unknown, false
}

func _scanKeywordINTERSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordINTERSEC(s)
	}
	return token.Unknown, false
}

func _scanKeywordINTERSEC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordINTERSECT(s)
	}
	return token.Unknown, false
}

func _scanKeywordINTERSECT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordIntersect, true
	}
	return token.Unknown, false
}

func _scanKeywordIND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordINDE(s)
	}
	return token.Unknown, false
}

func _scanKeywordINDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'X':
		return _scanKeywordINDEX(s)
	}
	return token.Unknown, false
}

func _scanKeywordINDEX(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordIndex, true
	}
	switch next {
	case 'E':
		return _scanKeywordINDEXE(s)
	}
	return token.Unknown, false
}

func _scanKeywordINDEXE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordINDEXED(s)
	}
	return token.Unknown, false
}

func _scanKeywordINDEXED(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordIndexed, true
	}
	return token.Unknown, false
}

func _scanKeywordINS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordINST(s)
	case 'E':
		return _scanKeywordINSE(s)
	}
	return token.Unknown, false
}

func _scanKeywordINST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordINSTE(s)
	}
	return token.Unknown, false
}

func _scanKeywordINSTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordINSTEA(s)
	}
	return token.Unknown, false
}

func _scanKeywordINSTEA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordINSTEAD(s)
	}
	return token.Unknown, false
}

func _scanKeywordINSTEAD(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordInstead, true
	}
	return token.Unknown, false
}

func _scanKeywordINSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordINSER(s)
	}
	return token.Unknown, false
}

func _scanKeywordINSER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordINSERT(s)
	}
	return token.Unknown, false
}

func _scanKeywordINSERT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordInsert, true
	}
	return token.Unknown, false
}

func _scanKeywordINN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordINNE(s)
	}
	return token.Unknown, false
}

func _scanKeywordINNE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordINNER(s)
	}
	return token.Unknown, false
}

func _scanKeywordINNER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordInner, true
	}
	return token.Unknown, false
}

func _scanKeywordIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordIs, true
	}
	switch next {
	case 'N':
		return _scanKeywordISN(s)
	}
	return token.Unknown, false
}

func _scanKeywordISN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordISNU(s)
	}
	return token.Unknown, false
}

func _scanKeywordISNU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordISNUL(s)
	}
	return token.Unknown, false
}

func _scanKeywordISNUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordISNULL(s)
	}
	return token.Unknown, false
}

func _scanKeywordISNULL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordIsnull, true
	}
	return token.Unknown, false
}

func _scanKeywordO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordOR(s)
	case 'N':
		return _scanKeywordON(s)
	case 'F':
		return _scanKeywordOF(s)
	case 'V':
		return _scanKeywordOV(s)
	case 'T':
		return _scanKeywordOT(s)
	case 'U':
		return _scanKeywordOU(s)
	}
	return token.Unknown, false
}

func _scanKeywordOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordOr, true
	}
	switch next {
	case 'D':
		return _scanKeywordORD(s)
	}
	return token.Unknown, false
}

func _scanKeywordORD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordORDE(s)
	}
	return token.Unknown, false
}

func _scanKeywordORDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordORDER(s)
	}
	return token.Unknown, false
}

func _scanKeywordORDER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordOrder, true
	}
	return token.Unknown, false
}

func _scanKeywordON(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordOn, true
	}
	return token.Unknown, false
}

func _scanKeywordOF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordOf, true
	}
	switch next {
	case 'F':
		return _scanKeywordOFF(s)
	}
	return token.Unknown, false
}

func _scanKeywordOFF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordOFFS(s)
	}
	return token.Unknown, false
}

func _scanKeywordOFFS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordOFFSE(s)
	}
	return token.Unknown, false
}

func _scanKeywordOFFSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordOFFSET(s)
	}
	return token.Unknown, false
}

func _scanKeywordOFFSET(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordOffset, true
	}
	return token.Unknown, false
}

func _scanKeywordOV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordOVE(s)
	}
	return token.Unknown, false
}

func _scanKeywordOVE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordOVER(s)
	}
	return token.Unknown, false
}

func _scanKeywordOVER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordOver, true
	}
	return token.Unknown, false
}

func _scanKeywordOT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordOTH(s)
	}
	return token.Unknown, false
}

func _scanKeywordOTH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordOTHE(s)
	}
	return token.Unknown, false
}

func _scanKeywordOTHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordOTHER(s)
	}
	return token.Unknown, false
}

func _scanKeywordOTHER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordOTHERS(s)
	}
	return token.Unknown, false
}

func _scanKeywordOTHERS(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordOthers, true
	}
	return token.Unknown, false
}

func _scanKeywordOU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordOUT(s)
	}
	return token.Unknown, false
}

func _scanKeywordOUT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordOUTE(s)
	}
	return token.Unknown, false
}

func _scanKeywordOUTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordOUTER(s)
	}
	return token.Unknown, false
}

func _scanKeywordOUTER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordOuter, true
	}
	return token.Unknown, false
}

func _scanKeywordA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordAU(s)
	case 'C':
		return _scanKeywordAC(s)
	case 'B':
		return _scanKeywordAB(s)
	case 'T':
		return _scanKeywordAT(s)
	case 'F':
		return _scanKeywordAF(s)
	case 'S':
		return _scanKeywordAS(s)
	case 'N':
		return _scanKeywordAN(s)
	case 'L':
		return _scanKeywordAL(s)
	case 'D':
		return _scanKeywordAD(s)
	}
	return token.Unknown, false
}

func _scanKeywordAU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordAUT(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordAUTO(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordAUTOI(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordAUTOIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordAUTOINC(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordAUTOINCR(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINCR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordAUTOINCRE(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINCRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordAUTOINCREM(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINCREM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordAUTOINCREME(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINCREME(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordAUTOINCREMEN(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINCREMEN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordAUTOINCREMENT(s)
	}
	return token.Unknown, false
}

func _scanKeywordAUTOINCREMENT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAutoincrement, true
	}
	return token.Unknown, false
}

func _scanKeywordAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordACT(s)
	}
	return token.Unknown, false
}

func _scanKeywordACT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordACTI(s)
	}
	return token.Unknown, false
}

func _scanKeywordACTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordACTIO(s)
	}
	return token.Unknown, false
}

func _scanKeywordACTIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordACTION(s)
	}
	return token.Unknown, false
}

func _scanKeywordACTION(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAction, true
	}
	return token.Unknown, false
}

func _scanKeywordAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordABO(s)
	}
	return token.Unknown, false
}

func _scanKeywordABO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordABOR(s)
	}
	return token.Unknown, false
}

func _scanKeywordABOR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordABORT(s)
	}
	return token.Unknown, false
}

func _scanKeywordABORT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAbort, true
	}
	return token.Unknown, false
}

func _scanKeywordAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordATT(s)
	}
	return token.Unknown, false
}

func _scanKeywordATT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordATTA(s)
	}
	return token.Unknown, false
}

func _scanKeywordATTA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordATTAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordATTAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordATTACH(s)
	}
	return token.Unknown, false
}

func _scanKeywordATTACH(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAttach, true
	}
	return token.Unknown, false
}

func _scanKeywordAF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordAFT(s)
	}
	return token.Unknown, false
}

func _scanKeywordAFT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordAFTE(s)
	}
	return token.Unknown, false
}

func _scanKeywordAFTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordAFTER(s)
	}
	return token.Unknown, false
}

func _scanKeywordAFTER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAfter, true
	}
	return token.Unknown, false
}

func _scanKeywordAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordAs, true
	}
	switch next {
	case 'C':
		return _scanKeywordASC(s)
	}
	return token.Unknown, false
}

func _scanKeywordASC(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAsc, true
	}
	return token.Unknown, false
}

func _scanKeywordAN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordANA(s)
	case 'D':
		return _scanKeywordAND(s)
	}
	return token.Unknown, false
}

func _scanKeywordANA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordANAL(s)
	}
	return token.Unknown, false
}

func _scanKeywordANAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordANALY(s)
	}
	return token.Unknown, false
}

func _scanKeywordANALY(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Z':
		return _scanKeywordANALYZ(s)
	}
	return token.Unknown, false
}

func _scanKeywordANALYZ(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordANALYZE(s)
	}
	return token.Unknown, false
}

func _scanKeywordANALYZE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAnalyze, true
	}
	return token.Unknown, false
}

func _scanKeywordAND(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAnd, true
	}
	return token.Unknown, false
}

func _scanKeywordAL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordALT(s)
	case 'L':
		return _scanKeywordALL(s)
	}
	return token.Unknown, false
}

func _scanKeywordALT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordALTE(s)
	}
	return token.Unknown, false
}

func _scanKeywordALTE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordALTER(s)
	}
	return token.Unknown, false
}

func _scanKeywordALTER(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAlter, true
	}
	return token.Unknown, false
}

func _scanKeywordALL(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAll, true
	}
	return token.Unknown, false
}

func _scanKeywordAD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordADD(s)
	}
	return token.Unknown, false
}

func _scanKeywordADD(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordAdd, true
	}
	return token.Unknown, false
}

func _scanKeywordD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordDA(s)
	case 'E':
		return _scanKeywordDE(s)
	case 'I':
		return _scanKeywordDI(s)
	case 'O':
		return _scanKeywordDO(s)
	case 'R':
		return _scanKeywordDR(s)
	}
	return token.Unknown, false
}

func _scanKeywordDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordDES(s)
	case 'L':
		return _scanKeywordDEL(s)
	case 'F':
		return _scanKeywordDEF(s)
	case 'T':
		return _scanKeywordDET(s)
	}
	return token.Unknown, false
}

func _scanKeywordDET(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordDETA(s)
	}
	return token.Unknown, false
}

func _scanKeywordDETA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordDETAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordDETAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordDETACH(s)
	}
	return token.Unknown, false
}

func _scanKeywordDETACH(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDetach, true
	}
	return token.Unknown, false
}

func _scanKeywordDES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordDESC(s)
	}
	return token.Unknown, false
}

func _scanKeywordDESC(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDesc, true
	}
	return token.Unknown, false
}

func _scanKeywordDEL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordDELE(s)
	}
	return token.Unknown, false
}

func _scanKeywordDELE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordDELET(s)
	}
	return token.Unknown, false
}

func _scanKeywordDELET(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordDELETE(s)
	}
	return token.Unknown, false
}

func _scanKeywordDELETE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDelete, true
	}
	return token.Unknown, false
}

func _scanKeywordDEF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordDEFE(s)
	case 'A':
		return _scanKeywordDEFA(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordDEFER(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordDEFERR(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFERR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordDEFERRE(s)
	case 'A':
		return _scanKeywordDEFERRA(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFERRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'B':
		return _scanKeywordDEFERRAB(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFERRAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordDEFERRABL(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFERRABL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordDEFERRABLE(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFERRABLE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDeferrable, true
	}
	return token.Unknown, false
}

func _scanKeywordDEFERRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordDEFERRED(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFERRED(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDeferred, true
	}
	return token.Unknown, false
}

func _scanKeywordDEFA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordDEFAU(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFAU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordDEFAUL(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFAUL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordDEFAULT(s)
	}
	return token.Unknown, false
}

func _scanKeywordDEFAULT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDefault, true
	}
	return token.Unknown, false
}

func _scanKeywordDI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordDIS(s)
	}
	return token.Unknown, false
}

func _scanKeywordDIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordDIST(s)
	}
	return token.Unknown, false
}

func _scanKeywordDIST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordDISTI(s)
	}
	return token.Unknown, false
}

func _scanKeywordDISTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordDISTIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordDISTIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordDISTINC(s)
	}
	return token.Unknown, false
}

func _scanKeywordDISTINC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordDISTINCT(s)
	}
	return token.Unknown, false
}

func _scanKeywordDISTINCT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDistinct, true
	}
	return token.Unknown, false
}

func _scanKeywordDO(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDo, true
	}
	return token.Unknown, false
}

func _scanKeywordDR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordDRO(s)
	}
	return token.Unknown, false
}

func _scanKeywordDRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordDROP(s)
	}
	return token.Unknown, false
}

func _scanKeywordDROP(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDrop, true
	}
	return token.Unknown, false
}

func _scanKeywordDA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordDAT(s)
	}
	return token.Unknown, false
}

func _scanKeywordDAT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordDATA(s)
	}
	return token.Unknown, false
}

func _scanKeywordDATA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'B':
		return _scanKeywordDATAB(s)
	}
	return token.Unknown, false
}

func _scanKeywordDATAB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordDATABA(s)
	}
	return token.Unknown, false
}

func _scanKeywordDATABA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordDATABAS(s)
	}
	return token.Unknown, false
}

func _scanKeywordDATABAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordDATABASE(s)
	}
	return token.Unknown, false
}

func _scanKeywordDATABASE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordDatabase, true
	}
	return token.Unknown, false
}

func _scanKeywordE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordEN(s)
	case 'L':
		return _scanKeywordEL(s)
	case 'A':
		return _scanKeywordEA(s)
	case 'S':
		return _scanKeywordES(s)
	case 'X':
		return _scanKeywordEX(s)
	}
	return token.Unknown, false
}

func _scanKeywordEX(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordEXC(s)
	case 'I':
		return _scanKeywordEXI(s)
	case 'P':
		return _scanKeywordEXP(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordEXCL(s)
	case 'E':
		return _scanKeywordEXCE(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordEXCLU(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCLU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordEXCLUD(s)
	case 'S':
		return _scanKeywordEXCLUS(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCLUD(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordEXCLUDE(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCLUDE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordExclude, true
	}
	return token.Unknown, false
}

func _scanKeywordEXCLUS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordEXCLUSI(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCLUSI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'V':
		return _scanKeywordEXCLUSIV(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCLUSIV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordEXCLUSIVE(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCLUSIVE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordExclusive, true
	}
	return token.Unknown, false
}

func _scanKeywordEXCE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordEXCEP(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCEP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordEXCEPT(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXCEPT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordExcept, true
	}
	return token.Unknown, false
}

func _scanKeywordEXI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordEXIS(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordEXIST(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXIST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordEXISTS(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXISTS(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordExists, true
	}
	return token.Unknown, false
}

func _scanKeywordEXP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordEXPL(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXPL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordEXPLA(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXPLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordEXPLAI(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXPLAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordEXPLAIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordEXPLAIN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordExplain, true
	}
	return token.Unknown, false
}

func _scanKeywordEN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordEND(s)
	}
	return token.Unknown, false
}

func _scanKeywordEND(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordEnd, true
	}
	return token.Unknown, false
}

func _scanKeywordEL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordELS(s)
	}
	return token.Unknown, false
}

func _scanKeywordELS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordELSE(s)
	}
	return token.Unknown, false
}

func _scanKeywordELSE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordElse, true
	}
	return token.Unknown, false
}

func _scanKeywordEA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordEAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordEAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordEACH(s)
	}
	return token.Unknown, false
}

func _scanKeywordEACH(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordEach, true
	}
	return token.Unknown, false
}

func _scanKeywordES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordESC(s)
	}
	return token.Unknown, false
}

func _scanKeywordESC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordESCA(s)
	}
	return token.Unknown, false
}

func _scanKeywordESCA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordESCAP(s)
	}
	return token.Unknown, false
}

func _scanKeywordESCAP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordESCAPE(s)
	}
	return token.Unknown, false
}

func _scanKeywordESCAPE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordEscape, true
	}
	return token.Unknown, false
}

func _scanKeywordW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordWI(s)
	case 'H':
		return _scanKeywordWH(s)
	}
	return token.Unknown, false
}

func _scanKeywordWI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordWIN(s)
	case 'T':
		return _scanKeywordWIT(s)
	}
	return token.Unknown, false
}

func _scanKeywordWIT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordWITH(s)
	}
	return token.Unknown, false
}

func _scanKeywordWITH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordWith, true
	}
	switch next {
	case 'O':
		return _scanKeywordWITHO(s)
	}
	return token.Unknown, false
}

func _scanKeywordWITHO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordWITHOU(s)
	}
	return token.Unknown, false
}

func _scanKeywordWITHOU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordWITHOUT(s)
	}
	return token.Unknown, false
}

func _scanKeywordWITHOUT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordWithout, true
	}
	return token.Unknown, false
}

func _scanKeywordWIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordWIND(s)
	}
	return token.Unknown, false
}

func _scanKeywordWIND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordWINDO(s)
	}
	return token.Unknown, false
}

func _scanKeywordWINDO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'W':
		return _scanKeywordWINDOW(s)
	}
	return token.Unknown, false
}

func _scanKeywordWINDOW(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordWindow, true
	}
	return token.Unknown, false
}

func _scanKeywordWH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordWHE(s)
	}
	return token.Unknown, false
}

func _scanKeywordWHE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordWHEN(s)
	case 'R':
		return _scanKeywordWHER(s)
	}
	return token.Unknown, false
}

func _scanKeywordWHER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordWHERE(s)
	}
	return token.Unknown, false
}

func _scanKeywordWHERE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordWhere, true
	}
	return token.Unknown, false
}

func _scanKeywordWHEN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordWhen, true
	}
	return token.Unknown, false
}

func _scanKeywordS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordSA(s)
	case 'E':
		return _scanKeywordSE(s)
	}
	return token.Unknown, false
}

func _scanKeywordSA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'V':
		return _scanKeywordSAV(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordSAVE(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAVE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordSAVEP(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAVEP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordSAVEPO(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAVEPO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordSAVEPOI(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAVEPOI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordSAVEPOIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAVEPOIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordSAVEPOINT(s)
	}
	return token.Unknown, false
}

func _scanKeywordSAVEPOINT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordSavepoint, true
	}
	return token.Unknown, false
}

func _scanKeywordSE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordSEL(s)
	case 'T':
		return _scanKeywordSET(s)
	}
	return token.Unknown, false
}

func _scanKeywordSET(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordSet, true
	}
	return token.Unknown, false
}

func _scanKeywordSEL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordSELE(s)
	}
	return token.Unknown, false
}

func _scanKeywordSELE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordSELEC(s)
	}
	return token.Unknown, false
}

func _scanKeywordSELEC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordSELECT(s)
	}
	return token.Unknown, false
}

func _scanKeywordSELECT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordSelect, true
	}
	return token.Unknown, false
}

func _scanKeywordP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordPR(s)
	case 'L':
		return _scanKeywordPL(s)
	case 'A':
		return _scanKeywordPA(s)
	}
	return token.Unknown, false
}

func _scanKeywordPR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordPRE(s)
	case 'A':
		return _scanKeywordPRA(s)
	case 'I':
		return _scanKeywordPRI(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordPREC(s)
	}
	return token.Unknown, false
}

func _scanKeywordPREC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordPRECE(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRECE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordPRECED(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRECED(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordPRECEDI(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRECEDI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordPRECEDIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRECEDIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordPRECEDING(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRECEDING(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordPreceding, true
	}
	return token.Unknown, false
}

func _scanKeywordPRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordPRAG(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRAG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordPRAGM(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRAGM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordPRAGMA(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRAGMA(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordPragma, true
	}
	return token.Unknown, false
}

func _scanKeywordPRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordPRIM(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordPRIMA(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRIMA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordPRIMAR(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRIMAR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordPRIMARY(s)
	}
	return token.Unknown, false
}

func _scanKeywordPRIMARY(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordPrimary, true
	}
	return token.Unknown, false
}

func _scanKeywordPL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordPLA(s)
	}
	return token.Unknown, false
}

func _scanKeywordPLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordPLAN(s)
	}
	return token.Unknown, false
}

func _scanKeywordPLAN(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordPlan, true
	}
	return token.Unknown, false
}

func _scanKeywordPA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordPAR(s)
	}
	return token.Unknown, false
}

func _scanKeywordPAR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordPART(s)
	}
	return token.Unknown, false
}

func _scanKeywordPART(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordPARTI(s)
	}
	return token.Unknown, false
}

func _scanKeywordPARTI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordPARTIT(s)
	}
	return token.Unknown, false
}

func _scanKeywordPARTIT(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordPARTITI(s)
	}
	return token.Unknown, false
}

func _scanKeywordPARTITI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'O':
		return _scanKeywordPARTITIO(s)
	}
	return token.Unknown, false
}

func _scanKeywordPARTITIO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordPARTITION(s)
	}
	return token.Unknown, false
}

func _scanKeywordPARTITION(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordPartition, true
	}
	return token.Unknown, false
}

func _scanKeywordL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordLA(s)
	case 'E':
		return _scanKeywordLE(s)
	case 'I':
		return _scanKeywordLI(s)
	}
	return token.Unknown, false
}

func _scanKeywordLE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'F':
		return _scanKeywordLEF(s)
	}
	return token.Unknown, false
}

func _scanKeywordLEF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordLEFT(s)
	}
	return token.Unknown, false
}

func _scanKeywordLEFT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordLeft, true
	}
	return token.Unknown, false
}

func _scanKeywordLI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'K':
		return _scanKeywordLIK(s)
	case 'M':
		return _scanKeywordLIM(s)
	}
	return token.Unknown, false
}

func _scanKeywordLIM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordLIMI(s)
	}
	return token.Unknown, false
}

func _scanKeywordLIMI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordLIMIT(s)
	}
	return token.Unknown, false
}

func _scanKeywordLIMIT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordLimit, true
	}
	return token.Unknown, false
}

func _scanKeywordLIK(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordLIKE(s)
	}
	return token.Unknown, false
}

func _scanKeywordLIKE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordLike, true
	}
	return token.Unknown, false
}

func _scanKeywordLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordLAS(s)
	}
	return token.Unknown, false
}

func _scanKeywordLAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordLAST(s)
	}
	return token.Unknown, false
}

func _scanKeywordLAST(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordLast, true
	}
	return token.Unknown, false
}

func _scanKeywordR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordRA(s)
	case 'I':
		return _scanKeywordRI(s)
	case 'E':
		return _scanKeywordRE(s)
	case 'O':
		return _scanKeywordRO(s)
	}
	return token.Unknown, false
}

func _scanKeywordRE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordRES(s)
	case 'G':
		return _scanKeywordREG(s)
	case 'F':
		return _scanKeywordREF(s)
	case 'I':
		return _scanKeywordREI(s)
	case 'C':
		return _scanKeywordREC(s)
	case 'P':
		return _scanKeywordREP(s)
	case 'N':
		return _scanKeywordREN(s)
	case 'L':
		return _scanKeywordREL(s)
	}
	return token.Unknown, false
}

func _scanKeywordREN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordRENA(s)
	}
	return token.Unknown, false
}

func _scanKeywordRENA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'M':
		return _scanKeywordRENAM(s)
	}
	return token.Unknown, false
}

func _scanKeywordRENAM(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordRENAME(s)
	}
	return token.Unknown, false
}

func _scanKeywordRENAME(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRename, true
	}
	return token.Unknown, false
}

func _scanKeywordREL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordRELE(s)
	}
	return token.Unknown, false
}

func _scanKeywordRELE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordRELEA(s)
	}
	return token.Unknown, false
}

func _scanKeywordRELEA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordRELEAS(s)
	}
	return token.Unknown, false
}

func _scanKeywordRELEAS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordRELEASE(s)
	}
	return token.Unknown, false
}

func _scanKeywordRELEASE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRelease, true
	}
	return token.Unknown, false
}

func _scanKeywordRES(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordREST(s)
	}
	return token.Unknown, false
}

func _scanKeywordREST(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordRESTR(s)
	}
	return token.Unknown, false
}

func _scanKeywordRESTR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordRESTRI(s)
	}
	return token.Unknown, false
}

func _scanKeywordRESTRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordRESTRIC(s)
	}
	return token.Unknown, false
}

func _scanKeywordRESTRIC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordRESTRICT(s)
	}
	return token.Unknown, false
}

func _scanKeywordRESTRICT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRestrict, true
	}
	return token.Unknown, false
}

func _scanKeywordREG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordREGE(s)
	}
	return token.Unknown, false
}

func _scanKeywordREGE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'X':
		return _scanKeywordREGEX(s)
	}
	return token.Unknown, false
}

func _scanKeywordREGEX(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'P':
		return _scanKeywordREGEXP(s)
	}
	return token.Unknown, false
}

func _scanKeywordREGEXP(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRegexp, true
	}
	return token.Unknown, false
}

func _scanKeywordREF(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordREFE(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordREFER(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFER(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordREFERE(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFERE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordREFEREN(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFEREN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordREFERENC(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFERENC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordREFERENCE(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFERENCE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordREFERENCES(s)
	}
	return token.Unknown, false
}

func _scanKeywordREFERENCES(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordReferences, true
	}
	return token.Unknown, false
}

func _scanKeywordREI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordREIN(s)
	}
	return token.Unknown, false
}

func _scanKeywordREIN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'D':
		return _scanKeywordREIND(s)
	}
	return token.Unknown, false
}

func _scanKeywordREIND(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordREINDE(s)
	}
	return token.Unknown, false
}

func _scanKeywordREINDE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'X':
		return _scanKeywordREINDEX(s)
	}
	return token.Unknown, false
}

func _scanKeywordREINDEX(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordReindex, true
	}
	return token.Unknown, false
}

func _scanKeywordREC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'U':
		return _scanKeywordRECU(s)
	}
	return token.Unknown, false
}

func _scanKeywordRECU(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'R':
		return _scanKeywordRECUR(s)
	}
	return token.Unknown, false
}

func _scanKeywordRECUR(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordRECURS(s)
	}
	return token.Unknown, false
}

func _scanKeywordRECURS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'I':
		return _scanKeywordRECURSI(s)
	}
	return token.Unknown, false
}

func _scanKeywordRECURSI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'V':
		return _scanKeywordRECURSIV(s)
	}
	return token.Unknown, false
}

func _scanKeywordRECURSIV(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordRECURSIVE(s)
	}
	return token.Unknown, false
}

func _scanKeywordRECURSIVE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRecursive, true
	}
	return token.Unknown, false
}

func _scanKeywordREP(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordREPL(s)
	}
	return token.Unknown, false
}

func _scanKeywordREPL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordREPLA(s)
	}
	return token.Unknown, false
}

func _scanKeywordREPLA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordREPLAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordREPLAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordREPLACE(s)
	}
	return token.Unknown, false
}

func _scanKeywordREPLACE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordReplace, true
	}
	return token.Unknown, false
}

func _scanKeywordRO(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'W':
		return _scanKeywordROW(s)
	case 'L':
		return _scanKeywordROL(s)
	}
	return token.Unknown, false
}

func _scanKeywordROW(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.KeywordRow, true
	}
	switch next {
	case 'S':
		return _scanKeywordROWS(s)
	}
	return token.Unknown, false
}

func _scanKeywordROWS(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRows, true
	}
	return token.Unknown, false
}

func _scanKeywordROL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'L':
		return _scanKeywordROLL(s)
	}
	return token.Unknown, false
}

func _scanKeywordROLL(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'B':
		return _scanKeywordROLLB(s)
	}
	return token.Unknown, false
}

func _scanKeywordROLLB(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'A':
		return _scanKeywordROLLBA(s)
	}
	return token.Unknown, false
}

func _scanKeywordROLLBA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'C':
		return _scanKeywordROLLBAC(s)
	}
	return token.Unknown, false
}

func _scanKeywordROLLBAC(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'K':
		return _scanKeywordROLLBACK(s)
	}
	return token.Unknown, false
}

func _scanKeywordROLLBACK(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRollback, true
	}
	return token.Unknown, false
}

func _scanKeywordRA(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'N':
		return _scanKeywordRAN(s)
	case 'I':
		return _scanKeywordRAI(s)
	}
	return token.Unknown, false
}

func _scanKeywordRAN(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordRANG(s)
	}
	return token.Unknown, false
}

func _scanKeywordRANG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordRANGE(s)
	}
	return token.Unknown, false
}

func _scanKeywordRANGE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRange, true
	}
	return token.Unknown, false
}

func _scanKeywordRAI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'S':
		return _scanKeywordRAIS(s)
	}
	return token.Unknown, false
}

func _scanKeywordRAIS(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordRAISE(s)
	}
	return token.Unknown, false
}

func _scanKeywordRAISE(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRaise, true
	}
	return token.Unknown, false
}

func _scanKeywordRI(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'G':
		return _scanKeywordRIG(s)
	}
	return token.Unknown, false
}

func _scanKeywordRIG(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'H':
		return _scanKeywordRIGH(s)
	}
	return token.Unknown, false
}

func _scanKeywordRIGH(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'T':
		return _scanKeywordRIGHT(s)
	}
	return token.Unknown, false
}

func _scanKeywordRIGHT(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordRight, true
	}
	return token.Unknown, false
}

func _scanKeywordK(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'E':
		return _scanKeywordKE(s)
	}
	return token.Unknown, false
}

func _scanKeywordKE(s RuneScanner) (token.Type, bool) {
	next, ok := s.Lookahead()
	if !ok {
		return token.Unknown, false
	}
	switch next {
	case 'Y':
		return _scanKeywordKEY(s)
	}
	return token.Unknown, false
}

func _scanKeywordKEY(s RuneScanner) (token.Type, bool) {
	_, ok := s.Lookahead()
	if !ok {
		return token.KeywordKey, true
	}
	return token.Unknown, false
}
