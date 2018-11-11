package main

// Mock Data
func createMockData() {
	movies = append(movies,
		Movie{
			ID:          "1",
			Title:       "Thor: Ragnarok",
			Description: "Thor is on the other side of the universe and finds himself in a race against time to get back to Asgard to stop Ragnarok, the prophecy of destruction to his homeworld and the end of Asgardian civilization, at the hands of an all-powerful new threat, the ruthless Hela.",
			RunningTime: "126",
			ReleaseDate: "2017",
			Director:    &Director{ID: "1", FirstName: "Eric", LastName: "Pearson"},
			Genre:       &Genre{ID: "1", Name: "Action"},
		})

	movies = append(movies,
		Movie{
			ID:          "2",
			Title:       "Captain America: Civil War ",
			Description: "Following the events of Age of Ultron, the collective governments of the world pass an act designed to regulate all superhuman activity. This polarizes opinion amongst the Avengers, causing two factions to side with Iron Man or Captain America, which causes an epic battle between former allies.",
			RunningTime: "109",
			ReleaseDate: "2016",
			Director:    &Director{ID: "2", FirstName: "Joe", LastName: "Russo"},
			Genre:       &Genre{ID: "2", Name: "Adventure"},
		})

	movies = append(movies,
		Movie{
			ID:          "3",
			Title:       "Blade II",
			Description: "A rare mutation has occurred within the vampire community - The Reaper. A vampire so consumed with an insatiable bloodlust that they prey on vampires as well as humans, transforming victims who are unlucky enough to survive into Reapers themselves. Blade is asked by the Vampire Nation for his help in preventing a nightmare plague that would wipe out both humans and vampires.",
			RunningTime: "94",
			ReleaseDate: "2002",
			Director:    &Director{ID: "3", FirstName: "Gene", LastName: "Colan"},
			Genre:       &Genre{ID: "3", Name: "Fantasy"},
		})
}
