package brackets

type BracketMap map[string]string

var br = BracketMap{
	"{": "}",
	"[": "]",
	"(": ")",
}

func Bracket(input string) (bool, error) {
	res := " "
	for _, char := range input {
		val := br[res[len(res)-1:]]
		res += string(char)
		if string(char) == val {
			res = res[:len(res)-2]
		}
	}
	return len(res) == 1, nil
}
