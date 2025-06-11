ARSG
====

Table of Content
----------------

-	[ARSG](#arsg)
	-	[Introduction](#introduction)
	-	[System Requirements](#system-requirements)
	-	[Installation](#installation)
	-	[User Interface Overview](#user-interface-overview)
	-	[Using ARSG](#using-arsg)
	-	[Frequently Asked Questions (FAQ)](#frequently-asked-questions-faq)
	-	[Troubleshooting](#troubleshooting)
	-	[Support](#support)

Introduction
------------

Anime Rater and Score Generator (ARSG) is a tool designed to generate anime scores based on both user experience and objective ratings. The goal of this project is to create a tool that helps users generate scores for sites such as MyAnimeList, AniList, or any other anime tracking platforms.

System Requirements
-------------------

-	GoLang 1.24.3+
-	Sqlite3 3.49.2+
-	Make 4.4.1+ (optional)

Installation
------------

### From Source

#### makefile

1.	Clone the repository
2.	Run `make install` to install the binary
3.	Run `arsg` to start the application

#### manual

1.	Clone the repository
2.	Run `go build -o ./build/${PROJECT_NAME} ./main/main.go`
3.	Copy the binary file to `$HOME/local/bin` directory
4.	Create and copy contents of docs to `$HOME/local/share/args` directory
5.	Run `arsg` to start the application

User Interface Overview
-----------------------

ARSG applications UI is consist two modes with three different views.

-	**Main View** - This view is the default view of the application consist form for the user.

-	**Ratings View** - This list view displays the generated score for the user.

-	**Docs view** - This view displays the documentation for the application.

### Main view

The main view consist of the input fields and buttons for the user along with help text consisting of key bindings.

User Input fields:

-	Text input (alphabets, numbers, and symbols)

	-	Name
	-	Comments

-	Number input (numbers only)

	-	Art/Animation
	-	Character/Cast
	-	Plot
	-	Bias

-	Buttons

	-	**Save**: Saves info, ratings and score to the database.
	-	**Restart**: Resets the input fields to their default values.
	-	**End**: Exits the application.

-	Output

	-	**Score**: Displays the generated score for the user. Can be copied to clipboard.

### Ratings view

Displays the store info, ratings and final scores of media from database. The user can filter the list by media name.

### Docs view

Displays this manual/documentation of application.

Using ARSG
----------

WIP

Frequently Asked Questions (FAQ)
--------------------------------

WIP

Troubleshooting
---------------

WIP

Support
-------

Contact me on github: [yashodhan](https://github.com/yashodhanketkar/)

Or email me at: [kykyashodhan@gmail.com](mailto:kykyashodhan@gmail.com)
