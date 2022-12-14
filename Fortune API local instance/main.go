package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func getFortune(filename string, fortunes int) string {
	// filename : the file the fortune is being selected from
	// fortunes : the total number of fortunes in the file
	// Returns : the fortune in the form of a formatted string

	if filename == "datfile/favicon.ico" {
		return ""
	}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	rand.Seed(time.Now().UnixNano())
	b := rand.Intn(fortunes - 1)
	fmt.Println("\nChose fortune", b, "of", fortunes, "total fortunes")

	count := 0
	fortune := ""

	for scanner.Scan() {
		if scanner.Text() != "%" && count == b {
			fortune = fortune + scanner.Text() + "\n"
		} else if scanner.Text() == "%" && (count > b) {
			break
		} else if scanner.Text() == "%" {
			count++
		}
	}

	return fortune
}

func getRandomFortune(w http.ResponseWriter, r *http.Request) {
	// Selects a random datfile to return a random fortune (not including offensive)
	// Returns : the fortune in a formatted string

	// setting header to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var files []string
	var datfiles []string

	// Get list of all fortune files in datfiles
	root := "datfiles"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !strings.Contains(file, "CMakeLists.txt") && !strings.Contains(file, "off") && !strings.Contains(file, "README.md") && file != "datfiles" {
			datfiles = append(datfiles, file)
		}
	}

	// Choose a random datfile
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(len(datfiles))
	fmt.Println("\nChose datfile", a, "of", len(datfiles), "total datfiles")

	// Select random file
	randFile := datfiles[a]

	fmt.Println("\ndatfile chosen:", randFile)

	file, err := os.Open(randFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	fortunes := 0

	for scanner.Scan() {
		if scanner.Text() == "%" {
			fortunes++
		}
	}

	fortune := getFortune(randFile, fortunes)
	json.NewEncoder(w).Encode(fortune)
}

func getSpecificFortuneType(w http.ResponseWriter, r *http.Request) {
	// Get a specific genre of fortune; must be a file within datfiles

	// setting header to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	if params["genre"] == "favicon.ico" {
		return
	}
	filePath := fmt.Sprintf("%s%s", "datfiles/", params["genre"])

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	fortunes := 0

	for scanner.Scan() {
		if scanner.Text() == "%" {
			fortunes++
		}
	}

	fortune := getFortune(filePath, fortunes)
	json.NewEncoder(w).Encode(fortune)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getRandomFortune).Methods("GET")
	router.HandleFunc("/{genre}", getSpecificFortuneType).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
