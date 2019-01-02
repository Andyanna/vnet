package log

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type PatternLogFormatter struct {
	Pattern string
}

const (
	pattern = "%{level} %{time} %{file}[%{line}] %{func}: %{message}\n"
)

func PatternLogFormatterFactory(params ...string) *PatternLogFormatter {
	var instance *PatternLogFormatter
	if params == nil {
		instance = &PatternLogFormatter{
			Pattern: pattern,
		}
	} else {
		instance = &PatternLogFormatter{
			Pattern: params[0],
		}
	}
	return instance
}

func (this *PatternLogFormatter) Format(message string, level string, params ...interface{}) string {

	file, funcName, line := GetRuntimeInfoShortFormat(5)

	replacer := strings.NewReplacer(
		LEVEL,
		level,
		TIME,
		time.Now().Format("2006-01-02 15:04:05"),
		FILE,
		file,
		FUNC,
		funcName,
		LINENO,
		strconv.Itoa(line),
		MESSAGE,
		message,
	)

	result := replacer.Replace(this.Pattern)
	return fmt.Sprintf(result, params...)
}
