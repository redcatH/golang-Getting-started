package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	// var messages map[string]string
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomForamt(), name)
	return message, nil
}

func init() {
	fmt.Println("test init")
	rand.Seed(time.Now().UnixNano())
}

func randomForamt() string {
	foramt := []string{
		"Hi,%v",
		"Hi HI HI ,%v",
		"hAIL, %v",
	}
	return foramt[rand.Intn(len(foramt))]
}
