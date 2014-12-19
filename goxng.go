package goxng

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"goxng/base64"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)


// Function creating a xng image.
// Takes a list of image path and a duration in ms
// Returns the xng content or an error
// The xng image size is the size of the first image of imgs
func GetXng(imgs []string, duration int) (string, error) {

	images := make([]imageStruct, len(imgs))
	set := make([]set, len(imgs))

	for i := 0; i < len(imgs); i++ {
		content, err := base64.ReadImage(imgs[i])
		if err != nil {
			return "", err
		}

		images[i].Ahref = content
		images[i].Id = strconv.Itoa(i)
		images[i].Height = "100%"

		set[i].Id = fmt.Sprintf("A%d", i)
		set[i].Ahref = fmt.Sprintf("#%d", i)
		set[i].AttributeName = "width"
		set[i].To = "100%"
		set[i].Dur = fmt.Sprintf("%dms", duration)
		set[i].Begin = getBegin(i, len(imgs))
	}

	// we use the size of the first image as xng's size
	wi, he, err := getImageSize(imgs[0])
	if err != nil {
		return "", err
	}

	document := &xng{Xmlns: "http://www.w3.org/2000/svg", Xmlnsa: "http://www.w3.org/1999/xlink", Width: wi, Height: he, Image: images, Set: set}

	xml, err := toXml(document)
	return xml, err
}

func getImageSize(filename string) (height, width int, err error) {
	reader, err := os.Open(filename)
	if err != nil {
		return
	}
	defer reader.Close()

	image, _, err := image.Decode(reader)
	if err != nil {
		return
	}
	bounds := image.Bounds()
	return bounds.Max.X, bounds.Max.Y, nil

}

func getBegin(i, size int) (begin string) {
	switch i {
	case 0:
		begin = fmt.Sprintf("A%d.end; 0s", size-1)
		break
	default:
		begin = fmt.Sprintf("A%d.end", i-1)
	}

	return begin

}

func toXml(document *xng) (xmlString string, err error) {
	var b = bytes.NewBuffer(nil)
	enc := xml.NewEncoder(b)
	enc.Indent("  ", "    ")
	if err := enc.Encode(document); err != nil {
		return "", err
	}
	return b.String(), nil
}

type xng struct {
	XMLName xml.Name      `xml:"svg"`
	Xmlns   string        `xml:"xmlns,attr"`
	Xmlnsa  string        `xml:"xmlns:A,attr"`
	Width   int           `xml:"width,attr"`
	Height  int           `xml:"height,attr"`
	Image   []imageStruct `xml:"image"`
	Set     []set         `xml:"set"`
}

type imageStruct struct {
	XMLName xml.Name `xml:"image"`
	Id      string   `xml:"id,attr"`
	Height  string   `xml:"height,attr"`
	Ahref   string   `xml:"A:href,attr"`
}

type set struct {
	Ahref         string `xml:"A:href,attr"`
	Id            string `xml:"id,attr"`
	AttributeName string `xml:"attributeName,attr"`
	To            string `xml:"to,attr"`
	Dur           string `xml:"dur,attr"`
	Begin         string `xml:"begin,attr"`
}

//        <svg xmlns="http://www.w3.org/2000/svg" xmlns:A="http://www.w3.org/1999/xlink" width="640" height="360">
//	<image id="000001" height="100%" A:href="data:image/jpeg;base64,/9j/4AAQSk…"/>
//	<image id="000002" height="100%" A:href="data:image/jpeg;base64,QbVFXX2Jn…"/>
//	<image id="000003" height="100%" A:href="data:image/jpeg;base64,/9j/4AAQ…"/>
//	<image id="000004" height="100%" A:href="data:image/jpeg;base64,1SOHxMt…"/>
//	<image id="000005" height="100%" A:href="data:image/jpeg;base64,4kjhtog…"/>
//
//	<set A:href="#000001" id="A000001" attributeName="width" to="100%" dur="33ms" begin="A000005.end; 0s"/>
//	<set A:href="#000002" id="A000002" attributeName="width" to="100%" dur="33ms" begin="A000001.end"/>
//	<set A:href="#000003" id="A000003" attributeName="width" to="100%" dur="33ms" begin="A000002.end"/>
//	<set A:href="#000004" id="A000004" attributeName="width" to="100%" dur="33ms" begin="A000003.end"/>
//	<set A:href="#000005" id="A000005" attributeName="width" to="100%" dur="33ms" begin="A000004.end"/>
//	</svg>
