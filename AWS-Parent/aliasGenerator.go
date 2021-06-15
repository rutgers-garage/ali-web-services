package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	adjectives := []string{
		"regular",
		"diligent",
		"flowery",
		"right",
		"truthful",
		"inquisitive",
		"greedy",
		"spurious",
		"volatile",
		"elegant",
		"dramatic",
		"befitting",
		"furtive",
		"clammy",
		"rough",
		"spotless",
		"malicious",
		"thirsty",
		"quixotic",
		"onerous",
		"little",
		"temporary",
		"elderly",
		"profuse",
		"noisy",
		"subsequent",
		"five",
		"physical",
		"relieved",
		"lopsided",
		"ethereal",
		"symptomatic",
		"whole",
		"mixed",
		"well",
		"spotted",
		"flagrant",
		"slim",
		"maniacal",
		"certain",
		"workable",
		"ate",
		"supreme",
		"reminiscent",
		"annoyed",
		"lively",
		"sparkling",
		"huge",
		"brave",
		"historical",
		"humorous",
		"wry",
		"imperfect",
		"successful",
		"rude",
		"lovely",
		"smiling",
		"hurt",
		"resonant",
		"overt",
		"small",
		"torpid",
		"frightened",
		"rightful",
		"aware",
		"chief",
		"callous",
		"somber",
		"cheap",
		"educated",
		"synonymous",
		"zonked",
		"chemical",
		"rich",
		"icky",
		"tremendous",
		"elastic",
		"thoughtless",
		"glistening",
		"efficacious",
		"animated",
		"pumped",
		"judicious",
		"cumbersome",
		"icy",
		"wicked",
		"handsome",
		"entertaining",
		"absent",
		"scattered",
		"unhealthy",
		"clumsy",
		"defeated",
		"superficial",
		"ultra",
		"bashful",
		"draconian",
		"tawdry",
		"itchy",
		"super",
	}

	nouns := []string{
		"aunt",
		"harmony",
		"magic",
		"bait",
		"sail",
		"drawer",
		"holiday",
		"earthquake",
		"rail",
		"stage",
		"zephyr",
		"prose",
		"plant",
		"property",
		"floor",
		"arch",
		"rate",
		"limit",
		"farm",
		"oatmeal",
		"fowl",
		"coil",
		"punishment",
		"flock",
		"sand",
		"ray",
		"love",
		"cup",
		"scarf",
		"sound",
		"crib",
		"ladybug",
		"harbor",
		"wire",
		"channel",
		"boundary",
		"airport",
		"top",
		"rule",
		"crayon",
		"giraffe",
		"wrist",
		"eggs",
		"north",
		"grain",
		"spring",
		"worm",
		"ice",
		"flight",
		"toothbrush",
		"advice",
		"sheep",
		"pleasure",
		"card",
		"spade",
		"dirt",
		"home",
		"basketball",
		"trousers",
		"potato",
		"belief",
		"whip",
		"voice",
		"parcel",
		"house",
		"linen",
		"record",
		"giants",
		"scene",
		"expert",
		"digestion",
		"quiet",
		"ocean",
		"duck",
		"reason",
		"income",
		"girls",
		"toe",
		"road",
		"pie",
		"queen",
		"cobweb",
		"apparel",
		"grade",
		"offer",
		"insect",
		"anger",
		"engine",
		"cart",
		"protest",
		"room",
		"hospital",
		"air",
		"breath",
		"knee",
		"chicken",
		"advertisement",
		"partner",
		"wealth",
		"monkey",
		"monkey",
	}

	rand.Seed(time.Now().UnixNano())

	newName := adjectives[rand.Intn(100)] + "-" + nouns[rand.Intn(100)]

	if _, err := os.Stat("takenAliases.json"); err == nil {
		var names []string
		readBytes, _ := ioutil.ReadFile("takenAliases.json")
		json.Unmarshal(readBytes, &names)

		count := 0
		found := 0

		for true {
			if found == count {
				break;
			}

			if names[count] == newName {
				newName = adjectives[rand.Intn(100)] + "-" + nouns[rand.Intn(100)]
				found = count
			}

			count++

			if found != 0 && count == len(names) {
				count = 0
			} else {
				break;
			}
    	}

		names = append(names, newName)
		marshaledDict, _ := json.Marshal(names)
		err := ioutil.WriteFile("takenAliases.json", marshaledDict, 0644)
    	check(err)
	} else if os.IsNotExist(err) {
		newAliasDict := []string{newName}
    	marshaledDict, _ := json.Marshal(newAliasDict)
		err := ioutil.WriteFile("takenAliases.json", marshaledDict, 0644)
    	check(err)
	}
}
