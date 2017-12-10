import "testing"

func BenchmarkIsPalindrome(b testing.B) bool {
	for i:=0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}