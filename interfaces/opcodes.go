package interfaces

type opcodes interface {
	Op() byte

	CLS()
	RET()
	JP(addr uint16)
	CALL(addr uint16)

	SE(register, val byte)
	SNE(register, val byte)
	RSE(x, y byte)
	LD(register, val byte)

	ADD(register, val byte)
	LDR(x, y byte)
	ORR(x, y byte)
	ANDR(x, y byte)

	XORR(x, y byte)
	ADDCR(x, y byte)
	SUBCR(x, y byte)
	SHR(register byte)

	SUBNCR(x, y byte)
	SHL(register byte)
	SNER(x, y byte)
	LDI(addr uint16)

	JPV0(addr uint16)
	RNDVX(register, val byte)
	DRW(x, y, height byte)
	SKPVX(x byte)

	SKNPVX(x byte)
	LDVXDT(x byte)
	LDVXK(x byte)
	LDDTVX(x byte)

	LDSTVX(x byte)
	ADDIVX(x byte)
	LDFVX(x byte)
	LDBVX(x byte)

	LDIVX(x byte)
	LDVXI(x byte)
}
