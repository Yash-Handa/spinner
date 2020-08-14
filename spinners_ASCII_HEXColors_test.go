package spinner

import (
	"fmt"
	"testing"
	"time"

	"github.com/gookit/color"
)

func TestASCII_HEXSpinners(t *testing.T) {
	var timmer time.Duration = 5
	if testing.Verbose() {
		t.Log(color.Info.Sprint("Testing ASCII spinners with HEX Colors (will work with terminals supporting Hex/ 256 colors or colors will have no effect)"))
	}

	if testing.Short() {
		timmer = 0
	}

	for k := range spinnerMap {
		if k > 99 {
			continue
		}

		t.Run(fmt.Sprintf("Spinner ID=%d", k), func(st *testing.T) {
			sp, err := New(k, 100*time.Millisecond, func() string { return "The starting text  " }, func() string { return fmt.Sprintf("  Spinner ID = %d", k) }, fmt.Sprintf("Hurray spinner no. %d done", k), "#F89F3D", HexBgNormal)
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
