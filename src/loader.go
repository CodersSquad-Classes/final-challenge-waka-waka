package src

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Player           string        `json:"player"`
	MaxLifestate     int           `json:"maxLifestate"`
	Ghost            string        `json:"ghost"`
	GhostBlue        string        `json:"ghost_blue"`
	Wall             string        `json:"wall"`
	Dot              string        `json:"dot"`
	Pill             string        `json:"pill"`
	Death            string        `json:"death"`
	Space            string        `json:"space"`
	UseEmoji         bool          `json:"use_emoji"`
	PillDurationSecs time.Duration `json:"pill_duration_secs"`
}

func LoadResources(mazefile, configFile string, maze *[]string, ghosts *[]*Ghost, player *Sprite, numDots *int, cfg *Config, ghostNum int) error {
	err := LoadMaze(mazefile, maze, ghosts, player, numDots, ghostNum)
	if err != nil {
		log.Println("failed to load maze:", err)
		return err
	}

	err = LoadConfig(configFile, cfg)
	if err != nil {
		log.Println("failed to load configuration:", err)
		return err
	}

	return nil
}

func LoadMaze(file string, maze *[]string, ghosts *[]*Ghost, player *Sprite, numDots *int, ghostsNum int) error {
	f, err := os.Open(file) // open file
	if err != nil {
		return err
	}
	defer f.Close() // close file on return

	scanner := bufio.NewScanner(f) // create scanner
	for scanner.Scan() {
		line := scanner.Text()      // get next line
		*maze = append(*maze, line) // append line to maze
	}

	for row, line := range *maze {
		for col, char := range line {
			switch char {
			case 'P':
				*player = Sprite{Row: row, Col: col, StartRow: row, StartCol: col}
			case 'G':
				if len(*ghosts) < ghostsNum {
					*ghosts = append(*ghosts, &Ghost{Position: Sprite{Row: row, Col: col, StartRow: row, StartCol: col}, Status: GhostStatusNormal})
				}
			case '.':
				*numDots++
			}
		}
	}

	return nil
}

func LoadConfig(file string, cfg *Config) error {
	f, err := os.Open(file) // open file
	if err != nil {
		return err
	}

	defer f.Close() // close file on return

	decoder := json.NewDecoder(f) // create json decoder
	err = decoder.Decode(&cfg)    // decode file into config struct
	if err != nil {
		return err
	}

	return nil
}
