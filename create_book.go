package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/a-h/templ"
)

func toUpperCase(s string) string {
	allWords := ""
	for _, word := range strings.Split(s, "-") {

		allWords += strings.ToUpper(word[0:1]) + strings.ToLower(word[1:]) + " "

	}
	return strings.TrimSpace(allWords)
}

func createFile(ebook EBookMiddle, dirName string, fileName string) error {
	file, _ := os.Create(path.Join(dirName, fileName))

	var component templ.Component

	if strings.Contains(fileName, "nav") {
		component = Nav(ebook)
	} else if strings.Contains(fileName, "content") {
		component = Content(ebook)
	}

	component.Render(context.Background(), file)

	return nil
}

func (e *EBookMiddle) fixPayload() {
	e.Interior.Payload = strings.ReplaceAll(e.Interior.Payload, "<br>", "")
	e.Interior.Payload = strings.ReplaceAll(e.Interior.Payload, "&nbsp;", "")
	e.Interior.Payload = strings.ReplaceAll(e.Interior.Payload, "<hr>", "<p class=\"pagebreak\">* * *</p>")

}

func sectionTitleName(sectionTitle string) string {
	return toUpperCase(sectionTitle[strings.Index(sectionTitle, "_")+1:])
}

func createEPUBFolder(ebook EBookMiddle, bookName string) error {
	epubPath := path.Join(bookName, "EPUB")
	os.Mkdir(epubPath, 0755)

	ebook.fixPayload()

	createFile(ebook, epubPath, "content_001.xhtml")
	createFile(ebook, epubPath, "nav.xhtml")

	file, _ := os.Create(path.Join(epubPath, "package.opf"))
	file.WriteString(PackageOPF(ebook))

	styles, _ := os.Create(path.Join(epubPath, "styles.css"))
	styles.WriteString(Styles(ebook))

	return nil
}

func createMETA_INFFolder(ebook EBookMiddle, bookName string) error {
	folderPath := path.Join(bookName, "META-INF")
	os.Mkdir(folderPath, 0755)

	file, _ := os.Create(path.Join(folderPath, "container.xml"))
	file.WriteString(ContainerXML(ebook))

	return nil
}

func createHTMLBook(ebook EBookMiddle) error {

	authors := ""

	if len(ebook.Meta.Authors) > 0 {
		authors = ebook.Meta.Authors[0].Name
	}

	bookName := strings.ReplaceAll(ebook.Meta.Title, " ", "-") + "-" + authors + ".epub"

	dir := "books"

	bookName = path.Join(dir, bookName)

	os.Mkdir(bookName, 0755)

	createEPUBFolder(ebook, bookName)
	createMETA_INFFolder(ebook, bookName)

	file, _ := os.Create(path.Join(bookName, "mimetype"))
	file.WriteString("application/epub+zip")

	command := "mv"
	args := []string{bookName, path.Join("~", "Desktop", "bookName")}

	cmd := exec.Command(command, args...)

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
