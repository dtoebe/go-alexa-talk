package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	Planet string `json:"planet"`
}

type response struct {
	Planet string `json:"panet"`
	Fact   string `json:"fact"`
}

var (
	// Source: alexaskillskitnodejsfact
	facts = map[string][]string{
		"mercury": []string{
			"A year on Mercury is just 88 days long.",
		},
		"venus": []string{
			"Despite being farther from the Sun, Venus experiences higher temperatures than Mercury.",
			"Venus rotates counter-clockwise, possibly because of a collision in the past with an asteroid.",
		},
		"mars": []string{
			"On Mars, the Sun appears about half the size as it does on Earth.",
		},
		"earth": []string{
			"Earth is the only planet not named after a god.",
		},
		"jupiter": []string{
			"Jupiter has the shortest day of all the planets.",
		},
		"sun": []string{
			"The Sun contains 99.86%% of the mass in the Solar System.",
			"The Sun is an almost perfect sphere.",
			"The temperature inside the Sun can reach 15 million degrees Celsius.",
		},
		"saturn": []string{
			"Saturn radiates two and a half times more energy into space than it receives from the sun.",
		},
		"moon": []string{
			"The Moon is moving approximately 3.8 cm away from our planet every year.",
		},
	}
)

func getFact(ctx context.Context, req request) (response, error) {
	rand.Seed(time.Now().Unix())

	planet := req.Planet
	if planet == "" {
		keys := make([]string, len(facts))
		i := 0
		for k := range facts {
			keys[i] = k
			i++
		}
		num := rand.Intn(len(keys))
		planet = keys[num]
	}

	num := rand.Intn(len(facts[planet]))

	return response{Planet: planet, Fact: facts[planet][num]}, nil
}

func main() {
	lambda.Start(getFact)
}
