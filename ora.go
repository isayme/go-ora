package ora

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Ora struct {
	spinner *spinner.Spinner
}

func New() *Ora {
	spinner := spinner.New(
		spinner.CharSets[14],
		time.Millisecond*100,
		spinner.WithColor("cyan"),
	)
	return &Ora{
		spinner: spinner,
	}
}

func (o *Ora) Start() {
	o.spinner.Start()
}

func (o *Ora) Stop() {
	o.spinner.Stop()
}

/**
 * set color of spinner
 */
func (o *Ora) Color(colors ...string) {
	o.spinner.Color(colors...)
}

/**
 * set text display after spinner, alias of @Suffix
 */
func (o *Ora) Text(text string) {
	o.Suffix(text)
}

/**
 * set text display before spinner
 */
func (o *Ora) Prefix(prefix string) {
	if prefix == "" {
		o.spinner.Prefix = ""
		return
	}
	o.spinner.Prefix = fmt.Sprintf("%s ", prefix)
}

/**
 * set text display after spinner,
 */
func (o *Ora) Suffix(suffix string) {
	if suffix == "" {
		o.spinner.Suffix = ""
		return
	}
	o.spinner.Suffix = fmt.Sprintf(" %s", suffix)
}

func (o *Ora) stopAndPersist(symbol, text string) {
	o.spinner.Start()
	o.spinner.FinalMSG = fmt.Sprintf("%s %s\n", symbol, text)
	o.spinner.Stop()
}

/**
 * stop spinner, set spinner to green "✔", display text after spinner
 */
func (o *Ora) Succeed(text string) {
	o.stopAndPersist(color.GreenString("✔"), text)
}

/**
 * stop spinner, set spinner to blue "ℹ", display text after spinner
 */
func (o *Ora) Info(text string) {
	o.stopAndPersist(color.BlueString("ℹ"), text)
}

/**
 * stop spinner, set spinner to yellow "⚠", display text after spinner
 */
func (o *Ora) Warn(text string) {
	o.stopAndPersist(color.YellowString("⚠"), text)
}

/**
 * stop spinner, set spinner to red "✖", display text after spinner
 */
func (o *Ora) Fail(text string) {
	o.stopAndPersist(color.RedString("✖"), text)
}
