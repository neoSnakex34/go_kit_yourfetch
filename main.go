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

	// Colori principali del tema Catppuccin Macchiato
	Text     = "\033[38;2;205;214;244m"
	Subtext0 = "\033[38;2;180;190;254m"
	Subtext1 = "\033[38;2;166;173;206m"
	Overlay0 = "\033[38;2;147;153;178m"
	Overlay1 = "\033[38;2;132;138;163m"
	Overlay2 = "\033[38;2;124;130;154m"
	Surface0 = "\033[38;2;89;95;123m"
	Surface1 = "\033[38;2;79;85;113m"
	Surface2 = "\033[38;2;70;76;104m"
	Base     = "\033[38;2;59;66;82m"
	Mantle   = "\033[38;2;48;54;68m"
	Crust    = "\033[38;2;37;43;56m"

	// Accenti
	Blue   = "\033[38;2;103;141;217m"
	Green  = "\033[38;2;166;218;149m"
	Pink   = "\033[38;2;243;139;168m"
	Red    = "\033[38;2;237;135;150m"
	Yellow = "\033[38;2;229;200;144m"
)

func main() {

	logo, err := utils.ReadFile("cat2.txt")
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

	shell := utils.GetShell()

	uptime, err := utils.GetUpTime()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{distroName, kernelVers, shell, uptime}

	maxLogoLineLength := utils.GetMaxLogoLineLenght(logo)
	max_info := len(info)

	fmt.Printf("%*s%s\n", maxLogoLineLength, "", bold+cattppuccin_purple+userAtHost[0]+reset+"@"+bold+cattppuccin_purple+userAtHost[1]+reset)
	// note that actual design is tied to 4 lines tall logo and for lines info, but it is kinda scalable
	// so i keep the checks for logo and info length
	for i, line := range logo {

		// print logo line
		// fmt.Print(line)

		// if info print info
		if i < max_info {
			padding := maxLogoLineLength - utils.LenLogoLine(line)

			fmt.Printf("%s%*s%s\n", line, padding, "", info[i])
		}

	}

}
