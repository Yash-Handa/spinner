package spinner

import (
	"fmt"
	"time"
)

func spinnerTemplate(s *Spinner) {
	// c can hold both s.color and s.colorRBG
	var c interface {
		Sprint(...interface{}) string
	}
	if s.colorRGB != nil {
		c = s.colorRGB
	} else {
		c = s.color
	}

	// the ticker
	it := time.NewTicker(s.interval)

	i := 0
	l := len(s.symbols)

	// length of symbol
	symbolLen := len(s.symbols[0])

	// backlength is the total length of the previously printed line which will be erases
	backLen := -symbolLen

	for {
		select {
		case <-*(s.stopper):
			fmt.Printf("\r%*s\r%v", backLen+symbolLen, "", s.doneText)
			it.Stop()
			return
		case <-*(s.calibrator):
			it.Reset(s.interval)
			if s.colorRGB != nil {
				c = s.colorRGB
			} else {
				c = s.color
			}
			s.wg2.Done()
		case <-it.C:
			i = i % l
			fmt.Printf("\r%*s", backLen+symbolLen, "")
			backLen = len([]rune(s.preText)) + len([]rune(s.postText))
			fmt.Printf("\r%s%s%s", s.preText, c.Sprint(s.symbols[i]), s.postText)
			i++
		}
	}
}
