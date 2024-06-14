package main

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/a-h/templ"
)

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

func createEPUBFolder(ebook EBookMiddle, bookName string) error {
	epubPath := path.Join(bookName, "EPUB")
	os.Mkdir(epubPath, 0755)

	ebook.Interior.Payload = strings.ReplaceAll(ebook.Interior.Payload, "<br>", "")

	createFile(ebook, epubPath, "content_001.xhtml")
	createFile(ebook, epubPath, "nav.xhtml")

	file, _ := os.Create(path.Join(epubPath, "package.opf"))
	file.WriteString(PackageOPF(ebook))

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
	bookName := strings.ReplaceAll(ebook.Meta.Title, " ", "-") + "-" + strings.ReplaceAll(ebook.Meta.Authors[0].Name, " ", "-")

	os.Mkdir(bookName, 0755)

	createEPUBFolder(ebook, bookName)
	createMETA_INFFolder(ebook, bookName)

	file, _ := os.Create(path.Join(bookName, "mimetype"))
	file.WriteString("application/epub+zip")

	return nil
}
