## ./pachctl set-branch

DEPRECATED Set a commit and its ancestors to a branch

### Synopsis


DEPRECATED Set a commit and its ancestors to a branch.

Examples:

```sh

# Set commit XXX and its ancestors as branch master in repo foo.
$ pachctl set-branch foo XXX master

# Set the head of branch test as branch master in repo foo.
# After running this command, "test" and "master" both point to the
# same commit.
$ pachctl set-branch foo test master
```

```
./pachctl set-branch repo-name commit-id/branch-name new-branch-name
```

### Options inherited from parent commands

```
      --no-metrics           Don't report user metrics for this command
      --no-port-forwarding   Disable implicit port forwarding
  -v, --verbose              Output verbose logs
```

