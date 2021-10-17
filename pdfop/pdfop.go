package pdfop

import (
	"io/ioutil"
	"os"
	"os/exec"
	"pdfannotations/annotations"
	"pdfannotations/draw"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
	"github.com/spf13/viper"

	"github.com/jung-kurt/gofpdf"
)

const MM_TO_PX_RATIO = 0.352778

//func that returns string value from key
func ConfVar(key string) string {
	log.WithFields(log.Fields{"configKey": key}).Info("Reading configuration key")
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value := v.GetString(key)
	return value
}

// create presentation pdf file with annotations on it
func CreateFinal(meet string, filename string) {
	// @fixme: check if the used presentation exists before doing any conversion
	presPath := ConfVar("BBBPresPath") + meet + "/" + filename
	dirName := filename + "-pages"

	log.WithFields(log.Fields{"presPath": presPath, "folderName": dirName}).Info("Got request to create a PDF")

	if PdfExist(presPath + "/" + filename + ".pdf") {
		log.WithFields(log.Fields{"presPath": presPath, "folderName": dirName}).Info("PDF found and selected for processing")
		SplitPdf(presPath+"/"+filename+".pdf", ConfVar("OutputPath")+dirName)
		AddAnnotations(meet, ConfVar("OutputPath")+dirName)
		MergePdf(ConfVar("OutputPath")+dirName+"-done", filename)
	} else {
		log.WithFields(log.Fields{"presPath": presPath, "folderName": dirName}).Info("PDF not found and falling back to generated SVG")
		SvgToPdf(presPath+"/svgs", ConfVar("OutputPath")+dirName, filename)
		AddAnnotations(meet, ConfVar("OutputPath")+dirName)
		MergePdf(ConfVar("OutputPath")+dirName+"-done", filename)
	}
}

//split one pdf file into multiple pdf pages saved in specefic folder
func SplitPdf(fileName string, folderName string) {
	// We split the PDF into pages because of single document multi-orentation issue
	log.WithFields(log.Fields{"fileName": fileName, "folderName": folderName}).Info("Splitting the PDF ")
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			panic(err)
		}
	}
	// @fixme: the absolute path to put in config
	cmd := exec.Command("python3", "tools/script/split.py", "-i", fileName, "-o", folderName)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//convert svg images in one folder to pdf files saved in specefic folder
