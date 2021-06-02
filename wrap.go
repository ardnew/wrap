package wrap

import (
	"strings"
	"unicode"
)

// Next scans the given string, starting at the given byte offset, and returns
// the first substring surrounded by whitespace (or string boundaries) and the
// number of runes in that substring.
//
// Both the returned string and the rune count include any leading whitespace
// if present.
//
// If the word is empty, byte offset is outside of string bounds, or the word
// consists only of whitespace runes, returns an empty string and 0.
func Next(s string, i int) (string, int) {
	b := strings.Builder{}
	n := 0
	inWord := false
	if i < len(s) {
		for _, c := range s[i:] {
			if inWord {
				if unicode.IsSpace(c) {
					break
				}
			} else {
				if unicode.IsSpace(c) {
					if strings.ContainsRune("\r\n\v\f", c) {
						// replace linebreaks with a single space
						c = ' '
					}
				} else {
					inWord = true
				}
			}
			b.WriteRune(c)
			n += 1
		}
	}
	if inWord {
		return b.String(), n
	} else {
		return "", 0
	}
}

// split returns a slice of each w-length substring in s, and any trailing
// string with length less than w. This does not take word-boundaries into
// account.
// Each substring (except for the trailing) is appended with h.
func split(s string, w int, h string) ([]string, string) {
	if w >= len(s) || w < 1 {
		return nil, s
	}
	var a []string
	var b strings.Builder
	w -= len(h)
	l, n := 0, 0
	b.Grow(w)
	for _, r := range s {
		b.WriteRune(r)
		l++
		n++
		if l == w && n < len([]rune(s))-1 {
			a = append(a, b.String()+h)
			l = 0
			b.Reset()
			b.Grow(w)
		}
	}
	r := ""
	if l > 0 {
		r = b.String()
	}
	return a, r
}

// String splits a string into substrings of length w, taking word boundaries
// into consideration, and joins any words whose length > w with h.
func String(s string, w int, h string) []string {
	cb := strings.Builder{}
	rs := []string{}
	rn := 0
	sn := 0
	for sn < len(s) {
		cs, cn := Next(s, sn)
		var ws []string
		var wr string
		if 0 == cn {
			break
		}
		if cb.Len() > 0 {
			if rn+cn >= w {
				ws, wr = split(cb.String(), w, h)
				rs = append(rs, ws...)
				rn = 0
				cb.Reset()
				// remove leading whitespace at beginning of row
				cs = strings.TrimSpace(wr + cs)
			}
		}
		sn += cn
		rn += len(cs)
		cb.WriteString(cs)
	}
	// Add leftover string remaining in buffer.
	if cb.Len() > 0 {
		ws, wr := split(cb.String(), w, h)
		rs = append(rs, ws...)
		rs = append(rs, wr)
	}
	return rs
}
