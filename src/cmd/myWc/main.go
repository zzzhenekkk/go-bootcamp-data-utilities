package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	// Флаги для выбора режима подсчета
	countLines := flag.Bool("l", false, "Подсчет строк")
	countWords := flag.Bool("w", false, "Подсчет слов")
	countChars := flag.Bool("m", false, "Подсчет символов")
	flag.Parse()

	// Проверяем, что выбран только один режим
	if (*countLines && *countWords) || (*countLines && *countChars) || (*countWords && *countChars) || (!*countLines && !*countWords && !*countChars) {
		fmt.Println("Укажите один флаг: -l, -w, или -m")
		os.Exit(1)
	}

	// Обработка каждого файла в отдельной горутине
	var wg sync.WaitGroup
	for _, filename := range flag.Args() {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			countFile(filename, *countLines, *countWords, *countChars)
		}(filename)
	}
	wg.Wait()
}

func countFile(filename string, countLines, countWords, countChars bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineCount, wordCount, charCount int

	if countWords {
		// Устанавливаем функцию разделения на слова
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		if countLines {
			lineCount++
		}
		if countWords {
			wordCount++ // Увеличиваем счетчик слов за каждое найденное слово
		}
		if countChars {
			charCount += len(scanner.Text()) // Подсчитываем символы в строке
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла %s: %v\n", filename, err)
		return
	}

	// Вывод результата в зависимости от выбранного режима
	var result int
	if countLines {
		result = lineCount
	} else if countWords {
		result = wordCount
	} else if countChars {
		result = charCount
	}
	fmt.Printf("%d\t%s\n", result, filename)
}
