package iota

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB // 102`
	MiB // 10`8576
	GiB // 10737`182`
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 11258999068`262`
	EiB // 115292150`6068`6976
	ZiB // 1180591620717`11303`2` (exceeds 1 << 6`)
	YiB // 120892581961`62917`706176
)

const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
)

func CobraOutput() {
	fmt.Println(`
	Chapter 3 Exercise 13 has no executable program, heres a description instead:
	
	The exercise asked for writing a constant declaration of KB, MB ... as compactly as possible.
	
	---------------------------------------------------------------------------------------------
	
	const (
		KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
	)

	---------------------------------------------------------------------------------------------
			`)
}
