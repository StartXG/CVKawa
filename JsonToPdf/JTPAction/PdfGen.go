package JTPAction

import (
	"JsonToPdf/common"
	"JsonToPdf/language"
	"embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
)

// PdfObj PdfObj struct
// orientationStr: P, L
// unitStr: mm, cm, in, pt
// sizeStr: A3, A4, A5, Letter, Legal
// fontDirStr: font directory
type PdfObj struct {
	OrientationStr string `default:"P"`
	UnitStr        string `default:"pt"`
	SizeStr        string `default:"A4"`
	FontDirStr     string `default:""`
}

type nextPoint struct {
	X float64
	Y float64
}

const (
	Title1Size  float64 = 18
	Title2Size  float64 = 14
	Title3Size  float64 = 12
	ContentSize float64 = 10
	RegularFont string  = "AlibabaPuHuiTi-Regular"
	BoldFont    string  = "AlibabaPuHuiTi-Bold"
)

//go:embed fontDir/*
var fontFS embed.FS

func (pObj PdfObj) New() *gofpdf.Fpdf {
	regularFontBytes, err := fontFS.ReadFile("fontDir/AlibabaPuHuiTi-Regular.ttf")
	if err != nil {
		panic(err)
	}
	boldFontBytes, err := fontFS.ReadFile("fontDir/AlibabaPuHuiTi-Bold.ttf")
	if err != nil {
		panic(err)
	}
	pdf := gofpdf.New(pObj.OrientationStr, pObj.UnitStr, pObj.SizeStr, pObj.FontDirStr)
	pdf.AddUTF8FontFromBytes(RegularFont, "", regularFontBytes)
	pdf.AddUTF8FontFromBytes(BoldFont, "", boldFontBytes)
	pdf.SetAutoPageBreak(true, common.MmToPt(10))
	pdf.AddPage()
	return pdf
}

func addBaseInfo(pdf *gofpdf.Fpdf, data *common.CvBaseInfo, lang string) (*gofpdf.Fpdf, *nextPoint) {
	divWidth := pdf.GetPageSizeStr("a4").Wd - 2*pdf.PointConvert(common.MmToPt(10)) - common.MmToPt(25.4)
	divTopHeight := (pdf.PointConvert(common.MmToPt(35)) - Title1Size - ContentSize) / 2
	pdf.SetXY(common.MmToPt(10), common.MmToPt(10)+divTopHeight)
	pdf.SetFont(BoldFont, "", Title1Size)
	pdf.CellFormat(
		divWidth,
		Title1Size,
		data.Name,
		"",
		1,
		"L",
		false,
		0,
		"",
	)
	pdf.Ln(common.MmToPt(2))
	// 添加第二行文本，不加粗
	pdf.SetFont(RegularFont, "", ContentSize)
	pdf.CellFormat(
		divWidth,
		ContentSize,
		strings.Join([]string{
			data.Email,
			strconv.Itoa(int(data.Mobile)),
			unixMilliToTimeStr(data.Birthday, language.GetLangFile(lang, "DateFullFormat")),
		}, " | "),
		"",
		0,
		"L",
		false,
		0,
		"",
	)

	imgX := pdf.GetX()

	pdf.Image(
		data.PhotoPath,
		imgX,
		common.MmToPt(10),
		common.MmToPt(25.4),
		common.MmToPt(35),
		false,
		"PNG",
		0,
		"",
	)
	pdf.SetXY(common.MmToPt(10), common.MmToPt(10)+common.MmToPt(35)+common.MmToPt(5))
	return pdf, &nextPoint{X: common.MmToPt(10), Y: common.MmToPt(10) + common.MmToPt(35) + common.MmToPt(5)}
}

func addEduInfo(pdf *gofpdf.Fpdf, data *[]common.CvEduInfo, c *nextPoint, lang string) (*gofpdf.Fpdf, *nextPoint) {
	if len(*data) == 0 {
		return pdf, c
	}
	pdf, c = blockTitle(pdf, language.GetLangFile(lang, "Education"), "a4", BoldFont)
	rowWidth := pdf.GetPageSizeStr("a4").Wd - 2*pdf.PointConvert(common.MmToPt(10))
	for _, v := range *data {
		pdf.SetFont(BoldFont, "", Title3Size)
		pdf.CellFormat(
			rowWidth,
			Title3Size,
			v.School,
			"",
			1,
			"L",
			false,
			0,
			"",
		)
		pdf.Ln(common.MmToPt(1))
		pdf.SetFont(RegularFont, "", ContentSize)
		pdf.CellFormat(
			rowWidth/2,
			ContentSize,
			strings.Join([]string{v.AcademicDegree, v.Major}, " - "),
			"",
			0,
			"L",
			false,
			0,
			"",
		)
		pdf.CellFormat(
			rowWidth/2,
			ContentSize,
			strings.Join([]string{unixMilliToTimeStr(v.StartTime, language.GetLangFile(lang, "DateFormat")), unixMilliToTimeStr(v.EndTime, language.GetLangFile(lang, "DateFormat"))}, " - "),
			"",
			0,
			"R",
			false,
			0,
			"",
		)
		pdf.Ln(common.MmToPt(3))
	}
	pdf.Ln(common.MmToPt(10))
	return pdf, &nextPoint{pdf.GetX(), pdf.GetY()}
}

