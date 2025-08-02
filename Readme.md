# Groupie Tracker

Groupie Tracker is a web application written in Go that receives data from a public API and displays information about music artists and bands. The site presents details such as:

- **Name**
- **First Album**
- **Creation Date**
- **Members**
- **Concerts** (locations and dates)

## Features

- Home page displays all artists as cards with images and names.
- Clicking on an artist card shows a detailed page with all available information.
- Concert locations and dates are shown for each artist.
- Error handling for invalid requests and missing data.
- Uses only Go standard packages.

## How to Run

1. Clone the repository.
2. Make sure you have Go installed.
3. Run the server:
   ```
   go run .
   ```
4. Open your browser and go to [http://localhost:8080](http://localhost:8080)

## Project Structure

- `main.go` - Starts the web server and sets up routes.
- `handlers/` - Contains Go files for HTTP handlers and API data fetching.
- `templates/` - HTML templates for the home page and artist detail pages.

## API

The application uses the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api) to fetch all artist, location, date, and relation data.


## License

This project is for educational purposes.
## Authors

boulhaj
mhilli