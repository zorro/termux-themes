package themes

import (
	_ "embed"
	"encoding/json"
	"flag"
	"os/exec"
	"strconv"
)

//go:embed data/themes.json
var th []byte

func Set() {
	flag.Parse()
	p := &T{}
	json.Unmarshal(th, p)
	s, _ := strconv.Atoi(flag.Arg(0))
	if s < 1 || s > 237 {
		p.list()
	} else {
		p.set(s)
		cmd := exec.Command("termux-reload-settings")
		cmd.Run()
	}
}
