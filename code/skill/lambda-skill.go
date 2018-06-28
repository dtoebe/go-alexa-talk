package main

import (
	"fmt"
	"math/rand"
	"time"

	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

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
			// "The temperature inside the Sun can reach 15 million degrees Celsius.",
		},
		"saturn": []string{
			"Saturn radiates two and a half times more energy into space than it receives from the sun.",
		},
		"moon": []string{
			"The Moon is moving approximately 3.8 cm away from our planet every year.",
		},
	}
)

func helpHandler(req alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Help for Planets",
		"To get a random fact about a random planet say \"ask Planets for random planet\"")
}

func randomPlanetHandler(req alexa.Request) alexa.Response {
	keys := make([]string, len(facts))
	i := 0
	for k := range facts {
		keys[i] = k
		i++
	}
	rand.Seed(time.Now().Unix())
	num := rand.Intn(len(keys))
	planet := keys[num]

	num = rand.Intn(len(facts[planet]))

	return alexa.NewSimpleResponse(
		fmt.Sprintf("Fun fact about planet %s", planet),
		facts[planet][num],
	)
}

func getPlanetHandler(req alexa.Request) alexa.Response {
	planet := req.Body.Intent.Slots["PlanetName"].Value

	if planet == "" {
		return randomPlanetHandler(req)
	}

	rand.Seed(time.Now().Unix())
	num := rand.Intn(len(facts[planet]))

	return alexa.NewSimpleResponse(
		fmt.Sprintf("Fun fact about planet %s", planet),
		facts[planet][num],
	)
}

func dispatchIntents(req alexa.Request) alexa.Response {
	var res alexa.Response
	switch req.Body.Intent.Name {
	case "RandomPlanetIntent":
		res = randomPlanetHandler(req)
	case "PickPlanetintent":
		res = getPlanetHandler(req)
	default:
		res = helpHandler(req)
	}

	return res
}

func handler(req alexa.Request) (alexa.Response, error) {
	return dispatchIntents(req), nil
}

func main() {
	lambda.Start(handler)
}
