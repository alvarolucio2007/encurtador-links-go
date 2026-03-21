// Package shortener faz a conversão pra base 62
package shortener

import (
	"slices"
	"strings"
)

const alfabeto string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(id uint64) string {
	if id == 0 {
		return "0"
	}
	var res []byte
	for id > 0 {
		resto := id % 62
		res = append(res, byte(alfabeto[resto]))
		id = id / 62
	}
	slices.Reverse(res)
	return string(res)
}

func Decode(source string) uint64 {
	var result uint64 = 0
	for index := range len(source) {
		posicao := strings.Index(alfabeto, string(source[index]))
		result *= 62
		result += uint64(posicao)
	}
	return result
}
