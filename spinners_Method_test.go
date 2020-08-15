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

func TestMethods(t *testing.T) {
	var id uint = 1
	sp, err := spinner.New(id, 100*time.Millisecond, spinner.Random16BitColor(), spinner.Normal)
	if err != nil {
		t.Fatal(color.Error.Sprint(err) + "\n")
	}

	sp.Start()
	defer sp.Stop()

	t.Run("it should set and get preText", func(st *testing.T) {
		want := "This is the pre-text"
		sp.SetPreText(want)
		got := sp.GetPreText()
		if got != want {
			st.Fatalf("PreText test: wanted %s but got %s", want, got)
		}
		sp.SetPreText("")
	})

	t.Run("it should set and get postText", func(st *testing.T) {
		want := "This is the post-text"
		sp.SetPostText(want)
		got := sp.GetPostText()
		if got != want {
			st.Fatalf("PostText test: wanted %s but got %s", want, got)
		}
		sp.SetPostText("")
	})

	t.Run("it should set and get doneText", func(st *testing.T) {
		want := "This is the done-text"
		sp.SetDoneText(want)
		got := sp.GetDoneText()
		if got != want {
			st.Fatalf("DoneText test: wanted %s but got %s", want, got)
		}
		sp.SetDoneText("")
	})

	t.Run("it should get ID", func(st *testing.T) {
		got := sp.GetID()
		if got != id {
			st.Fatalf("IDTest test: wanted %d but got %d", id, got)
		}
	})

	t.Run("it is spinning", func(st *testing.T) {
		if sp.IsSpinning() == false {
			st.Fatal("the spinner should be spinning but is not")
		}
	})

	t.Run("it should set and get interval", func(st *testing.T) {
		want := 300 * time.Millisecond
		sp.SetInterval(want)
		got := sp.GetInterval()
		if got != want {
			st.Fatalf("IntervalTest: wanted %v but got %v", want, got)
		}
	})
}

func TestOneSpinner(t *testing.T) {
	if testing.Short() == false {
		t.SkipNow()
	}
	var id uint = 1

	log.SetOutput(os.Stdout)
	log.Println(color.Info.Sprint("\nRunning this spinner if the tests are run with -short flag this spinner will still run"))

	sp, err := spinner.New(id, 100*time.Millisecond, spinner.Random16BitColor(), spinner.Normal)
	if err != nil {
		t.Fatal(color.Error.Sprint(err) + "\n")
	}
	sp.SetDoneText(fmt.Sprintf("Hurray spinner no. %d done", id))

	sp.Start()
	time.Sleep(3 * time.Second)
	sp.Stop()
	fmt.Println()
}
