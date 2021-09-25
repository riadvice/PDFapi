package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"pdfannotations/pdfop"
	"strconv"

	"github.com/gorilla/mux"
)

//  download created pdf file with annotaions on it
func ExportDone(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	PresentationID := vars["file"]
	MeetingID := vars["meeting"]
	Filename := pdfop.ConfVar("TempPath") + PresentationID + "-final/" + PresentationID + ".pdf"
	if !pdfop.Pdf_exist(Filename) {
		pdfop.CreateFinal(MeetingID, PresentationID)
	}
	//Check if file exists and open
	Openfile, err := os.Open(Filename)
	if err != nil {
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
	writer.Header().Set("Content-Disposition", "attachment; filename="+PresentationID+".pdf")
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
	Filename := pdfop.ConfVar("TempPath") + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	if !pdfop.Pdf_exist(Filename) {
		pdfop.CreateFinal(MeetingID, PresentationID)
	}
	//Check if file exists and open
	Openfile, err := os.Open(Filename)
	if err != nil {
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
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/{meeting}/{file}", ExportDone).Methods("GET")
	myRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.Printf("Serving on HTTP port " + pdfop.ConfVar("port"))
	log.Fatal(http.ListenAndServe(pdfop.ConfVar("port"), myRouter))
}
