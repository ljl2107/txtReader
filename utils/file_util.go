package utils

import (
	"bufio"
	"io"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func ReadTxtByPage(filePath string, page int, perPage int) (string, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var allLines []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// 将读取到的每行内容从GBK编码转换为UTF-8编码
		utf8Line, _, _ := transform.String(simplifiedchinese.GBK.NewDecoder(), string(line))
		allLines = append(allLines, utf8Line)
	}

	totalLines := len(allLines)
	totalPages := (totalLines + perPage - 1) / perPage
	startIndex := (page - 1) * perPage
	endIndex := page * perPage
	if endIndex > totalLines {
		endIndex = totalLines
	}

	content := ""
	for i := startIndex; i < endIndex; i++ {
		content += allLines[i] + "\n"
	}

	return content, totalPages, nil
}
