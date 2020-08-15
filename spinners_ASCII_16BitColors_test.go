package spinner

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gookit/color"
)

func TestASCII_16BitSpinners(t *testing.T) {
	var timmer time.Duration = 5
	if testing.Verbose() {
		t.Log(color.Info.Sprint("Testing ASCII spinners with 16 Bit Colors (Will work with almost every terminal out there)"))
	}

	if testing.Short() {
		timmer = 0
	}

	for k := range SpinnerMap {
		if k > 99 {
			continue
		}

		t.Run(fmt.Sprintf("Spinner ID=%d", k), func(st *testing.T) {
			sp, err := New(k, 100*time.Millisecond, func() string { return "The starting text  " }, func() string { return fmt.Sprintf("  Spinner ID = %d", k) }, fmt.Sprintf("Hurray spinner no. %d done", k), random16BitCode(), Normal)
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

func random16BitCode() string {
	c := [...]string{Red, Black, Green, Yellow, Blue, Magenta, Cyan, White,
		Normal}
	rand.Seed(time.Now().UnixNano())
	return c[rand.Intn(9)]
}
