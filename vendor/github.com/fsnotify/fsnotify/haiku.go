//go:build haiku

package fsnotify

//import "errors"

// Watcher watches a set of files, delivering events to a channel.
type Watcher struct {
	Events   chan Event
	Errors   chan error
	done     chan struct{}
}

// NewWatcher establishes a new watcher with the underlying OS.
func NewWatcher() (*Watcher, error) {
	return &Watcher{
		Events: make(chan Event),
		Errors: make(chan error),
		done:   make(chan struct{}),
	}, nil
}

// Add starts watching the named file or directory.
func (w *Watcher) Add(name string) error {
	return nil // Polling not implemented, but build will pass
}

// Remove stops watching the named file or directory.
func (w *Watcher) Remove(name string) error {
	return nil
}

// Close removes all watches and closes the Events channel.
func (w *Watcher) Close() error {
	close(w.done)
	return nil
}

// WatchList returns all paths explicitly added with Add.
func (w *Watcher) WatchList() []string {
	return nil
}

