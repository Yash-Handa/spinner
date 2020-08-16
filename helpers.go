package spinner

import (
	"fmt"
	"math/rand"
	"time"
)

// Random16BitColor returns the string representation of a random 16 Bit color
func Random16BitColor() string {
	c := [...]string{Red, Black, Green, Yellow, Blue, Magenta, Cyan, White,
		Normal}
	rand.Seed(time.Now().UnixNano())
	return c[rand.Intn(9)]
}

// RandomHexColor return the string representation of the hex code ("#rrggbb") of a a random color
func RandomHexColor() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(256)
	rand.Seed(13 * time.Now().UnixNano())
	g := rand.Intn(256)
	rand.Seed(271 * time.Now().UnixNano())
	b := rand.Intn(256)

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
