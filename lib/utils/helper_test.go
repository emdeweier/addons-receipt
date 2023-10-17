package utils

import (
	"fmt"
	"log"
	"testing"
)

func Test_CutStringReceiptV2(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "test 50 character",
			text: "Akulah Cinta Di Langit Prudence Lovely Princess Of Awanamp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CutStringReceiptV2(tt.text, 20, 1, false)
			fmt.Println("result : ", res)
		})
	}
}

func Test_cutDotWord(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name: "test 50 character",
			text: "AkulahCintaDiLangit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CutDotWord(tt.text, 20)
			fmt.Println("result : ", res)
		})
	}
}

func Test_textStructuring(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected []string
		cutDOt   bool
		length   int
		height   int
	}{
		{
			name:     "one word cutdot",
			text:     "AkulahCintaDiLangittiga",
			cutDOt:   true,
			length:   20,
			expected: []string{"AkulahCintaDiLangit..."},
		},
		{
			name:     "one word newline single line",
			text:     "AkulahCintaDiLangittiga",
			cutDOt:   false,
			length:   20,
			height:   1,
			expected: []string{"AkulahCintaDiLangitt"},
		},
		{
			name:     "one word 2 line",
			text:     "AkulahCintaDiLangittiga",
			cutDOt:   false,
			length:   20,
			height:   2,
			expected: []string{"AkulahCintaDiLangitt", "iga"},
		},
		{
			name:     "one word 5 line",
			text:     "AkulahCintaDiLangittiga",
			cutDOt:   false,
			length:   20,
			height:   5,
			expected: []string{"AkulahCintaDiLangitt", "iga"},
		},
		{
			name:     "normal",
			text:     "Akulah Cinta Di Langit Prudence Lovely Princess Of Awanamp",
			cutDOt:   false,
			length:   20,
			height:   5,
			expected: []string{"Akulah Cinta Di", "Langit Prudence", "Lovely Princess Of", "Awanamp"},
		},
		{
			name:     "normal cutdot",
			text:     "Akulah Cinta Di Langit Prudence Lovely Princess Of Awanamp",
			cutDOt:   true,
			length:   20,
			height:   1,
			expected: []string{"Akulah Cinta Di..."},
		},
		{
			name:     "normal have word long",
			text:     "Akulah Cinta Di AkulahCintaDiLangittiga Langit Prudence Lovely Princess Of Awanamp",
			cutDOt:   false,
			length:   20,
			height:   5,
			expected: []string{"Akulah Cinta Di", "AkulahCintaDiLangitt", "iga Langit Prudence", "Lovely Princess Of", "Awanamp"},
		},
		{
			name:     "normal have word long with 3 line",
			text:     "Akulah Cinta Di AkulahCintaDiLangittiga Langit Prudence Lovely Princess Of Awanamp",
			cutDOt:   false,
			length:   20,
			height:   3,
			expected: []string{"Akulah Cinta Di", "AkulahCintaDiLangitt", "iga Langit Prudence"},
		},
		{
			name:     "normal have word long with 2 line",
			text:     "Akulah Cinta Di AkulahCintaDiLangittiga Langit Prudence Lovely Princess Of Awanamp",
			cutDOt:   false,
			length:   20,
			height:   2,
			expected: []string{"Akulah Cinta Di", "AkulahCintaDiLangitt"},
		},
		{
			name:     "many word just get first word",
			text:     "Akulah Cinta Di AkulahCintaDiLangittiga Langit Prudence Lovely Princess Of Awanamp",
			cutDOt:   false,
			length:   5,
			height:   2,
			expected: []string{"Akula", "h"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := textStructuring(tt.text, tt.length, tt.height, tt.cutDOt)

			if len(tt.expected) != len(res) {
				t.Fail()
			}

			for i, v := range res {
				if tt.expected[i] != v {
					t.Fail()
				}
				log.Println("result :", v)
			}

		})
	}
}
