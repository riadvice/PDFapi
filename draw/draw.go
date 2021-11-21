package draw

import (
	"encoding/json"
	"fmt"
	"math"
	"pdfannotations/annotations"
	"unicode"

	"github.com/01walid/goarabic"
	log "github.com/sirupsen/logrus"

	log "github.com/sirupsen/logrus"

	"github.com/jung-kurt/gofpdf"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dpdf"
)

// coordinates array in form of string
type Coordinates struct {
	str string
}

// function to return array of integers in string to an array of float
func (str Coordinates) GetPoints() []float64 {
	var points []float64
	str.str = "[" + str.str + "]"
	err := json.Unmarshal([]byte(str.str), &points)
	if err != nil {
		panic(err)
	}
	return points
}

// function to return array of integers in string to an array of integer
func (str Coordinates) GetCommands() []int {
	var points []int
	str.str = "[" + str.str + "]"
	err := json.Unmarshal([]byte(str.str), &points)
	if err != nil {
		panic(err)
	}
	return points
}

// function to draw line to pdf in exact position
func DrawLine(pdf *gofpdf.Fpdf, line annotations.ShapeDetails, size gofpdf.SizeType) {
	log.Info("Drawing line")
	pdf.SetDrawColor(line.Color.Red, line.Color.Green, line.Color.Blue)
	pdf.SetLineWidth(GetStrokeWidth(line.Thickness, size.Wd))
	points := (Coordinates{line.DataPoints}).GetPoints()
	p0 := DenormalizeCoord(points[0], size.Wd)
	p1 := DenormalizeCoord(points[1], size.Ht)
	p2 := DenormalizeCoord(points[2], size.Wd)
	p3 := DenormalizeCoord(points[3], size.Ht)
	pdf.Line(p0, p1, p2, p3)
}

// function to draw ellipse to pdf in proportional coordinates based on the page size
func DrawEllipse(pdf *gofpdf.Fpdf, ellipse annotations.ShapeDetails, size gofpdf.SizeType) {
	log.Info("Drawing ellipse")
	pdf.SetDrawColor(ellipse.Color.Red, ellipse.Color.Green, ellipse.Color.Blue)
	pdf.SetLineWidth(GetStrokeWidth(ellipse.Thickness, size.Wd))
	points := (Coordinates{ellipse.DataPoints}).GetPoints()
	// rx - horizontal radius, ry - vertical radius .. cx and cy - coordinates of the ellipse's center

	rx := (points[2] - points[0]) / 2
	ry := (points[3] - points[1]) / 2
	cx := DenormalizeCoord(rx+points[0], size.Wd)
	cy := DenormalizeCoord(ry+points[1], size.Ht)
	rx = DenormalizeCoord(math.Abs((points[2]-points[0])/2), size.Wd)
	ry = DenormalizeCoord(math.Abs((points[3]-points[1])/2), size.Ht)

	pdf.Ellipse(cx, cy, rx, ry, 0, "")
}

// function to draw triangle to pdf in proportional coordinates based on the page size
func DrawTriangle(pdf *gofpdf.Fpdf, triangle annotations.ShapeDetails, size gofpdf.SizeType) {
	log.Info("Drawing triangle")
	pdf.SetDrawColor(triangle.Color.Red, triangle.Color.Green, triangle.Color.Blue)
	pdf.SetLineWidth(GetStrokeWidth(triangle.Thickness, size.Wd))
	points := (Coordinates{triangle.DataPoints}).GetPoints()

	xTop, yTop := DenormalizeCoord(((points[2]-points[0])/2)+points[0], size.Wd), DenormalizeCoord(points[1], size.Ht)
	xBottomLeft, yBottomLeft := DenormalizeCoord(points[0], size.Wd), DenormalizeCoord(points[3], size.Ht)
	xBottomRight, yBottomRight := DenormalizeCoord(points[2], size.Wd), DenormalizeCoord(points[3], size.Ht)

	pdf.Polygon([]gofpdf.PointType{{X: xBottomLeft, Y: yBottomLeft}, {X: xBottomRight, Y: yBottomRight}, {X: xTop, Y: yTop}, {X: xBottomLeft, Y: yBottomLeft}, {X: xBottomRight, Y: yBottomRight}}, "")
}

