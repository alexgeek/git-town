# propose-no-web

This setting configures whether `git town propose` opens a browser after
creating a proposal.

## options

When set to `false` (the default value), `git town propose` opens a browser
after creating the proposal. When set to `true`, `git town propose` does not
open the browser. For CLI-based connectors (`gh`, `glab`), the proposal is still
created on the command line but the browser is not opened.

## via CLI flag

The [propose](../commands/propose.md) command has a `--no-web` CLI flag to
disable opening the browser for that invocation.

## in config file

The [config file](../configuration-file.md) can disable opening the browser
permanently like this:

```toml
[propose]
no-web = true
```
