# sync-fork
> Small utility to keep a fork in sync with its upstream repository

### Usage

```
$ sync-fork
```

By default, sync-fork will push a successful merge to the origin repository. It is possible to avoid this by specifying the `--np` / `--no-push` switch:
```
$ sync-fork --np
```

If you happen to use a remote name for your upstream repository other than "upstream" you can specify this via the `--u` / `--upstream` switch:
```
$ sync-fork --upstream actual-upstream-name
```

### Random bits and pieces

The usual process for syncing a fork involves:

```
$ git fetch upstream
...
$ git checkout master
...
$ git merge upstream/master
...
# and optionally:
$ git push
```

Things to consider:
- ensure we are running it on a git repository
- make sure the assumptions for the fork exist (e.g. remote that we can read from, called upstream)
- what happens in case of merge conflicts? We certainly cannot push, but do we need any sort of interaction at any point?
- how do submodules take part in this?
