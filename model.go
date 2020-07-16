package main

import "time"

type theContent struct {
	// has to be unique
	DirPath  string       `json:"dir_path"`
	Commands []theCommand `json:"commands"`
	/// BeforeCommandExecuteOnChanges
	// in seconds
	Delay       time.Duration `json:"delay"`
	IgnoreFiles []string      `json:"ignore_files"`
}

type theCommand struct {
	Command     string `json:"command"`
	CommandArgs string `json:"command_args"`
}
