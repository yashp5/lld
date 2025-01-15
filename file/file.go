package file

/*
Level 1 – Basic File I/O
Implement basic operations for reading from and writing to files:

WRITE_TO_FILE(fileName string, content string):
	Write the specified content to a file.
	If the file already exists, overwrite it.

READ_FROM_FILE(fileName string) string:
	Read the entire content of the specified file and return it as a string.
	If the file does not exist, return an error.

APPEND_TO_FILE(fileName string, content string):
	Append the specified content to the end of the file.
*/

/*
Level 2 – Structured File Handling
Handle structured data stored in files, such as CSV or JSON:

WRITE_JSON_TO_FILE(fileName string, jsonData map[string]interface{}):
	Write a JSON object to a file.

READ_JSON_FROM_FILE(fileName string) map[string]interface{}:
	Read a JSON object from a file and parse it into a dictionary.

PARSE_CSV(fileName string) [][]string:
	Parse a CSV file into a 2D array of strings (rows and columns).

WRITE_CSV(fileName string, data [][]string):
	Write a 2D array of strings into a CSV file.
*/

/*
Level 3 – File-Based Searching and Processing
Add functionality to search and process data within files:

SEARCH_IN_FILE(fileName string, query string) []int:
	Return the line numbers where the query string appears in the file.

REPLACE_IN_FILE(fileName string, old string, new string):
	Replace all occurrences of a string in the file with another string.

SUM_COLUMN_IN_CSV(fileName string, columnIndex int) int:
	Sum all numeric values in a specified column of a CSV file.

Input CSV:
Name, Age
Alice, 30
Bob, 25
SUM_COLUMN_IN_CSV("data.csv", 1) -> 55
*/

/*
Level 4 – Large File Handling
Refactor the system to handle large files efficiently:

STREAM_READ_FILE(fileName string, chunkSize int) []string:
	Read a file in chunks of a specified size and return the chunks as a list of strings.

STREAM_WRITE_FILE(fileName string, lines []string):
	Write a file line by line from a list of strings.

FILTER_LARGE_FILE(fileName string, predicate func(string) bool):
	Filter lines in a large file based on a predicate function and write the result to a new file.
*/

/*
Level 5 – Advanced Features
Introduce file versioning, compression, and error handling:

ENABLE_VERSIONING(fileName string):
	Maintain versions of a file whenever it’s modified (e.g., fileName_v1, fileName_v2).

COMPRESS_FILE(fileName string, compressedFileName string):
	Compress the file using a simple algorithm (e.g., gzip).

DECOMPRESS_FILE(compressedFileName string, outputFileName string):
	Decompress the file back to its original state.

HANDLE_FILE_ERRORS(func() error):
	Provide a wrapper function to catch and handle file-related errors, ensuring graceful recovery.
*/

/*
Constraints
	Each file can have a maximum size of 1GB.
	The system should support simultaneous read and write operations (e.g., with locks).
	For Level 4, memory usage should not exceed 100MB while processing large files.
*/
