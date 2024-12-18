package vietqr

import (
	"fmt"
	"strconv"
	"strings"
)

func hashCrc(s string) string {
	h := NewCrc16(CRC16_CCITT_FALSE)
	h.Write([]byte(s))

	return paddingString(fmt.Sprintf("%X", h.Sum(nil)), paddingCrc)
}

func validCrcContent(s string) bool {
	checkContent := s[:len(s)-paddingCrc]
	crcCode := strings.ToUpper(s[len(s)-paddingCrc:])

	genCrcCode := hashCrc(checkContent)
	return crcCode == genCrcCode
}

// returns id (2 bytes) + length (2 bytes) + value (length bytes)
func genFieldData(id, value string) string {
	if len(id) != 2 || len(value) <= 0 {
		return ""
	}

	return joinString(id, paddingNumber(len(value), 2), value)
}

func joinString(ss ...string) string {
	return strings.Join(ss, "")
}

func paddingNumber(n, fl int) string {
	return paddingString(strconv.Itoa(n), fl)
}

func paddingString(s string, fl int) string {
	if fl <= 0 {
		return s
	}

	if len(s) >= fl {
		return s[len(s)-fl:]
	}

	for len(s) < fl {
		s = "0" + s
	}
	return s
}

// slideContent parses a content string and returns the id, value and nextContent.
// id (2 bytes) + length (2 bytes) + value (length bytes) + nextContent (rest of string)
func slideContent(s string) (id string, value string, nextContent string) {
	id = s[:2]
	length, _ := strconv.Atoi(s[2:4])
	value = s[4 : 4+length]
	nextContent = s[4+length:]
	return
}