// function to draw rectangle to pdf in proportional coordinates based on the page size
func DrawRectangle(pdf *gofpdf.Fpdf, rectangle annotations.ShapeDetails, size gofpdf.SizeType) {
	log.Info("Drawing rectangle")
	pdf.SetDrawColor(rectangle.Color.Red, rectangle.Color.Green, rectangle.Color.Blue)
	pdf.SetLineWidth(GetStrokeWidth(rectangle.Thickness, size.Wd))
	points := (Coordinates{rectangle.DataPoints}).GetPoints()

	x := DenormalizeCoord(points[0], size.Wd)
	y := DenormalizeCoord(points[1], size.Ht)
	width := DenormalizeCoord((points[2] - points[0]), size.Wd)
	height := DenormalizeCoord((points[3] - points[1]), size.Ht)

	pdf.Rect(x, y, width, height, "")
}

<<<<<<< HEAD
<<<<<<< HEAD
// checkifArabic is Arabic characters or not.
func checkifArabic(input string) bool {

	var isArabic = false

	for _, v := range input {
		if unicode.In(v, unicode.Arabic) {
			isArabic = true
		} else {
			isArabic = false
		}
	}
	return isArabic
}

||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
// function to write text to pdf in  proportional coordinates based on the page size
<<<<<<< HEAD

//move cursor to x and y
//split text sting into multiple lines based on text box height
//insert each line as text cell and check if out of box boundaries

||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
//function to write text to pdf in  proportional coordinates based on the page size
=======
// function to write text to pdf in  proportional coordinates based on the page size
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
//function to write text to pdf in  proportional coordinates based on the page size
=======
// checkifArabic is Arabic characters or not.
func checkifArabic(input string) bool {

	var isArabic = false

	for _, v := range input {
		if unicode.In(v, unicode.Arabic) {
			isArabic = true
		} else {
			isArabic = false
		}
	}
	return isArabic
}

// function to write text to pdf in  proportional coordinates based on the page size

//move cursor to x and y
//split text sting into multiple lines based on text box height
//insert each line as text cell and check if out of box boundaries

>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
func WriteText(pdf *gofpdf.Fpdf, text annotations.TextDetails, size gofpdf.SizeType) {
<<<<<<< HEAD
	log.Info("Writing text")
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
=======
	log.Info("Writing text")
	pdf.SetTextColor(text.Color.Red, text.Color.Green, text.Color.Blue)
	pdf.SetFont("Arial-0", "", 0)
	pdf.SetFontUnitSize(denormalizeFont(text.CalcedSize, size.Ht))
	x := DenormalizeCoord(text.X, size.Wd)
	y := DenormalizeCoord(text.Y, size.Ht)
	BoxWidth := DenormalizeCoord(text.Width, size.Wd)
	BoxHeight := DenormalizeCoord(text.Height, size.Ht)
	pdf.SetAutoPageBreak(false, 0)
	/*
		pdf.SetFillColor(255, 0, 0)
		pdf.Circle(x, y, 1.0, "F")
		pdf.SetDrawColor(0, 255, 0)
		pdf.Rect(x, y, BoxWidth, BoxHeight, "D")
	*/
	if text.Text != "" {
		if checkifArabic(text.Text) {
			pdf.RTL()
			lineHt, _ := pdf.GetFontSize()
			pdf.MoveTo(x-1, y+1.7)
			lines := pdf.SplitText((text.Text), size.Wd)
			for i, line := range lines {
				if (float64(i) * lineHt) < BoxHeight {
					pdf.CellFormat(BoxWidth, lineHt, goarabic.ToGlyph(string(line)), "", 2, "LT", false, 0, "")
				} else {
					break
				}
			}
		} else {
			pdf.LTR()
			lineHt, _ := pdf.GetFontSize()
			lineHt = lineHt * 0.4055
			lines := pdf.SplitLines([]byte(text.Text), size.Wd)
			pdf.MoveTo(x-1, y+1.7)
			for i, line := range lines {
				if (float64(i) * lineHt) < BoxHeight {
					pdf.CellFormat(BoxWidth, lineHt, string(line), "", 2, "LT", false, 0, "")
				} else {
					break
				}

			}
		}
	}
}

