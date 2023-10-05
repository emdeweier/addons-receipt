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
