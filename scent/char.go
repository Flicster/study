package scent

import (
	"strconv"
	"strings"
)

const (
	zeroWidthSpace     = "\u200B"
	zeroWidthNonJoiner = "\u200C"
	runeLen            = 3
)

var (
	encodeZeroWidthBinaryMap = map[rune]string{
		'0': zeroWidthSpace,
		'1': zeroWidthNonJoiner,
	}
	decodeZeroWidthBinaryMap = map[string]string{
		zeroWidthSpace:     "0",
		zeroWidthNonJoiner: "1",
	}
)

// EncodeWithZeroWidth encodes a receiver ID into the data using zero-width characters.
func EncodeWithZeroWidth(data string, receiverID int64) string {
	binaryID := strconv.FormatInt(receiverID, 2)

	var encodedID strings.Builder
	for _, bit := range binaryID {
		encodedID.WriteString(encodeZeroWidthBinaryMap[bit])
	}
	return data + encodedID.String()
}

// DecodeWithZeroWidth decodes the receiver ID from the data containing zero-width characters.
func DecodeWithZeroWidth(data string) (int64, error) {
	var binaryID strings.Builder
	for i := len(data) - runeLen; i >= 0; i -= runeLen {
		char := data[i : i+runeLen]
		binaryDigit, ok := decodeZeroWidthBinaryMap[char]
		if !ok {
			break
		}
		binaryID.WriteString(binaryDigit)
	}
	receiverID, err := strconv.ParseInt(reverseString(binaryID.String()), 2, 64)
	if err != nil {
		return 0, err
	}
	return receiverID, nil
}

// reverseString reverses a string.
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
