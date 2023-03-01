package main

import (
	"flag"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Parse the command-line arguments
	contentPtr := flag.String("content", "", "the content to generate QR code for")
	filenamePtr := flag.String("filename", "qrcode.png", "the name of the output file")
	flag.Parse()

	// Check if the content is provided
	if *contentPtr == "" {
		log.Fatal("Please provide a content using the -content flag")
	}

	err := qrcode.WriteFile(*contentPtr, qrcode.Highest, 512, *filenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("QR code saved to %s\n", *filenamePtr)
}
