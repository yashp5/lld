package notificationsystem

/*
Level 1 – Initial Design & Basic Functions
Implement a basic notification system:

SEND_NOTIFICATION(userID string, message string):
	Send a notification to the specified user.
	Store the notification in the user’s inbox.

GET_NOTIFICATIONS(userID string) []string:
	Retrieve all notifications for the specified user in the order they were received (FIFO).

MARK_AS_READ(userID string, notificationID string):
	Mark a specific notification as read.
*/

/*
Level 2 – Data Structures & Data Processing
Enhance the system to support advanced features:

GET_UNREAD_COUNT(userID string) int:
	Return the number of unread notifications for a user.

DELETE_NOTIFICATION(userID string, notificationID string):
	Remove a specific notification from the user’s inbox.

BULK_SEND_NOTIFICATION(userIDs []string, message string):
	Send the same notification to a group of users.

SEND_NOTIFICATION("user1", "Welcome to the platform!")
GET_UNREAD_COUNT("user1") -> 1
MARK_AS_READ("user1", "notif1")
GET_UNREAD_COUNT("user1") -> 0
*/

/*
Level 3 – Refactoring & Encapsulation
Add support for scheduling and filtering notifications:

SCHEDULE_NOTIFICATION(userID string, message string, timestamp int):
	Schedule a notification to be sent at a specific time.

FILTER_NOTIFICATIONS(userID string, criteria map[string]string) []string:
	Filter notifications based on criteria (e.g., type, read/unread, time range).

EXPIRATION(notificationID string, ttl int):
	Automatically expire a notification after its TTL (in seconds).

SCHEDULE_NOTIFICATION("user1", "Your event starts soon!", 1678901234)
FILTER_NOTIFICATIONS("user1", {"read": "false"}) -> ["Your event starts soon!"]
*/

/*
Level 4 – Extending Design & Functionality
Introduce multi-channel support and user preferences:

SET_PREFERENCES(userID string, channel string, enabled bool):
	Allow users to enable or disable specific notification channels (e.g., email, SMS, push).

SEND_CHANNEL_NOTIFICATION(userID string, message string, channel string):
	Send a notification via the specified channel.

TRACK_DELIVERY(notificationID string) map[string]string:
	Track the delivery status of a notification (e.g., delivered, failed).

SET_PREFERENCES("user1", "email", true)
SEND_CHANNEL_NOTIFICATION("user1", "Your order is shipped!", "email")
TRACK_DELIVERY("notif1") -> {"status": "delivered"}
*/

/*
Level 5 – Concurrency and Scalability
Refactor the system to handle high traffic and concurrent requests:

HANDLE_CONCURRENT_NOTIFICATIONS():
	Ensure that multiple notifications can be sent and processed concurrently without conflicts.

LOAD_BALANCE_NOTIFICATIONS():
	Distribute notification processing across multiple servers using a load balancer.

RETRY_FAILED_NOTIFICATIONS(notificationID string):
	Retry sending a failed notification after a backoff period.

BULK_TRACK_DELIVERY(notificationIDs []string) []map[string]string:
	Track delivery statuses for multiple notifications simultaneously.
*/
