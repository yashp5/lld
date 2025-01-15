package keyvaluestore

/*
Level 1 – Initial Design & Basic Functions
Implement a basic in-memory key-value store with essential operations:

SET(key string, value string):
	Store the given value associated with the key.
	If the key already exists, overwrite the value.

GET(key string) string:
	Retrieve the value associated with the key.
	If the key does not exist, return "NOT_FOUND".

DELETE(key string):
	Remove the key-value pair from the store.
	If the key does not exist, do nothing.
*/

/*
Level 2 – Data Structures & Data Processing
Enhance the key-value store with advanced data operations:

EXISTS(key string) bool:
	Check if a key exists in the store.

SET_WITH_TTL(key string, value string, ttl int):
	Store the key-value pair with a time-to-live (TTL) in seconds.
	The key should expire and be automatically removed from the store after the TTL.

GET_KEYS_BY_PREFIX(prefix string) []string:
	Return a list of keys that start with the given prefix.

SET("user:1", "Alice")
SET("user:2", "Bob")
GET_KEYS_BY_PREFIX("user:") -> ["user:1", "user:2"]
*/

/*
Level 3 – Refactoring & Encapsulation
Add transaction support to the key-value store:

BEGIN():
Start a new transaction.

ROLLBACK():
	Undo all changes made in the current transaction.
	If no transaction is active, do nothing.

COMMIT():
	Save all changes made in the current transaction.
	If no transaction is active, do nothing.

BEGIN()
SET("key1", "value1")
ROLLBACK()
GET("key1") -> "NOT_FOUND"
*/

/*
Level 4 – Extending Design & Functionality
Introduce multi-node support for a distributed key-value store:

SET_REPLICA(nodeID string, key string, value string):
	Store a key-value pair on a replica node for redundancy.

SYNC():
	Synchronize the key-value store with all replica nodes.

SHARD(key string) string:
	Determine which node is responsible for storing a given key (hashing-based sharding).

FAILOVER(nodeID string):
	Handle the failure of a replica node by promoting another node as a primary.

*/

/*
Level 5 – Refactor to Support Concurrency
	Refactor the key-value store to handle multiple concurrent operations safely and efficiently.

New Requirements
	Thread Safety:
		Ensure all operations (SET, GET, DELETE, etc.) are thread-safe.
		Use appropriate synchronization mechanisms (e.g., mutexes, read-write locks) to prevent race conditions.
		Fine-Grained Locking:
		Use fine-grained locking (lock at the key level) to allow higher concurrency instead of locking the entire store.

	Atomic Operations:
		Ensure atomicity for operations like INCREMENT or multi-key updates in transactions.
		Use atomic primitives (e.g., sync/atomic in Go) where applicable.

	Reader-Writer Locks:
		Optimize for scenarios where read operations dominate. Use a read-write lock to allow multiple readers but ensure exclusive access for writers.

	Deadlock Prevention:
		Ensure that operations requiring locks on multiple keys (e.g., transactions) are deadlock-free. Use strategies like:
		Ordering keys for lock acquisition.
		Timeouts for acquiring locks.

Modified Methods
	SET(key string, value string):
	Ensure thread-safe writes by locking only the specific key.

	GET(key string):
	Use a read lock to allow multiple concurrent readers for the same key.

	DELETE(key string):
	Ensure thread-safe deletion by locking the specific key during the operation.

	BEGIN_TRANSACTION():
	Start a transaction and acquire locks on all keys involved.
	Ensure that no other thread can modify the keys until the transaction is committed or rolled back.

	ROLLBACK_TRANSACTION():
	Release all locks acquired during the transaction.
	Undo changes made during the transaction.

	COMMIT_TRANSACTION():
	Apply all changes atomically and release locks.
*/
