//go:build !linux

package main

func switchDataPathV2() bool {
	return true
}

func checkKernelVersion(k, major, minor int) bool {
	return false
}
