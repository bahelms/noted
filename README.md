# Noted

Distributed notes using your favorite editor

## Current Features

- Create/open local files
- Delete local files
- List local files

#### TODO

- opening files:
  - handle remote storage
  - watch file changes and update remote file
- `del` command:
  - delete from remote storage
- `list` command:
  - sync local/remote files
- add `track` command:
  - add file to local and remote storage
- add `rename` command:
  - change name of a file locally/remotely
- add `archive` command:
  - tags a file as "archived"
  - "archived" files don't show on `list` by default
  - add option to `list` to show "archived" files

sync workflow

- remote files are downloaded
- remotes with same name overwrite locals
- locals not in remote are pushed

Refactor

- move Config to cmd package
- take config from yaml file
- customize `list` logging
