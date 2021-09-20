package pdfop

import (
	"io/ioutil"
	"os"
	"os/exec"
	"pdfannotations/annotations"
	"pdfannotations/draw"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf/contrib/gofpdi"

	"github.com/jung-kurt/gofpdf"
)

// create presentation pdf file with annotations on it
func CreateFinal(meet string, filename string) {
	PresPath := "/var/bigbluebutton/" + meet + "/" + filename
	foldername := filename + "-pages"

	if Pdf_exist(PresPath + "/" + filename + ".pdf") {
		split_pdf(PresPath+"/"+filename+".pdf", "/tmp/"+foldername)
		AddAnnotations(meet, "/tmp/"+foldername)
		merge_pdf("/tmp/"+foldername+"-done", filename)
	} else {
		svg_to_pdf(PresPath+"/svgs", "/tmp/"+foldername, filename)
		AddAnnotations(meet, "/tmp/"+foldername)
		merge_pdf("/tmp/"+foldername+"-done", filename)
	}
}

//split one pdf file into multiple pdf pages saved in specefic folder
func split_pdf(filename string, foldername string) {
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		err := os.Mkdir(foldername, 0755)
		if err != nil {
			panic(err)
		}
	}
	cmd := exec.Command("tools/split", "-i", filename, "-o", foldername)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//convert svg images in one folder to pdf files saved in specefic folder
func svg_to_pdf(file string, output_dir string, prefix string) {
	if _, err := os.Stat(output_dir); os.IsNotExist(err) {
		err := os.Mkdir(output_dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	cmd := exec.Command("tools/svgtopdf", "-n", prefix, "-o", output_dir, "-p", file)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//add to all files in folder annotations
func AddAnnotations(meeting string, folder string) {
	files, _ := ioutil.ReadDir(folder)
	if _, err := os.Stat(folder + "-done"); os.IsNotExist(err) {
		err := os.Mkdir(folder+"-done", 0755)
		if err != nil {
			panic(err)
		}
	}
	for _, f := range files {
		pageN := GetIntInBetweenStr(f.Name(), "_", ".pdf")
		err := InsertPage(meeting, f.Name(), folder+"-done"+"/"+f.Name(), pageN)
		if err != nil {
			panic(err)
		}
	}
}

//add to one pdf file it's specefic annotations
func InsertPage(MeetingId string, input string, output string, pageNUM int) error {
	var ThisPage []annotations.Event
	presID := GetStringBeforeChar(input, "_")
	var s gofpdf.SizeType
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitMillimeter, gofpdf.PageSizeA4, "")
	tpl := gofpdi.ImportPage(pdf, "/tmp/"+presID+"-pages/"+input, 1, "/MediaBox")
	pageSizes := gofpdi.GetPageSizes()
	s.Wd, s.Ht = pageSizes[1]["/MediaBox"]["w"]*0.352778, pageSizes[1]["/MediaBox"]["h"]*0.352778
	pdf.AddPageFormat(gofpdf.OrientationPortrait, s) // Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl, 0, 0, s.Wd, 0)
	ThisPage = annotations.PageShapes(MeetingId, presID, pageNUM)
	for _, element := range ThisPage {
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
	return pdf.OutputFileAndClose(output)
}

//merge multiple pdf pages in a specefic directory in one pdf file
func merge_pdf(foldername string, presID string) {
	out_dir := foldername[:len(foldername)-11] + "-final"
	if _, err := os.Stat(out_dir); os.IsNotExist(err) {
		err := os.Mkdir(out_dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	cmd := exec.Command("tools/merge", "-p", foldername, "-o", out_dir+"/"+presID+".pdf", "-n", presID)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//check if a file exists
func Pdf_exist(filename string) bool {
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

//used to get the presentation id from filename
func GetStringBeforeChar(str string, end string) string {
	e := strings.Index(str, end)
	sub := string(str[0:e])
	return sub
}
