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

func TestASCII_16BitSpinners(t *testing.T) {
	var timmer time.Duration = 3

	log.SetOutput(os.Stdout)
	log.Println(color.Info.Sprint("Testing ASCII spinners with 16 Bit Colors (Will work with almost every terminal out there)"))

	if testing.Short() {
		timmer = 0
	}

	for k := range spinner.SpinnerMap {
		if k > 999 {
			continue
		}

		t.Run(fmt.Sprintf("Spinner ID=%d", k), func(st *testing.T) {
			sp, err := spinner.New(k, 100*time.Millisecond, spinner.Random16BitColor(), spinner.Normal)
			if err != nil {
				st.Fatal(color.Error.Sprint(err) + "\n")
			}

			sp.SetPreText("The starting text  ")
			sp.SetPostText(fmt.Sprintf("  Spinner ID = %d", k))
			sp.SetDoneText(fmt.Sprintf("Hurray spinner no. %d done", k))

			sp.Start()
			if timmer != 0 {
				time.Sleep(timmer * time.Second)
				sp.SetInterval(300 * time.Millisecond)
				sp.SetColor(spinner.Random16BitColor(), "")
				time.Sleep(timmer * time.Second)
			} else {
				sp.SetInterval(300 * time.Millisecond)
				sp.SetColor(spinner.Random16BitColor(), "")
			}
			sp.Stop()
			fmt.Println()
		})
	}
}
