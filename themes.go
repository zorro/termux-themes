package themes

import (
	"fmt"
	"os"
	"path/filepath"
)

type T struct {
	Themes []Themes `json:"themes"`
}
type Themes struct {
	Name       string `json:"name"`
	Foreground string `json:"foreground"`
	Background string `json:"background"`
	Cursor     string `json:"cursor"`
	Color0     string `json:"color0"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	Color4     string `json:"color4"`
	Color5     string `json:"color5"`
	Color6     string `json:"color6"`
	Color7     string `json:"color7"`
	Color8     string `json:"color8"`
	Color9     string `json:"color9"`
	Color10    string `json:"color10"`
	Color11    string `json:"color11"`
	Color12    string `json:"color12"`
	Color13    string `json:"color13"`
	Color14    string `json:"color14"`
	Color15    string `json:"color15"`
}

var ppp string = `
# Name: %v

foreground=%v
background=%v
cursor=%v

color0=%v
color1=%v
color2=%v
color3=%v
color4=%v
color5=%v
color6=%v
color7=%v
color8=%v
color9=%v
color10=%v
color11=%v
color12=%v
color13=%v
color14=%v
color15=%v
`

func (t *T) list() {
	exepath, _ := os.Executable()
	exe := filepath.Base(exepath)
	for i, l := range t.Themes {
		fmt.Printf("  %v %v\n", i+1, l.Name)
	}
	fmt.Printf("Usage: %v <theme number>\n", exe)
}
func (t *T) set(n int) {
	f, _ := os.Create("/data/data/com.termux/files/home/.termux/colors.properties")
	for i, l := range t.Themes {
		if i+1 == n {
			colors := fmt.Sprintf(ppp, l.Name, l.Foreground,
				l.Background, l.Cursor, l.Color0,
				l.Color1, l.Color2, l.Color3,
				l.Color4, l.Color5, l.Color6,
				l.Color7, l.Color8, l.Color9,
				l.Color10, l.Color11, l.Color12,
				l.Color13, l.Color14, l.Color15)
			f.Write([]byte(colors))
		}
	}
}
