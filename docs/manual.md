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

Using ARSG
----------

Frequently Asked Questions (FAQ)
--------------------------------

Troubleshooting
---------------

Support
-------
