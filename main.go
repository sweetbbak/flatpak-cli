package main

import (
	"bufio"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jessevdk/go-flags"
	fzf "github.com/ktr0731/go-fuzzyfinder"
)

var opts struct {
	Download  bool   `short:"d" long:"download" description:"download App info file from flathub"`
	Appstream string `short:"a" long:"appstream" description:"path to appstream.xml file [default: uses the one on the system]"`
	Link      bool   `short:"l" long:"link" description:"export flatpaks into their CLI tool names by creating a shim"`
	Verbose   bool   `short:"v" long:"verbose" description:"print debugging information and verbose output"`
}

var Debug = func(string, ...interface{}) {}

var url = "https://hub.flathub.org/flathub/appstream/x86_64/appstream.xml.gz"

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

func install_it(pkgz []string) error {
	fmt.Println("Installing: ", pkgz)

	// idk why this happens but sometimes the names falsely have an extension
	for i := 0; i < len(pkgz); i++ {
		if strings.Contains(pkgz[i], ".desktop") {
			pkgz[i] = strings.ReplaceAll(pkgz[i], ".desktop", "")
		}
	}

	args := append([]string{"install", "-y"}, pkgz...)
	cmd := exec.Command("flatpak", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not install %s: \n%s\n", pkgz, err)
	}
	return nil
}

// Multi package install -----------------------

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [Y/n]: ", s)

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

func appstream_fallback() error {
	BASEURL := "https://flathub.org/repo/appstream"
	ARCH := "x86_64"
	url := strings.Join([]string{BASEURL, ARCH, "appstream.xml.gz"}, "/")
	fmt.Println("Downloading Flatpak database: ", url)
	out, err := os.Create("appstream.xml.gz")
	if err != nil {
		return err
	}

	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	appstream, err := os.Create("appstream.xml")
	if err != nil {
		return err
	}

	r, err := gzip.NewReader(resp.Body)
	defer r.Close()
	_, err = io.Copy(appstream, r)
	return err
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func create_shims() error {
	// list flatpak apps and ID's to auto-generate file executable shims
	cmd := exec.Command("flatpak", "list", "--app", "--columns=name")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	x := []string{}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line != "" {
			x = append(x, line)
		}
	}

	cmd2 := exec.Command("flatpak", "list", "--app", "--columns=app")
	output2, err := cmd2.Output()
	if err != nil {
		return err
	}

	y := []string{}

	lines2 := strings.Split(string(output2), "\n")
	for _, line := range lines2 {
		if line != "" {
			y = append(y, line)
			// fmt.Println(line)
		}
	}

	flatpak_bin := "flatpak-bin"
	_, err = os.Stat(flatpak_bin)
	if err != nil {
		os.MkdirAll(flatpak_bin, 0755)

	} else {
		fmt.Println("Directory already exists...")
		fmt.Println("\x1b[33mThis will overwrite already existing exported flatpak shims.\x1b[0m")
		c := askForConfirmation("Overwrite? ")
		if !c {
			return fmt.Errorf("Please rename bin folder, or manually resolve this issue, then re-run.")
		} else {
			err := os.RemoveAll(flatpak_bin)
			if err != nil {
				return err
			}
			os.MkdirAll(flatpak_bin, 0755)
		}
	}

	workdir, _ := os.Getwd()
	for i := 0; i < len(x); i++ {
		exe := strings.ToLower(x[i])
		exe = strings.ReplaceAll(exe, " ", "_")
		exe = strings.ReplaceAll(exe, "'", "")
		exe = strings.ReplaceAll(exe, "\\", "")
		exe = trimQuotes(exe)

		// remove flatpak bin from env first
		// before looking for conflicting binary names
		in_path, _ := exec.LookPath(exe)
		if in_path != "" {
			fmt.Println("Conflicting binary names")
			fmt.Println(exe)
		}

		path := filepath.Join(workdir, flatpak_bin, exe)
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("File created successfully: ", file.Name())
		defer file.Close()

		script := []string{"!#/bin/bash", y[i]}
		data_string := strings.Join(script, "\n")
		data := []byte(data_string)

		err = os.WriteFile(path, data, 0755)
		os.Chmod(path, 0755)
	}
	return nil
}

func curl_xml() {
	c := exec.Command("curl", "-Sl", "https://raw.githubusercontent.com/sweetbbak/flatpak-cli/main/appstream.xml")
	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
	os.WriteFile("appstream.xml", b, 0644)
}

func merge(ms ...map[string]string) map[string][]string {
	res := map[string][]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = append(res[k], v)
		}
	}
	return res
}

func get_apps() {
	cmd := exec.Command("flatpak", "search", "--columns", "name", "")
	fmt.Println(cmd.Args)
	output2, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		// os.Exit(0)
	}

	y := []string{}

	lines2 := strings.Split(string(output2), "\n")
	for _, line := range lines2 {
		if line != "" {
			y = append(y, line)
		}
	}

	cmd1 := exec.Command("flatpak", "search", "--columns", "application", "")
	output1, err := cmd1.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		// os.Exit(0)
	}

	x := []string{}
	lines := strings.Split(string(output1), "\n")
	for _, line := range lines {
		if line != "" {
			x = append(x, line)
		}
	}

	ms := make(map[string]string)
	for i := 0; i < len(x); i++ {
		ms[y[i]] = x[i]
	}
	fmt.Println(ms)

}

func findXml() (string, error) {
	base := "/var/lib/flatpak/appstream"
	var fpath string

	filepath.WalkDir(base, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if !d.Type().IsRegular() {
			return nil
		}

		f := filepath.Base(path)
		if f == "appstream.xml" {
			fpath = path
		}
		return nil
	})
	if fpath == "" {
		return "", fmt.Errorf("appstream.xml not found")
	}

	return fpath, nil
}

func Flatpak(args []string) error {
	if opts.Link {
		err := create_shims()
		if err != nil {
			return err
		}
		return nil
	}

	var xml_file *os.File
	flatpakFile := "/var/lib/flatpak/appstream/flathub/x86_64/active/appstream.xml"
	if opts.Appstream != "" {
		flatpakFile = opts.Appstream
	}

	_, err := os.Stat(flatpakFile)
	if err == nil {
		xml_file, err = os.Open(flatpakFile)
	} else {
		return fmt.Errorf("Flatpak Database not found")
	}

	defer xml_file.Close()
	var component Comp

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xml_file)
	err = xml.Unmarshal(byteValue, &component)
	if err != nil {
		return err
	}

	idx, err := fzf.FindMulti(
		component.App,
		func(i int) string {
			return component.App[i].Name
		},
		fzf.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return "Hello there :)"
			}
			return fmt.Sprintf("Name: %s (%s)\nSummary: %s\n%s\n",
				component.App[i].Name,
				component.App[i].ID,
				component.App[i].Description,
				component.App[i].Summary)
		}))
	if err == fzf.ErrAbort {
		os.Exit(0)
	}

	package_list_from_index := []string{}
	for i := 0; i < len(idx); i++ {
		package_list_from_index = append(package_list_from_index, component.App[idx[i]].ID)
	}

	fmt.Println(arrayToString(package_list_from_index))

	c := askForConfirmation("Would you like to install: ")
	if c {
		fmt.Println("OKAY :)")
		install_it(package_list_from_index)
	} else {
		fmt.Println("OKAY Maybe next time :)")
		os.Exit(0)
	}

	return nil
}

func main() {
	args, err := flags.Parse(&opts)
	if flags.WroteHelp(err) {
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}

	if opts.Verbose {
		Debug = log.Printf
	}

	if err := Flatpak(args); err != nil {
		log.Fatal(err)
	}
}
