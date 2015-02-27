package main

import (
	"flag"
	"fmt"
	"github.com/masayukioguni/go-webp-sample/webp"
	"os"
)

func main() {
	jpgfile := flag.String("i", "", "jpeg file name")
	webpfile := flag.String("o", "", "webp file name")
	lossless := flag.Bool("l", true, "Lossless true/false")
	quality := flag.Int("q", 50, "quality")

	flag.Parse()

	fmt.Printf("%v %v %v %v", *jpgfile, *webpfile, *lossless, float32(*quality))

	f, err := os.Open(*jpgfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open failed %v  %v", *jpgfile, err)
		return
	}

	defer f.Close()

	m, _ := webp.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "webp.Decode failed %v", err)
		return
	}

	toimg, _ := os.Create(*webpfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Create failed %v %v", *webpfile, err)
		return
	}
	defer toimg.Close()

	err = webp.Encode(toimg, m, &webp.Options{Lossless: *lossless, Quality: float32(*quality)})
	if err != nil {
		fmt.Fprintf(os.Stderr, "webp.Encode failed %v %v %v", *lossless, *quality, err)
		return
	}

}
