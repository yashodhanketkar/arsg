ARSG
====

Anime Rater and Score Generator (ARSG) is a tool designed to generate anime scores based on both user experience and objective ratings. The goal of this project is to create a tool that helps users generate scores for sites such as MyAnimeList, AniList, or any other anime tracking platforms.

Table of Content
----------------

-	[ARSG](#arsg)
	-	[Table of Content](#table-of-content)
	-	[Features](#features)
	-	[Keybinds](#keybinds)
	-	[Installation and Execution](#installation-and-execution)
	-	[TODO](#todo)
		-	[Phase 1](#phase-1)
		-	[Phase 2](#phase-2)
		-	[Phase 3](#phase-3)
	-	[License](#license)

Features
--------

-	Scoring Based on Multiple Parameters: This tool provides users with a well-rounded scoring system, incorporating various parameters to ensure accurate and thoughtful ratings while keeping the process simple and easy.

-	Personalized & Biased Scoring: A unique feature of this tool is its ability to include user biases based on personal tastes, allowing the generated scores to reflect the personality and preferences of the list author, making each list truly unique.

Keybinds
--------

| Key                                  | Action                                              |
|--------------------------------------|-----------------------------------------------------|
| <kbd>F1</kbd>                        | Show help message                                   |
| <kbd>UP</kbd> or <kbd>PGUP</kbd>     | Move cursor up                                      |
| <kbd>DOWN</kbd> or <kbd>PGDOWN</kbd> | Move cursor down                                    |
| <kbd>HOME</kbd>                      | Move cursor to the first field                      |
| <kbd>END</kbd>                       | Move cursor to the last field                       |
| <kbd>CTRL</kbd> <kbd>c</kbd>         | Copy score to clipboard                             |
| <kbd>CTRL</kbd> <kbd>t</kbd>         | Switch content type                                 |
| <kbd>DEL</kbd>                       | Reset focused field                                 |
| <kbd>ESC</kbd>                       | Reset all fields and move cursor to the first field |
| <kbd>q</kbd>                         | Quit and exit application                           |
| <kbd>CTRL</kbd> <kbd>e</kbd>         | Export ratings into export.json                     |
| <kbd>CTRL</kbd> <kbd>s</kbd>         | Switch scoring system                               |
| <kbd>CTRL</kbd> <kbd>r</kbd>         | Switch cursor mode                                  |

Installation and Execution
--------------------------

1.	Install the latest version of Go (Golang):

	-	Download and install the latest version of Go from the official Golang website.

2.	Run the application:

	-	After installing Go, navigate to the project directory and run the following command based on your operating system:

		```sh
		# For Windows
		go run main/main.go

		# For Linux/Mac
		go run ./main/main.go
		```

	-	If you have make installed, you can run the following command to build the application:

		```sh
		make run
		```

3.	Build the executable:

	-	If the application works as intended, you can build the executable file for future use. Run the following command to compile the application:

		```sh
		# For Windows

		go build main/main.go

		# For Linux/Mac

		go build ./main/main.go
		```

	-	If you have make installed, you can run the following command to build the application:

		```sh
		make build
		```

4.	Use the compiled executable:

Once the application is successfully compiled, an executable file will be generated, which you can use for future runs without needing to recompile.

TODO
----

### Phase 1

-	[X] Build a simple 4-parameter scoring system, with bias as the fourth parameter. (Completed)
-	[X] Create test cases and a score converter for sites based on user preferences. (Completed)

### Phase 2

-	[X] Enable the application to export JSON files
-	[X] Add a UI-based control system.
-	[X] Provide executable releases.

### Phase 3

-	[ ] Make the system more customizable by allowing users to add or remove parameters.
-	[X] Include in-application documentation.

### Phase 4

-	[ ] Add API support for major tracking platforms such as MyAnimeList, AniList, etc.
-	[ ] Enable the application to import and export JSON files wrt tracking platforms, allowing users to upload their list with minimal hassle.

License
-------

[MIT](LICENSE)
