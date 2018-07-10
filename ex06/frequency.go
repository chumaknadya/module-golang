package letter

type Map map[rune]int

func Frequency(s string) Map {
	m := Map{}
	for _, val := range s {
		m[val]++
	}
	return m
}

func ConcurrentFrequency(arrStr []string) Map {
	m := Map{}
	freq := make(chan Map, len(arrStr))
	for _, str := range arrStr {
		go func(s string) {
			freq <- Frequency(s)
		}(str)
	}

	for range arrStr {
		for let, count := range <-freq {
			m[let] += count
		}
	}

	return m
}
