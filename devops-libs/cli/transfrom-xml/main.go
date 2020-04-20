package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Pom struct {
	Modules struct {
		Module []string `xml:"module"`
	} `xml:"modules"`
}

type Response struct {
	Name string `json:"name"`
}

func main() {
	app := &cli.App{
		Name:  "transform",
		Usage: "transform pom.xml to json",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Value:   "pom.xml",
				Usage:   "specific pom.xml file",
			},
		},
		Action: func(c *cli.Context) error {
			var res []Response
			fileLocation := c.String("path")
			data, err := ioutil.ReadFile(fileLocation)
			if err != nil {
				return err
			}
			pom := &Pom{}
			_ = xml.Unmarshal([]byte(data), &pom)
			for _, v := range pom.Modules.Module {
				tmp := Response{
					strings.Split(v, "/")[0],
				}
				res = append(res, tmp)
			}
			b, err := json.Marshal(&res)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
