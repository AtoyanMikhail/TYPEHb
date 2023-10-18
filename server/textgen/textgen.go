package textgen

import (
	"math/rand"
	"os"
	"strings"
)

// пакеты всегда лежат в одноименной папке

// Удовлетворяй эту хуйню
// Одинаковый seed должен возвращать одинаковый текст
// В интерфейсе только объявления методов
type Generator interface {
	Generate(seed int64) string
}

// Все тута :)
// Local это частный случай Generator
// Соответственно, т.к Local, то словарь слов получает из файла
type Local struct {
	root  string
	words []string
}

func NewLocal(root string) Local {
	data, err := os.ReadFile(root)

	if err != nil {
		panic(err)
	}

	words := strings.Split(string(data), "\n")

	local := Local{
		root:  root,
		words: words,
	}

	return local
}

func (local Local) Generate(seed int64) string {
	rand.Seed(seed)
	words_copy := local.words[:]
	rand.Shuffle(len(words_copy), func(i, j int) {
		words_copy[i], words_copy[j] = words_copy[j], words_copy[i]
	})

	return strings.Join(words_copy[:200], " ")
}
