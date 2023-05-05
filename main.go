package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var sum int

type Bi struct {
	count int
	prob  float64
}

func ReadFile(fileName string) map[string]Bi {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Я не смог прочесть файл")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	birgrams := make(map[string]Bi)

	for scanner.Scan() {
		word := "^" + scanner.Text() + "$"
		for i := 0; word[i] != '$'; i++ {
			sum++
			bigram := string(word[i]) + string(word[i+1])
			b := birgrams[bigram]
			b.count++
			birgrams[bigram] = Bi{b.count, 0}
		}
	}
	return birgrams
}

func probBi(bigrams map[string]Bi) []string {
	lenbi := len(bigrams)

	f := make([]string, 0, lenbi)

	for i := range bigrams {
		b := bigrams[i]

		b.prob = float64(b.count) / float64(sum)
		bigrams[i] = Bi{b.count, b.prob}
		if contains(i) {
			f = append(f, i)
		}
	}
	return (f)
}

func contains(s string) bool {
	if s[0] == '^' {
		return true
	}
	return false
}

func randFirstLet(f []string) string {
	rand.Seed(time.Now().UnixNano())
	randKey := f[rand.Intn(len(f))]
	return randKey
}

func countName(f string, big map[string]Bi) string {
	nextL := f[len(f)-1]

	if nextL == '$' {
		return f
	}

	var bigi []string
	var probi []float64
	for i := range big {
		if i[0] == nextL {
			bigi = append(bigi, i)
			b := big[i]
			probi = append(probi, b.prob)
		}
	}

	var newP float64

	for _, j := range probi {
		newP += j
	}

	newP = float64(1) / newP

	probi2 := make([]float64, len(probi))

	for i := 0; i < len(probi); i++ {
		probi2[i] = newP * probi[i]
	}

	cumProbi := make([]float64, len(probi))
	cumProbi[0] = probi2[0]

	for i := 1; i < len(probi2); i++ {
		cumProbi[i] = cumProbi[i-1] + probi2[i]
	}
	randFloat := rand.Float64()

	var index int

	for i := 0; i < len(cumProbi); i++ {
		if randFloat < cumProbi[i] {
			index = i
			break
		}
	}

	f += string(bigi[index][1])

	return countName(f, big)
}

func main() {
	bigrams := ReadFile("name.txt")
	firstLetters := probBi(bigrams)
	firstLetter := randFirstLet(firstLetters)

	for i := 0; i < 10; i++ {
		name := countName(firstLetter, bigrams)
		name = name[:len(name)-1]
		name = name[1:]
		if len(name) < 2 || len(name) > 12 {
			for len(name) < 2 || len(name) > 12 {
				name = countName(firstLetter, bigrams)
				name = name[:len(name)-1]
				name = name[1:]
			}
		}
		fmt.Println(name)
	}

	fmt.Println("Вы хотите увидеть вероятности всех биграм?\nY-Да N-Нет\nЯ на всякий случай все запишу в отдельный output.txt")

	var answer string

	fmt.Scan(&answer)

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if answer == "Y" || answer == "y" {
		text := "|Биграмма|Вероятность в %"
		fmt.Println(text)
		_, err = file.WriteString(text + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := range bigrams {
			text = "  " + i + "   |"
			ftext := fmt.Sprintf("%f", bigrams[i].prob*100) + "%"
			fmt.Println(text + ftext)
			_, err = file.WriteString(text + ftext + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}
