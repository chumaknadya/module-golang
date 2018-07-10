package downcase

func Downcase(initStr string) (string, error) {
  arr := make([]byte, len(initStr))
  for i := range arr {
    letter := initStr[i]
    if letter >= 'A' && letter <= 'Z' {
      letter += 'a' - 'A'
    }
    arr[i] = letter
 }
 return string(arr), nil
}
