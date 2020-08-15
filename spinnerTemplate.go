package spinner

import (
	"fmt"
	"strings"
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
		case <-it.C:
			i = i % l
			fmt.Printf("\r%*s", backLen+symbolLen, "")
			pre, post, temp := prePostSanitizer(s.preTextGen, s.postTextGen)
			backLen = temp
			fmt.Printf("\r%s%s%s", pre, c.Sprint(s.symbols[i]), post)
			// fmt.Println(post)
			i++
		}
	}
}

// A helper function to sanitize pre and post strings
func prePostSanitizer(pre, post func() string) (preStr string, postStr string, l int) {
	preStr = strings.Trim(pre(), "\n\r")
	postStr = strings.Trim(post(), "\n\r")
	l = len([]rune(preStr)) + len([]rune(postStr))
	return
}
