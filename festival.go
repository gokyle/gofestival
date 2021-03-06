// package festival provides a Go interface to festival's text-to-speech
// functionality.
package festival

import (
	"os/exec"
	"strings"
)

func setupCmd() (festival *exec.Cmd, err error) {
	path, err := exec.LookPath("festival")
	if err != nil {
		return
	}
	festival = exec.Command(path, "--tts")
	return
}

// Speak passes the text to festival, speaking it. It will block until
// the command finishes and the text is done being spoken.
func Speak(text string) (err error) {
	var festival *exec.Cmd
	if festival, err = setupCmd(); err != nil {
		return
	}
	reader := strings.NewReader(text)

	festival.Stdin = reader
	err = festival.Run()
	return
}
