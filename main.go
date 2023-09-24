package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	// "unsafe"

	fzf "github.com/ktr0731/go-fuzzyfinder"
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

	if strings.Contains(pkg, ".desktop") {
		pkg = strings.ReplaceAll(pkg, ".desktop", "")
	}

	fmt.Println("Installing: ", pkg)

	// r, w := io.Pipe()
	// cmd.Stdin = r
	cmd := exec.Command("flatpak", "install", "-y", "--noninteractive", pkg)

	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("could not install %s: \n%s\n", pkg, err)
	}
}

// Multi package install -----------------------
func arrayToString(arr []string) string {
	return strings.Join([]string(arr), "\n")
}

func install_it(pkgz []string) {
	fmt.Println("Installing: ", pkgz)

	for i := 0; i < len(pkgz); i++ {
		// fmt.Println(pkgz[i])
		if strings.Contains(pkgz[i], ".desktop") {
			pkgz[i] = strings.ReplaceAll(pkgz[i], ".desktop", "")
		}
	}

	// pkg := strings.Join([]string(pkgz), " ")
	// cmdstr := strings.Join([]string{"install", "-y", "--noninteractive", pkg}, " ")
	// cmd := exec.Command("flatpak", "install", pkgz...)
	args := append([]string{"install", "-y"}, pkgz...)
	// cmdd := exec.Cmd{
	// 	"flatpak",
	// 	args,
	// }
	// cmd := exec.Cmd("flatpak", "install", pkgz...)
	cmd := exec.Command("flatpak", args...)
	// cmd := exec.Cmd(cmdd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("could not install %s: \n%s\n", pkgz, err)
	}
}

// Multi package install -----------------------

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if response == "" {
			return true
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" || response == "\n" || response == "" {
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

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func create_shims() {
	path := os.Getenv("PATH")
	fmt.Println(path)

	cmd := exec.Command("flatpak", "list", "--app", "--columns=name")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	x := []string{}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line != "" {
			x = append(x, line)
			fmt.Println(line)
		}
	}

	fmt.Println(x)

	cmd2 := exec.Command("flatpak", "list", "--app", "--columns=app")
	output2, err := cmd2.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	y := []string{}

	lines2 := strings.Split(string(output2), "\n")
	for _, line := range lines2 {
		if line != "" {
			y = append(y, line)
			fmt.Println(line)
		}
	}

	fmt.Println(y)

	os.MkdirAll("bin", 0755)

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
		fmt.Println(y[i])

		workdir, _ := os.Getwd()

		exe := strings.ToLower(x[i])
		exe = strings.ReplaceAll(exe, " ", "_")
		exe = trimQuotes(exe)

		path := filepath.Join(workdir, "bin", exe)

		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("File created successfully: ", file)
		defer file.Close()

		script := []string{"!#/bin/bash", y[i]}
		data_string := strings.Join(script, "\n")
		data := []byte(data_string)

		err = os.WriteFile(path, data, 0755)
		os.Chmod(path, 0755)

	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--link" {
			create_shims()
			os.Exit(0)
		} else if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println("go-flatpak")
			fmt.Println("\nRun go-flatpak with no arguments to install a package using a fzf")
			fmt.Println("--help|-h\t\tshow this help message")
			fmt.Println("--link\t\t\tcreate a bin directory and export shorthand executable files")
			fmt.Println("\nex: org.blender.Blender - creates an executable file called \"blender\"")
			os.Exit(0)
		}
	}

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
		log.Fatal(err)
	}

	package_list_from_index := []string{}
	for i := 0; i < len(idx); i++ {
		package_list_from_index = append(package_list_from_index, component.App[idx[i]].ID)
	}

	fmt.Println(arrayToString(package_list_from_index))
	c := askForConfirmation("Would you like to install: ")
	if c {
		fmt.Println("OKAY :)")
		// install(choice)
		// package_list_from_index := []string{}
		// for i := 0; i < len(idx); i++ {
		// 	package_list_from_index = append(package_list_from_index, x[idx[i]])
		// }
		install_it(package_list_from_index)

	} else {
		fmt.Println("OKAY Maybe next time :)")
		os.Exit(0)
	}
}
