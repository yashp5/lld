package json

/*
Level 1 – Basic JSON Parsing
Implement a system to parse and extract data from JSON strings.

PARSE_JSON(input string) map[string]interface{}:
	Parse a JSON string into a dictionary.
	Return an error if the input is not valid JSON.

GET_VALUE(json map[string]interface{}, key string) interface{}:
	Retrieve the value associated with a specific key from the parsed JSON.
	Support nested keys using dot notation (e.g., "user.address.city").

LIST_KEYS(json map[string]interface{}) []string:
	Return a list of all keys in the JSON object, including nested keys.

Input JSON: {"user": {"name": "Alice", "age": 30}}
GET_VALUE(json, "user.name") -> "Alice"
LIST_KEYS(json) -> ["user.name", "user.age"]
*/

/*
Level 2 – JSON Manipulation
Enhance the system to modify JSON objects:

SET_VALUE(json map[string]interface{}, key string, value interface{}):
	Set or update the value for a specific key in the JSON object.
	Create nested keys if they do not exist.

DELETE_KEY(json map[string]interface{}, key string):
	Remove a specific key from the JSON object.
	If the key is nested, remove only that key.

MERGE_JSON(json1, json2 map[string]interface{}) map[string]interface{}:
	Merge two JSON objects, with values from json2 overwriting those from json1 for overlapping keys.
	Original JSON: {"user": {"name": "Alice", "age": 30}}

SET_VALUE(json, "user.address.city", "New York")
DELETE_KEY(json, "user.age")
MERGE_JSON({"a": 1}, {"b": 2, "a": 3}) -> {"a": 3, "b": 2}

*/

/*
Level 3 – Validation and Querying
Add support for validating and querying JSON data:

VALIDATE_JSON(json map[string]interface{}, schema map[string]string) bool:
	Validate the JSON object against a schema, where:
	Keys in the schema represent JSON keys.
	Values in the schema represent expected data types (e.g., "string", "int").

QUERY_JSON(json map[string]interface{}, query string) []interface{}:
	Query the JSON object using a basic query language (e.g., SELECT user.name WHERE user.age > 25).

Schema: {"user.name": "string", "user.age": "int"}
VALIDATE_JSON({"user": {"name": "Alice", "age": 30}}, schema) -> true
QUERY_JSON({"users": [{"name": "Alice", "age": 30}, {"name": "Bob", "age": 20}]}, "SELECT name WHERE age > 25") -> ["Alice"]
*/

/*
Level 4 – Advanced Features
Introduce advanced manipulation and transformation features:

FLATTEN_JSON(json map[string]interface{}) map[string]interface{}:
	Flatten a nested JSON object into a single level using dot notation for keys.

UNFLATTEN_JSON(flattened map[string]interface{}) map[string]interface{}:
	Convert a flattened JSON object back into its nested structure.

TRANSFORM_JSON(json map[string]interface{}, transformations map[string]string):
	Apply a set of transformations to the JSON object. Transformations are specified as:
	"user.name -> user.fullName"
	"user.age -> user.details.age"

Input JSON: {"user": {"name": "Alice", "address": {"city": "NY"}}}
FLATTEN_JSON(json) -> {"user.name": "Alice", "user.address.city": "NY"}
UNFLATTEN_JSON({"user.name": "Alice"}) -> {"user": {"name": "Alice"}}
TRANSFORM_JSON(json, {"user.name -> user.fullName"}) -> {"user": {"fullName": "Alice"}}
*/

/*
Constraints
	JSON objects have a maximum depth of 10.
	Each JSON object can have up to 1,000,000 keys.
	Streaming operations should handle files up to 1GB in size efficiently.
*/
