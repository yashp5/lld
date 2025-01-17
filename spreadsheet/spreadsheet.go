package spreadsheet

/*
Design an Excel-like spreadsheet
Its more like modified system design questions in a leetcode format
I rmb my final being creating a spreadsheet that supports formula applications across different rows, and you have to make sure there’s no cycles in formula applications
*/

/*
Level 1 – Initial Design & Basic Functions
Implement the following methods for basic spreadsheet operations:

SET(cell string, value string):
	Sets the value of the specified cell.
	If the cell already contains a value, overwrite it.

GET(cell string) string:
	Retrieves the value of the specified cell.
	If the cell is empty, return "0".

CLEAR(cell string):
	Resets the cell to its default value ("0").

SET("A1", "5")
SET("B1", "10")
*/

/*
Level 2 – Data Structures & Data Processing
Expand the spreadsheet to support formulas:

SET_FORMULA(cell string, formula string):
	Sets a formula in a cell (e.g., "=A1+B1").
	The formula references other cells, and the value should be calculated dynamically when needed.

EVALUATE(cell string) int:
	Computes the value of a cell with a formula.
	Supports basic operations (+, -, *, /).
	If referenced cells contain formulas, resolve them recursively.

SET("A1", "5")
SET("B1", "10")
SET_FORMULA("C1", "=A1+B1")
EVALUATE("C1") -> 15
*/

/*
Level 3 – Refactoring & Encapsulation
Add the ability to define ranges and perform operations over them:

SET_RANGE(startCell string, endCell string, value int):
	Sets a value for all cells within the specified range.

SUM_RANGE(startCell string, endCell string) int:
	Computes the sum of all values within a range of cells.

SET_RANGE("A1", "A3", 5)
SUM_RANGE("A1", "A3") -> 15
*/

/*
Level 4 – Extending Design & Functionality
Introduce snapshotting and rollback for the spreadsheet:

SNAPSHOT():
Save the current state of the spreadsheet.

ROLLBACK():
	Restore the spreadsheet to the last saved state.
	All subsequent changes after the snapshot will be undone.

SET("A1", "5")
SNAPSHOT()
SET("A1", "10")
ROLLBACK()
GET("A1") -> "5"
*/

/*
Constraints
	Cell names follow Excel-like conventions (e.g., "A1", "B2").
	The spreadsheet has a fixed size of 100x100 (cells A1 to J100).
	Formulas and operations only involve cells within the spreadsheet bounds.
*/
