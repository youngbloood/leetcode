package recursion

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/
func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

var phoneDigits = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return letterCombination(nil, "", digits)
}

func letterCombination(result []string, src, digits string) []string {
	if len(digits) == 0 {
		result = append(result, src)
		return result
	}
	waitStr := phoneDigits[string(digits[0])]
	for _, letter := range waitStr {
		src += string(letter)
		result = letterCombination(result, src, digits[1:])
		src = src[:len(src)-1]
	}
	return result
}
