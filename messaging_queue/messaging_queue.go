package messagingqueue

/*
Level 1 – Initial Design & Basic Functions
Implement a basic messaging queue to handle producers and consumers:

PUBLISH(queueName string, message string):
Add a message to the specified queue.
If the queue does not exist, create it.

SUBSCRIBE(queueName string) string:
Fetch the next message from the queue in the order it was added (FIFO).
If the queue is empty or does not exist, return "EMPTY".

DELETE_QUEUE(queueName string):
Remove all messages and delete the specified queue.
*/

/*
Level 2 – Data Structures & Data Processing
Enhance the messaging queue with advanced features:

PEEK(queueName string) string:
	View the next message in the queue without removing it.
	If the queue is empty or does not exist, return "EMPTY".

LIST_QUEUES() []string:
	Return a list of all existing queues sorted lexicographically.

MESSAGE_COUNT(queueName string) int:
	Return the number of messages in the specified queue.

PUBLISH("queue1", "Message1")
PUBLISH("queue1", "Message2")
MESSAGE_COUNT("queue1") -> 2
PEEK("queue1") -> "Message1"
LIST_QUEUES() -> ["queue1"]
*/

/*
Level 3 – Refactoring & Encapsulation
Introduce features for message acknowledgment and retries:

ACK(queueName string, messageID string):
Acknowledge that a consumer has successfully processed the message.
The message should then be removed from the queue.

RETRY(queueName string, messageID string):
Re-add a message to the end of the queue if processing fails.

GET_UNACKED_MESSAGES(queueName string) []string:
	Return all unacknowledged messages in the queue.

PUBLISH("queue1", "Message1")
SUBSCRIBE("queue1") -> "Message1"
GET_UNACKED_MESSAGES("queue1") -> ["Message1"]
ACK("queue1", "Message1")
GET_UNACKED_MESSAGES("queue1") -> []
*/

/*
Level 4 – Extending Design & Functionality
Add support for distributed systems and delayed processing:

SET_DELAY(queueName string, delay int):
	Add a delay (in seconds) to a queue.
	Messages in this queue should not be visible to consumers until the delay has elapsed.

FETCH_HISTORY(queueName string) []string:
	Return a list of all messages that have been published to the queue (processed or unprocessed).

PURGE_QUEUE(queueName string):
	Remove all messages from the queue without deleting the queue itself.

ROLLBACK(queueName string, timestamp int):
	Roll back the queue to a previous state based on a given timestamp.
	Messages added after the timestamp should be removed.

PUBLISH("queue1", "Message1")
SET_DELAY("queue1", 10)
SUBSCRIBE("queue1") -> "EMPTY" (since delay has not elapsed)
PURGE_QUEUE("queue1")
MESSAGE_COUNT("queue1") -> 0
*/

/*
Constraints
Messages are identified by unique IDs (auto-generated or provided during publishing).
Each queue can store up to 10,000 messages.
Support up to 1,000 queues concurrently.
Delayed messages should not be visible to consumers until their delay has elapsed.
*/
