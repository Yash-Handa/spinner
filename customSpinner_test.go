package spinner_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Yash-Handa/spinner"
	"github.com/gookit/color"
)

func TestCustomSpinner(t *testing.T) {
	var timmer time.Duration = 3

	log.SetOutput(os.Stdout)
	log.Println(color.Info.Sprint("Testing A custom spinner"))

	if testing.Short() {
		timmer = 0
	}

	symbols := []string{"N   ", "IN  ", "PIN ", "SPIN", " SPI", "  SP", "   S", "    "}
	sp, err := spinner.Custom(symbols, 100*time.Millisecond, spinner.Random16BitColor(), spinner.Normal)
	if err != nil {
		t.Fatal(color.Error.Sprint(err) + "\n")
	}

	sp.SetPreText("The starting text  ")
	sp.SetPostText(fmt.Sprintf("  the ending text"))
	sp.SetDoneText(fmt.Sprintf("Hurray spinner with symbols %#v done", symbols))

	sp.Start()
	if timmer != 0 {
		time.Sleep(timmer * time.Second)
		// sp.SetInterval(300 * time.Millisecond)
		sp.SetColor(spinner.Random16BitColor(), "")
		time.Sleep(timmer * time.Second)
	} else {
		// sp.SetInterval(300 * time.Millisecond)
		sp.SetColor(spinner.Random16BitColor(), "")
	}
	sp.Stop()

	fmt.Println()
}
