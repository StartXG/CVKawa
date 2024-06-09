package commands

import (
	"JsonToPdf/JTPAction"
	"JsonToPdf/common"
	"github.com/spf13/cobra"
)

func NewJsonToPdfCmd() *cobra.Command {
	var jsonFilePath string
	var lang string
	var outPut string
	var rootCmd = &cobra.Command{
		Use: "jtp",
		Run: func(cmd *cobra.Command, args []string) {
			JTPAction.LoadData(&jsonFilePath)
			obj := JTPAction.PdfObj{
				OrientationStr: "P",
				UnitStr:        "pt",
				SizeStr:        "A4",
				FontDirStr:     "",
			}
			pdf := obj.New()
			JTPAction.Gen(pdf, common.JsonData, outPut, lang)
		},
		Version: "alpha-0.1",
	}
	rootCmd.Flags().StringVarP(&jsonFilePath, "file", "f", "", "The path of the json file")
	rootCmd.Flags().StringVarP(&lang, "lang", "l", "cn", "The language of the pdf[default: en]")
	rootCmd.Flags().StringVarP(&outPut, "out", "o", "cv.pdf", "The output path of the pdf[default: cv.pdf]")
	if rootCmd.MarkFlagRequired("file") != nil {
		panic("Failed to mark flag required")
	}
	return rootCmd
}
