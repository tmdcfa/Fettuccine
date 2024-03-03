# Fettuccine

```# Restaurant CLI Application
#
## Overview
This Command-Line Interface (CLI) application fetches restaurant data based on a provided postcode and presents it in the console. It retrieves essential information such as the restaurant's name, cuisines, star rating, and address.

#
## Building and Running the Application

## Clone the Repository
`git clone <repository-url>`
## Navigate to the Project Directory
`cd restaurant-cli`
## Compile and Execute the Application
`go run main.go`
## Usage Instructions
- Upon execution, the application will prompt you to input a postcode.
- Enter the postcode as requested and press Enter.
- The application will then fetch and display restaurant information based on the provided postcode.

#
## Assumptions and Clarifications
- The application is designed to be interactive, prompting users to input a postcode for targeted restaurant information    retrieval.
- To prevent overwhelming the console, only data for the first 10 restaurants is displayed.
- Cuisine information is formatted as a comma-separated string to enhance readability.

#
## Potential Enhancements
- **Sorting:** Allow users to sort restaurant results based on different criteria such as star rating, distance, or alphabetical order.
- **Filtering:** Implement filtering options to refine restaurant results based on specific criteria such as cuisine type, price range, or dietary preferences.
- **Interactive Menu:** Create an interactive menu system within the CLI interface, enabling users to navigate through different options and features more intuitively.
- **Data Visualization:** Incorporate data visualization techniques to present restaurant information graphically, providing users with a visual representation of the data alongside text-based information.```