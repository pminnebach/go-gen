package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "go-gen"
	app.Usage = "Generate a golang project folder."
	app.Authors = []cli.Author{
		cli.Author{
			Name: "pminnebach",
		},
	}
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Copyright = "MIT (c) 2017"

	var repo string
	var projectName string
	var license string
	var readme bool
	var gitignore bool
	var overwrite bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "repository,repo",
			Usage:       "Name of your repo. (github.com/<username>)",
			Destination: &repo,
		},
		cli.StringFlag{
			Name:        "name,n,project",
			Value:       "myapp",
			Usage:       "Name of the project you want to create.",
			Destination: &projectName,
		},
		cli.StringFlag{
			Name:        "license,l",
			Usage:       "Add a LICENSE file. (MIT, Apache)",
			Destination: &license,
		},
		cli.BoolFlag{
			Name:        "readme,r",
			Usage:       "Add a README.md file.",
			Destination: &readme,
		},
		cli.BoolFlag{
			Name:        "gitignore,g",
			Usage:       "Add a .gitignore file.",
			Destination: &gitignore,
		},
		cli.BoolFlag{
			Name:        "overwrite,o",
			Usage:       "Overwrite existing files.",
			Destination: &overwrite,
			Hidden:      true,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("+--------------------------+\n")
		fmt.Printf("| Golang Project Generator |\n")
		fmt.Printf("+--------------------------+\n")

		if len(projectName) > 0 {
			fmt.Printf("ProjectName: %s\n", projectName)

			path := projectName + "/hello"
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.MkdirAll(path, 0700)
			}

			h := filepath.Join(projectName, "hello/hello.go")
			fmt.Printf("INFO: Creating file: %s\n", h)
			data, err := Asset("templates/hello/_hello.go")
			check(err)
			ioutil.WriteFile(h, data, 0644)

			ht := filepath.Join(projectName, "hello/hello_test.go")
			fmt.Printf("INFO: Creating file: %s\n", ht)
			data, err = Asset("templates/hello/_hello_test.go")
			check(err)
			ioutil.WriteFile(ht, data, 0644)

			m := filepath.Join(projectName, "main.go")
			fmt.Printf("INFO: Creating file: %s\n", m)
			data, err = Asset("templates/_main.go")
			check(err)
			v1 := bytes.Replace([]byte(data), []byte("<%=myrepoUrl%>"), []byte(repo), 1)
			v2 := bytes.Replace([]byte(v1), []byte("<%=myappName%>"), []byte(projectName), 1)
			ioutil.WriteFile(m, v2, 0644)

			mf := filepath.Join(projectName, "Makefile")
			fmt.Printf("INFO: Creating file: %s\n", mf)
			data, err = Asset("templates/_Makefile")
			check(err)
			v1 = bytes.Replace([]byte(data), []byte("<%=myrepoUrl%>"), []byte(repo), 1)
			v2 = bytes.Replace([]byte(v1), []byte("<%=myappName%>"), []byte(projectName), 2)
			ioutil.WriteFile(mf, v2, 0644)
		}

		fmt.Printf("License:     %v \n", license)
		fmt.Printf("Readme:      %v \n", readme)
		fmt.Printf("Gitignore:   %v \n", gitignore)

		if len(license) > 0 {
			path := filepath.Join(projectName, "LICENSE")
			fmt.Println(path)

			switch strings.ToUpper(license) {
			case "MIT":
				data, err := Asset("templates/_LICENSE_MIT")
				check(err)
				ioutil.WriteFile(path, data, 0644)
			case "APACHE":
				data, err := Asset("templates/_LICENSE_APACHE")
				check(err)
				ioutil.WriteFile(path, data, 0644)
			default:
				err := fmt.Errorf("ERROR: No valid license type specified.\n")
				if err != nil {
					fmt.Print(err)
				}
			}
		}

		if readme == true {
			path := filepath.Join(projectName, "README.md")
			fmt.Println(path)

			data, err := Asset("templates/_README.md")
			check(err)
			ioutil.WriteFile(path, data, 0644)
		}

		if gitignore == true {
			path := filepath.Join(projectName, ".gitignore")
			fmt.Println(path)

			data, err := Asset("templates/_gitignore")
			check(err)
			ioutil.WriteFile(path, data, 0644)
		}

		return nil
	}

	app.Run(os.Args)
}
