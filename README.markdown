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

go mod init pdfannotations
go mod tidy

go get -u github.com/gorilla/mux
go get -u github.com/jung-kurt/gofpdf
go get -u github.com/llgcode/draw2d

Run with `go run main.go`

+ start local server on port 8100 
    - wait for get requests on :
        http://127.0.0.1/8100/{meetingID}/{presentationID}
        Response: Download pdf file (local location : /tmp/presentationID-final/presentationID.pdf)

+ Create pdf file of the desired presentation with annotations on it 
    + go fetch on /var/bigbluebutton/meetingID/presentationID
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
 