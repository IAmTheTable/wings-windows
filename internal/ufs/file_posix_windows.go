//go:build windows

package ufs

func ignoringEINTR(fn func() error) error {
	return fn()
}

func syscallMode(i FileMode) (o FileMode) {
	return i.Perm()
}
