# Noted

Distributed notes using your favorite editor (Neovim, of course)

## Current Features

- Create/open local/remote files
- Sync files
- Delete local files
- List local files

#### TODO

- `sync`
    - locals not in remote are pushed
- `del` command:
  - delete from remote storage
- `list` command:
    - sync first
    - sort by latest modified first
- add `rename` command:
  - change name of a file locally/remotely
- add `archive` command:
  - tags a file as "archived"
  - "archived" files don't show on `list` by default
  - add option to `list` to show "archived" files
