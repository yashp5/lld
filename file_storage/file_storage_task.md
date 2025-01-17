# Simplified File Storage Service

You are tasked with implementing a simplified file storage service that allows users to upload, retrieve, copy, and manage files. Each file has a name, size, and optional time-to-live (TTL) value.

---

## Level 1 – Initial Design & Basic Functions

Implement the following methods:

- **`FILE_UPLOAD(fileName string, fileSize int)`**
  Upload a file to the storage system.  
  - If a file with the same name already exists, throw an error.

- **`FILE_GET(fileName string) int`**  
  Retrieve the size of a file by its name.  
  - If the file does not exist, return `-1`.

- **`FILE_COPY(sourceFile string, destFile string)`**  
  Copy a file to a new location with a different name.  
  - If the source file does not exist, throw an error.  
  - If the destination file already exists, overwrite it.

---

## Level 2 – Data Structures & Data Processing

Add the following method:

- **`FILE_SEARCH(prefix string) []string`**  
  Return the top 10 files whose names start with the provided prefix.  
  - Sort the results in descending order by size.  
  - In case of a tie, sort by file name in ascending order.

---

## Level 3 – Refactoring & Encapsulation

Extend the previous functionality to include TTL (Time to Live). Update the methods as follows:

- **`FILE_UPLOAD_AT(timestamp int, fileName string, fileSize int, ttl int)`**  
  Upload a file with an optional TTL.  
  - Files with TTL will expire after the specified duration.

- **`FILE_GET_AT(timestamp int, fileName string) int`**  
  Retrieve the size of a file at a given timestamp.  
  - Return `-1` if the file has expired or does not exist.

- **`FILE_COPY_AT(timestamp int, sourceFile string, destFile string)`**  
  Copy a file with all its metadata (including TTL) at a given timestamp.

- **`FILE_SEARCH_AT(timestamp int, prefix string) []string`**  
  Return the top 10 files that are still "alive" at the specified timestamp.

---

## Level 4 – Extending Design & Functionality

Add rollback functionality:

- **`ROLLBACK(timestamp int)`**  
  Roll back the state of the file storage system to the specified timestamp.  
  - Ensure that TTLs are recalculated accordingly, and expired files are removed from the system.