/*
func oldWrite(pdf *gofpdf.Fpdf, text annotations.TextDetails, size gofpdf.SizeType) {
	log.Info("Writing text")
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
	pdf.SetTextColor(text.Color.Red, text.Color.Green, text.Color.Blue)
	pdf.SetFont("Arial-0", "", 0)
	pdf.SetFontUnitSize(denormalizeFont(text.CalcedSize, size.Ht))
	x := DenormalizeCoord(text.X, size.Wd)
	y := DenormalizeCoord(text.Y, size.Ht)
	BoxWidth := DenormalizeCoord(text.Width, size.Wd)
	BoxHeight := DenormalizeCoord(text.Height, size.Ht)
	pdf.SetAutoPageBreak(false, 0)
	/*
		pdf.SetFillColor(255, 0, 0)
		pdf.Circle(x, y, 1.0, "F")
		pdf.SetDrawColor(0, 255, 0)
		pdf.Rect(x, y, BoxWidth, BoxHeight, "D")
	*/
	if text.Text != "" {
		if checkifArabic(text.Text) {
			pdf.RTL()
			lineHt, _ := pdf.GetFontSize()
			pdf.MoveTo(x-1, y+1.7)
			lines := pdf.SplitText((text.Text), size.Wd)
			for i, line := range lines {
				if (float64(i) * lineHt) < BoxHeight {
					pdf.CellFormat(BoxWidth, lineHt, goarabic.ToGlyph(string(line)), "", 2, "LT", false, 0, "")
				} else {
					break
				}
			}
		} else {
			pdf.LTR()
			lineHt, _ := pdf.GetFontSize()
			lineHt = lineHt * 0.4055
			lines := pdf.SplitLines([]byte(text.Text), size.Wd)
			pdf.MoveTo(x-1, y+1.7)
			for i, line := range lines {
				if (float64(i) * lineHt) < BoxHeight {
					pdf.CellFormat(BoxWidth, lineHt, string(line), "", 2, "LT", false, 0, "")
				} else {
					break
				}

			}
		}
	}
}

/*
func oldWrite(pdf *gofpdf.Fpdf, text annotations.TextDetails, size gofpdf.SizeType) {
	log.Info("Writing text")
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
=======
	log.Info("Writing text")
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
	pdf.SetTextColor(text.Color.Red, text.Color.Green, text.Color.Blue)
	pdf.SetFont("Arial", "", 0)
	pdf.SetFontUnitSize((text.CalcedSize / 100) * size.Ht)
	BoxWidth := DenormalizeCoord(text.Width, size.Wd)
	BoxHeight := DenormalizeCoord(text.Height, size.Ht)
	pdf.MoveTo(DenormalizeCoord(text.X, size.Wd), DenormalizeCoord(text.Y, size.Ht))
	pdf.SetAutoPageBreak(false, 0)
	if checkifArabic(text.Text) {
		pdf.MultiCell(BoxWidth, BoxHeight, goarabic.ToGlyph(text.Text), gofpdf.BorderNone, gofpdf.AlignTop, false)
	} else {
		pdf.LTR()
		pdf.MultiCell(BoxWidth, BoxHeight, text.Text, gofpdf.BorderNone, gofpdf.AlignTop, false)
	}
}
*/

