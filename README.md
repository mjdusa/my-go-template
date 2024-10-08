# my-go-template

## Description

- TODO: Description of project goes here.

## Customizing Template after coping

- TODO: Replace 'mjdusa' with your GitHub org or user
- TODO: Replace 'my-go-template' with your app name
- TODO: Rename the cmd/my-go-template directory to your app name

## Contributing

Please see our [Contributing](./CONTRIBUTING.md) for how to contribute to the project.

## Setting up for development

1. Clone repository

```bash
git clone <URL>
```

2. Setup Pre-commit Hooks
   When you clone this repository to your workstation, make sure to install the [pre-commit](https://pre-commit.com/) hooks. [GitHub Repository](https://github.com/pre-commit/pre-commit)

- Installing tools

```bash
brew install pre-commit
```

- Check installed versions.

```bash
$ pre-commit --version
pre-commit 3.3.2
```

- Update configured pre-commit plugins. Updates repository versions in .pre-commit-config.yaml to the latest.

```bash
pre-commit autoupdate
```

- Install pre-commit into the local Git config.

```bash
pre-commit install --install-hooks
```

- Run pre-commit checks manually.

```bash
pre-commit run --all-files
```

## Running

```bash
make release
...
./dist/my-go-template
```

## Maintaining, Housekeeping, Greenkeeping, etc

### Upgrade Go Version

```bash
go mod edit -go=<go_version> && go mod tidy
```

### Upgrade Dependency Versions

```bash
go get -u && go mod tidy
```

### Running GitHub Super-Linter Locally

```bash
docker run --rm -e RUN_LOCAL=true --env-file ".github/super-linter.env" -v $PWD:/tmp/lint github/super-linter:latest
```

### Running golangci-lint Locally

```bash
golangci-lint run --config .github/linters/.golangci.yml --issues-exit-code 0 --out-format=checkstyle
```
