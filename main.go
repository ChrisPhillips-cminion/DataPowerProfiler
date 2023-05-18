package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	deepAttributes := make(map[string][]string)
	deepAttributes["action"] = append(deepAttributes["action"], "aaa-policy")
	deepAttributes["action"] = append(deepAttributes["action"], "type")
	deepAttributes["all"] = append(deepAttributes["all"], "admin-state")

	path := flag.String("path", "", "path to the cfg file")
	flag.Parse()
	fmt.Println("--------------------------------------")
	fmt.Println("Loading file ::", *path)
	fmt.Println("--------------------------------------")

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]int)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	var inAction string
	var cryptoFlag bool
	cryptoFlag = false
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		var prefix string
		prefix = ""
		line = strings.TrimSpace(line)
		if len(line) > 2 {
			if strings.HasPrefix(line, "no ") {
				line  = strings.Replace(line," ","-",2)
			}
			if cryptoFlag {
				prefix = "crypto"
			}
			if strings.HasPrefix(line, "top") {
			} else if strings.HasPrefix(line, "exit") {
				// fmt.Println("Found Exit ending ",inAction)
				if inAction == "" {
					cryptoFlag = false
				}
				inAction = ""
			} else if line[0] == "#"[0] {
			} else if line[0] == "%"[0] {
			} else if line == "" {
			} else if line == "\t" {
				// } else if line[0] == " "[0] {
			} else if line == "crypto" {


				inAction = ""

				cryptoFlag = true
			} else if inAction != "" {

				// fmt.Println("In Action  ",inAction,"->",line)

				for v := range deepAttributes[inAction] {
					// fmt.Println(v)
					line = strings.TrimSpace(line)
					if strings.HasPrefix(line, deepAttributes[inAction][v]) {
						key := inAction + "->" + deepAttributes[inAction][v] + ":" + strings.Split(line, " ")[1]

						m[key]++

					}
				}
				for v := range deepAttributes["all"] {

					if strings.HasPrefix(line, deepAttributes["all"][v]) {
						key := inAction + "->" + deepAttributes["all"][v] + ":" + strings.Split(line, " ")[1]

						m[key]++

					}
				}
			} else {
				line = strings.TrimSpace(line)
				if cryptoFlag {
					key := fmt.Sprint(prefix, "->", strings.Split(line, " ")[0])
					if len(strings.Split(line, " ")) < 2 {
						inAction = key
					}

					if key == "crypto->" {
						key = "crypto"
					}
					m[key]++

				} else {
					line = strings.TrimSpace(line)
					key := fmt.Sprint(prefix, strings.Split(line, " ")[0])
					inAction = key

					m[key]++

				}
			}
		}
	}
	mk := make([]string, len(m))
	for k, v := range m {
		mk = append(mk, fmt.Sprint("'", k, "' : ", v))
	}
	sort.Strings(mk)
	for v := range mk {
		if mk[v] != "" {
			fmt.Println(mk[v])
		}
	}
}
