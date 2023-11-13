# avm
Armory Version Manager - A utility to manage the Armory CLI.

# Installation
go to https://github.com/armory/avm/releases/latest download the release for your OS and Arch and mark it executable and place it in your path.

# Usage
See the [generated docs](docs/avm.md) or the help output of the CLI

# Release Process

The AVM install script pulls artifacts from S3. Run the following command to kick off the S3 release workflow:

```shell
export RELEASE_TAG=vx.y.z
curl -X POST \
  -H "Authorization: token $(gh auth token)" \
  -H "Accept: application/vnd.github.everest-preview+json" 
  -H "Content-Type: application/json" \
  https://api.github.com/repos/armory-io/armory-cli-releaser/dispatches \
  --data "{\"event_type\": \"ReleaseAvm\", \"client_payload\": {\"release_tag\": \"$RELEASE_TAG\"}}"
```