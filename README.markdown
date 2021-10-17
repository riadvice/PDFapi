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

<<<<<<< HEAD
`go get github.com/sirupsen/logrus`

`go get github.com/01walid/goarabic`

`go get github.com/pdfcpu/pdfcpu`

||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
=======
`go get github.com/sirupsen/logrus`

>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)

Required packages for Python scripts:

`pip install CairoSVG`

`pip install PyPDF2`


Run with `go run main.go`


+ start local server on port 8100 
<<<<<<< HEAD
    - Listen for GET request on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
    - Listen on :
        http://127.0.0.1/8100/{meetingID}/{presentationID}
=======
    - Listen on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)

        Response: Download pdf file (local location : /tmp/presentationID-final/presentationID.pdf)
<<<<<<< HEAD

    - Listen for POST request on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}
        example request (post.sh)
        Response: Download pdf file (local location : /tmp/presentationID-final/presentationID.pdf)

    - Listen for GET request on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}/{PageNumber}
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
    - Listen on :
        http://127.0.0.1/8100/{meetingID}/{presentationID}/{PageNumber}
=======
    - Listen on :
        http://127.0.0.1:8100/{meetingID}/{presentationID}/{PageNumber}
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
        
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
 