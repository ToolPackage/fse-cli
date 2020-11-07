package version

import (
	c "github.com/fatih/color"
	"strings"
)

const Logo = `███████╗███████╗███████╗
██╔════╝██╔════╝██╔════╝
█████╗  ███████╗█████╗  
██╔══╝  ╚════██║██╔══╝  
██║     ███████║███████╗
╚═╝     ╚══════╝╚══════╝`
const Name = "fse"
const Version = "0.1.0"
const Description = "The fse command line client"
const Website = "github.com/ToolPackage/fse-cli"

var (
	Build string
)

func ColorLogo() string {
	var buf strings.Builder
	buf.WriteRune('\n')
	buf.WriteRune('\n')
	buf.WriteString(c.RedString("███████╗███████╗███████╗\n"))
	buf.WriteString(c.RedString("██╔════╝██╔════╝██╔════╝\n"))
	buf.WriteString(c.RedString("█████╗  ███████╗█████╗  \n"))
	buf.WriteString(c.RedString("██╔══╝  ╚════██║██╔══╝  "))
	buf.WriteString(c.GreenString(" Version: %s\n", Version))
	buf.WriteString(c.RedString("██║     ███████║███████╗"))
	buf.WriteString(c.GreenString(" Description: %s\n", Description))
	buf.WriteString(c.RedString("╚═╝     ╚══════╝╚══════╝"))
	buf.WriteString(c.GreenString(" Source Code: %s\n", Website))
	return buf.String()
}