func SvgToPdf(fileName string, outputDir string, prefix string) {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	log.WithFields(log.Fields{"fileName": fileName, "outputDir": outputDir, "prefix": prefix}).Info("Converting SVG files to a single PDF")
	cmd := exec.Command("python3", "tools/script/svgtopdf.py", "-n", prefix, "-o", outputDir, "-p", fileName)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//add to all files in folder annotations
func AddAnnotations(meetingId string, dirDone string) {
	log.WithFields(log.Fields{"meetingId": meetingId, "dirDone": dirDone}).Info("Adding annotations to single pages")
	files, _ := ioutil.ReadDir(dirDone)
	if _, err := os.Stat(dirDone + "-done"); os.IsNotExist(err) {
		err := os.Mkdir(dirDone+"-done", 0755)
		if err != nil {
			panic(err)
		}
	}
	for _, f := range files {
		pageNumber := GetIntInBetweenStr(f.Name(), "_", ".pdf")
		err := InsertPage(meetingId, f.Name(), dirDone+"-done"+"/"+f.Name(), pageNumber)
		if err != nil {
			panic(err)
		}
	}
}

// add to one pdf file it's specefic annotations
func InsertPage(meetingId string, pageFileName string, outputDir string, pageNum int) error {
	var currentPage []annotations.Event
	presID := strings.Split(pageFileName, "_")[0]
	var s gofpdf.SizeType
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitMillimeter, gofpdf.PageSizeA4, "")
	tpl := gofpdi.ImportPage(pdf, ConfVar("OutputPath")+presID+"-pages/"+pageFileName, 1, "/MediaBox")
	pageSizes := gofpdi.GetPageSizes()
	s.Wd, s.Ht = pageSizes[1]["/MediaBox"]["w"]*MM_TO_PX_RATIO, pageSizes[1]["/MediaBox"]["h"]*MM_TO_PX_RATIO
	pdf.AddPageFormat(gofpdf.OrientationPortrait, s) // Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl, 0, 0, s.Wd, 0)
	currentPage = annotations.PageShapes(meetingId, presID, pageNum)
	log.WithFields(log.Fields{"meetingId": meetingId, "pageFileName": pageFileName, "pageNum": pageNum + 1}).Info("Generating page with annotation")
	for _, element := range currentPage {
		switch {
		case element.Type == "text":
			{
				text := annotations.TextDetails{
					X: element.X,
					Y: element.Y,
					Color: annotations.DEC{
						Decimal_color: int64(element.FontColor),
					}.Dec2RGB(),
					Width:      element.TextBoxWidth,
					Height:     element.TextBoxHeight,
					Text:       element.Text,
					CalcedSize: element.CalcedFontSize}
				draw.WriteText(pdf, text, s)
			}
		case element.Type == "line":
			{
				LineDetails := annotations.ShapeDetails{
					DataPoints: element.DataPoints,
					Color: annotations.DEC{
						Decimal_color: int64(element.Color),
					}.Dec2RGB(),
					Thickness: element.Thickness,
				}
				draw.DrawLine(pdf, LineDetails, s)
			}
		case element.Type == "ellipse":
			{
				EllipseDetails := annotations.ShapeDetails{
					DataPoints: element.DataPoints,
					Color: annotations.DEC{
						Decimal_color: int64(element.Color),
					}.Dec2RGB(),
					Thickness: element.Thickness,
				}
				draw.DrawEllipse(pdf, EllipseDetails, s)
			}
		case element.Type == "triangle":
			{
				TriangleDetails := annotations.ShapeDetails{
					DataPoints: element.DataPoints,
					Color: annotations.DEC{
						Decimal_color: int64(element.Color),
					}.Dec2RGB(),
					Thickness: element.Thickness,
				}
				draw.DrawTriangle(pdf, TriangleDetails, s)

			}
		case element.Type == "rectangle":
			{
				RectangleDetails := annotations.ShapeDetails{
					DataPoints: element.DataPoints,
					Color: annotations.DEC{
						Decimal_color: int64(element.Color),
					}.Dec2RGB(),
					Thickness: element.Thickness,
				}
				draw.DrawRectangle(pdf, RectangleDetails, s)
			}
		case element.Type == "pencil":
			{
				MyPencil := annotations.PencilDetails{
					DataPoints: element.DataPoints,
					Color: annotations.DEC{
						Decimal_color: int64(element.Color),
					}.Dec2RGBA(),
					Commands:  element.Commands,
					Thickness: element.Thickness,
				}
				draw.DrawPencil(pdf, MyPencil, s)
			}
		}
	}
	return pdf.OutputFileAndClose(outputDir)
}

//merge multiple pdf pages in a specefic directory in one pdf file
func MergePdf(folderName string, presId string) {
	log.WithFields(log.Fields{"presId": presId}).Info("Merging pdf pages")
	out_dir := folderName[:len(folderName)-11] + "-final"
	if _, err := os.Stat(out_dir); os.IsNotExist(err) {
		err := os.Mkdir(out_dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	cmd := exec.Command("python3", "tools/script/merge.py", "-p", folderName, "-o", out_dir+"/"+presId+".pdf", "-n", presId)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//check if a file exists
func PdfExist(filename string) bool {
	log.WithFields(log.Fields{"filename": filename}).Info("Checking PDF file exists")
	if _, err := os.Stat(filename); err == nil {
		return true
	} else /* os.IsNotExist(err)*/ {
		return false
	}
}

//used to get the page number of the selected filename
func GetIntInBetweenStr(str string, start string, end string) int {
	s := strings.Index(str, start)
	s += len(start)
	e := strings.Index(str, end)
	n, _ := strconv.Atoi(string(str[s:e]))
	return n
}
