package main

import (
	"fmt"
	"unicode/utf8"
)

// examineRune takes a rune and prints a message if it matches
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func main() {

	// A UTF-8 string in Thai
	const s = "สวัสดี"

	// 1️⃣ Length in bytes
	fmt.Println("Len:", len(s))

	// 2️⃣ Print each byte in hexadecimal
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// 3️⃣ Count runes in the string
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 4️⃣ Iterate using range (gives rune and starting byte index)
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 5️⃣ Iterate using utf8.DecodeRuneInString explicitly
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Pass the rune to a function
		examineRune(runeValue)
	}
}
