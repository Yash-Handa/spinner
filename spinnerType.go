/*Package spinner is Command line spinner library.

Have 100+ spinners, Segregated into ASCII(id: 1-999) and Unitcode(id >= 1000) spinners and even custom (user-definned) spinners !!!

Support rich color rendering output for the spinners (in both 16bit colors and Hexcodes), universal API method, compatible with Windows system.

Made with concurrency handling in mind.

Source code and other details for the project are available at GitHub:

	https://github.com/Yash-Handa/spinner

More usage please see README and tests.
*/
package spinner

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
)

// Supported 16-bit colors, these are the most supported colors and will render on any version of cmd/ powershell.
// These can be used as both forground and background colors.
const (
	Red        = "Red"
	Black      = "Black"
	Green      = "Green"
	Yellow     = "Yellow"
	Blue       = "Blue"
	Magenta    = "Magenta"
	Cyan       = "Cyan"
	White      = "White"
	Normal     = "Normal" // Normal represents No color
	defInerval = 100 * time.Millisecond
)

// HexBgNormal is a special color that sets background for HexColor forground to nothing.
// Basically a spinner with only Hex (256) code forground color
const HexBgNormal = "HexBgNormal"

// Spinner type represents the spinner itself.
// Note it does not export anything, Use getters and setters on this type to manipulate the spinner.
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
	wg         sync.WaitGroup // for stopper
	wg2        sync.WaitGroup // for calibrator
	color      color.Style
	colorRGB   *color.RGBStyle

	// if different go routines update the value use mux
	prePostMux  sync.Mutex
	spinningMux sync.Mutex
}

// GetID returns the ID of a spinner.
//
// Note:
//  If a spinner is ASCII based its ID will be from 1 to 999.
//  If a spinner is Unicode based then its ID will be >= 1000.
//  For custom/ user-defined spinners the ID is 0.
func (s *Spinner) GetID() uint {
	return s.id
}

// GetInterval returns the interval of the spinner
func (s *Spinner) GetInterval() time.Duration {
	return s.interval
}

// GetPreText returns the current value of the text which will be printed before the spinner
func (s *Spinner) GetPreText() string {
	return s.preText
}

// GetPostText returns the current value of the text which will be printed after the spinner
func (s *Spinner) GetPostText() string {
	return s.postText
}

// GetDoneText returns the current value of the text which will be printed after the spinner has been stopped, by calling stop() method on spinner variable
func (s *Spinner) GetDoneText() string {
	return s.doneText
}

// IsSpinning returns true if the spinner is active i.e., started() has been called but the corresponding stop() method hasn't
func (s *Spinner) IsSpinning() bool {
	s.spinningMux.Lock()
	defer s.spinningMux.Unlock()
	return s.spinning
}

// SetPreText takes a string that will be printed before the spinner.
//
// SetPreText can be called multiple time during the life-time of the spinner to change the string value (default is an empty string). The spinner will check and print the value of the value of this pre-text at each interval duration (refresh time of the spinner)
//
// Note: any leading or trialing newlines ('\n', '\r\n') will be trimmed
func (s *Spinner) SetPreText(pre string) {
	pre = strings.Trim(pre, "\n\r")
	s.prePostMux.Lock()
	s.preText = pre
	s.prePostMux.Unlock()
}

// SetPostText takes a string that will be printed after the spinner.
//
// SetPostText can be called multiple time during the life-time of the spinner to change the string value (default is an empty string). The spinner will check and print the value of the value of this pre-text at each interval duration (refresh time of the spinner)
//
// Note: any leading or trialing newlines ('\n', '\r\n') will be trimmed
func (s *Spinner) SetPostText(post string) {
	post = strings.Trim(post, "\n\r")
	s.prePostMux.Lock()
	s.postText = post
	s.prePostMux.Unlock()
}

// SetDoneText takes a string that will be printed after the spinners has stopped spinning.
func (s *Spinner) SetDoneText(done string) {
	s.doneText = done
}

