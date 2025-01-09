document.addEventListener('DOMContentLoaded', () => {
    // Create and append search input
    const searchInput = document.createElement('input');
    searchInput.type = 'text';
    searchInput.placeholder = 'Search by artist-name, location-name, date, first album, creation date...';
    searchInput.classList.add('search-input');

    const searchResults = document.createElement('div');
    searchResults.classList.add('search-results');

    const header = document.querySelector('header');
    if (!header) {
        console.error("Header element not found. Ensure your HTML has a <header> element.");
        return;
    }
    header.appendChild(searchInput);
    header.appendChild(searchResults);

    // Function to fetch search results from API
    const fetchSearchResults = async (searchTerm) => {
        try {
            const response = await fetch(`/search?q=${encodeURIComponent(searchTerm)}`);
            if (!response.ok) {
                console.error("Failed to fetch search results:", response.statusText);
                return [];
            }
            const data = await response.json();
            return data;
        } catch (error) {
            console.error("Error fetching search results:", error);
            return [];
        }
    };

    // Function to handle search results display
    const displayResults = (suggestions) => {
        searchResults.innerHTML = ''; // Clear previous results

        if (suggestions.length === 0) {
            searchResults.style.display = 'none';
            return;
        }

        const resultsFragment = document.createDocumentFragment();

        suggestions.forEach(suggestion => {
            const resultItem = document.createElement('div');
            resultItem.classList.add('search-result-item');

            // Single line display with the requested format
            const resultText = document.createElement('div');
            resultText.classList.add('result-text');
            resultText.textContent = `${suggestion.Value}-${suggestion.Match}-${suggestion.Value}-By Artist:${suggestion.ArtistName}`;

            resultItem.appendChild(resultText);

            resultItem.addEventListener('click', () => {
                window.location.href = `/artist/?id=${suggestion.ArtistId}`;
            });

            resultsFragment.appendChild(resultItem);
        });

        searchResults.appendChild(resultsFragment);
        searchResults.style.display = 'block';
    };

    // Debounced search handler
    let debounceTimeout;
    const handleSearch = async (e) => {
        const searchTerm = e.target.value.trim();

        // Fetch immediately when input is 2 characters
        if (searchTerm.length === 1) {
            const suggestions = await fetchSearchResults(searchTerm);
            displayResults(suggestions);
            return;
        }

        // Apply debounce for longer input
        clearTimeout(debounceTimeout);
        debounceTimeout = setTimeout(async () => {
            if (searchTerm.length > 1) {
                const suggestions = await fetchSearchResults(searchTerm);
                displayResults(suggestions);
            } else {
                searchResults.style.display = 'Value not found';
            }
        }, 600); // 600ms debounce delay
    };

    // Add event listener for input
    searchInput.addEventListener('input', handleSearch);

    // Add event listener for Enter key press
    searchInput.addEventListener('keydown', async (e) => {
        if (e.key === 'Enter') {
            const searchTerm = searchInput.value.trim();
            if (searchTerm.length > 0) {
                const suggestions = await fetchSearchResults(searchTerm);

                // If suggestions are found, display all outputs for the searched artist
                if (suggestions.length > 0) {
                    searchResults.innerHTML = ''; // Clear previous results

                    const resultsFragment = document.createDocumentFragment();

                    suggestions.forEach(suggestion => {
                        const resultItem = document.createElement('div');
                        resultItem.classList.add('search-result-item');

                        // Display full artist details
                        const resultText = document.createElement('div');
                        resultText.classList.add('result-text');
                        resultText.innerHTML = `
                        ${suggestion.Value} -
                        ${suggestion.Match} - By Artist : 
                        ${suggestion.ArtistName}
                          
                            <br>
                        `;

                        resultItem.appendChild(resultText);

                        resultItem.addEventListener('click', () => {
                            window.location.href = `/artist/?id=${suggestion.ArtistId}`;
                        });

                        resultsFragment.appendChild(resultItem);
                    });

                    searchResults.appendChild(resultsFragment);
                    searchResults.style.display = 'block';
                } else {
                    searchResults.style.display = 'none';
                }
            }
        }
    });

    // Close search results when clicking outside
    document.addEventListener('click', (e) => {
        if (!searchResults.contains(e.target) && !searchInput.contains(e.target)) {
            searchResults.style.display = 'none';
        }
    });

    // Prevent the search results from disappearing when clicking inside them
    searchResults.addEventListener('click', (e) => {
        e.stopPropagation();
    });
});
