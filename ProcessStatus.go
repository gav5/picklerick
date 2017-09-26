package main

type ProcessStatus int

const (
	procNew ProcessStatus = iota
	procReady
	procBlocked
	procRunning
)
