package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/paldraken/sqldebugwatch/internal/types"
)

func Lines(line chan string, debug chan *types.SqlDebugRow) {

	re, err := regexp.Compile(`^TIME: \d+\.\d`)
	if err != nil {
		fmt.Printf("Error compiling regex pattern: %v", err)
		return
	}

	insideBlock := false
	block := []string{}
	for {
		line := <-line

		if re.MatchString(line) {
			insideBlock = true
		}

		if checkSeparator(line) {
			insideBlock = false
			debugRow := blockToSqlDebugRow(block)

			debug <- debugRow
			block = []string{}
		}

		if insideBlock {
			block = append(block, line)
		}
	}
}

func blockToSqlDebugRow(block []string) *types.SqlDebugRow {
	debugRow := &types.SqlDebugRow{}

	traseRe, _ := regexp.Compile(`^ from\s.*\/modules\/.*$`)

	insideSql := false
	insideTrase := false
	for n, line := range block {

		if n == 0 {
			extractTiemAndSession(debugRow, line)
			continue
		}

		if n == 2 {
			insideSql = true
		}

		if traseRe.MatchString(line) {
			insideTrase = true
			insideSql = false
		}

		if insideSql {
			debugRow.Sql += " " + strings.TrimSpace(line)
		}

		if insideTrase {
			debugRow.Trace = append(debugRow.Trace, line)
		}

	}

	debugRow.Sql = strings.TrimSpace(debugRow.Sql)
	debugRow.Table = extractTableName(debugRow.Sql)
	debugRow.Modules = extractModules(debugRow.Trace)

	return debugRow
}

func extractTableName(sql string) string {
	re := regexp.MustCompile(`(?i)(?:insert into|update|delete from|from)\s` + "`" + `?([a-z0-9_]+)` + "`" + `?\b`)
	matches := re.FindStringSubmatch(sql)
	if len(matches) != 2 {
		return ""
	}
	tableName := matches[1]
	return tableName
}

func extractModules(trace []string) []string {
	re := regexp.MustCompile(`/modules/([^/]+)/?`)
	result := []string{}
	for _, s := range trace {
		match := re.FindStringSubmatch(s)
		if len(match) > 1 {
			result = append(result, match[1])
		}
	}
	return uniqueStrings(result)
}

func extractTiemAndSession(sqlDebugRow *types.SqlDebugRow, line string) {
	regexTime := regexp.MustCompile(`TIME:\s+(\d+.\d*.)SES`)
	regexSession := regexp.MustCompile(`SESSION: ([\w\d]+)`)

	var timeStr string
	timeSubMatch := regexTime.FindStringSubmatch(line)

	if len(timeSubMatch) > 1 {
		timeStr = timeSubMatch[1]
		timeStr = strings.Trim(timeStr, " ")
	} else {
		timeStr = "0"
	}

	time, _ := strconv.ParseFloat(timeStr, 32)
	sqlDebugRow.Time = float32(time)

	if regexSession.MatchString(line) {
		sqlDebugRow.Session = regexSession.FindStringSubmatch(line)[1]
	}
}

func checkSeparator(line string) bool {
	blockSeparator := "-----------------------"
	return strings.HasPrefix(line, blockSeparator)
}

func uniqueStrings(strs []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, str := range strs {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}
