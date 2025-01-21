package filestorage

import (
	"fmt"
	"slices"
	"strings"
)

type file struct {
	Name         string
	Size         int
	CreationTime int
	TTL          int
}

func (f *file) isAlive(at int) bool {
	if f.TTL <= 0 {
		return true
	}
	return f.CreationTime+f.TTL > at
}

type operationType int

const (
	opUpload operationType = iota
	opCopy
	opDelete
)

type operation struct {
	Timestamp   int
	Type        operationType
	FileName    string
	OldFile     *file
	NewFile     *file
	AuxFileName string
	AuxOldFile  *file
	AuxNewFile  *file
}

// fileStorage is our in-memory storage system with rollback functionality
type fileStorage struct {
	files   map[string]*file // currrent state
	history []operation      // log of all opeartions for rollback
}

func newFileStorage() *fileStorage {
	return &fileStorage{
		files:   make(map[string]*file),
		history: make([]operation, 0),
	}
}

// uploads a file at a given timestamp, optionally with ttl
func (fs *fileStorage) fileUploadAt(timestamp int, fileName string, fileSize int, ttl int) error {
	if existing, ok := fs.files[fileName]; ok && existing.isAlive(timestamp) {
		return fmt.Errorf("file %s already exists (alive at %d)", fileName, timestamp)
	}

	newFile := &file{
		Name:         fileName,
		Size:         fileSize,
		CreationTime: timestamp,
		TTL:          ttl,
	}

	// Record old state in operation (if the file existed at all, alive or not)
	var oldFile *file
	if existing, found := fs.files[fileName]; found {
		// We keep the oldFile for rollback, even if it was expired
		temp := *existing
		oldFile = &temp
	}

	fs.files[fileName] = newFile

	// Log operation for rollback
	op := operation{
		Timestamp: timestamp,
		Type:      opUpload,
		FileName:  fileName,
		OldFile:   oldFile,
		NewFile:   newFile,
	}
	fs.history = append(fs.history, op)

	return nil
}

func (fs *fileStorage) filetGetAt(timestamp int, fileName string) int {
	f, exists := fs.files[fileName]
	if !exists || !f.isAlive(timestamp) {
		return 0
	}
	return f.Size
}

func (fs *fileStorage) fileDeleteAt(timestamp int, fileName string) error {
	f, exists := fs.files[fileName]
	if !exists || !f.isAlive(timestamp) {
		return fmt.Errorf("file %s does not exist or expired at %d", fileName, timestamp)
	}

	// Record delete op in history
	op := operation{
		Timestamp: timestamp,
		Type:      opDelete,
		FileName:  fileName,
		OldFile:   f,
	}

	delete(fs.files, fileName)

	fs.history = append(fs.history, op)

	return nil
}

func (fs *fileStorage) fileCopyAt(timestamp int, sourceFile string, destFile string) error {
	src, ok := fs.files[sourceFile]
	if !ok || src.isAlive(timestamp) {
		return fmt.Errorf("source file %s does not exist or expired at %d", sourceFile, timestamp)
	}

	// Save old state for destination (in case it existed)
	var oldDest *file
	if existing, found := fs.files[destFile]; found {
		tmp := *existing
		oldDest = &tmp
	}

	newCopy := &file{
		Name:         destFile,
		Size:         src.Size,
		CreationTime: src.CreationTime,
		TTL:          src.TTL,
	}
	fs.files[destFile] = newCopy

	op := operation{
		Timestamp:   timestamp,
		Type:        opCopy,
		FileName:    destFile,
		OldFile:     oldDest,
		NewFile:     newCopy,
		AuxFileName: sourceFile,
	}
	fs.history = append(fs.history, op)

	return nil
}

func (fs *fileStorage) fileSearchAt(timestamp int, prefix string) []string {
	var candidates []file
	for _, f := range fs.files {
		if f.isAlive(timestamp) && strings.HasPrefix(f.Name, prefix) {
			candidates = append(candidates, *f)
		}
	}

	// Sort by size desc, then name asc
	slices.SortFunc(candidates, func(i, j file) int {
		if i.Size != j.Size {
			return j.Size - i.Size
		}
		return strings.Compare(i.Name, j.Name)
	})

	// Return top 10 names
	limit := 10
	if len(candidates) < 10 {
		limit = len(candidates)
	}

	result := []string{}
	for i := 0; i < limit; i++ {
		result = append(result, candidates[i].Name)
	}
	return result
}

// rollback reverts the state of the storage system to the specified timestamp
// We look at the history (which is appended in chronological order)
// For any opeartion with timestamp > target, we revert it in reverse order
func (fs *fileStorage) rollback(target int) error {
	for i := len(fs.history) - 1; i >= 0; i-- {
		op := fs.history[i]
		if op.Timestamp <= target {
			break
		}

		// Revert
		switch op.Type {
		case opUpload:
			// An upload replaced oldFile with newFile in fs.files
			// So we restore oldFile (which could be nil if it didn't exist)
			if op.OldFile == nil {
				delete(fs.files, op.FileName)
			} else {
				// Restore the old file
				oldCopy := *op.OldFile
				fs.files[op.FileName] = &oldCopy
			}
		case opCopy:
			// A copy replaced oldDest with newDest
			if op.OldFile == nil {
				// The dest file didn't exist previously, so remove it
				delete(fs.files, op.FileName)
			} else {
				oldCopy := *op.OldFile
				fs.files[op.FileName] = &oldCopy
			}
		case opDelete:
			// A file was deleted, restore it
			fs.files[op.FileName] = op.OldFile
		}

		// Remove this operation from history
		fs.history = fs.history[:i]
	}
	return nil
}
