# CloudQuery Fincode Source Plugin

[![test](https://github.com/dataqueen-center/cq-source-fincode/actions/workflows/test.yml/badge.svg)](https://github.com/dataqueen-center/cq-source-fincode/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dataqueen-center/cq-source-fincode)](https://goreportcard.com/report/github.com/dataqueen-center/cq-source-fincode)

A [fincode](https://fincode.com/) source plugin for CloudQuery that loads data from the [fincode API](https://fincode.com/docs/api/) to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/docs/quickstart), such as PostgreSQL, BigQuery, Athena, and many more.

## Supported Resources

For a full list of supported resources, see [the tables documentation](./docs/tables/README.md).

## Configuration

The following source configuration file will sync [supported data points](./docs/tables/README.md) to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "fincode"
  path: "dataqueen-center/fincode"
  version: "v1.1.0"
  destinations: [postgresql]
  spec:
    # plugin spec section
    client_id: ${fincode_CLIENT_ID}
    secret: ${fincode_SECRET}
    access_token: ${fincode_ACCESS_TOKEN}
    environment: sandbox
```

### Plugin Spec

- `api_key` (string, required):

  A fincode secret from your dashboard.

- `environment` (string, optional):

  The fincode environment to use. ["test", "live"] Defaults to `test`.

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```
### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub  

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
