package main

// play.golang.org to play around with go without having to set it up

import "fmt"
import "os"
import "log"
import "bufio"
import "strings"
import "regexp"
import "strconv"
import "math/rand"
import "time"

func main() {
	// gather args and read in text source //
	var shingle_num int = 3
	var output_len int = 150
	var err error = nil

	if len(os.Args) < 2 || len(os.Args) > 4 {
		log.Fatal("Usage: babel sourcetext.txt (shingle_size) (output_length)")
	} else if len(os.Args) == 3 {
		shingle_num, err = strconv.Atoi(os.Args[2])
	} else if len(os.Args) == 4 {
		output_len, err = strconv.Atoi(os.Args[3])
	}

	text_source := os.Args[1]
	file, err := os.Open(text_source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string
	// split text lines into single words remove punctuation and lowercase //

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line_words := strings.Split(scanner.Text(), " ")
		for _, new_word := range line_words {
			no_punctuation_word := reg.ReplaceAllString(new_word, "")
			lowered_word := strings.ToLower(no_punctuation_word)
			if lowered_word != "" {
				words = append(words, lowered_word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// add first shingle_num words onto front of words array
	for i := 0; i<shingle_num; i++ {
		words = append(words, words[i])
	}

	// make words into ngrams

	var lookup_table map[string][]string = make(map[string][]string)

	for i := 0; i < len(words)-shingle_num; i++ {
		new_key := words[i]
		for j:= 1; j < shingle_num-1; j++ {
			new_key = new_key + " " + words[i+j]
		}
		endwords, ok := lookup_table[new_key]
		if ok {
			endwords = append(endwords, words[i+shingle_num-1])
			lookup_table[new_key] = endwords
		} else {
			lookup_table[new_key] = []string{words[i+shingle_num-1]}
		}
	}

	// time to babble out things
	var babble_key_array []string
	rand.Seed(time.Now().Unix())

	random_num := rand.Intn(len(words)-shingle_num-1)
	babble_key_array = append(babble_key_array, words[random_num])
	for i := 1; i < shingle_num-1; i++{
		babble_key_array = append(babble_key_array, words[random_num+i])
	}

	for i := 0; i < output_len; i++ {
		var babble_key string = babble_key_array[0]
		for j := 1; j < len(babble_key_array); j++ {
			babble_key = babble_key + " " + babble_key_array[j]
		}
		output_words, ok := lookup_table[babble_key]
		if (ok){
			output_word := output_words[rand.Intn(len(output_words))]
			fmt.Printf(output_word+" ")
			babble_key_array = babble_key_array[1:]
			babble_key_array = append(babble_key_array, output_word)
		}
	}
}
