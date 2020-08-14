package spinner

var spinnerMap = map[uint]func(*Spinner){
	1:  spinnerTemplate1([]string{"|", "/", "-", "|", "\\"}),
	2:  spinnerTemplate1([]string{"[   ]", "[=  ]", "[ = ]", "[  =]"}),
	3:  spinnerTemplate1([]string{"[=   ]", "[==  ]", "[ == ]", "[  ==]", "[   =]"}),
	4:  spinnerTemplate1([]string{"[    ]", "[#   ]", "[##  ]", "[### ]", "[####]", "[ ###]", "[  ##]", "[   #]"}),
	5:  spinnerTemplate1([]string{"   ", ".  ", ".. ", "..."}),
	6:  spinnerTemplate1([]string{"   ", ".  ", ".. ", "...", " ..", "  ."}),
	7:  spinnerTemplate1([]string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[====]"}),
	8:  spinnerTemplate1([]string{"(-*--)", "(--*-)", "(---*)", "(--*-)", "(-*--)", "(*---)"}),
	9:  spinnerTemplate1([]string{"( *  )", "(  * )", "(   *)", "(  * )", "( *  )", "(*   )"}),
	10: spinnerTemplate1([]string{"|   ", "||  ", "||| ", "||||", " |||", "  ||", "   |"}),

	100: spinnerTemplate1([]string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}),
	102: spinnerTemplate1([]string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"}),
	101: spinnerTemplate1([]string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}),
	103: spinnerTemplate1([]string{"▖", "▘", "▝", "▗"}),
	104: spinnerTemplate1([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}),
	105: spinnerTemplate1([]string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}),
	106: spinnerTemplate1([]string{"◢", "◣", "◤", "◥"}),
	107: spinnerTemplate1([]string{"◰", "◳", "◲", "◱"}),
	108: spinnerTemplate1([]string{"◴", "◷", "◶", "◵"}),
	109: spinnerTemplate1([]string{"◐", "◓", "◑", "◒"}),
	110: spinnerTemplate1([]string{"◡◡", "⊙⊙", "◠◠"}),
	111: spinnerTemplate1([]string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}),
}
