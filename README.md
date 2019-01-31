Noted actions:
  * help
  * open
  * list
  * delete
  * track
  * sync -private

CLI examples:
$ noted [--help,-h]
  display help

$ noted some_file[.ext]
  sync
  if file doesn't exist:
    create some_file.ext  // ext defaults to .txt
  open file

$ noted list [--no-sync]
  // maybe feature to display local/remote diff
  sync
  display all local files

$ noted del some_file
  delete some_file.* locally
  delete some_file.* remotely

$ noted track /path/to/file.ex
  copy file.ex to local
  store file.ex to remote

sync workflow
  remote files are downloaded
    * remotes with same name overwrite locals
    * locals not in remote are pushed
