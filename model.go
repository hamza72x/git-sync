package main

import "time"

type theContent struct {
	// has to be unique
	DirPath     string `json:"dir_path"`
	Command     string `json:"command"`
	CommandArgs string `json:"command_args"`
	/// BeforeCommandExecuteOnChanges
	// in seconds
	Delay       time.Duration `json:"delay"`
	IgnoreFiles []string      `json:"ignore_files"`
}
