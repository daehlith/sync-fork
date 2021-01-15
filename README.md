# sync-fork
> Small utility to keep a fork in sync with its upstream repository

![Build status](https://github.com/daehlith/sync-fork/workflows/Build/badge.svg)

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

Similarly it is also possible to specify the name of the master branch via the `--m` / `--master` switch, should it not be "master":
```
$ sync-fork --m the-real-master-branch-name
```

### Contributing

#### Creating a new release

Simply tag the `master` branch commit to release from. Then push said tag to GitHub. The CI pipeline will pick it up and generate a new release, including changelog generation.
```bash
$ git checkout master
$ git tag 0.3.4
$ git push origin 0.3.4
```

Things to consider:
- ensure we are running it on a git repository
- make sure the assumptions for the fork exist (e.g. remote that we can read from, called upstream)
- what happens in case of merge conflicts? We certainly cannot push, but do we need any sort of interaction at any point?
- how do submodules take part in this?
