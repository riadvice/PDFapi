package server

import (
	"io"
<<<<<<< HEAD
<<<<<<< HEAD
	"io/ioutil"
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
	"log"
=======
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
	"log"
=======
	"io/ioutil"
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	"net/http"
	"os"
	"pdfannotations/config"
	"pdfannotations/pdfop"

	"strconv"

<<<<<<< HEAD
<<<<<<< HEAD
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	log "github.com/sirupsen/logrus"

<<<<<<< HEAD
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
=======
	log "github.com/sirupsen/logrus"

>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
=======
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	log "github.com/sirupsen/logrus"

>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	"github.com/gorilla/mux"
)

//  download created pdf file with annotaions on it
func ExportDocument(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
<<<<<<< HEAD
<<<<<<< HEAD
	presentationID := vars["file"]
	meetingID := vars["meeting"]
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Request to export full presentation with annotations")
	// @fixed: ConfVar should go global, store into a dict then load variables
	Filename := config.OUTPUT + presentationID + "-final/" + presentationID + ".pdf"
	pdfop.CreateFinal(meetingID, presentationID)
	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)

||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
	PresentationID := vars["file"]
	MeetingID := vars["meeting"]
	Filename := pdfop.ConfVar("TempPath") + PresentationID + "-final/" + PresentationID + ".pdf"
	if !pdfop.Pdf_exist(Filename) {
		pdfop.CreateFinal(MeetingID, PresentationID)
	}
	//Check if file exists and open
=======
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	presentationID := vars["file"]
	meetingID := vars["meeting"]
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Request to export full presentation with annotations")
	// @fixme: ConfVar should go global, store into a dict then load variables
	Filename := pdfop.ConfVar("OutputPath") + presentationID + "-final/" + presentationID + ".pdf"
	pdfop.CreateFinal(meetingID, presentationID)
	//Check if file exists and open
<<<<<<< HEAD
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
	PresentationID := vars["file"]
	MeetingID := vars["meeting"]
	Filename := pdfop.ConfVar("TempPath") + PresentationID + "-final/" + PresentationID + ".pdf"
	if !pdfop.Pdf_exist(Filename) {
		pdfop.CreateFinal(MeetingID, PresentationID)
	}
	//Check if file exists and open
=======
	presentationID := vars["file"]
	meetingID := vars["meeting"]
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Request to export full presentation with annotations")
	// @fixed: ConfVar should go global, store into a dict then load variables
	Filename := config.OUTPUT + presentationID + "-final/" + presentationID + ".pdf"
	pdfop.CreateFinal(meetingID, presentationID)
	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)

>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	Openfile, err := os.Open(Filename)
	if err != nil {
		log.WithFields(log.Fields{"file": Filename}).Panic("The file could not be generated")
		http.Error(writer, "File not found.", 404)
	}
	defer Openfile.Close() //Close after function return
<<<<<<< HEAD
<<<<<<< HEAD
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Returning the generated file in HTTP response")
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
=======
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	//File is found, create and send the correct headers
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Returning the generated file in HTTP response")
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
<<<<<<< HEAD
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
=======
	log.WithFields(log.Fields{"meetingId": meetingID, "presId": presentationID}).Info("Returning the generated file in HTTP response")
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
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
<<<<<<< HEAD
<<<<<<< HEAD
	Filename := config.OUTPUT + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	pdfop.CreateFinal(MeetingID, PresentationID)

	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
	Filename := pdfop.ConfVar("TempPath") + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	if !pdfop.Pdf_exist(Filename) {
		pdfop.CreateFinal(MeetingID, PresentationID)
	}
	//Check if file exists and open
=======
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	Filename := pdfop.ConfVar("OutputPath") + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	pdfop.CreateFinal(MeetingID, PresentationID)
	//Check if file exists and open
<<<<<<< HEAD
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
	Filename := pdfop.ConfVar("TempPath") + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	if !pdfop.Pdf_exist(Filename) {
		pdfop.CreateFinal(MeetingID, PresentationID)
	}
	//Check if file exists and open
=======
	Filename := config.OUTPUT + PresentationID + "-pages-done/" + PresentationID + "_" + strconv.Itoa(Pagenum-1) + ".pdf"
	pdfop.CreateFinal(MeetingID, PresentationID)

	log.Info("Optimizing PDF")
	pdfcpu.OptimizeFile(Filename, "", nil)
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
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
<<<<<<< HEAD
<<<<<<< HEAD
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/{meeting}/{file}", ExportDocument).Methods("GET")
	apiRouter.HandleFunc("/{meeting}/{file}", ExportFromRaw).Methods("POST")
	apiRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.WithFields(log.Fields{"port": config.PORT}).Info("Serving HTTP")
	log.Fatal(http.ListenAndServe(":"+config.PORT, apiRouter))
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/{meeting}/{file}", ExportDone).Methods("GET")
	myRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.Printf("Serving on HTTP port " + pdfop.ConfVar("port"))
	log.Fatal(http.ListenAndServe(pdfop.ConfVar("port"), myRouter))
=======
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/{meeting}/{file}", ExportDocument).Methods("GET")
	apiRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.WithFields(log.Fields{"port": pdfop.ConfVar("port")}).Info("Serving HTTP")
	log.Fatal(http.ListenAndServe(":"+pdfop.ConfVar("port"), apiRouter))
<<<<<<< HEAD
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/{meeting}/{file}", ExportDone).Methods("GET")
	myRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.Printf("Serving on HTTP port " + pdfop.ConfVar("port"))
	log.Fatal(http.ListenAndServe(pdfop.ConfVar("port"), myRouter))
=======
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/{meeting}/{file}", ExportDocument).Methods("GET")
	apiRouter.HandleFunc("/{meeting}/{file}", ExportFromRaw).Methods("POST")
	apiRouter.HandleFunc("/{meeting}/{file}/{pagenum}", GetPage).Methods("GET")
	log.WithFields(log.Fields{"port": config.PORT}).Info("Serving HTTP")
	log.Fatal(http.ListenAndServe(":"+config.PORT, apiRouter))
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
}
