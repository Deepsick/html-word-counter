package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Deepsick/page-loader/lib"

	"github.com/PuerkitoBio/goquery"
)

const (
	EXTENSION          = "html"
	OUTPUT_FOLDER_NAME = "output"
)

func main() {
	link, err := lib.GetLink(&os.Args)
	if err != nil {
		lib.LogToFile("Can't get a link from args", lib.LOGGER_FILE_PATH, err)
		log.Fatal(err)
	}

	validateErr := lib.ValidateUrl(link)
	if validateErr != nil {
		lib.LogToFile("Link is invalid", lib.LOGGER_FILE_PATH, validateErr)
		log.Fatal(validateErr)
	}

	htmlReader, htmlErr := lib.GetHtml(link)
	if htmlErr != nil {
		lib.LogToFile("Can't read html", lib.LOGGER_FILE_PATH, htmlErr)
		log.Fatal(htmlErr)
	}
	copiedHtmlReader, htmlErr := lib.GetHtml(link)
	if htmlErr != nil {
		lib.LogToFile("Can't read html", lib.LOGGER_FILE_PATH, htmlErr)
		log.Fatal(htmlErr)
	}
	doc, err := goquery.NewDocumentFromReader(*htmlReader)
	if err != nil {
		lib.LogToFile("Can't read html", lib.LOGGER_FILE_PATH, err)
		log.Fatal(err)
	}

	title := fmt.Sprintf("%s/%s.%s", OUTPUT_FOLDER_NAME, doc.Find("title").First().Text(), EXTENSION)
	lib.CreateFolder(OUTPUT_FOLDER_NAME)
	lib.WriteFile(*copiedHtmlReader, &title)

	wordCounters := lib.GetWordCounters(doc)
	pairs := lib.RankByWordCount(wordCounters)

	for _, pair := range pairs {
		println(pair.Key, pair.Value)
	}
}
