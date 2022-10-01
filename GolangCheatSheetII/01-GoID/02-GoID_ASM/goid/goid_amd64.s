#include "textflag.h"

TEXT ·ID(SB), NOSPLIT, $0-8
	MOVQ	(TLS), AX     // AX = getg()
	MOVQ	0x98(AX), AX   // AX = AX.goid
	MOVQ	AX, ret+0(FP) // ret = AX
	RET
