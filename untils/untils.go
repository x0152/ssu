package untils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Id int

func GenerationId() Id {
	id := Id(rand.Int())
	return id
}

func WriteMsgLog(msg string) {
	fmt.Printf("%s: %s\n", time.Now().String(), msg)
	log.Printf("%s\n", msg)
}

func WriteMsgLogError(msg error) {
	fmt.Printf("%s: %s\n", time.Now().String(), msg)
	log.Printf("%s\n", msg)
}

func ConvStringToId(id string) (Id, error) {
	intId, err := strconv.Atoi(id)
	return Id(intId), err
}

func ConvStringToInt(act string) (int, error) {
	action, err := strconv.Atoi(act)
	return action, err
}
