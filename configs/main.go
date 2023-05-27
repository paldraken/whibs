package configs

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	flag "github.com/spf13/pflag"
)

var once sync.Once

var k = koanf.New(".")

func initConfig() {
	once.Do(func() {

		f := flag.NewFlagSet("config", flag.ContinueOnError)
		f.Usage = func() {
			fmt.Println(f.FlagUsages())
			os.Exit(0)
		}
		// Path to one or more config files to load into koanf along with some config params.
		f.StringSlice("conf", []string{"whibs.yaml"}, "path to one or more .yaml config files")

		f.StringP("path", "p", "", "Путь к логу sql (mysql_debug.sql)")
		f.StringP("truncate_log", "t", "", "Обрезать лог айл при достижении размера. (2MB, 1GB)")

		f.BoolP("server.enable", "w", false, "enable ws server")
		f.IntP("server.port", "P", 8080, "server port")

		f.StringP("filter.module", "M", "", "Фильтровать запросы в консоле по имени модуля")
		f.StringP("filter.table", "T", "", "Фильтровать запросы в консоле по таблице")
		f.StringP("filter.sql", "S", "", "Фильтровать запросы в консоле по вхожнию в sql запрос")
		f.StringP("filter.trace", "D", "", "Фильтровать запросы в консоле по вхождению в stack trace")
		f.BoolP("filter.show_trace", "d", false, "показывать trace в консоле")
		f.BoolP("filter.pause", "x", false, "Не выводить в консоль")

		f.Parse(os.Args[1:])

		// Load the config files provided in the commandline.
		cFiles, _ := f.GetStringSlice("conf")
		for _, c := range cFiles {
			if err := k.Load(file.Provider(c), yaml.Parser()); err != nil {
				log.Fatalf("error loading file: %v", err)
			}
		}

		// "time" and "type" may have been loaded from the config file, but
		// they can still be overridden with the values from the command line.
		// The bundled posflag.Provider takes a flagset from the spf13/pflag lib.
		// Passing the Koanf instance to posflag helps it deal with default command
		// line flag values that are not present in conf maps from previously loaded
		// providers.
		if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
	})
}
