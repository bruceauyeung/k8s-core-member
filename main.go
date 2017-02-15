package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	yaml "gopkg.in/yaml.v2"
)

type User struct {
	Name       string
	ReviewerIn []string
	ApproverIn []string
}
type Owners struct {
	Reviewers []string
	Approvers []string
}

type Config struct {
	Users []string
}

var cfg Config = Config{[]string{}}
var Users []*User = []*User{}

func StringSliceContainsAny(s []string, str ...string) bool {
	if len(str) == 0 {
		return false
	}
	for _, elem := range str {
		for _, e := range s {
			if e == elem {
				return true
			}
		}
	}

	return false
}
func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}
	if strings.EqualFold(f.Name(), "OWNERS") {
		owners := Owners{}
		parentPath := strings.TrimSuffix(path, f.Name())
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = yaml.Unmarshal(b, &owners)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		for _, u := range Users {
			if StringSliceContainsAny(owners.Reviewers, u.Name) {
				u.ReviewerIn = append(u.ReviewerIn, parentPath)
			}
			if StringSliceContainsAny(owners.Approvers, u.Name) {
				u.ApproverIn = append(u.ApproverIn, parentPath)
			}
		}

	}

	return nil
}

func readConfig() {
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, u := range cfg.Users {
		Users = append(Users, &User{Name: u})
	}
}
func main() {
	flag.Parse()
	root := flag.Arg(0)
	readConfig()
	err := filepath.Walk(root, visit)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, u := range Users {
		if len(u.ReviewerIn) > 0 {
			fmt.Fprintf(color.Output, "user %s is reviewer in :\n", color.RedString(u.Name))
			for _, r := range u.ReviewerIn {
				fmt.Printf("\t%s\n", r)
			}
		}
		if len(u.ApproverIn) > 0 {
			fmt.Fprintf(color.Output, "user %s is approver in :\n", color.RedString(u.Name))
			for _, a := range u.ApproverIn {
				fmt.Printf("\t%s\n", a)
			}
		}
	}
}
