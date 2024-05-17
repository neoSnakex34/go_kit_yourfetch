package main

import (
	"fmt"
	"log"

	"github.com/neoSnakex34/kit_fetch/utils"
)

const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"

	Rosewater = "\033[38;2;244;219;214m"
	Flamingo  = "\033[38;2;240;198;198m"
	Pink      = "\033[38;2;245;189;230m"
	Mauve     = "\033[38;2;198;160;246m"
	Red       = "\033[38;2;237;135;150m"
	Maroon    = "\033[38;2;238;153;160m"
	Peach     = "\033[38;2;245;169;127m"
	Yellow    = "\033[38;2;238;212;159m"
	Green     = "\033[38;2;166;218;149m"
	Teal      = "\033[38;2;139;213;202m"
	Sky       = "\033[38;2;145;215;227m"
	Sapphire  = "\033[38;2;125;196;228m"
	Blue      = "\033[38;2;138;173;244m"
	Lavender  = "\033[38;2;183;189;248m"
	Text      = "\033[38;2;202;211;245m"
	Subtext1  = "\033[38;2;184;192;224m"
	Subtext0  = "\033[38;2;165;173;203m"
	Overlay2  = "\033[38;2;147;154;183m"
	Overlay1  = "\033[38;2;128;135;162m"
	Overlay0  = "\033[38;2;110;115;141m"
	Surface2  = "\033[38;2;91;96;120m"
	Surface1  = "\033[38;2;73;77;100m"
	Surface0  = "\033[38;2;54;58;79m"
	Base      = "\033[38;2;36;39;58m"
	Mantle    = "\033[38;2;30;32;48m"
	Crust     = "\033[38;2;24;25;38m"
)

func main() {

	// FETCH

	// NOTE logo should be parsed in utils
	logo, err := utils.ReadFile("./cat2.txt")
	if err != nil {
		log.Fatal(err)
	}

	if len(logo) != 4 {
		log.Fatal("Logo must be 4 lines tall")
	}

	userAtHost, err := utils.GetUserAtHost()
	if err != nil {
		log.Fatal(err)
	}

	distroName, err := utils.GetDistroName()
	if err != nil {
		log.Fatal(err)
	}

	kernelVers, err := utils.GetKernelVers()
	if err != nil {
		log.Fatal(err)
	}

	shellName, err := utils.GetShell()
  if err != nil {
    log.Fatal(err)
  }

	uptimeValue, err := utils.GetUpTime()
	if err != nil {
		log.Fatal(err)
	}

	// ORGANIZE AND STYLE
	home := "" + Bold + Yellow + userAtHost[0] + Reset + Bold + "@" + Green + userAtHost[1] + Reset
	distro := "on: " + Bold + Teal + distroName + Reset
	kernel := "kl: " + Bold + Sky + kernelVers + Reset
	shell := "sh: " + Bold + Sapphire + shellName + Reset
	uptime := "up: " + Bold + Blue + uptimeValue + Reset

	info := []string{distro, kernel, shell, uptime}

	// TODO prettify logo

	maxLenLogoLine := utils.GetMaxLogoLineLenght(logo)
	max_info := len(info)

	// PRINT
	utils.ClearTerm()
	fmt.Println()
	fmt.Printf("  %*s    %s\n\n", maxLenLogoLine, "", home)

	// note that actual design is tied to 4 lines tall logo and for lines info, but it is kinda scalable
	// so i keep the checks for logo and info length
	for i, line := range logo {

		if i < max_info {
			padding := maxLenLogoLine - utils.LenLogoLine(line)

			fmt.Printf("  %s%*s%s\n", line, padding, "", info[i])
		} // else this would be a mess, need to add some controls
		// and modularize in order to add support for more lines
		// in logo and info

	}
	fmt.Println()

}
