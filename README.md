A mock Git server that create disposable Git repos for testing.

## Usage

Start the server:

```bash
podman run --name mockgitserver -p 8080:8080 --name git-server mockgitserver
```

Clone the repository:

```bash
cd /tmp
git clone @http://myusername@localhost:8080/platformhubrepo
```

## Adding new sample repos

1. Run `unpacktemplate.sh` to unpack the template repo into `repotemplate`
2. Add a directory in `repotemplate`
3. Run `git init` in the new directory
4. Run `git config http.receivepack true` in the new directory
5. Add template files
6. Run `git add .` and `git commit -m "Add new sample repo"`
7. Run `git config --bool core.bare true`
8. Run `packtemplate.sh` to pack the template repo into `repotemplate.tar.gz`

## Update sample repo

1. Run `unpacktemplate.sh` to unpack the template repo into `repotemplate`
2. Run `git config --bool core.bare false`
3. Make changes to the repo
4. Run `git add .` and `git commit -m "Update sample repo"`
5. Run `git config --bool core.bare true`
6. Run `packtemplate.sh` to pack the template repo into `repotemplate.tar.gz`