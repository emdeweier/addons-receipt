package utils

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func CutStringReceiptV2(req string, length int, height int, cutDot bool) []string {
	var arrayString []string
	if height > 1 {
		height = 1
	}
	lenOfString := len(strings.Split(req, ""))
	longLoopString := 0
	if lenOfString > length {
		d := float64(lenOfString) / float64(length)
		longLoopString = int(math.Ceil(d))
	}

	format := fmt.Sprintf("%s%d%s", "%.", length, "s")
	if length <= 0 {
		format = fmt.Sprintf("%s%s", "%.", "s")
	}

	log.Printf("lenOfString: %d", lenOfString)
	log.Printf("longLoopString: %d | %s", longLoopString, req)

	if longLoopString > 1 {
		for i := 0; i < longLoopString; i++ {
			log.Println("i: ", i)
			log.Println("longLoopString-1: ", longLoopString-1)
			arrayString = append(arrayString, fmt.Sprintf(format, req))
			log.Println("arrayString: ", arrayString)
			req = strings.Replace(req, fmt.Sprintf(format, req), "", 1)
			req = strings.TrimLeft(req, " ")
			if height > 0 {
				if i == longLoopString-(longLoopString-1) {
					if cutDot {
						lenArray := len(arrayString) - 1
						log.Println("lenArray: ", lenArray)
						lastWord := arrayString[lenArray]
						log.Println("lastWord: ", lastWord)
						lenLastWord := len(lastWord)
						log.Println("lenLastWord: ", lenLastWord)
						if lenLastWord > length-1 {
							lastIndex := strings.LastIndex(lastWord, lastWord[lenLastWord-3:lenLastWord])
							threeLastWord := lastWord[lenLastWord-3 : lenLastWord]
							arrayString[lenArray] = lastWord[:lastIndex] + strings.Replace(lastWord[lastIndex:], threeLastWord, "...", 1)
						}
					}
				}
			} else {
				if i == longLoopString-longLoopString {
					if cutDot {
						lenArray := len(arrayString) - 1
						log.Println("lenArray: ", lenArray)
						lastWord := arrayString[lenArray]
						log.Println("lastWord: ", lastWord)
						lenLastWord := len(lastWord)
						log.Println("lenLastWord: ", lenLastWord)
						if lenLastWord > length-1 {
							lastIndex := strings.LastIndex(lastWord, lastWord[lenLastWord-3:lenLastWord])
							threeLastWord := lastWord[lenLastWord-3 : lenLastWord]
							arrayString[lenArray] = lastWord[:lastIndex] + strings.Replace(lastWord[lastIndex:], threeLastWord, "...", 1)
						}
					}
				}
			}
			if i > height-1 {
				break
			}
		}
	} else {
		arrayString = append(arrayString, req)
	}

	return arrayString
}

func NewLineWord(word string, maxLength int, lengthWord int, height int) (result []string) {

	for i := 0; i < lengthWord; i += maxLength {

		if i+maxLength > lengthWord {
			result = append(result, word[i:])
			return
		}

		result = append(result, word[i:i+(maxLength)])

		if height > 0 {
			if len(result) >= height {
				return
			}
		}

	}

	return
}

// this param word must be contains maxlength character
// if the text not true value return add three dot
func CutDotWord(word string, maxLength int) []string {
	lengthWord := len(word)

	if lengthWord < maxLength {
		return []string{word + "..."}
	}

	word = word[:maxLength-1]

	return []string{word + "..."}
}

// this function implement nextline/neter every maxlength
func TextStructuring(req string, maxLength int, height int, cutDot bool) (result []string) {

	arrString := strings.Split(req, " ")

	if len(arrString) == 0 {
		return
	}

	// TODO if text just have one word
	if len(arrString) == 1 {

		word := arrString[0]
		lengthWord := len(word)

		if lengthWord <= maxLength {
			return arrString
		}

		if cutDot {
			return CutDotWord(arrString[0], maxLength)
		}

		return NewLineWord(word, maxLength, lengthWord, height)
	}

	var tempArr []string
	var tempNumberCharacter int
	for _, v := range arrString {

		if len(result) >= height {
			goto FINISH
		}

		lenV := len(v)

		if len(tempArr) > 0 {
			numberCharacter := (len(tempArr) - 1) + tempNumberCharacter + lenV

			if numberCharacter > maxLength {

				if cutDot {
					return CutDotWord(strings.Join(tempArr, " "), maxLength)
				}

				result = append(result, strings.Join(tempArr, " "))
				tempArr = []string{}
				tempNumberCharacter = 0
			}
		}

		if lenV > maxLength {
			newLineWordRes := NewLineWord(v, maxLength, lenV, height)

			for _, v2 := range newLineWordRes[:len(newLineWordRes)-1] {
				if len(result) >= height {
					goto FINISH
				}
				result = append(result, v2)
			}

			if len(result) >= height {
				goto FINISH
			}

			lenLastnewLineWordRes := len(newLineWordRes[len(newLineWordRes)-1])
			if lenLastnewLineWordRes == maxLength {
				result = append(result, newLineWordRes[len(newLineWordRes)-1])
				continue
			}

			tempNumberCharacter += lenLastnewLineWordRes
			tempArr = append(tempArr, newLineWordRes[len(newLineWordRes)-1])
			continue
		}

		tempNumberCharacter += lenV
		tempArr = append(tempArr, v)
	}

	result = append(result, strings.Join(tempArr, " "))

FINISH:

	return
}
