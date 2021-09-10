package draw

import (
	"encoding/json"
	"math"
	"pdfannotations/annotations"

	"github.com/jung-kurt/gofpdf"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dpdf"
)

// coordinates array in form of string
type Coordinates struct {
	str string
}

// function to return array of integers in string to an array of float
func (str Coordinates) getPoints() []float64 {
	var points []float64
	str.str = "[" + str.str + "]"
	err := json.Unmarshal([]byte(str.str), &points)
	if err != nil {
		panic(err)
	}
	return points
}

// function to return array of integers in string to an array of integer
func (str Coordinates) getCommands() []int {
	var points []int
	str.str = "[" + str.str + "]"
	err := json.Unmarshal([]byte(str.str), &points)
	if err != nil {
		panic(err)
	}
	return points
}

//function to draw line to pdf in exact position
func DrawLine(pdf *gofpdf.Fpdf, line annotations.ShapeDetails, size gofpdf.SizeType) {
	pdf.SetDrawColor(line.Color.Red, line.Color.Green, line.Color.Blue)
	pdf.SetLineWidth(getStrokeWidth(line.Thickness, size.Wd))
	points := (Coordinates{line.DataPoints}).getPoints()
	p0 := denormalizeCoord(points[0], size.Wd)
	p1 := denormalizeCoord(points[1], size.Ht)
	p2 := denormalizeCoord(points[2], size.Wd)
	p3 := denormalizeCoord(points[3], size.Ht)
	pdf.Line(p0, p1, p2, p3)
}

//function to draw ellipse to pdf in proportional coordinates based on the page size
func DrawEllipse(pdf *gofpdf.Fpdf, ellipse annotations.ShapeDetails, size gofpdf.SizeType) {
	pdf.SetDrawColor(ellipse.Color.Red, ellipse.Color.Green, ellipse.Color.Blue)
	pdf.SetLineWidth(getStrokeWidth(ellipse.Thickness, size.Wd))
	points := (Coordinates{ellipse.DataPoints}).getPoints()
	// rx - horizontal radius, ry - vertical radius .. cx and cy - coordinates of the ellipse's center

	rx := (points[2] - points[0]) / 2
	ry := (points[3] - points[1]) / 2
	cx := denormalizeCoord(rx+points[0], size.Wd)
	cy := denormalizeCoord(ry+points[1], size.Ht)
	rx = denormalizeCoord(math.Abs((points[2]-points[0])/2), size.Wd)
	ry = denormalizeCoord(math.Abs((points[3]-points[1])/2), size.Ht)

	pdf.Ellipse(cx, cy, rx, ry, 0, "")
}

//function to draw triangle to pdf in proportional coordinates based on the page size
func DrawTriangle(pdf *gofpdf.Fpdf, triangle annotations.ShapeDetails, size gofpdf.SizeType) {
	pdf.SetDrawColor(triangle.Color.Red, triangle.Color.Green, triangle.Color.Blue)
	pdf.SetLineWidth(getStrokeWidth(triangle.Thickness, size.Wd))
	points := (Coordinates{triangle.DataPoints}).getPoints()

	xTop, yTop := denormalizeCoord(((points[2]-points[0])/2)+points[0], size.Wd), denormalizeCoord(points[1], size.Ht)
	xBottomLeft, yBottomLeft := denormalizeCoord(points[0], size.Wd), denormalizeCoord(points[3], size.Ht)
	xBottomRight, yBottomRight := denormalizeCoord(points[2], size.Wd), denormalizeCoord(points[3], size.Ht)

	pdf.Polygon([]gofpdf.PointType{{X: xBottomLeft, Y: yBottomLeft}, {X: xBottomRight, Y: yBottomRight}, {X: xTop, Y: yTop}, {X: xBottomLeft, Y: yBottomLeft}, {X: xBottomRight, Y: yBottomRight}}, "")
}

//function to draw rectangle to pdf in proportional coordinates based on the page size
func DrawRectangle(pdf *gofpdf.Fpdf, rectangle annotations.ShapeDetails, size gofpdf.SizeType) {
	pdf.SetDrawColor(rectangle.Color.Red, rectangle.Color.Green, rectangle.Color.Blue)
	pdf.SetLineWidth(getStrokeWidth(rectangle.Thickness, size.Wd))
	points := (Coordinates{rectangle.DataPoints}).getPoints()

	x := denormalizeCoord(points[0], size.Wd)
	y := denormalizeCoord(points[1], size.Ht)
	width := denormalizeCoord((points[2] - points[0]), size.Wd)
	height := denormalizeCoord((points[3] - points[1]), size.Ht)

	pdf.Rect(x, y, width, height, "")
}

//function to write text to pdf in  proportional coordinates based on the page size
func WriteText(pdf *gofpdf.Fpdf, text annotations.TextDetails, size gofpdf.SizeType) {
	pdf.SetTextColor(text.Color.Red, text.Color.Green, text.Color.Blue)
	pdf.SetFont("Arial", "", 0)
	pdf.SetFontUnitSize((text.CalcedSize / 100) * size.Ht)
	BoxWidth := denormalizeCoord(text.Width, size.Wd)
	BoxHeight := denormalizeCoord(text.Height, size.Ht)
	pdf.MoveTo(denormalizeCoord(text.X, size.Wd), denormalizeCoord(text.Y, size.Ht))
	pdf.SetAutoPageBreak(false, 0)
	pdf.MultiCell(BoxWidth, BoxHeight, text.Text, gofpdf.BorderNone, gofpdf.AlignTop, false)
}

//function to draw pencil shape to pdf in proportional coordinates based on the page size
func DrawPencil(pdf *gofpdf.Fpdf, pencil annotations.PencilDetails, size gofpdf.SizeType) {
	commands := (Coordinates{pencil.Commands}).getCommands()
	points := (Coordinates{pencil.DataPoints}).getPoints()
	gc := draw2dpdf.NewGraphicContext(pdf)
	gc.SetStrokeColor(pencil.Color)
	gc.SetLineWidth(getStrokeWidth(pencil.Thickness, size.Wd))
	gc.SetLineCap(draw2d.RoundCap)

	for i, j := 0, 0; i < len(commands); i += 1 {
		switch commands[i] {
		case 1:
			gc.MoveTo(denormalizeCoord(points[j], size.Wd), denormalizeCoord(points[j+1], size.Ht))
			j += 2
		case 2:
			gc.LineTo(denormalizeCoord(points[j], size.Wd), denormalizeCoord(points[j+1], size.Ht))
			j += 2
		case 3:
			gc.QuadCurveTo(
				denormalizeCoord(points[j], size.Wd),
				denormalizeCoord(points[j+1], size.Ht),
				denormalizeCoord(points[j+2], size.Wd),
				denormalizeCoord(points[j+3], size.Ht))
			j += 4
		case 4:
			gc.CubicCurveTo(
				denormalizeCoord(points[j], size.Wd),
				denormalizeCoord(points[j+1], size.Ht),
				denormalizeCoord(points[j+2], size.Wd),
				denormalizeCoord(points[j+3], size.Ht),
				denormalizeCoord(points[j+4], size.Wd),
				denormalizeCoord(points[j+5], size.Ht))
			j += 6
		default:
		}
	}
	gc.Stroke()
	gc.Close()
}

//get calculated coordinates based on percantage of coordinates on page size
func denormalizeCoord(normCoord float64, sideLength float64) float64 {
	return ((normCoord / 100) * sideLength)
}

//get calculated stroke width based on percantage of coordinates on page size
func getStrokeWidth(thickness float64, slideWidth float64) float64 {
	return (thickness * slideWidth) / 100
}
