package spinner

import (
	"fmt"
	"sync"
	"time"

	"github.com/gookit/color"
)

// Supported colors
const (
	Red         = "Red"
	Black       = "Black"
	Green       = "Green"
	Yellow      = "Yellow"
	Blue        = "Blue"
	Magenta     = "Magenta"
	Cyan        = "Cyan"
	White       = "White"
	Normal      = "Normal"
	HexBgNormal = "HexBgNormal"
)

// Spinner yo
type Spinner struct {
	id          uint
	interval    time.Duration
	preTextGen  func() string
	postTextGen func() string
	doneText    string

	stopper  *chan bool
	spinning bool
	wg       sync.WaitGroup
	color    color.Style
	colorRGB *color.RGBStyle
}

// GetID yo
func (s *Spinner) GetID() uint {
	return s.id
}

// GetInterval yo
func (s *Spinner) GetInterval() time.Duration {
	return s.interval
}

// IsSpinning yo
func (s *Spinner) IsSpinning() bool {
	return s.spinning
}

// Start yo
func (s *Spinner) Start() {
	if s.spinning {
		return
	}
	f := spinnerMap[s.id]

	s.wg.Add(1)
	s.spinning = true
	go func() {
		defer s.wg.Done()
		f(s)
	}()
	return
}

// Stop yo
func (s *Spinner) Stop() {
	if s.spinning == false {
		return
	}
	*s.stopper <- true
	s.wg.Wait()
	s.spinning = false
}

// New yo
func New(id uint, interval time.Duration, preTextGen, postTextGen func() string, doneText string, forground, background string) (*Spinner, error) {
	temp := new(Spinner)

	_, ok := spinnerMap[id]
	if !ok {
		return nil, fmt.Errorf("%d is not a valid id", id)
	}
	temp.id = id
	temp.interval = interval

	if preTextGen == nil {
		temp.preTextGen = func() string { return "" }
	} else {
		temp.preTextGen = preTextGen
	}

	if postTextGen == nil {
		temp.postTextGen = func() string { return "" }
	} else {
		temp.postTextGen = postTextGen
	}

	temp.doneText = doneText

	fg, bg := color.FgDefault, color.BgDefault
	// setting empty fgRBG and bgRBG
	fgRBG, bgRBG := color.HEX("ggg"), color.HEX("ggg")

	switch forground {
	case "Red":
		fg = color.FgRed
	case "Black":
		fg = color.FgBlack
	case "Green":
		fg = color.FgGreen
	case "Blue":
		fg = color.FgBlue
	case "Magenta":
		fg = color.FgMagenta
	case "Yellow":
		fg = color.FgYellow
	case "Cyan":
		fg = color.FgCyan
	case "White":
		fg = color.BgWhite
	case "Normal", "":
	default:
		// check if it is a hex code...
		fgRBG = color.HEX(forground, false)
		if fgRBG.IsEmpty() {
			return nil, fmt.Errorf("%s is not a valid forground color", forground)
		}
	}

	switch background {
	case "Red":
		bg = color.BgRed
	case "Black":
		bg = color.BgBlack
	case "Green":
		bg = color.BgGreen
	case "Blue":
		bg = color.BgBlue
	case "Magenta":
		bg = color.BgMagenta
	case "Yellow":
		bg = color.BgYellow
	case "Cyan":
		bg = color.BgCyan
	case "White":
		bg = color.BgWhite
	case "HexBgNormal":
		bgRBG = color.HEX("000000")
	case "Normal", "":
	default:
		// check if it is a hex code...
		bgRBG = color.HEX(background, true)
		if bgRBG.IsEmpty() {
			return nil, fmt.Errorf("%s is not a valid background color", background)
		}
	}

	if bgRBG.IsEmpty() != fgRBG.IsEmpty() {
		return nil, fmt.Errorf("both forground and background should be ether Hex or Constant colors not a combination")
	}

	if bgRBG.IsEmpty() == false && color.IsSupport256Color() {
		if background == "HexBgNormal" {
			temp.colorRGB = color.NewRGBStyle(fgRBG)
		} else {
			temp.colorRGB = color.NewRGBStyle(fgRBG, bgRBG)
		}
	} else {
		temp.color = color.New(fg, bg)
		temp.colorRGB = nil
	}

	ch := make(chan bool)
	temp.stopper = &ch
	temp.spinning = false
	return temp, nil
}
