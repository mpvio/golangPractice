package basic

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringExamples() {
	// BasicInfo()
	// BytesFromString("TEST")
	// CreateStringFromSlices()
	// IteratingOverString()
	// ComparingStrings()
	// ConcatenatingString()
	// StringTrimming()
	// StringSplitting()
	// StringContains()
	otherStringFunctions()
}

func BasicInfo() {
	str1 := "Strings:\nare read-only slices of bytes"
	var str2 string = `if created with backticks, they can span
	multiple lines BUT can't use escape characters. for example: \n does nothing.
	line breaks occur where they're placed right here
	like
	this`

	fmt.Println(str1)
	fmt.Println(str2)
}

// foreach loop: declare index AND character separately
func IteratingOverString() {
	for index, character := range "String To Be Iterated Over" {
		// print index number, then character. character type is rune
		if index < 5 {
			fmt.Printf("char #%d is character %c\n", index, character)
		}
	}
}

// basic for loop AND byte example
func BytesFromString(input string) {
	for c := range len(input) {
		// above is abbreviation of `c := 0; c < len(input); c++`
		// %c is character, %x is byte value (hexadecimal)
		if c < 5 {
			fmt.Printf("Character is %c, Byte value is %x\n", input[c], input[c])
		}
	}

	fmt.Printf("String length can also be found with RuneCountInString(): %d\n", utf8.RuneCountInString(input))
}

// can create string from a slice of bytes and slice of runes
func CreateStringFromSlices() {
	sliceOfBytes := []byte{0x47, 0x65}
	sliceofRunes := []rune{0x0047, 0x0065}

	str1 := string(sliceOfBytes)
	str2 := string(sliceofRunes)

	fmt.Println(str1, str2)
}

func ComparingStrings() {
	s1 := "Hello"
	s2 := "World"
	s3 := "Hello"

	equality1 := s1 == s2
	equality2 := s1 == s3

	fmt.Println(equality1, equality2)

	// comparison: using >, <, >=, <= to check alphabetical order
	// returns boolean
	is_s1_earlierThan_s2 := s1 < s2

	//compare function in strings module returns int:
	// 0 if the same, 1 if 1st is bigger
	// -1 if 2nd is bigger
	comparison := strings.Compare(s1, s2)

	fmt.Println(is_s1_earlierThan_s2, comparison)
}

func ConcatenatingString() {
	s1 := "Hello, "
	s2 := "World!"

	// multiple ways to concatenate them
	fmt.Println("with +:", s1+s2)

	var b bytes.Buffer
	b.WriteString(s1)
	b.WriteString(s2)
	fmt.Println("Create a bytes.Buffer and add strings to it:", b.String())

	var bu strings.Builder
	bu.WriteString(s1)
	bu.WriteString(s2)
	fmt.Println("strings.Builder is similar to bytes.Buffer", bu.String())

	result := fmt.Sprintf("Create formatted string: %s%s", s1, s2)
	fmt.Println(result)

	joinedString := strings.Join([]string{s1, s2}, "separator")
	fmt.Println(
		"Join takes a list of strings as parameter, then a separator (can be '')",
		joinedString)

	s1 += s2
	fmt.Println("+= replaces s1 with a new string that combines them:", s1)
}

func StringTrimming() {
	fmt.Println(strings.Trim("@!@!ALL chars in cutset are removed regardless of arrangement!@!@", "!@"))
	fmt.Println(strings.TrimLeft("@!Only chars from left are removed!@", "@!"))
	fmt.Println(strings.TrimRight("@!Only chars from right are removed!@", "@!"))
	fmt.Println(strings.TrimSpace("  trim all leading/ trailing spaces only   "))
	fmt.Println(strings.TrimPrefix("@@removes EXACT substring from string start ONLY@@", "@@"))
	fmt.Println(strings.TrimSuffix("@@removes EXACT substring from string end ONLY@@", "@@"))
}

func StringSplitting() {
	str := "Hello,World,!"
	fmt.Println("OG String:", str)

	var result []string = strings.Split(str, ",")
	fmt.Println("Split on , (, is removed) produces", result, len(result))

	result = strings.SplitAfter(str, ",")
	fmt.Println("SplitAfter on , (, is kept) produces", result, len(result))

	result = strings.SplitN(str, ",", 1)
	fmt.Println("SplitN on , & 1 (i.e. max 1 substring) produces", result, len(result))

	result = strings.SplitAfterN(str, ",", 1)
	fmt.Println("SplitAfterN on , & 1 (i.e. only split once) produces", result, len(result))
}

func StringContains() {
	str := "Hello World"

	contains := strings.Contains(str, "World")
	containsAny := strings.ContainsAny("str", "H&W")
	cA2 := strings.ContainsAny("str", "H&F")
	cA3 := strings.ContainsAny("str", "H|F")
	cA4 := strings.ContainsAny("str", "HW")
	cA5 := strings.ContainsAny("str", "HF")

	fmt.Println("contains checks for a specific substring", contains)
	fmt.Println("containsAny checks if any of listed chars are in string", containsAny)
	fmt.Println("true & false, true OR false, two characters (true), two characters (false)", cA2, cA3, cA4, cA5)
}

func otherStringFunctions() {
	str := "Hello World"

	repeatedStr := strings.Repeat(str, 2)
	fmt.Println("Repeat repeats string N times, here N = 2:", repeatedStr)

	fmt.Println("Index finding:")
	validIndex := strings.Index(str, "Hello")
	invalidIndex := strings.Index(str, "ABC")
	firstValidIndex := strings.Index(str, "o")
	fmt.Println(validIndex, invalidIndex, firstValidIndex)

	fmt.Println("IndexAny returns first instance of ANY chars in search param:")
	fmt.Println(strings.IndexAny(str, "abcde"))

	fmt.Println("IndexByte is identical to Index but takes byte/ char as search param:")
	fmt.Println(strings.IndexByte(str, 'H'))
	fmt.Println(strings.IndexByte(str, 'X'))

	fmt.Println("Counting substrings:")
	countO := strings.Count(repeatedStr, "o")
	countX := strings.Count(repeatedStr, "X")
	countEmpty := fmt.Sprintf("Counting '' returns ALL characters, i.e. length of string: %d", strings.Count(repeatedStr, ""))
	fmt.Println(countO, countX, countEmpty)
}
