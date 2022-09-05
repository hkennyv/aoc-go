package main

import "testing"

type Case1 struct {
	input                 string
	hasThreeVowels        bool
	containsDoubleLetters bool
	doesNotContainPairs   bool
}

type Case2 struct {
	input                     string
	hasNonOverlappingPairs    bool
	hasRepeatingLetterBetween bool
}

func TestRules(t *testing.T) {

	cases1 := []Case1{
		{"ugknbfddgicrmopn", true, true, true},
		{"aaa", true, true, true},
		{"jchzalrnumimnmhp", true, false, true},
		{"haegwjzuvuyypxyu", true, true, false},
		{"dvszwmarrgswjxmb", false, true, true},
	}

	t.Run("test has three vowels", func(t *testing.T) {
		for _, c := range cases1 {
			got := hasThreeVowels(c.input)

			if got != c.hasThreeVowels {
				t.Errorf("%s - got %v want %v", c.input, got, c.hasThreeVowels)
			}
		}
	})

	t.Run("test contains double letters", func(t *testing.T) {
		for _, c := range cases1 {
			got := containsDoubleLetter(c.input)

			if got != c.containsDoubleLetters {
				t.Errorf("%s - got %v want %v", c.input, got, c.containsDoubleLetters)
			}
		}
	})

	t.Run("test does not contain forbidden pairs", func(t *testing.T) {
		for _, c := range cases1 {
			got := doesNotContainPairs(c.input)

			if got != c.doesNotContainPairs {
				t.Errorf("%s - got %v want %v", c.input, got, c.doesNotContainPairs)
			}
		}
	})

	cases2 := []Case2{
		{"qjhvhtzxzqqjkmpb", true, true},
		{"aabcdefgaa", true, false},
		{"uurcxstgmygtbstg", true, false},
		{"ieodomkazucvgmuy", false, true},
		{"jjjrprbnlijzatuw", false, true},
		{"jcaqyaqvsefwtaya", true, true},
	}

	t.Run("test contains non overlapping pairs", func(t *testing.T) {
		for _, c := range cases2 {
			got := hasNonOverlappingPair(c.input)

			if got != c.hasNonOverlappingPairs {
				t.Errorf("%s - got %v want %v", c.input, got, c.hasNonOverlappingPairs)
			}
		}
	})

	t.Run("test has repeating letter between", func(t *testing.T) {
		for _, c := range cases2 {
			got := hasRepeatingLetterBetween(c.input)

			if got != c.hasRepeatingLetterBetween {
				t.Errorf("%s - got %v want %v", c.input, got, c.hasRepeatingLetterBetween)
			}
		}
	})

}
