## ./pachctl auth login

Log in to Pachyderm

### Synopsis


Login to Pachyderm. Any resources that have been restricted to the account you have with your ID provider (e.g. GitHub, Okta) account will subsequently be accessible.

```
./pachctl auth login
```

### Options

```
  -o, --one-time-password   If set, authenticate with a Dash-provided One-Time Password, rather than via GitHub
```

### Options inherited from parent commands

```
      --no-metrics           Don't report user metrics for this command
      --no-port-forwarding   Disable implicit port forwarding
  -v, --verbose              Output verbose logs
```

