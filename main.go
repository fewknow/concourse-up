package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/EngineerBetter/concourse-up/bosh"
	"github.com/EngineerBetter/concourse-up/director"
	"github.com/EngineerBetter/concourse-up/fly"
	"github.com/EngineerBetter/concourse-up/terraform"
	"github.com/fatih/color"
)

import "gopkg.in/urfave/cli.v1"

// ConcourseUpVersion is a compile-time variable set with -ldflags
var ConcourseUpVersion = "COMPILE_TIME_VARIABLE_main_concourseUpVersion"
var nonInteractive bool
var yellow = color.New(color.FgYellow).SprintFunc()
var blue = color.New(color.FgCyan, color.Bold).SprintfFunc()

func NonInteractiveModeEnabled() bool {
	return nonInteractive
}

func main() {
	app := cli.NewApp()
	app.Name = "Concourse-Up"
	app.Usage = "A CLI tool to deploy Concourse CI"
	app.Version = ConcourseUpVersion
	app.Commands = []cli.Command{
		{
			Name: "deploy",
		},
		{
			Name: "destroy",
		},
		{
			Name: "info",
		},
	}
	app.Flags = []cli.Flag{
	cli.BoolFlag{
		Name:        "non-interactive, n",
		EnvVar:      "NON_INTERACTIVE",
		Usage:       "Non interactive",
		Destination: &nonInteractive,
		},
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%s

See 'concourse-up help <command>' to read about a specific command.

Built by %s %s

`, cli.AppHelpTemplate, blue("EngineerBetter"), blue("http://engineerbetter.com"))

	if err := checkCompileTimeArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func checkCompileTimeArgs() error {
	boshCompileTimeArgs := map[string]string{
		"bosh.ConcourseReleaseSHA1":      bosh.ConcourseReleaseSHA1,
		"bosh.ConcourseReleaseURL":       bosh.ConcourseReleaseURL,
		"bosh.ConcourseReleaseVersion":   bosh.ConcourseReleaseVersion,
		"bosh.ConcourseStemcellSHA1":     bosh.ConcourseStemcellSHA1,
		"bosh.ConcourseStemcellURL":      bosh.ConcourseStemcellURL,
		"bosh.ConcourseStemcellVersion":  bosh.ConcourseStemcellVersion,
		"bosh.DirectorCPIReleaseSHA1":    bosh.DirectorCPIReleaseSHA1,
		"bosh.DirectorCPIReleaseURL":     bosh.DirectorCPIReleaseURL,
		"bosh.DirectorCPIReleaseVersion": bosh.DirectorCPIReleaseVersion,
		"bosh.DirectorReleaseSHA1":       bosh.DirectorReleaseSHA1,
		"bosh.DirectorReleaseURL":        bosh.DirectorReleaseURL,
		"bosh.DirectorReleaseVersion":    bosh.DirectorReleaseVersion,
		"bosh.DirectorStemcellSHA1":      bosh.DirectorStemcellSHA1,
		"bosh.DirectorStemcellURL":       bosh.DirectorStemcellURL,
		"bosh.DirectorStemcellVersion":   bosh.DirectorStemcellVersion,
		"bosh.GardenReleaseSHA1":         bosh.GardenReleaseSHA1,
		"bosh.GardenReleaseURL":          bosh.GardenReleaseURL,
		"bosh.GardenReleaseVersion":      bosh.GardenReleaseVersion,
		"bosh.GrafanaReleaseSHA1":        bosh.GrafanaReleaseSHA1,
		"bosh.GrafanaReleaseURL":         bosh.GrafanaReleaseURL,
		"bosh.GrafanaReleaseVersion":     bosh.GrafanaReleaseVersion,
		"bosh.InfluxDBReleaseSHA1":       bosh.InfluxDBReleaseSHA1,
		"bosh.InfluxDBReleaseURL":        bosh.InfluxDBReleaseURL,
		"bosh.InfluxDBReleaseVersion":    bosh.InfluxDBReleaseVersion,
		"bosh.RiemannReleaseSHA1":        bosh.RiemannReleaseSHA1,
		"bosh.RiemannReleaseURL":         bosh.RiemannReleaseURL,
		"bosh.RiemannReleaseVersion":     bosh.RiemannReleaseVersion,
		"director.DarwinBinaryURL":       director.DarwinBinaryURL,
		"director.LinuxBinaryURL":        director.LinuxBinaryURL,
		"director.WindowsBinaryURL":      director.WindowsBinaryURL,
		"fly.DarwinBinaryURL":            fly.DarwinBinaryURL,
		"fly.LinuxBinaryURL":             fly.LinuxBinaryURL,
		"fly.WindowsBinaryURL":           fly.WindowsBinaryURL,
		"terraform.DarwinBinaryURL":      terraform.DarwinBinaryURL,
		"terraform.LinuxBinaryURL":       terraform.LinuxBinaryURL,
		"terraform.WindowsBinaryURL":     terraform.WindowsBinaryURL,
	}

	if ConcourseUpVersion == "" || strings.HasPrefix(ConcourseUpVersion, "COMPILE_TIME_VARIABLE") {
		return errors.New("Compile-time variable main.ConcourseUpVersion not set, please build with: `go build -ldflags \"-X main.ConcourseUpVersion=0.0.0\"`")
	}

	for key, value := range boshCompileTimeArgs {
		if value == "" || strings.HasPrefix(value, "COMPILE_TIME_VARIABLE") {
			return fmt.Errorf("Compile-time variable %s not set, please build with: `go build -ldflags \"-X github.com/EngineerBetter/concourse-up/%s=SOME_VALUE\"`", key, key)
		}
	}

	return nil
}
