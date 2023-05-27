package configs

import (
	"math/big"

	"github.com/dustin/go-humanize"
)

var TruncateLog *big.Int

func init() {
	initConfig()
	truncateLog := k.String("truncate_log")

	if truncateLog != "" {
		TruncateLog = prepareSize(truncateLog)
	}
}

func prepareSize(size string) *big.Int {
	s, err := humanize.ParseBigBytes(size)
	if err != nil {
		panic(err)
	}

	return s
}
