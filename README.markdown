# BigBlueButton PDF Annotations Server

## Development environment
If you want to use virtualisation to run the project, you need to have `vagrant` installed.

## Env vars

Add the following lines to your `.profile`

```bash
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
```

`source .profile`

## How to run?
`cd /app` or `app` alias if you are under vagrant

`go get github.com/gorilla/mux`

`go get github.com/jung-kurt/gofpdf`

`go get github.com/llgcode/draw2d`

`go get github.com/spf13/viper`

`go get github.com/sirupsen/logrus`

`go get github.com/01walid/goarabic`

`go get github.com/pdfcpu/pdfcpu`


Required packages for Python scripts:

`pip install CairoSVG`

`pip install PyPDF2`


Run with `go run main.go`


+ start local server on port 8100 
    - Listen for GET request on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}

        Response: Download pdf file (local location : /tmp/presentationID-final/presentationID.pdf)

    - Listen for POST request on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}
        example request (post.sh)
        Response: Download pdf file (local location : /tmp/presentationID-final/presentationID.pdf)

    - Listen for GET request on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}/{PageNumber}
        
        Response: Download pdf file of the wanted page only (local location : /tmp/presentationID-pages-done/presentationID_PageNumber.pdf)


+ Create pdf file of the desired presentation with annotations on it 
    + go fetch on /var/bigbluebutton/meetingID/meetingID/presentationID
        + if pdf file (presentationID.pdf) exists 
            - split the pdf file into single pdf pages and save them on /tmp/presentationID-pages
            + add annotations on every page 
                - get the annotations details of each page from /var/bigbluebutton/MeetingID/events.xml 
                - draw on each page it's annotations and save the page on /tmp/presentationID-done
            - merge the pdf pages on /tmp/presentationID-done on one pdf file and save in /tmp/presentationID-final
        + if the source presentation is in other format than pdf 
            - convert the svgs /var/bigbluebutton/meetingID/presentationID/svgs to pdf files 
              and save them to /tmp/presentationID-pages 
            + add annotations on every page 
                - get the annotations details of each page from /var/bigbluebutton/MeetingID/events.xml 
                - draw on each page it's annotations and save the page on /tmp/presentationID-done
            - merge the pdf pages on /tmp/presentationID-done on one pdf file and save in /tmp/presentationID-final
 

# PDF export API

Run API with  `go run main.go`

## Main options 
+ Download Presentation PDF
	* [From events.xml]   : `GET /MeetingID/PresentationID`
	* [From JSON body] : `GET /MeetingID/PresentationID`
  
+ Download one page from Presentation PDF
	* [from events.xml] : `GET /MeetingID/PresentationID/PageNumber`

  
## Configuration variables
* port (listening port for api , `8100`)
* OutputPath: (temp path , `/tmp/`)
* BBBPresPath: (presentation location , `/var/bigbluebutton/`)
* ScriptDir: (python scripts location , `/usr/share/bbb-api-pyscript/`)
* EventsPath: (events.xml file location , `/var/bigbluebutton/`)
* FontPath: ( Arial.ttf font location , `/usr/share/fonts/`)

## Example Requests:


* GET http://127.0.0.1:8100/a49e87847170b7bfa6dbd633004657683499034d-1634463938342/07bfb86c086066090f39d810096a329151261aec-1634463976664
  * > In the meeting `a49e87847170b7bfa6dbd633004657683499034d-1634463938342` the presentation pdf file `07bfb86c086066090f39d810096a329151261aec-1634463976664` will be downloaded with it's annotations extracted from:
  `EventsPath/a49e87847170b7bfa6dbd633004657683499034d-1634463938342/events.xml` 

* GET http://127.0.0.1:8100/a49e87847170b7bfa6dbd633004657683499034d-1634463938342/07bfb86c086066090f39d810096a329151261aec-1634463976664/1
  * > In the meeting `a49e87847170b7bfa6dbd633004657683499034d-1634463938342` the page number `1` from the presentation pdf  `07bfb86c086066090f39d810096a329151261aec-1634463976664` will be downloaded with it's annotations extracted from:
  `EventsPath/a49e87847170b7bfa6dbd633004657683499034d-1634463938342/events.xml` 


* POST exemple
  * >curl --request POST \
  --url http://127.0.0.1:8100/a49e87847170b7bfa6dbd633004657683499034d-1634463938342/07bfb86c086066090f39d810096a329151261aec-1634463976664 \
  --header 'Content-Type: application/json' \
  --data 
``` json '[
 {
  "_eventname": "AddShapeEvent",
  "presentation": "07bfb86c086066090f39d810096a329151261aec-1634463976664",
  "whiteboardId": "07bfb86c086066090f39d810096a329151261aec-1634463976664/1",
  "pageNumber": 0,
  "type": "pencil",
  "position": 0,
  "dataPoints": "20.768291,44.483734,28.794786,43.68109,36.86562,44.239834,44.91463,44.239834",
  "color": 16711680,
  "thickness": 0.3658536585365854,
  "dimensions": "List(547, 410)",
  "commands": "1,4"
 },
 {
  "_eventname": "AddShapeEvent",
  "presentation": "07bfb86c086066090f39d810096a329151261aec-1634463976664",
  "whiteboardId": "07bfb86c086066090f39d810096a329151261aec-1634463976664/2",
  "pageNumber": 1,
  "type": "text",
  "x": 52.42053561740452,
  "y": 52.905454282407405,
  "fontColor": 255,
  "textBoxWidth": 31.723717583550346,
  "textBoxHeight": 8.80195900245949,
  "text": "Test text",
  "fontSize": 20,
  "calcedFontSize": 4.88997555012225,
  "position": 0,
  "dataPoints": "52.42053561740452,52.905454282407405"
 }
]' 