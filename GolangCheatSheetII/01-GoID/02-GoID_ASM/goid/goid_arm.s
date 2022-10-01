#include "textflag.h"

TEXT ·ID(SB), NOSPLIT, $0-8
	MOVD	0x98(R14), F1 // F1 = g.goid
	MOVD	F1, ret+0(FP) // ret F1
	RET
