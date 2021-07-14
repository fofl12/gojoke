package gojoke
import (
	"testing"
	"fmt"
)

func TestAnyJoke(t *testing.T) {
	fmt.Println(Any())
}
func TestSafeJoke(t *testing.T) {
	fmt.Println(Safe())
}
func TestJokeInPunCategory(t *testing.T) {
	fmt.Println(Category("Pun"))
}
func TestJokeInGermanLanguage(t *testing.T) {
	fmt.Println(Language("de"))
}
