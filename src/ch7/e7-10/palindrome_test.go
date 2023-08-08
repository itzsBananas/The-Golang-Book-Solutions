package palindrome

import "testing"

type stringVariant []byte

func (x stringVariant) Len() int           { return len(x) }
func (x stringVariant) Less(i, j int) bool { return x[i] < x[j] }
func (x stringVariant) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type intVariant []int

func (x intVariant) Len() int           { return len(x) }
func (x intVariant) Less(i, j int) bool { return x[i] < x[j] }
func (x intVariant) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestString(t *testing.T) {

	pal := stringVariant([]byte("abccba"))

	if !IsPalindrome(pal) {
		t.Errorf("%s is a palindrome, but reported otherwise", pal)
	}

	pal = stringVariant([]byte("abcba"))

	if !IsPalindrome(pal) {
		t.Errorf("%s is a palindrome, but reported otherwise", pal)
	}

	pal = stringVariant([]byte("abcdf"))

	if IsPalindrome(pal) {
		t.Errorf("%s is not a palindrome, but reported otherwise", pal)
	}
}

func TestIntArray(t *testing.T) {
	pal := intVariant([]int{1, 2, 2, 1})

	if !IsPalindrome(pal) {
		t.Errorf("%x is a palindrome, but reported otherwise", pal)
	}

	pal = intVariant([]int{1, 2, 1})

	if !IsPalindrome(pal) {
		t.Errorf("%x is a palindrome, but reported otherwise", pal)
	}

	pal = intVariant([]int{1, 2, 2, 4, 1})

	if IsPalindrome(pal) {
		t.Errorf("%x is not a palindrome, but reported otherwise", pal)
	}

}
