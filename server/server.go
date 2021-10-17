package server

import (
	"io"
	"net/http"
	"os"
	"pdfannotations/pdfop"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

//  download created pdf file with annotaions on it
func ExportDocument(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	presentationID := vars["file"]
	meetingID := vars["meeting"]
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Request to export full presentation with annotations")
	// @fixme: ConfVar should go global, store into a dict then load variables
	Filename := pdfop.ConfVar("OutputPath") + presentationID + "-final/" + presentationID + ".pdf"
	pdfop.CreateFinal(meetingID, presentationID)
	//Check if file exists and open
	Openfile, err := os.Open(Filename)
	if err != nil {
		log.WithFields(log.Fields{"file": Filename}).Panic("The file could not be generated")
		http.Error(writer, "File not found.", 404)
	}
	defer Openfile.Close() //Close after function return
	//File is found, create and send the correct headers
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Returning the generated file in HTTP response")
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+presentationID+".pdf")
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
}

//  download created pdf file with annotaions on it
func GetPage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	PresentationID := vars["file"]
	MeetingID := vars["meeting"]
	Pagenum, _ := strconv.Atoi(vars["pagenum"])
	Filename := pdfop.ConfVar("OutputPath") + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	pdfop.CreateFinal(MeetingID, PresentationID)
	//Check if file exists and open
	Openfile, err := os.Open(Filename)
	if err != nil {
		log.WithFields(log.Fields{"file": Filename}).Panic("The file could not be generated")
		http.Error(writer, "File not found.", 404)
	}
	defer Openfile.Close() //Close after function return

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+PresentationID+"_"+strconv.Itoa(Pagenum)+".pdf")
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
}

//handle server requsests
func HandleRequests() {
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/{meeting}/{file}", ExportDocument).Methods("GET")
	apiRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.WithFields(log.Fields{"port": pdfop.ConfVar("port")}).Info("Serving HTTP")
	log.Fatal(http.ListenAndServe(":"+pdfop.ConfVar("port"), apiRouter))
}
