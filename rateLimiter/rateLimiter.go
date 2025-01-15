package ratelimiter

/*
Design a Rate Limiter
*/

/*
Level 1 – Initial Design & Basic Functions
Implement the following methods for a basic rate limiter:

SET_LIMIT(userID string, maxRequests int, timeWindow int):
	Set the rate limit for a user.
	maxRequests specifies the maximum number of requests allowed.
	timeWindow is the time frame (in seconds) for which the limit applies.

ALLOW(userID string, timestamp int) bool:
	Check if a user’s request at a given timestamp is allowed.
	Return true if the request is allowed, false otherwise.

SET_LIMIT("user1", 5, 60)
ALLOW("user1", 1) -> true
ALLOW("user1", 2) -> true
...
ALLOW("user1", 61) -> true

*/

/*
Level 2 – Data Structures & Data Processing
Enhance the rate limiter to track detailed request logs and allow analytics:

GET_REMAINING_QUOTA(userID string, timestamp int) int:
	Return the remaining request quota for the user at the specified timestamp.

USER_STATS(userID string) map[string]int:
	Return statistics for the user, including:
	Total requests made.
	Requests allowed.
	Requests denied.

SET_LIMIT("user1", 5, 60)
ALLOW("user1", 1) -> true
ALLOW("user1", 61) -> true
GET_REMAINING_QUOTA("user1", 61) -> 4
USER_STATS("user1") -> {"total": 2, "allowed": 2, "denied": 0}
*/

/*
Level 3 – Refactoring & Encapsulation
Extend the functionality to support multiple APIs with independent rate limits:

SET_API_LIMIT(userID string, apiEndpoint string, maxRequests int, timeWindow int):
	Set a rate limit for a specific API endpoint for a user.

ALLOW_API(userID string, apiEndpoint string, timestamp int) bool:
	Check if a user’s request to a specific API endpoint is allowed.

GET_API_STATS(userID string, apiEndpoint string):
	Return statistics for the user for a specific API endpoint.

SET_API_LIMIT("user1", "/login", 5, 60)
ALLOW_API("user1", "/login", 1) -> true
GET_API_STATS("user1", "/login") -> {"total": 1, "allowed": 1, "denied": 0}

*/

/*
Level 4 – Extending Design & Functionality
Add advanced features such as burst handling, dynamic limits, and rollbacks:

ENABLE_BURST(userID string, maxBurst int):
Allow users to make up to maxBurst extra requests in a burst period (e.g., a few seconds).

DYNAMIC_LIMIT(userID string, currentLoad int):
Adjust a user’s rate limit dynamically based on system load (currentLoad).

ROLLBACK(userID string, timestamp int):
	Roll back the request logs for a user to a previous state at the given timestamp.

SET_LIMIT("user1", 5, 60)
ALLOW("user1", 1) -> true
ENABLE_BURST("user1", 2)
ALLOW("user1", 10) -> true (from burst quota)
ROLLBACK("user1", 5)
ALLOW("user1", 6) -> true (burst quota restored)
*/
