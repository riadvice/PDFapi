package annotations

import (
	"encoding/json"
	"encoding/xml"
	"image/color"
	"io/ioutil"
	"log"
	"pdfannotations/config"
	"strconv"
)

// Parent node of events
type Recording struct {
	Meeting   Meeting `xml:"meeting" json:"meeting"`
	MeetingID string  `xml:"meeting_id,attr" json:"_meeting_id"`
	Event     []Event `xml:"event" json:"event"`
}

// First child of recording := meeting
type Meeting struct {
	ID string `xml:"id,attr" json:"_id"`
}

// Second child of recording := Event
type Event struct {
	Eventname      string  `xml:"eventname,attr" json:"_eventname"`
	Presentation   string  `xml:"presentation" json:"presentation"`
	WhiteboardID   string  `xml:"whiteboardId" json:"whiteboardId"`
	PageNumber     int     `xml:"pageNumber" json:"pageNumber"`
	Type           string  `xml:"type" json:"type"`
	X              float64 `xml:"x" json:"x,omitempty"`
	Y              float64 `xml:"y" json:"y,omitempty"`
	FontColor      int     `xml:"fontColor" json:"fontColor,omitempty"`
	TextBoxWidth   float64 `xml:"textBoxWidth" json:"textBoxWidth,omitempty"`
	TextBoxHeight  float64 `xml:"textBoxHeight" json:"textBoxHeight,omitempty"`
	Text           string  `xml:"text" json:"text,omitempty"`
	FontSize       int     `xml:"fontSize" json:"fontSize,omitempty"`
	CalcedFontSize float64 `xml:"calcedFontSize" json:"calcedFontSize,omitempty"`
	Position       int     `xml:"position" json:"position"`
	DataPoints     string  `xml:"dataPoints" json:"dataPoints"`
	Color          int     `xml:"color" json:"color,omitempty"`
	Thickness      float64 `xml:"thickness" json:"thickness,omitempty"`
	Dimensions     string  `xml:"dimensions" json:"dimensions,omitempty"`
	Commands       string  `xml:"commands" json:"commands,omitempty"`
}

// one event elements shape structure (line, ellipse, triangle, rectangle)
type ShapeDetails struct {
	DataPoints string
	Color      RGB
	Thickness  float64
}

// pencil event details
type PencilDetails struct {
	Commands   string
	DataPoints string
	Color      color.RGBA
	Thickness  float64
}

// text event details
type TextDetails struct {
	X          float64
	Y          float64
	Color      RGB
	Width      float64
	Height     float64
	Text       string
	FontSize   int
	CalcedSize float64
}

//colors are in decimal format so we have to adapt them to ather formats
type DEC struct {
	Decimal_color int64
}

// red green blue format of color
type RGB struct {
	Red, Green, Blue int
}

// function to transform colors from decimal format to RGB format
func (Dec_c DEC) Dec2RGB() RGB {

	hex := string(strconv.FormatInt(Dec_c.Decimal_color, 16))
	for len(hex) < 6 {
		hex = "0" + hex
	}
	R, _ := strconv.ParseInt(hex[:2], 16, 10)
	G, _ := strconv.ParseInt(hex[2:4], 16, 18)
	B, _ := strconv.ParseInt(hex[4:], 16, 10)
	return RGB{int(R), int(G), int(B)}
}

// function to transform colors from decimal format to RGBA format
func (Dec_c DEC) Dec2RGBA() (c color.RGBA) {
	hex := string(strconv.FormatInt(Dec_c.Decimal_color, 16))
	for len(hex) < 6 {
		hex = "0" + hex
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}
	c.R = hexToByte(hex[0])<<4 + hexToByte(hex[1])
	c.G = hexToByte(hex[2])<<4 + hexToByte(hex[3])
	c.B = hexToByte(hex[4])<<4 + hexToByte(hex[5])
	c.A = 0xff
	return c
}

//i/o events.xml
func PageShapes(MeetingID string, PresentationID string, PageNum int) []Event {
	var data Recording
	var InPage []Event
<<<<<<< HEAD
<<<<<<< HEAD
	rawXmlData, err := ioutil.ReadFile(config.EVENTS + MeetingID + "/events.xml")
	if err != nil {
		log.Fatal("Can't read events.xml file")
	}
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
	rawXmlData, _ := ioutil.ReadFile(ConfVar("BBB") + MeetingID + "/events.xml")
=======
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	rawXmlData, _ := ioutil.ReadFile(ConfVar("BBBPresPath") + MeetingID + "/events.xml")
<<<<<<< HEAD
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
	rawXmlData, _ := ioutil.ReadFile(ConfVar("BBB") + MeetingID + "/events.xml")
=======
	rawXmlData, err := ioutil.ReadFile(config.EVENTS + MeetingID + "/events.xml")
	if err != nil {
		log.Fatal("Can't read events.xml file")
	}
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	xml.Unmarshal([]byte(rawXmlData), &data)
	var k = 0
	for _, found := range data.Event {
		if (found.Eventname) == "AddShapeEvent" && found.Presentation == PresentationID && found.PageNumber == PageNum {
			InPage = append(InPage, found)
			k++
		}
		if (found.Eventname) == "UndoAnnotationEvent" && found.Presentation == PresentationID && found.PageNumber == PageNum {
			InPage = append(InPage[:k-1], InPage[k:]...)
			k -= 2
		}
		if (found.Eventname) == "ClearWhiteboardEvent" && found.Presentation == PresentationID && found.PageNumber == PageNum {
			InPage = nil //InPage[:0]
			k = 0
		}
	}
	return (InPage)
}

//i/o events.xml
func PageShapesFromRaw(MeetingID string, PresentationID string, PageNum int, Raw []byte) []Event {
	var InPage []Event
	var result []Event

	json.Unmarshal([]byte(Raw), &InPage)

	var k = 0
	for _, found := range InPage {
		if (found.Eventname) == "AddShapeEvent" && found.Presentation == PresentationID && found.PageNumber == PageNum {
			result = append(result, found)
			k++
		}
		if (found.Eventname) == "UndoAnnotationEvent" && found.Presentation == PresentationID && found.PageNumber == PageNum {
			result = append(result[:k-1], result[k:]...)
			k -= 2
		}
		if (found.Eventname) == "ClearWhiteboardEvent" && found.Presentation == PresentationID && found.PageNumber == PageNum {
			result = nil //InPage[:0]
			k = 0
		}
	}
	return (result)
}
