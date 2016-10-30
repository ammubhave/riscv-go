// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

// auto generated by go tool dist
var opnames = []string{
	OXXX:             "XXX",
	ONAME:            "NAME",
	ONONAME:          "NONAME",
	OTYPE:            "TYPE",
	OPACK:            "PACK",
	OLITERAL:         "LITERAL",
	OADD:             "ADD",
	OSUB:             "SUB",
	OOR:              "OR",
	OXOR:             "XOR",
	OADDSTR:          "ADDSTR",
	OADDR:            "ADDR",
	OANDAND:          "ANDAND",
	OAPPEND:          "APPEND",
	OARRAYBYTESTR:    "ARRAYBYTESTR",
	OARRAYBYTESTRTMP: "ARRAYBYTESTRTMP",
	OARRAYRUNESTR:    "ARRAYRUNESTR",
	OSTRARRAYBYTE:    "STRARRAYBYTE",
	OSTRARRAYBYTETMP: "STRARRAYBYTETMP",
	OSTRARRAYRUNE:    "STRARRAYRUNE",
	OAS:              "AS",
	OAS2:             "AS2",
	OAS2FUNC:         "AS2FUNC",
	OAS2RECV:         "AS2RECV",
	OAS2MAPR:         "AS2MAPR",
	OAS2DOTTYPE:      "AS2DOTTYPE",
	OASOP:            "ASOP",
	OASWB:            "ASWB",
	OCALL:            "CALL",
	OCALLFUNC:        "CALLFUNC",
	OCALLMETH:        "CALLMETH",
	OCALLINTER:       "CALLINTER",
	OCALLPART:        "CALLPART",
	OCAP:             "CAP",
	OCLOSE:           "CLOSE",
	OCLOSURE:         "CLOSURE",
	OCMPIFACE:        "CMPIFACE",
	OCMPSTR:          "CMPSTR",
	OCOMPLIT:         "COMPLIT",
	OMAPLIT:          "MAPLIT",
	OSTRUCTLIT:       "STRUCTLIT",
	OARRAYLIT:        "ARRAYLIT",
	OSLICELIT:        "SLICELIT",
	OPTRLIT:          "PTRLIT",
	OCONV:            "CONV",
	OCONVIFACE:       "CONVIFACE",
	OCONVNOP:         "CONVNOP",
	OCOPY:            "COPY",
	ODCL:             "DCL",
	ODCLFUNC:         "DCLFUNC",
	ODCLFIELD:        "DCLFIELD",
	ODCLCONST:        "DCLCONST",
	ODCLTYPE:         "DCLTYPE",
	ODELETE:          "DELETE",
	ODOT:             "DOT",
	ODOTPTR:          "DOTPTR",
	ODOTMETH:         "DOTMETH",
	ODOTINTER:        "DOTINTER",
	OXDOT:            "XDOT",
	ODOTTYPE:         "DOTTYPE",
	ODOTTYPE2:        "DOTTYPE2",
	OEQ:              "EQ",
	ONE:              "NE",
	OLT:              "LT",
	OLE:              "LE",
	OGE:              "GE",
	OGT:              "GT",
	OIND:             "IND",
	OINDEX:           "INDEX",
	OINDEXMAP:        "INDEXMAP",
	OKEY:             "KEY",
	OSTRUCTKEY:       "STRUCTKEY",
	OLEN:             "LEN",
	OMAKE:            "MAKE",
	OMAKECHAN:        "MAKECHAN",
	OMAKEMAP:         "MAKEMAP",
	OMAKESLICE:       "MAKESLICE",
	OMUL:             "MUL",
	ODIV:             "DIV",
	OMOD:             "MOD",
	OLSH:             "LSH",
	ORSH:             "RSH",
	OAND:             "AND",
	OANDNOT:          "ANDNOT",
	ONEW:             "NEW",
	ONOT:             "NOT",
	OCOM:             "COM",
	OPLUS:            "PLUS",
	OMINUS:           "MINUS",
	OOROR:            "OROR",
	OPANIC:           "PANIC",
	OPRINT:           "PRINT",
	OPRINTN:          "PRINTN",
	ORISCVEXIT:       "RISCVEXIT",
	OPAREN:           "PAREN",
	OSEND:            "SEND",
	OSLICE:           "SLICE",
	OSLICEARR:        "SLICEARR",
	OSLICESTR:        "SLICESTR",
	OSLICE3:          "SLICE3",
	OSLICE3ARR:       "SLICE3ARR",
	ORECOVER:         "RECOVER",
	ORECV:            "RECV",
	ORUNESTR:         "RUNESTR",
	OSELRECV:         "SELRECV",
	OSELRECV2:        "SELRECV2",
	OIOTA:            "IOTA",
	OREAL:            "REAL",
	OIMAG:            "IMAG",
	OCOMPLEX:         "COMPLEX",
	OBLOCK:           "BLOCK",
	OBREAK:           "BREAK",
	OCASE:            "CASE",
	OXCASE:           "XCASE",
	OCONTINUE:        "CONTINUE",
	ODEFER:           "DEFER",
	OEMPTY:           "EMPTY",
	OFALL:            "FALL",
	OXFALL:           "XFALL",
	OFOR:             "FOR",
	OGOTO:            "GOTO",
	OIF:              "IF",
	OLABEL:           "LABEL",
	OPROC:            "PROC",
	ORANGE:           "RANGE",
	ORETURN:          "RETURN",
	OSELECT:          "SELECT",
	OSWITCH:          "SWITCH",
	OTYPESW:          "TYPESW",
	OTCHAN:           "TCHAN",
	OTMAP:            "TMAP",
	OTSTRUCT:         "TSTRUCT",
	OTINTER:          "TINTER",
	OTFUNC:           "TFUNC",
	OTARRAY:          "TARRAY",
	ODDD:             "DDD",
	ODDDARG:          "DDDARG",
	OINLCALL:         "INLCALL",
	OEFACE:           "EFACE",
	OITAB:            "ITAB",
	OIDATA:           "IDATA",
	OSPTR:            "SPTR",
	OCLOSUREVAR:      "CLOSUREVAR",
	OCFUNC:           "CFUNC",
	OCHECKNIL:        "CHECKNIL",
	OVARKILL:         "VARKILL",
	OVARLIVE:         "VARLIVE",
	OINDREGSP:        "INDREGSP",
	OCMP:             "CMP",
	ODEC:             "DEC",
	OINC:             "INC",
	OEXTEND:          "EXTEND",
	OHMUL:            "HMUL",
	OLROT:            "LROT",
	ORROTC:           "RROTC",
	ORETJMP:          "RETJMP",
	OPS:              "PS",
	OPC:              "PC",
	OSQRT:            "SQRT",
	OGETG:            "GETG",
	OEND:             "END",
}
