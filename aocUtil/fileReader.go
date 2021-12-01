package aocUtil

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ReadLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func LoadInput(day int, part string) io.Reader {
	path := getPath(day, part)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := downloadInput(day, path)
		if err != nil {
			log.Fatalf("Could not download the inputfile: %v", err)
		}
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	return bufio.NewReader(file)
}

func downloadInput(day int, path string) error {
		url := strings.Builder{}
	_, err := fmt.Fprintf(&url, "https://adventofcode.com/2021/day/%d/input", day)
	if err != nil {
		return err
	}

	var client http.Client
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	viper.SetConfigFile("credentials.yml")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	sessionID := viper.GetString("sessionID")

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatalf("Could not close file: %v", err)
		}
	}(out)

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getPath(day int, part string) string {
	path := strings.Builder{}
	path.WriteString("inputs/day")

	if day < 10 {
		path.WriteString("0")
	}

	_, err := fmt.Fprintf(&path, "%d%s.txt", day, part)
	if err != nil {
		log.Fatalf("Could not build filename: %v", err)
	}

	return path.String()
}