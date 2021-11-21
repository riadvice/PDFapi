package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

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

// convert xml to json
func PageShapes(MeetingID string, PresentationID string /*, PN string*/) /*[]Event*/ {
	var data Recording
	var InPage []Event
	//PageNum, _ := strconv.Atoi(PN)
	rawXmlData, _ := ioutil.ReadFile("/var/bigbluebutton/" + MeetingID + "/events.xml")
	xml.Unmarshal([]byte(rawXmlData), &data)
	//spew.Dump(data)
	var k = 0
	for _, found := range data.Event {
		if (found.Eventname) == "AddShapeEvent" && found.Presentation == PresentationID /*&& found.PageNumber == PageNum*/ {
			InPage = append(InPage, found)
			k++
		}
		if (found.Eventname) == "UndoAnnotationEvent" && found.Presentation == PresentationID /*&& found.PageNumber == PageNum*/ {
			InPage = append(InPage[:k-1], InPage[k:]...)
			k -= 2
		}
		if (found.Eventname) == "ClearWhiteboardEvent" && found.Presentation == PresentationID /*&& found.PageNumber == PageNum*/ {
			InPage = nil //InPage[:0]
			k = 0
		}
	}
	file, _ := json.MarshalIndent(InPage, "", " ")
	ioutil.WriteFile("AddShape.json", file, 0644)
	//return (InPage)
}

func main() {
	//fmt.Println(os.Args)
	PageShapes(os.Args[1], os.Args[2])
}

/*
xml data input form :
<Recording>
	<meeting id="xxxxxxxxxxxxxxxxxxx-xxxxx" />
	<event eventname="------------">
		<presentation>		</presentation>
		<whiteboardid>		</whiteboardid>
		<pagenumber>		</pagenumber>
		<type>				</type>
		<x>					</x>
		<y>					</y>
		<fontcolor>			</fontcolor>
		<textboxwidth>		</textboxwidth>
		<textboxheight>		<textboxheight>
		<text>				</text>
		<fontsize>      	</fontsize>
		<calcedfontsize>	</calcedfontsize>
		<position>      	</position>
		<datapoints>    	</datapoints>
		<color>         	</color>
		<thickness>			</thickness>
		<dimensions>		</dimensions>
		<commands>			</commands>
	</event>
</recording>

expected output :
{
  "pages": [
    {
      "number": "1",
      "property": "value",
      "annotations": [
        {
          "id": "test",
          "x": "85.1",
          "y": "11.7",
          "property": "...."
        }
      ]
    }
  ]
}

working output :
[
	{
		"_eventname"
		"presentation"
		"whiteboardid"
		"pagenumber"
		"type"
		"x"
		"y"
		"fontcolor"
		"textboxwidth"
		"textboxheight"
		"text"
		"fontsize"
		"calcedfontsize"
		"position"
		"datapoints"
		"color"
		"thickness"
		"dimensions"
		"commands"

	},
	{

	},
	{

	}
]
*/
