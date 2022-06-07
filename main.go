package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	pacman "PacmanGo/pacman"
)

func main() {
	var (
		configFile = flag.String("config-file", "config.json", "path to custom configuration file")
		mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
	)

	var ghostsStatusMx sync.RWMutex // guards ghosts and ghostStatus
	var pillMx sync.Mutex           // guards pillTimer

	var cfg pacman.Config      // global configuration
	var player pacman.Sprite   // global player
	var ghosts []*pacman.Ghost // global ghosts
	var maze []string          // global maze
	var score int              // global score
	var numDots int            // global number of dots
	var lives = 3              // global number of lives

	var pillTimer *time.Timer

	flag.Parse() // parse the command line arguments

	var ghostNum int
	if len(os.Args) != 2 {
		log.Println("No number of enemies provided or too many arguments. Correct usage: go run main.go [number of enemies]")
		return
	}

	ghostNum, _ = strconv.Atoi(os.Args[1])

	if ghostNum < 1 || ghostNum > 12 {
		log.Println("Invalid number of enemies. It must be between 1 and 12")
		return
	}

	pacman.Initialise()
	defer pacman.Cleanup()

	err := pacman.LoadResources(*mazeFile, *configFile, &maze, &ghosts, &player, &numDots, &cfg, ghostNum)
	if err != nil {
		log.Fatal(err)
		return
	}

	// run the game
	pacman.Run(&player, &maze, &numDots, &score, &lives, &pillMx, &ghostsStatusMx, &ghosts, pillTimer, &cfg)
}
