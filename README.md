# Noted

Distributed notes using your favorite editor

## Features

- Create/open local files
- Delete local files
- List local files

#### TODO

- opening files:
  - handle remote storage
  - watch file changes and update remote file
- deleting files:
  - delete from remote storage
- add `list` command:
  - sync local/remote files
- add `track` command:
  - add file to local and remote storage

sync workflow

- remote files are downloaded
- remotes with same name overwrite locals
- locals not in remote are pushed

Refactor

- use `homedir "github.com/mitchellh/go-homedir"` instead of `"os/user"`
- move Config to cmd package
- take config from yaml file
