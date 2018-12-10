package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/code560/audigo/player"
)

type clientRepl struct {
	c Client
}

func ClientRepl(domain string, id string) {
	// init
	c := &clientRepl{
		c: NewClient(domain, id),
	}
	c.printInit()

	// // pipeline
	// if terminal.IsTerminal(0) {
	// 	src, _ := ioutil.ReadAll(os.Stdin)
	// 	c.play(string(src))
	// 	return
	// }

	// wait input
	c.cli()
}

func (r *clientRepl) play(src string) {
	r.c.Play(&player.PlayArgs{
		Src: src,
	})
}

func (r *clientRepl) stop() {
	r.c.Stop()
}

func (r *clientRepl) pause() {
	r.c.Pause()
}

func (r *clientRepl) resume() {
	r.c.Resume()
}

func (r *clientRepl) volume(a string) {
	v, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Warn(fmt.Sprintf("invalid format volume string: %s", a))
		return
	}
	r.c.Volume(&player.VolumeArgs{
		Vol: v,
	})
}

func (r *clientRepl) list() {
	files, err := ioutil.ReadDir("./asset/audio/")
	if err != nil {
		log.Warn("dont walk asset directory")
		return
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fmt.Println(f.Name())
	}
}

func (r *clientRepl) cli() {
	s := bufio.NewScanner(os.Stdin)

	for {
		r.printHead()
		s.Scan() // <- wait input
		inputs := strings.Split(s.Text(), " ")
		switch strings.ToLower(inputs[0]) {
		case "play":
			r.play(inputs[1])
		case "stop":
			r.stop()
		case "pause":
			r.pause()
		case "resume":
			r.resume()
		case "volume":
			r.volume(inputs[1])
		case "exit":
			return
		case "help":
			r.help()
		case "list":
			r.list()
		default:
			fmt.Println()
		}
	}
}

func (r *clientRepl) printInit() {
	fmt.Println("welcome to audigo client 1.0")
}

func (r *clientRepl) help() {
	fmt.Print(`
        play <file name>    play sound file
        stop                stop music
        pause               pause music
        resume              resume music
        volume <new vol>    change volume (float)

        list                show sound files
        help                this is it
        exit                finish audigo client
    `)
	fmt.Println()
}

func (r *clientRepl) printHead() {
	fmt.Print("audigo client >>> ")
}

func (r *clientRepl) printf(f string, a ...interface{}) {
	r.printHead()
	fmt.Printf(f, a...)
}

func (r *clientRepl) println(s ...interface{}) {
	r.printHead()
	fmt.Println(s...)
}
