package server

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"pdfannotations/config"
	"pdfannotations/pdfop"

	"strconv"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

//  download created pdf file with annotaions on it
func ExportDocument(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	presentationID := vars["file"]
	meetingID := vars["meeting"]
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Request to export full presentation with annotations")
	// @fixed: ConfVar should go global, store into a dict then load variables
	Filename := config.OUTPUT + presentationID + "-final/" + presentationID + ".pdf"
	pdfop.CreateFinal(meetingID, presentationID)
	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)

	Openfile, err := os.Open(Filename)
	if err != nil {
		log.WithFields(log.Fields{"file": Filename}).Panic("The file could not be generated")
		http.Error(writer, "File not found.", 404)
	}
	defer Openfile.Close() //Close after function return
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Returning the generated file in HTTP response")
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+presentationID+".pdf")
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
}

//  download created pdf file with annotaions on it
func GetPage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	PresentationID := vars["file"]
	MeetingID := vars["meeting"]
	Pagenum, _ := strconv.Atoi(vars["pagenum"])
	Filename := config.OUTPUT + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	pdfop.CreateFinal(MeetingID, PresentationID)

	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)
	Openfile, err := os.Open(Filename)
	if err != nil {
		log.WithFields(log.Fields{"file": Filename}).Panic("The file could not be generated")
		http.Error(writer, "File not found.", 404)
	}
	defer Openfile.Close() //Close after function return

	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
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

func ExportFromRaw(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	PresentationID := vars["file"]
	meetingID := vars["meeting"]
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": PresentationID}).Info("POST Request to export full presentation with annotations")
	// @fixed: ConfVar should go global, store into a dict then load variables
	Filename := config.OUTPUT + PresentationID + "-final/" + PresentationID + ".pdf"
	RawData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("Can't read data from POST request body")
	}

	pdfop.CreateFinalFromRaw(meetingID, PresentationID, RawData)
	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)
	Openfile, err := os.Open(Filename)
	if err != nil {
		log.WithFields(log.Fields{"file": Filename}).Panic("The file could not be generated")
		http.Error(writer, "File not found.", 404)
	}
	defer Openfile.Close() //Close after function return
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": PresentationID}).Info("Returning the generated file in HTTP response")
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+PresentationID+".pdf")
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
}

//handle server requsests
func HandleRequests() {
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/{meeting}/{file}", ExportDocument).Methods("GET")
	apiRouter.HandleFunc("/{meeting}/{file}", ExportFromRaw).Methods("POST")
	apiRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.WithFields(log.Fields{"port": config.PORT}).Info("Serving HTTP")
	log.Fatal(http.ListenAndServe(":"+config.PORT, apiRouter))
}
