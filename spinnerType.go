package spinner

import (
	"fmt"
	"strings"
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
	defInerval  = 100 * time.Millisecond
)

// Spinner yo
type Spinner struct {
	id       uint
	interval time.Duration
	preText  string
	postText string
	doneText string
	symbols  []string
	bg       string

	stopper    *chan bool
	calibrator *chan bool
	spinning   bool
	wg         sync.WaitGroup
	wg2        sync.WaitGroup
	color      color.Style
	colorRGB   *color.RGBStyle

	// if different go routines update the value use mux
	prePostMux  sync.Mutex
	spinningMux sync.Mutex
}

// GetID yo
func (s *Spinner) GetID() uint {
	return s.id
}

// GetInterval yo
func (s *Spinner) GetInterval() time.Duration {
	return s.interval
}

// GetPreText yo
func (s *Spinner) GetPreText() string {
	return s.preText
}

// GetPostText yo
func (s *Spinner) GetPostText() string {
	return s.postText
}

// GetDoneText yo
func (s *Spinner) GetDoneText() string {
	return s.doneText
}

// IsSpinning yo
func (s *Spinner) IsSpinning() bool {
	s.spinningMux.Lock()
	defer s.spinningMux.Unlock()
	return s.spinning
}

// SetPreText yo
func (s *Spinner) SetPreText(pre string) {
	pre = strings.Trim(pre, "\n\r")
	s.prePostMux.Lock()
	s.preText = pre
	s.prePostMux.Unlock()
}

// SetPostText yo
func (s *Spinner) SetPostText(post string) {
	post = strings.Trim(post, "\n\r")
	s.prePostMux.Lock()
	s.postText = post
	s.prePostMux.Unlock()
}

// SetDoneText yo
func (s *Spinner) SetDoneText(done string) {
	s.doneText = done
}

// SetInterval yo
func (s *Spinner) SetInterval(t time.Duration) {
	if t == 0 {
		s.interval = defInerval
	}
	s.interval = t

	// calibrate the spinner
	s.calibrate()
}

// SetColor yo
func (s *Spinner) SetColor(forground, background string) {
	if background == "" {
		background = s.bg
	}
	colorSetter(s, forground, background)

	// calibrate the spinner
	s.calibrate()
}

// Restart yo
func (s *Spinner) Restart() {
	s.Stop()
	s.Start()
}

func (s *Spinner) calibrate() {
	s.wg2.Add(1)
	*s.calibrator <- true
	s.wg2.Wait()
}

// Start yo
func (s *Spinner) Start() {
	s.spinningMux.Lock()
	defer s.spinningMux.Unlock()
	if s.spinning {
		return
	}

	s.wg.Add(1)
	s.spinning = true
	go func() {
		defer s.wg.Done()
		spinnerTemplate(s)
	}()
	return
}

// Stop yo
func (s *Spinner) Stop() {
	s.spinningMux.Lock()
	defer s.spinningMux.Unlock()
	if s.spinning == false {
		return
	}
	*s.stopper <- true
	s.wg.Wait()
	s.spinning = false
}

// New yo
func New(id uint, interval time.Duration, forground, background string) (*Spinner, error) {
	temp := new(Spinner)

	_, ok := SpinnerMap[id]
	if !ok {
		return nil, fmt.Errorf("%d is not a valid id", id)
	}
	temp.id = id

	if interval == 0 {
		temp.interval = defInerval
	} else {
		temp.interval = interval
	}
	temp.symbols = SpinnerMap[id]

	temp.preText = ""
	temp.postText = ""
	temp.doneText = ""

	temp.bg = background
	colorSetter(temp, forground, background)

	ch := make(chan bool)
	temp.stopper = &ch

	ch2 := make(chan bool)
	temp.calibrator = &ch2
	temp.spinning = false
	return temp, nil
}

func colorSetter(temp *Spinner, forground, background string) {
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
			fg = color.FgDefault
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
			bg = color.BgDefault
		}
	}

	if bgRBG.IsEmpty() != fgRBG.IsEmpty() {
		if bgRBG.IsEmpty() == false {
			bgRBG = color.HEX("ggg")
			bg = color.BgDefault
		} else {
			fgRBG = color.HEX("ggg")
			fg = color.FgDefault
		}
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
}
