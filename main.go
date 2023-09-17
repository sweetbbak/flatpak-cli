package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Comp struct {
	XMLName xml.Name `xml:"components"`
	App     []struct {
		ID          string `xml:"id"`
		Name        string `xml:"name"`
		Summary     string `xml:"summary"`
		Description string `xml:"description"`
	} `xml:"component"`
}

type track struct {
	ID      string
	Name    string
	Summary string
}

func fuzzy(component *Comp) {
	idx, err := fzf.FindMulti(
		component.App,
		func(i int) string {
			return component.App[i].Name
		},
		fzf.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Name: %s (%s)\nSummary: %s",
				component.App[i].Name,
				component.App[i].ID,
				component.App[i].Summary)
		}))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("selected: %v\n", idx)
	fmt.Printf("flatpak install %s\n", component.App[idx[0]].ID)
}

func install(pkg string) {
	fmt.Println("Installing: ", pkg)

	// r, w := io.Pipe()
	cmd := exec.Command("flatpak", "install", "-y", "--noninteractive", pkg)
	// cmd.Stdin = r

	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Sprintf("could not install %s: \n%s\n", pkg, err)
	}
}

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" || response == "\n" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func parse_xml() {
	xml_file, err := os.Open("/var/lib/flatpak/appstream/flathub/x86_64/active/appstream.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xml_file.Close()

	var component Comp

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xml_file)
	xml.Unmarshal(byteValue, &component)

	idx, err := fzf.FindMulti(
		component.App,
		func(i int) string {
			return component.App[i].Name
		},
		fzf.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Name: %s (%s)\nSummary: %s\n%s\n",
				component.App[i].Name,
				component.App[i].ID,
				component.App[i].Description,
				component.App[i].Summary)
		}))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(idx)
	choice := component.App[idx[0]].ID
	fmt.Println(choice)
}

func main() {
	home := os.Getenv("HOME")
	fmt.Println(home)

	xml_file, err := os.Open("/var/lib/flatpak/appstream/flathub/x86_64/active/appstream.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xml_file.Close()

	var component Comp

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xml_file)

	// fmt.Println(xml.Unmarshal(byteValue, &component))
	xml.Unmarshal(byteValue, &component)

	// fuzzy(&component)

	// for i := 0; i < len(component.App); i++ {
	// 	fmt.Println("Name: " + component.App[i].Name)
	// 	fmt.Println("ID: " + component.App[i].ID)
	// 	fmt.Println("Summary: " + component.App[i].Summary)
	// 	fmt.Println("")
	// }

	idx, err := fzf.FindMulti(
		component.App,
		func(i int) string {
			return component.App[i].Name
		},
		fzf.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Name: %s (%s)\nSummary: %s\n%s\n",
				component.App[i].Name,
				component.App[i].ID,
				component.App[i].Description,
				component.App[i].Summary)
		}))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("selected: %v\n", idx)
	// fmt.Printf("flatpak install %s\n", component.App[idx[0]].ID)
	fmt.Println(idx)
	choice := component.App[idx[0]].ID
	fmt.Println(choice)

	c := askForConfirmation("Would you like to install: " + choice)
	if c {
		fmt.Println("OKAY :)")
		install(choice)
	} else {
		fmt.Println("OKAY Maybe next time :)")
		os.Exit(0)
	}
}
