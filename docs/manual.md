ARSG
====

Table of Contents
-----------------

-	[Introduction](#introduction)
-	[Installation](#installation)
	-	[From Source](#from-source)
-	[User Interface Overview](#user-interface-overview)
	-	[Main View](#main-view)
	-	[Ratings View](#ratings-view)
	-	[Docs View](#docs-view)
-	[Using ARSG](#using-arsg)
	-	[Keyboard Shortcuts](#keyboard-shortcuts)
-	[Contact](#contact)

Introduction
------------

Anime Rater and Score Generator (ARSG) is a simple, terminal-based tool designed to help users generate anime scores based on both subjective experience and objective criteria. It's especially useful for rating entries on platforms like MyAnimeList, AniList, or similar anime tracking services. System Requirements

-	GoLang 1.24.3 or later
-	SQLite3 3.49.2 or later
-	Make 4.4.1 or later (optional)

Installation
------------

### From Source

#### Using make

-	Clone the repository.
-	Run make install to build and install the binary.
-	Run arsg to start the application.

#### Manual Installation

-	Clone the repository.
-	Run go build -o ./build/${PROJECT_NAME} ./main/main.go
-	Copy the binary to your $HOME/local/bin directory.
-	Create the docs directory at $HOME/local/share/arsg and copy the documentation files there.
-	Run arsg to launch the application.

User Interface Overview
-----------------------

The ARSG application offers two modes with three different views:

-	Main View – The default input form where users enter ratings.
-	Ratings View – Displays saved entries and their calculated scores.
-	Docs View – Displays this user manual inside the application.

### Main View

This is the primary view, featuring:

```
- Text Inputs
  - Name
  - Comments (optional)

- Numerical Inputs (accepts numbers only):
  - Art/Animation
  - Character/Cast
  - Plot
  - Bias

- Buttons:
  - Save – Stores the data in the database.
  - Restart – Clears all fields and resets to defaults.
  - End – Exits the application.

- Output:
  - Score – Automatically generated from your inputs; can be copied to clipboard.
  - Help Box: Located at the bottom; displays available key bindings.
```

### Ratings View

Displays stored entries (name, ratings, score). A search allows you to filter entries by media name.

Navigation key:

<kbd>F1</kbd> – Opens Docs View

<kbd>F3</kbd> – Returns to Main View

### Docs View

Displays this manual.

Navigation key:

<kbd>F1</kbd> – Returns to the previously visited view (Main or Ratings)

Using ARSG
----------

In the Main View, enter the required details in the input fields. After filling in the numerical fields, a score is automatically calculated and shown in the output field.

#### To save a rating:

-	Name and all numerical fields are required.
-	Comments is optional.
-	Click Save to store the entry. A confirmation dialog appears with a Confirm button to return to the main view.

#### To reset:

-	Click Restart to clear all fields and reset focus to the Name field.

#### To exit:

-	Click End to close the application.

### Keyboard Shortcuts:

<kbd>F1</kbd> – Switch to Docs View

<kbd>F3</kbd> – Switch to Ratings View (from Main or Docs)

Note: Refer to the help box at the bottom of the screen for a full list of key bindings. Support

Contact
-------

Github: [yashodhan](https://github.com/yashodhanketkar/)

Email: [kykyashodhan@gmail.com](mailto:kykyashodhan@gmail.com)
