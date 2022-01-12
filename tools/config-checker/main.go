package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	if err := checkNGWords("ok_config"); err != nil {
		return err
	}
	if err := checkNGWords("ng_config"); err != nil {
		return err
	}
	return nil
}

func checkNGWords(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	fs := bufio.NewScanner(file)
	i := 0
	for fs.Scan() {
		i++
		for _, v := range ngWords() {
			if strings.Contains(fs.Text(), v) {
				return fmt.Errorf("%s found on line:%d@%s", v, i, name)
			}
		}
	}
	return fs.Err()
}

func ngWords() []string {
	return []string{
		"example.com",
	}
}
