# Auto sync(aka commit) on changes (in a directory) to your git repo.

You can run other commands too!

1. go get -u github.com/thejini3/git-sync
2. Create a file called `.config.git-sync.json` in your $HOME directory
3. Keep running in background
    - `$ nohup git-sync &>/dev/null &`

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
