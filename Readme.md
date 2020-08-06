# Auto sync(aka commit) on changes (in a directory) to your git repo.

You can run other commands too!

1. go get -u github.com/thejini3/git-sync
2. Create a file called `.config.git-sync.json` in your $HOME directory
3. Keep running in background - check bottom section (LaunchAgents)

Example `.config.git-sync.json` file

```json
[
    {
        "dir_path": "/Users/nix/shell",
        "commands": [
            {
                "command": "git",
                "args": ["add", "."]
            },
            {
                "command": "git",
                "args": ["commit", "-m", "[auto] git-sync $CURRENT_TIME$"]
            },
            {
                "command": "git",
                "args": ["push", "origin", "master"]
            }
        ],
        "delay": 5,
        "ignore_files": [
            ".DS_Store",
            ".env"
        ]
    }
]
```

- `delay`, in seconds
- `ignore_files`, files at which (on changes) command won't execute
- special args `$CURRENT_TIME$` = `2006-01-02 15:04:05.999999999 -0700 MST`, `$RANDOM$` = `random 10 characters`,

# Launch Agents

For mac: put `git.sync.runner.plist` in `$HOME/Library/LaunchAgents` directory

For linux: put `git-sync.service` in `/lib/systemd/system` directory
	- sudo systemctl start git-sync
	- sudo systemctl enable git-sync (for auto start at boot)

:: Make sure you made your path/directory changes to your launch agents ::


