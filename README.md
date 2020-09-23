# defuse

CS:GO Demo Analyzer written in Go.



# TODO

- Finish HLTV scraping library, under hltv-scrape.

- Create a Team struct with useful data

- Create a Map struct with useful data, including coords of nade landing spots. Then create an array of Maps with all metadata in each one.

- Create a Game struct with all data, including data scraped from HLTV, aswell as data scraped from the demo.

- Fork cs-demoinfo library, add mappings of coords to map callout names.

- Add functionality to extract where grenade was thrown from and landing location.

- Use this to implement a simple execute detect.




# Roadmap

- Use fyne as a GUI library.


- Creation of a library to track executes and strats off of patterns from player movements and nade landing spots & nade trajectories.

- Designing ML to predict success of round/strat from this training data.

- Using CV/AI to create serialized data so model can be applied live without access to live demo feed.

