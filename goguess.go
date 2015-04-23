package goguess

import "bytes"

type BruteForceSetup struct {
	Letters   []rune
	MinLength int
	MaxLength int
}

func NewState(fuzzer *BruteForceSetup) []int {
	return make([]int, fuzzer.MinLength)
}

func HasNext(state []int, fuzzer *BruteForceSetup) bool {
	cMax := len(fuzzer.Letters) - 1

	if len(state) < fuzzer.MaxLength {
		return true
	}

	for _, v := range state {
		if v < cMax {
			return true
		}
	}
	return false
}

func MoveNext(state []int, fuzzer *BruteForceSetup) []int {
	cCount := len(fuzzer.Letters)
	cMax := cCount - 1
	carry := false

	for i := range state {
		if state[i] < cMax {
			state[i] += 1
			return state
		} else {
			state[i] = 0
			carry = true
		}
	}

	if carry && len(state) < fuzzer.MaxLength {
		for i := range state {
			state[i] = 0
		}
		state = append(state, 0)
	} else {
		for i := range state {
			state[i] = cMax
		}
	}
	return state
}

func StateToRawByte(state []int, fuzzer *BruteForceSetup) []byte {
	w := &bytes.Buffer{}
	for _, v := range state {
		w.WriteByte(byte(v))
	}
	return w.Bytes()
}

func StateToRuneBytes(state []int, fuzzer *BruteForceSetup) []byte {
	w := &bytes.Buffer{}
	for _, v := range state {
		w.WriteRune(fuzzer.Letters[v])
	}
	return w.Bytes()
}

func StateToString(state []int, fuzzer *BruteForceSetup) string {
	w := &bytes.Buffer{}
	for _, v := range state {
		w.WriteRune(fuzzer.Letters[v])
	}
	return w.String()
}
