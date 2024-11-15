# ARSG

Anime Rater and Score Generator (ARSG) is a tool designed to generate anime scores based on both user experience and objective ratings.
The goal of this project is to create a tool that helps users generate scores for sites such as MyAnimeList, AniList, or any other anime tracking platforms.

## Table of Content

- [ARSG](#arsg)
  - [Table of Content](#table-of-content)
  - [Features](#features)
  - [Installation and Execution](#installation-and-execution)
  - [TODO](#todo)
    - [Phase 1](#phase-1)
    - [Phase 2](#phase-2)
    - [Phase 3](#phase-3)
  - [License](#license)

## Features

- Scoring Based on Multiple Parameters: This tool provides users with a well-rounded scoring system, incorporating various parameters to ensure accurate and thoughtful ratings while keeping the process simple and easy.

- Personalized & Biased Scoring: A unique feature of this tool is its ability to include user biases based on personal tastes, allowing the generated scores to reflect the personality and preferences of the list author, making each list truly unique.

## Installation and Execution

1. Install the latest version of Go (Golang):

   - Download and install the latest version of Go from the official Golang website.

2. Run the application:

   - After installing Go, navigate to the project directory and run the following command based on your operating system:

     ```sh
     # For Windows
     go run main/main.go

     # For Linux/Mac
     go run ./main/main.go
     ```

3. Build the executable:

   - If the application works as intended, you can build the executable file for future use. Run the following command to compile the application:

     ```sh
     # For Windows

     go build main/main.go

     # For Linux/Mac

     go build ./main/main.go
     ```

4. Use the compiled executable:

Once the application is successfully compiled, an executable file will be generated, which you can use for future runs without needing to recompile.

## TODO

### Phase 1

- Build a simple 4-parameter scoring system, with bias as the fourth parameter.
- Create test cases and a score converter for sites based on user preferences.

### Phase 2

- Make the system more customizable by allowing users to add or remove parameters.
- Enable the application to import and export JSON files from tracking platforms, allowing users to upload their list with minimal hassle.

### Phase 3

- Add API support for major tracking platforms such as MyAnimeList, AniList, etc.
- Add a UI-based control system.
- Include in-application documentation.
- Provide executable releases.

## License

[MIT](LICENSE)
