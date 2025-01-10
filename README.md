# Groupie Trackers Search Bar

**Groupie Trackers Search Bar** is a web-based platform designed to receive, manipulate, and search data from a given API, providing comprehensive information about bands and artists. The platform features a search bar for quickly finding bands, artists, locations, and dates, alongside visualizing key details such as band names, images, founding years, concert locations, and dates in an interactive and user-friendly manner.
## Project Overview

The project focuses on displaying the following key elements retrieved from an API:

- **Artists**: Information about various bands and artists, including their name(s), image, start year, first album release date, and members.
- **Locations**: A list of last and upcoming concert locations.
- **Dates**: Concert dates, both past and upcoming.
- **Relation**: Data that links the artists, concert dates, and locations together.

## New Feature: Search Bar Integration
A powerful search bar has been implemented on the main page, allowing users to:
- Quickly search for bands or artists by name.
- Filter results by concert locations.
- Narrow down searches to specific dates.
This feature enhances the user experience by enabling efficient navigation and easy access to desired information.

### Key Features
- **Search Functionality**: Seamless searching capability across artists, bands, locations, and dates.
- **Data Visualization**: Information is displayed using various methods such as blocks, cards, tables, lists, pages, and graphs.
- **Client-Server Communication**: Event-driven features trigger client-server requests, ensuring that data is updated dynamically.
- **Error Handling**: The website is robust and handles errors gracefully, ensuring smooth operation at all times.
- **Backend in Go**: The backend of the platform is written entirely in Go, utilizing its strong concurrency and performance features.
- **Good Code Practices**: The code adheres to best practices, including the use of unit tests for critical functionality.

## Getting Started

### Prerequisites

- **Go** (1.16 or later)
- **HTML/CSS/JavaScript** for frontend development

### Installation

1. Clone the repository:

   ```bash
   git clone https://learn.zone01kisumu.ke/git/hshikuku/groupie-tracker-search-bar.git
   ```

2. Navigate to the project directory:

   ```bash
   cd groupie-trackers-search-bar
   ```

3. Run the Go server:

   ```bash
     cd cmd/server
      go run .
   ```

4. Open the website in your browser:

   ```bash
   http://localhost:8080
   ```

## Usage

Once the server is running, the homepage will display a list of bands and artists. The following features are available:
- **Search Functionality**: Use the search bar to easily find bands, artists, locations, or specific dates. The search results are displayed dynamically for quick access to the desired information.
- **View Band Details**: Click on any band or artist to view detailed information such as:
- Band Members
- Start Year and First Album Release Date
- Upcoming and Past Concert Locations
- Concert Dates

You can also trigger specific events that will send client requests to the server and update the data in real-time.

## API

The application communicates with an external RESTful API that provides the data for artists, locations, and concert dates. The relationship between these entities is maintained by the API.

## Contributing

We welcome contributions from the community. If you want to contribute:

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## Authors

- **Hezborn Shikuku** 
- **Cliff Omollo** 
- **Quinter Ochieng** 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Special thanks to all the contributors and the open-source community.
```
