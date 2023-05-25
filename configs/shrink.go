package configs

import (
	"math/big"

	"github.com/dustin/go-humanize"
	"github.com/spf13/viper"
)

var ShrinkLog *big.Int

func init() {
	initViper()
	shrinkLog := viper.GetString("shrink_log")

	if shrinkLog != "" {
		ShrinkLog = prepareShrinkSize(shrinkLog)
	}
}

func prepareShrinkSize(size string) *big.Int {
	s, err := humanize.ParseBigBytes(size)
	if err != nil {
		panic(err)
	}

	return s
}
