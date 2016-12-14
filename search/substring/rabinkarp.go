// Implemting a substring search using Rabin-Karp algorithm
// Usage:
// rk := NewRabinKarp("ex")
// rk.search("my text") returns the index of the first match

package substring

import (
	"crypto/rand"
)

type RabinKarp struct {
	PatternHash int32
	PatternLength int
	Radix int32 // 256 for extended ASCII strings
	Mod int32 // BigPrime
	BasePowered int32 // Base^(PatternLength - 1) % Modulo
}

// Precompute a few attributes
func NewRabinKarp(pattern string) *RabinKarp {
	rk := RabinKarp{
		Radix: int32(256),
		Mod: bigPrime(),
		PatternLength: len(pattern)}

	// Precomputing for efficiency
	basePowered := int32(1)
	for i := 1; i < rk.PatternLength; i++ {
		basePowered = (basePowered * rk.Radix) % rk.Mod
	}

	rk.BasePowered = basePowered
	rk.PatternHash = rk.hash(pattern)

	return &rk
}

// Returns the index of the first match
// Monte-Carlo version: there might be a hash collision
func (rk *RabinKarp) search(input string) int {
	start := input[0:rk.PatternLength]
	h := rk.hash(string(start))
	if h == rk.PatternHash {
		return 0;
	}
	for i := rk.PatternLength; i < len(input); i++ {
		h = (h + rk.Mod - rk.BasePowered * int32(input[i-rk.PatternLength]) % rk.Mod ) % rk.Mod
		h = (rk.Radix * h + int32(input[i]) ) % rk.Mod
		if h == rk.PatternHash {
			return i - rk.PatternLength + 1
		}
	}
	return -1;
}

func (rk *RabinKarp) hash(key string) int32 {
	res := int32(0)
	for _, runeValue := range key {
		res = (runeValue + rk.Radix * res) % rk.Mod
	}

	return res
}

func bigPrime() int32 {
	p, _ := rand.Prime(rand.Reader, 31)
	return int32(p.Int64())
}