// SetInterval takes a parameter of time.Duration type and sets the interval / refersh rate of the spinner to it. It can be called any number of time to update the interval value.
//
// the spinner symbol, pre-text and post-text are all refreshed at each interval duraton
//
// Default: The default value or zero value is 100 * time.Millisecond
func (s *Spinner) SetInterval(t time.Duration) {
	if t == 0 {
		s.interval = defInerval
	}
	s.interval = t

	// calibrate the spinner
	s.calibrate()
}

// SetColor sets the forground and background of a spinner. It can be called any number of time to update the colors.
//
// Default: The default value or zero values are forground = spinner.Normal and background = spinner.Normal. Incase of any erroneous value these defaults are used.
//
// Note:
//  -> Both the forground and background should be either of 16-Bit type (Constants are defined at the starting of the package like spinner.Red, etc)
//     or Hex type (like "00e600", "#00E600", "ccc", etc) but not a combination of both
//
//  -> If an empty string is provided as the value of background then the priviously used background value is used
func (s *Spinner) SetColor(forground, background string) {
	if background == "" {
		background = s.bg
	} else {
		s.bg = background
	}
	colorSetter(s, forground, background)

	// calibrate the spinner
	s.calibrate()
}

// Restart stops the spinner (if not stopped already) and then restarts it.
func (s *Spinner) Restart() {
	s.Stop()
	s.Start()
}

func (s *Spinner) calibrate() {
	s.wg2.Add(1)
	*s.calibrator <- true
	s.wg2.Wait()
}

// Start starts the spinner. It is concurrency safe. Start can be called multiple times even on an already started spinner without any problem
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

// Stop stops a spinner. It is concurrency safe. Stop can be called multiple times even on an already stopped spinner without any problem
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

// New creates a new spinner based on the ID provided.
//
// Parameters:
//  id: it is a uint mapped a pre-defined spinners. If the id is invalid an error will occur
//      id from 1 to 999 represent ASCII spinners (supported on all terminals)
//      id >= 1000 represent Unicode point spinners (supported by terminals which are UTF-8 encoded)
//
//  interval: it is of type time.Duration and sets the refresh rate of the spinner
//            (the time frame after which spinner will be updated). If 0 is provided the
//            default value of 100 * time.Millisecond will be
//
//  forground: it could be a 16-Bit type (16-bit colors are defined as constants at the
//             starting of the package like spinner.Red, etc) or Hex type (strings like
//             "00e600", "#00E600", "ccc", etc). Empty string will use spinner.Normal
//
//  background: it could be a 16-Bit type (16-bit colors are defined as constants at the
//             starting of the package like spinner.Red, etc) or Hex type (strings like
//             "00e600", "#00E600", "ccc", etc). Empty string will use spinner.Normal
//
//  Note- forground and background should be of the same form either both will be 16-bit
//  colors or both will be hex. Not a combination. In case of any error in resolving
//  colors both forground and background will use spinner.Normal
//
// Example:
//  package main
//
//  import (
//  	"time"
//  	"log"
//
//  	spinner "github.com/Yash-Handa/spinner"
//  )
//
//  func main() {
//  	sp, err := spinner.Custom(4, 50 * time.Millisecond, spinner.Cyan, spinner.Normal)
//  	if err != nil {
//  		log.Fatal(err)
//  	}
//
//  	sp.SetPostText("  A custom spinner")
//  	sp.SetDoneText("Hurray spinner worked\n")
//
//  	sp.Start() // the spinner starts
//  	time.Sleep(3 * time.Second) // after 3 seconds
//  	sp.SetColor(spinner.Magenta, "") // use previous background color
//  	sp.SetInterval(100 * time.Millisecond)
//  	sp.SetPostText("  The color and speed Changed !!")
//  	time.Sleep(3 * time.Second)
//  	sp.Stop() // the spinner stops
//  }
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

