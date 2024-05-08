package art

import (
	"log"
)

func Check(s string) {
	for _, i := range s {
		if i > 126 || i < 32 {
			log.Fatal("there are invalide agre in text ")
		}
	}
}