func addSkills(pdf *gofpdf.Fpdf, data *common.CvSkills, c *nextPoint, lang string) (*gofpdf.Fpdf, *nextPoint) {
	if len(*data) == 0 {
		return pdf, c
	}
	pdf, c = blockTitle(pdf, language.GetLangFile(lang, "Skills"), "a4", BoldFont)
	pdf.SetFont(RegularFont, "", ContentSize)
	for k, v := range *data {
		pdf.CellFormat(
			0,
			ContentSize,
			fmt.Sprintf("%d. %s", k+1, v),
			"",
			1,
			"L",
			false,
			0,
			"",
		)
		pdf.Ln(common.MmToPt(2))
	}
	pdf.Ln(common.MmToPt(8))
	return pdf, &nextPoint{pdf.GetX(), pdf.GetY()}
}

func addResume(pdf *gofpdf.Fpdf, data *[]common.CvResume, c *nextPoint, lang string) (*gofpdf.Fpdf, *nextPoint) {
	if len(*data) == 0 {
		return pdf, c
	}
	pdf, c = blockTitle(pdf, language.GetLangFile(lang, "Resume"), "a4", BoldFont)
	rowWidth := pdf.GetPageSizeStr("a4").Wd - (2 * pdf.PointConvert(common.MmToPt(10)))
	for _, v := range *data {
		pdf.SetFont(BoldFont, "", Title3Size)
		pdf.CellFormat(
			rowWidth,
			Title3Size,
			v.Company,
			"",
			1,
			"L",
			false,
			0,
			"",
		)
		pdf.Ln(common.MmToPt(1))
		pdf.SetFont(BoldFont, "", ContentSize)
		pdf.CellFormat(
			rowWidth,
			Title3Size,
			language.GetLangFile(lang, "ResumeSteps"),
			"",
			1,
			"L",
			false,
			0,
			"",
		)
		pdf.Ln(common.MmToPt(1))
		for _, step := range v.ResumeSteps {
			pdf.SetFont(RegularFont, "", ContentSize)
			pdf.CellFormat(
				rowWidth/2,
				ContentSize,
				strings.Join([]string{step.Position, step.Department}, " | "),
				"",
				0,
				"L",
				false,
				0,
				"",
			)
			pdf.CellFormat(
				rowWidth/2,
				ContentSize,
				strings.Join([]string{unixMilliToTimeStr(step.StartTime, language.GetLangFile(lang, "DateFormat")), unixMilliToTimeStr(step.EndTime, language.GetLangFile(lang, "DateFormat"))}, " - "),
				"",
				1,
				"R",
				false,
				0,
				"",
			)
			pdf.Ln(common.MmToPt(1))
			// manual line break
			c := splitStringRune(step.Responsibilities, int(math.Floor(rowWidth/ContentSize)))
			pdf.MultiCell(
				rowWidth,
				ContentSize+2,
				c,
				"",
				"L",
				false,
			)
			pdf.Ln(common.MmToPt(3))
		}
		pdf.Ln(common.MmToPt(5))
		pdf.SetFont(BoldFont, "", ContentSize)
		pdf.CellFormat(
			rowWidth,
			ContentSize,
			language.GetLangFile(lang, "Achievement"),
			"",
			1,
			"L",
			false,
			0,
			"",
		)
		pdf.Ln(common.MmToPt(2))
		rowFonts := math.Floor(rowWidth / ContentSize)
		c := splitStringRune(v.Achievement, int(rowFonts))
		pdf.SetFont(RegularFont, "", ContentSize)
		pdf.MultiCell(
			rowWidth,
			ContentSize+2,
			c,
			"",
			"L",
			false,
		)
		pdf.Ln(common.MmToPt(5))
	}
	return pdf, c
}

func Gen(pdf *gofpdf.Fpdf, data *common.JsonSchema, fileFullPath string, lang string) {
	c := &nextPoint{X: 0, Y: 0}
	pdf, c = addBaseInfo(pdf, &data.CvBaseInfo, lang)
	pdf, c = addEduInfo(pdf, &data.CvEduInfo, c, lang)
	pdf, c = addSkills(pdf, &data.CvSkills, c, lang)
	pdf, _ = addResume(pdf, &data.CvResume, c, lang)

	err := pdf.OutputFileAndClose(fileFullPath)
	if err != nil {
		panic(err)
	}
}
