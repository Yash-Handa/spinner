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

func TestASCII_HEXSpinners(t *testing.T) {
	var timmer time.Duration = 3

	log.SetOutput(os.Stdout)
	log.Println(color.Info.Sprint("Testing ASCII spinners with HEX Colors (will work with terminals supporting Hex/ 256 colors or colors will have no effect)"))

	if testing.Short() {
		timmer = 0
	}

	for k := range spinner.SpinnerMap {
		if k > 999 {
			continue
		}

		t.Run(fmt.Sprintf("Spinner ID=%d", k), func(st *testing.T) {
			sp, err := spinner.New(k, 100*time.Millisecond, spinner.RandomHexColor(), spinner.HexBgNormal)
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
				sp.SetColor(spinner.RandomHexColor(), "")
				time.Sleep(timmer * time.Second)
			} else {
				sp.SetInterval(300 * time.Millisecond)
				sp.SetColor(spinner.RandomHexColor(), "")
			}
			sp.Stop()

			if timmer != 0 {
				fmt.Println()
			}
		})
	}

	if timmer == 0 {
		fmt.Println()
	}
}
