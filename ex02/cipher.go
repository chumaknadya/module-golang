package cipher

import "strings"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

type shift int

func NewShift(distance int) Cipher {
	key := shift(distance)
	switch {
	case key > 0 && key < 26:
		return key
	case key > -26 && key < 0:
		key += 26
		return key
	}
	return nil
}

func enc(val rune, shift int) rune {
	if val >= 'a' && val <= 'z' {
		return (val-'a'+rune(shift))%26 + 'a'
	}
	return -1
}

func dec(val rune, shift int) rune {
	if val >= 'a' && val <= 'z' {
		return (val-'a'+rune(26-shift))%26 + 'a'
	}
	return -1
}

func (sh shift) Encode(s string) string {
	downcaseStr := Downcase(s)
	return strings.Map(
		func(r rune) rune { return enc(r, int(sh)) }, downcaseStr)
}

func (sh shift) Decode(s string) string {
	downcaseStr := Downcase(s)
	return strings.Map(
		func(r rune) rune { return dec(r, int(sh)) },
		downcaseStr)
}

type vigenere string

func NewVigenere(key string) Cipher {
	flag := false

	for _, val := range key {
		if val < 'a' || val > 'z' {
			return nil
		}
		if val > 'a' {
			flag = true
		}
	}

	if !flag {
		return nil
	}
	return vigenere(key)
}

func Downcase(initStr string) string {
	arr := make([]byte, len(initStr))
	for i := range arr {
		letter := initStr[i]
		if letter >= 'A' && letter <= 'Z' {
			letter += 'a' - 'A'
		}
		arr[i] = letter
	}
	return string(arr)
}

func (v vigenere) Encode(s string) string {
	downcaseStr := Downcase(s)
	k := 0
	return strings.Map(
		func(r rune) rune {
			if r = enc(r, int(v[k]-'a')); r >= 0 {
				k = (k + 1) % len(v)
			}
			return r
		}, downcaseStr)
}

func (v vigenere) Decode(s string) string {
	downcaseStr := Downcase(s)
	k := 0
	return strings.Map(
		func(r rune) rune {
			if r = dec(r, int(v[k]-'a')); r >= 0 {
				k = (k + 1) % len(v)
			}
			return r
		}, downcaseStr)
}
