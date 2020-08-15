package spinner_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Yash-Handa/spinner"
	"github.com/gookit/color"
)

func TestUnicode_16BitSpinners(t *testing.T) {
	var timmer time.Duration = 5
	if testing.Verbose() {
		t.Log(color.Info.Sprint("Testing Unicode spinners with 16 Bit Colors (will work with utf-8 enabled terminals)"))
	}

	if testing.Short() {
		timmer = 0
	}

	for k := range spinner.SpinnerMap {
		if k < 1000 {
			continue
		}

		t.Run(fmt.Sprintf("Spinner ID=%d", k), func(st *testing.T) {
			sp, err := spinner.New(k, 100*time.Millisecond, func() string { return "The starting text  " }, func() string { return fmt.Sprintf("  Spinner ID = %d", k) }, fmt.Sprintf("Hurray spinner no. %d done", k), spinner.Random16BitColor(), spinner.Normal)
			if err != nil {
				st.Fatal(color.Error.Sprint(err) + "\n")
			}

			sp.Start()
			if timmer != 0 {
				time.Sleep(timmer * time.Second)
			}
			sp.Stop()
			fmt.Println()
		})
	}
}
