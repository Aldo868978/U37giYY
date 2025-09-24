// 代码生成时间: 2025-09-24 09:51:22
package main

import (
    "archive/zip"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// ExcelGenerator represents the structure for generating excel files.
type ExcelGenerator struct {
    // Define necessary fields for the Excel generator
}

// NewExcelGenerator creates a new instance of ExcelGenerator.
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{}
}

// GenerateExcel creates an Excel file (.xlsx) with a given name and data.
func (e *ExcelGenerator) GenerateExcel(fileName string, data [][]string) error {
    // Define the output file path
    outputPath := fmt.Sprintf("%s.xlsx", fileName)
    
    // Create a new zip file for the Excel file
    f, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer f.Close()
    
    zipWriter := zip.NewWriter(f)
    defer zipWriter.Close()
    
    // Create the necessary files inside the zip archive
    files := []struct {
        path, content string
    }{
        {
            path:     " xl/_rels/workbook.xml.rels",
            content: xmlWorkbookRels,
        },
        {
            path:     " xl/workbook.xml",
            content: xmlWorkbook,
        },
        {
            path:     " xl/worksheets/sheet1.xml",
            content: xmlSheet,
        },
        {
            path:     " xl/sharedStrings.xml",
            content: sharedStringsXML(data),
        },
    }
    
    for _, file := range files {
        zipFile, err := zipWriter.Create(file.path)
        if err != nil {
            return err
        }
        _, err = zipFile.Write([]byte(file.content))
        if err != nil {
            return err
        }
    }
    
    // Create the .csv file for data
    csvFile, err := zipWriter.Create(" xl/worksheets/sheet1.csv")
    if err != nil {
        return err
    }
    writer := csv.NewWriter(csvFile)
    if err := writer.Write(data[0]); err != nil {
        return err
    }
    for _, record := range data[1:] {
        if err := writer.Write(record); err != nil {
            return err
        }
    }
    writer.Flush()
    if err := writer.Error(); err != nil {
        return err
    }
    
    return nil
}

// xmlWorkbookRels contains the XML content for the workbook.xml.rels file.
var xmlWorkbookRels = "<?xml version=\"\""1.0\"\"" encoding=\"\""UTF-8\"\"" standalone=\"\""yes\"\""?>\r\
<Relationships xmlns=\"\""http://schemas.openxmlformats.org/package/2006/relationships\"\"">\r\
    <Relationship Id=\"\""rId1\"\"" Type=\"\""http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument\"\"" Target=\"\""xl/workbook.xml\"\""/>\r\
</Relationships>"

// xmlWorkbook contains the XML content for the workbook.xml file.
var xmlWorkbook = "<?xml version=\"\""1.0\"\"" encoding=\"\""UTF-8\"\"" standalone=\"\""yes\"\""?>\r\
<workbook xmlns=\"\""http://schemas.openxmlformats.org/spreadsheetml/2006/main\"\"" xmlns:r=\"\""http://schemas.openxmlformats.org/package/2006/relationships\"\"">\r\
    <sheets>\r\
        <sheet name=\"\""Sheet1\"\"" sheetId=\"\""1\"\"" r:id=\"\""rId1\"\"" />\r\
    </sheets>\r\
</workbook>"

// xmlSheet contains the XML content for the sheet1.xml file.
var xmlSheet = "<?xml version=\"\""1.0\"\"" encoding=\"\""UTF-8\"\"" standalone=\"\""yes\"\""?>\r\
<worksheet xmlns=\"\""http://schemas.openxmlformats.org/spreadsheetml/2006/main\"\"">\r\
    <dimension ref=\"\""A1\"\"" />\r\
    <sheetViews>\r\
        <sheetView tabSelected=\"\""1\"\"" workbookViewId=\"\""0\"\"" />\r\
    </sheetViews>\r\
    <sheetFormatPr defaultRowHeight=\"\""15\"\"" />\r\
    <sheetData>\r\
        <!-- Data will be inserted here -->\r\
    </sheetData>\r\
</worksheet>"

// sharedStringsXML generates the XML content for the sharedStrings.xml file.
func sharedStringsXML(data [][]string) string {
    sb := strings.Builder{}
    sb.WriteString("<?xml version=\"\""1.0\"\"" encoding=\"\""UTF-8\"\"" standalone=\"\""yes\"\""?>\r\
<sst xmlns=\"\""http://schemas.openxmlformats.org/spreadsheetml/2006/main\"\"" count=\"\""" + fmt.Sprintf("%d", len(data)*len(data[0]) + 1) + "\"\"" uniqueCount=\"\""" + fmt.Sprintf("%d", len(data)*len(data[0]) + 1) + "\