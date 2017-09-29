package main

// ProcessStatus describes the status of a given process
type ProcessStatus int

const (
	//
	procNew ProcessStatus = iota
	procReady
	procBlocked
	procRunning
)
