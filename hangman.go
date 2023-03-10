package Hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Game(file string) {
	attempts := 10
	var StockUserChoise []string
	var Redondant bool
	data, err := os.Open(file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	var str []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	random := str[y1.Intn(len(str))]
	n := len(random)/2 - 1
	var randomSplitted []string
	var ToShow []string
	RandomRune := []rune(random)
	for i := range RandomRune {
		RandomRune[i] = RandomRune[i] - 32
	}
	var RandomUpper string
	for i := range RandomRune {
		RandomUpper += string(RandomRune[i])
	}
	randomSplitted = strings.Split(RandomUpper, "")
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(RandomRune))
		pick := RandomRune[randomIndex]
		if RandomRune[randomIndex] != 0 {
			ToShow = append(ToShow, string(pick))
		} else {
			i--
		}
		for k := range RandomRune {
			if RandomRune[k] == pick {
				RandomRune[k] = 0
			}
		}
	}
	res := make([]string, len(randomSplitted))
	for i := 0; i < len(randomSplitted); i++ {
		res[i] = "_"
	}
	for y := 0; y <= len(ToShow)-1; y++ {
		count := 0
		for _, i := range randomSplitted {
			if ToShow[y] == i {
				res[count] = i
			}
			count++
		}
	}
	dataPosition, err := os.Open("hangman.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fileScanner := bufio.NewScanner(dataPosition)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	print("Good Luck, you have 10 attempts.\n")
	for _, i := range res {
		print(i)
		print(" ")
	}
	for i := 0; i < 2; i++ {
		print("\n")
	}
	countFinish := 0
	for x := attempts; x > 0; x-- {
		Redondant = false
		if attempts < 1 {
			break
		}
		countFinish = 0
		for elm := range res {
			if res[elm] != "_" {
				countFinish++
			}
		}
		if countFinish == len(res) {
			print("Congrats !")
			break
		}
		var UserChoice string
		found := false
		fmt.Print("Choose: ")
		fmt.Scan(&UserChoice)
		if UserChoice == "STOP" {
			type ElementSaved struct {
				Attempts int
				Word     []string
				Result   []string
			}
			ElmtSaved := ElementSaved{
				Attempts: attempts,
				Word:     randomSplitted,
				Result:   res,
			}
			save, err := json.Marshal(ElmtSaved)
			if err != nil {
				fmt.Println("error:", err)
			}
			fileSave, err := os.Create("save.txt")
			if err != nil {
				log.Fatal(err)
			}
			errWrite := ioutil.WriteFile("save.txt", save, 0777)
			if errWrite != nil {
				fmt.Println(errWrite)
			}
			defer fileSave.Close()
			fmt.Println("Game Saved in save.txt.")
			break
		}
		for _, i := range StockUserChoise {
			if UserChoice == i {
				Redondant = true
			}
		}
		if Redondant == false {
			StockUserChoise = append(StockUserChoise, UserChoice)
		}
		if len(UserChoice) == 1 && UserChoice > string(rune(64)) && UserChoice < string(rune(91)) && Redondant == false {
			for i := range randomSplitted {
				if UserChoice == randomSplitted[i] {
					res[i] = UserChoice
					x++
					found = true
				}
			}
			if found == true {
				for _, i := range res {
					print(i)
					print(" ")
				}
				for j := 0; j < 2; j++ {
					print("\n")
				}
			}
			if found == false {
				attempts--
				print("Not present in the word, ", attempts, " attempts remaining\n")
				if attempts == 9 {
					for i := 0; i < 8; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 8 {
					for i := 8; i < 16; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 7 {
					for i := 16; i < 24; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 6 {
					for i := 24; i < 32; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 5 {
					for i := 32; i < 40; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 4 {
					for i := 40; i < 48; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 3 {
					for i := 48; i < 56; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 2 {
					for i := 56; i < 64; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 1 {
					for i := 64; i < 72; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts < 1 {
					for i := 72; i < 80; i++ {
						fmt.Println(lines[i])
					}
				}
			}
		} else if len(UserChoice) > 1 && Redondant == false {
			count := 0
			for _, i := range UserChoice {
				if string(i) > string(rune(64)) && string(i) < string(rune(91)) {
					count++
				}
			}
			if count == len(UserChoice) {
				var StrRandomSlitted string
				for _, i := range randomSplitted {
					StrRandomSlitted += i
				}
				if UserChoice == StrRandomSlitted {
					for i := range randomSplitted {
						res[i] = randomSplitted[i]
					}
					for _, i := range res {
						print(i)
						print(" ")
					}
					for j := 0; j < 2; j++ {
						print("\n")
					}
				} else {
					attempts -= 2
					if attempts < 0 {
						attempts = 0
					}
					print("Not present in the word, ", attempts, " attempts remaining\n")
					if attempts == 9 {
						for i := 0; i < 8; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 8 {
						for i := 8; i < 16; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 7 {
						for i := 16; i < 24; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 6 {
						for i := 24; i < 32; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 5 {
						for i := 32; i < 40; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 4 {
						for i := 40; i < 48; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 3 {
						for i := 48; i < 56; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 2 {
						for i := 56; i < 64; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 1 {
						for i := 64; i < 72; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts < 1 {
						for i := 72; i < 80; i++ {
							fmt.Println(lines[i])
						}
					}
					x++
				}
			}
		} else if Redondant == true {
			fmt.Println("Already used ! ")
			x++
		}
	}
	if countFinish != len(res) && attempts < 1 {
		fmt.Println("You lose ! The result was ", RandomUpper, ".")
	}
}
