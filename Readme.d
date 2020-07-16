# Auto sync(aka commit) on changes (in a directory) to your git repo.

You can run other commands too!

1. go get -u github.com/thejini3/git-sync
2. Create a file called `.config.git-sync.json` in your $HOME directory
3. Keep running the command in background

Example `.config.git-sync.json` file
```
[
    {
        "dir_path": "/Users/nix/shell",
        "commands": [
			{"command": "git", "command_args": "add ."},
            {"command": "git", "command_args": "commit -m \"git-sync $(date \"+%Y-%m-%d_%H-%M-%S\")\""},
            {"command": "git", "command_args": "push origin master"}
		],
        "delay": 10,
        "ignore_files": [
            ".DS_Store",
            ".env"
        ]
    }
]
```

- `delay`, in seconds
- `ignore_files`, files at which (on changes) command won't execute