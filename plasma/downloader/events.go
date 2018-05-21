package downloader

type PlsDoneEvent struct{}
type PlsStartEvent struct{}
type PlsFailedEvent struct{ Err error }