// Custom creates a new spinner based on user defined symbols (a slice of strings).
//
// Parameters:
//  symbols: it is slice of stings used to produce a spinner. All strings in the slice
//           should be of same size or an error will be returned
//
//  interval: it is of type time.Duration and sets the refresh rate of the spinner
//            (the time frame after which spinner will be updated). If 0 is provided the
//            default value of 100 * time.Millisecond will be
//
//  forground: it could be a 16-Bit type (16-bit colors are defined as constants at the
//             starting of the package like spinner.Red, etc) or Hex type (strings like
//             "00e600", "#00E600", "ccc", etc). Empty string will use spinner.Normal
//
//  background: it could be a 16-Bit type (16-bit colors are defined as constants at the
//             starting of the package like spinner.Red, etc) or Hex type (strings like
//             "00e600", "#00E600", "ccc", etc). Empty string will use spinner.Normal
//
//  Note- forground and background should be of the same form either both will be 16-bit
//  colors or both will be hex. Not a combination. In case of any error in resolving
//  colors both forground and background will use spinner.Normal
//
// Example:
//  package main
//
//  import (
//  	"time"
//  	"log"
//
//  	spinner "github.com/Yash-Handa/spinner"
//  )
//
//  func main() {
//  	symbols := []string{"N   ", "IN  ", "PIN ", "SPIN", " SPI", "  SP", "   S", "    "}
//  	sp, err := spinner.Custom(symbols, 0, spinner.Red, spinner.Normal)
//  	if err != nil {
//  		log.Fatal(err)
//  	}
//
//  	sp.SetPostText("  A custom spinner")
//
//  	sp.Start() // the spinner starts
//  	time.Sleep(3 * time.Second) // after 3 seconds the color changes to lime green
//  	sp.SetColor("00e600", spinner.HexBgNormal) // spinner.HexBgNormal is used with Hex forground to indicate that no background color to be used
//  	sp.SetPostText("  The color Changed !!")
//  	time.Sleep(3 * time.Second)
//  	sp.Stop() // the spinner stops
//  }
//
// all custom spinners have ID of 0
func Custom(symbols []string, interval time.Duration, forground, background string) (*Spinner, error) {
	l := len(symbols[0])
	for _, str := range symbols {
		if l != len(str) {
			return nil, fmt.Errorf("all symbols of the spinner should be of the same length")
		}
	}
	temp := new(Spinner)
	temp.id = 0

	if interval == 0 {
		temp.interval = defInerval
	} else {
		temp.interval = interval
	}
	temp.symbols = symbols

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
	case Red:
		fg = color.FgRed
	case Black:
		fg = color.FgBlack
	case Green:
		fg = color.FgGreen
	case Blue:
		fg = color.FgBlue
	case Magenta:
		fg = color.FgMagenta
	case Yellow:
		fg = color.FgYellow
	case Cyan:
		fg = color.FgCyan
	case White:
		fg = color.BgWhite
	case Normal, "":
	default:
		// check if it is a hex code...
		fgRBG = color.HEX(forground, false)
		if fgRBG.IsEmpty() {
			fg = color.FgDefault
		}
	}

	switch background {
	case Red:
		bg = color.BgRed
	case Black:
		bg = color.BgBlack
	case Green:
		bg = color.BgGreen
	case Blue:
		bg = color.BgBlue
	case Magenta:
		bg = color.BgMagenta
	case Yellow:
		bg = color.BgYellow
	case Cyan:
		bg = color.BgCyan
	case White:
		bg = color.BgWhite
	case HexBgNormal:
		bgRBG = color.HEX("000000")
	case Normal, "":
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
		if background == HexBgNormal {
			temp.colorRGB = color.NewRGBStyle(fgRBG)
		} else {
			temp.colorRGB = color.NewRGBStyle(fgRBG, bgRBG)
		}
	} else {
		temp.color = color.New(fg, bg)
		temp.colorRGB = nil
	}
}
