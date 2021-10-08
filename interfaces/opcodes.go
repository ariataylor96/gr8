package interfaces

type opcode interface {
	Op() byte

	CLS()
	RET()
	JP(addr uint16)
	CALL(addr uint16)

	SE(register, val uint8)
	SNE(register, val uint8)
	RSE(x, y uint8)
	LD(register, val uint8)

	ADD(register, val uint8)
	LDR(x, y uint8)
	ORR(x, y uint8)
	ANDR(x, y uint8)

	XORR(x, y uint8)
	ADDCR(x, y uint8)
	SUBCR(x, y uint8)
	SHR(register uint8)

	SUBNCR(x, y uint8)
	SHL(register uint8)
	SNER(x, y uint8)
	LDI(addr uint16)

	JPV0(addr uint16)
	RNDVX(register, val uint8)
	DRW(x, y, height uint8)
	SKPVX(x uint8)

	SKNPVX(x uint8)
	LDVXDT(x uint8)
	LDVXK(x uint8)
	LDDTVX(x uint8)

	LDSTVX(x uint8)
	ADDIVX(x uint8)
	LDFVX(x uint8)
	LDBVX(x uint8)

	LDIVX(x uint8)
	LDVXI(x uint8)
}
