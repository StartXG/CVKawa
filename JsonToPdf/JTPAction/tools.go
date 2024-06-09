package JTPAction

import (
	"JsonToPdf/common"
	"github.com/jung-kurt/gofpdf/v2"
	"strings"
	"time"
)

func unixMilliToTimeStr(timestamp int64, format string) string {
	secs := timestamp / 1000
	t := time.Unix(secs, 0)
	return t.Format(format)
}

func blockTitle(pdf *gofpdf.Fpdf, title string, pageSize string, fontName string) (*gofpdf.Fpdf, *nextPoint) {
	blockWidth := pdf.GetPageSizeStr(pageSize).Wd - 2*pdf.PointConvert(common.MmToPt(10))
	pdf.SetFont(fontName, "", Title2Size)
	pdf.CellFormat(blockWidth, Title2Size, title, "", 1, "L", false, 0, "")
	pdf.Ln(common.MmToPt(2))
	pdf.Line(pdf.GetX(), pdf.GetY(), pdf.GetX()+blockWidth, pdf.GetY()+2)
	pdf.Ln(common.MmToPt(2))
	return pdf, &nextPoint{pdf.GetX(), pdf.GetY()}
}

func splitStringRune(s string, size int) string {
	size = size + 1
	sArray := strings.Split(s, "\n")
	var result []string
	for _, v := range sArray {
		r := []rune(v)
		for i := 0; i < len(r); i += size {
			end := i + size
			if end > len(r) {
				end = len(r)
			}
			result = append(result, string(r[i:end]))
		}
	}
	return strings.Join(result, "\n")
}
