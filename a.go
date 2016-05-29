package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	destDirFmt string
	doRun      bool
)

const (
	FMT_YEAR  = "year"
	FMT_MONTH = "month"
	FMT_DAY   = "day"
)

func main() {
	flag.Parse()

	var (
		err     error
		baseDir string
		base    *os.File
	)

	baseDir = flag.Arg(0)
	if baseDir == "" {
		baseDir, err = os.Getwd()
		if err != nil {
			abort(err)
		}
	}

	base, err = os.Open(baseDir)
	if err != nil {
		abort(err)
	}
	files, err := base.Readdir(-1)
	if err != nil {
		abort(err)
	}
	for _, info := range files {
		// Skip dir.
		if info.IsDir() {
			continue
		}

		m := NewMove(
			filepath.Join(baseDir, info.Name()),
			destPathf(destDirFmt, info),
		)

		fmt.Println(m.Describe(doRun))
		if doRun {
			if err := m.Run(); err != nil {
				abort(err)
			}
		}
	}
}

func init() {
	flag.StringVar(&destDirFmt, "fmt", "year-month", "dest dir name format")
	flag.BoolVar(&doRun, "x", false, "run?")
}

func abort(err error) {
	fmt.Print(err)
	os.Exit(1)
}

func destPathf(pathFmt string, info os.FileInfo) string {
	modTime := info.ModTime()
	year := fmt.Sprintf("%d", modTime.Year())
	month := fmt.Sprintf("%d", modTime.Month())
	day := fmt.Sprintf("%d", modTime.Day())
	pathFmt = strings.Replace(pathFmt, FMT_YEAR, year, -1)
	pathFmt = strings.Replace(pathFmt, FMT_MONTH, month, -1)
	pathFmt = strings.Replace(pathFmt, FMT_DAY, day, -1)
	return pathFmt
}

type Move interface {
	Describe(doRun bool) string
	Run() error
}

type move struct {
	srcPath string
	destDir string
}

func NewMove(srcPath, destDir string) *move {
	return &move{srcPath, destDir}
}

func (m move) Describe(doRun bool) string {
	verb := "Will move"
	if doRun {
		verb = "Moving"
	}
	return fmt.Sprintf("%s `%s` to `%s/`", verb, m.srcName(), m.destDir)
}

func (m move) Run() error {
	destPath, err := m.destPath()
	if err != nil {
		return err
	}
	return os.Rename(m.srcPath, destPath)
}

func (m move) srcName() string {
	return filepath.Base(m.srcPath)
}

func (m move) destPath() (string, error) {
	return filepath.Abs(filepath.Join(m.destDir, m.srcName()))
}