// function to draw pencil shape to pdf in proportional coordinates based on the page size
func DrawPencil(pdf *gofpdf.Fpdf, pencil annotations.PencilDetails, size gofpdf.SizeType) {
	log.Info("Drawing pencil")
	commands := (Coordinates{pencil.Commands}).GetCommands()
	points := (Coordinates{pencil.DataPoints}).GetPoints()
	gc := draw2dpdf.NewGraphicContext(pdf)
	gc.SetStrokeColor(pencil.Color)
	gc.SetLineWidth(GetStrokeWidth(pencil.Thickness, size.Wd))
	gc.SetLineCap(draw2d.RoundCap)

	for i, j := 0, 0; i < len(commands); i += 1 {
		switch commands[i] {
		case 1:
			gc.MoveTo(DenormalizeCoord(points[j], size.Wd), DenormalizeCoord(points[j+1], size.Ht))
			j += 2
		case 2:
			gc.LineTo(DenormalizeCoord(points[j], size.Wd), DenormalizeCoord(points[j+1], size.Ht))
			j += 2
		case 3:
			gc.QuadCurveTo(
				DenormalizeCoord(points[j], size.Wd),
				DenormalizeCoord(points[j+1], size.Ht),
				DenormalizeCoord(points[j+2], size.Wd),
				DenormalizeCoord(points[j+3], size.Ht))
			j += 4
		case 4:
			gc.CubicCurveTo(
				DenormalizeCoord(points[j], size.Wd),
				DenormalizeCoord(points[j+1], size.Ht),
				DenormalizeCoord(points[j+2], size.Wd),
				DenormalizeCoord(points[j+3], size.Ht),
				DenormalizeCoord(points[j+4], size.Wd),
				DenormalizeCoord(points[j+5], size.Ht))
			j += 6
		default:
		}
	}
	gc.Stroke()
	gc.Close()
}

// get calculated coordinates based on percantage of coordinates on page size
func DenormalizeCoord(normCoord float64, sideLength float64) float64 {
	return ((normCoord * 0.01) * sideLength)
}

//get calculated stroke width based on percantage of coordinates on page size
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
func GetStrokeWidth(thickness float64, slideWidth float64) float64 {
	return (thickness * slideWidth) * 0.01
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
func getStrokeWidth(thickness float64, slideWidth float64) float64 {
	return (thickness * slideWidth) / 100
=======
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
func GetStrokeWidth(thickness float64, slideWidth float64) float64 {
	return (thickness * slideWidth) * 0.01
}
<<<<<<< HEAD

//get normalized font size based on percantage of page size
func denormalizeFont(calced float64, slideHeight float64) float64 {
	return (calced * slideHeight) / 100
}

//draw grid as visual refernce for move to
func DrawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(0, 0, 100)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) { //del Letter size
		pdf.Line(x, 0, x, h)
		_, lineHeight := pdf.GetFontSize()
		pdf.Text(x, lineHeight, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y = y + (w / 20.0) { //del Letter size
		pdf.Line(0, y, w, y)
		//_, lineHeight := pdf.GetFontSize()
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}

||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
func getStrokeWidth(thickness float64, slideWidth float64) float64 {
	return (thickness * slideWidth) / 100
=======
func GetStrokeWidth(thickness float64, slideWidth float64) float64 {
	return (thickness * slideWidth) * 0.01
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
}
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======

//get normalized font size based on percantage of page size
func denormalizeFont(calced float64, slideHeight float64) float64 {
	return (calced * slideHeight) / 100
}

//draw grid as visual refernce for move to
func DrawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(0, 0, 100)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) { //del Letter size
		pdf.Line(x, 0, x, h)
		_, lineHeight := pdf.GetFontSize()
		pdf.Text(x, lineHeight, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y = y + (w / 20.0) { //del Letter size
		pdf.Line(0, y, w, y)
		//_, lineHeight := pdf.GetFontSize()
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}

>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
}
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
