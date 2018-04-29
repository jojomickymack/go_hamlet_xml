package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Play struct {
	Title      string       `xml:"title"`
	Playwright string       `xml:"playwright"`
	Pubdate    string       `xml:"pubdate"`
	Personae   PersonaeType `xml:"personae"`
	Acts       []Act        `xml:"act"`
}

type PersonaeType struct {
	Cast []PersonaType `xml:"persona"`
}

type PersonaType struct {
	Persname PersnameType `xml:"persname"`
}

type PersnameType struct {
	Name  string `xml:",chardata"`
	Short string `xml:"short,attr"`
}

type Act struct {
	Acttitle string      `xml:"acttitle"`
	Scenes   []SceneType `xml:"scene"`
}

type SceneType struct {
	Scenetitle    string       `xml:"scenetitle"`
	Scenelocation string       `xml:"scenelocation"`
	Scenepersonae []string     `xml:"scenepersonae"`
	Speeches      []SpeechType `xml:"speech"`
}

type SpeechType struct {
	Speaker string   `xml:"speaker"`
	Lines   []string `xml:"line"`
}

func (p Play) printPlayDetails() {
	fmt.Printf("the play is %s\n", p.Title)
}

func main() {
	xmlFile, err := os.Open("hamlet.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	data, _ := ioutil.ReadAll(xmlFile)

	var doc Play
	xml.Unmarshal(data, &doc)

	fmt.Printf("%s\n", doc.Title)
	fmt.Printf("%s\n", doc.Playwright)
	fmt.Printf("%s\n", doc.Pubdate)

	println("\nCast\n")

	for _, persona := range doc.Personae.Cast {
		fmt.Printf("%-20s\t%s\n", persona.Persname.Short, persona.Persname.Name)
	}

	println("\n")

	for _, act := range doc.Acts {
		println(act.Acttitle + "\n")
		for _, scene := range act.Scenes {
			for _, speech := range scene.Speeches {
				println(speech.Speaker)
				for _, line := range speech.Lines {
					println(line)
				}
				println()
			}
			println("\n")
		}
		println("\n\n")
	}
}
