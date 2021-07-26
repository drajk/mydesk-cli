# MyDesk CLI

Command line app to search users and tickets <br/>

## Prerequisites

- [Go 1.16.6](https://golang.org)
- [GVM - Go Version Manager](https://github.com/moovweb/gvm) - Optional

## Common commands

These commands will require [CNP Secure Pipeline](https://github.latitudefinancial.com/Latitude/cnp-secure-pipeline) to be installed.

Run unit tests - `make test`

Format the code - `make format`

Lint- `make lint`

Build the service - `make build`

## Project Structure

- [internal](./internal) - internal go packages that are specific to this service and are non exportable.
- [internal/config](./internal/config) - cli setup related functions, constants and helpers.
- [internal/commands](./internal/commands) - commands and handlers.
- [internal/stores](./internal/stores) - user and ticket stores, read from local JSON files and can be extended.
- [pkg](./pkg) - exportable/sharable packages that the service exposes (if needed).
- [mocks](./mocks) - mock files required for cli or unit tests

## Sample CLI Usage

To be used with `mydesk-cli` once built or `go run main.go` 

- `user --id 3`
- `user --id 3`
- `user --name "Francisca Rasmussen"`
- `ticket --id "2614576f-98fb-4031-9e13-beca7a6a73ee"`
- `ticket --type "question"`
- `ticket --tag "Wisconsin"`

## Optional - Debug using Visual Studio Code

1. Install `Go` extention.
2. Open `Command Palette` > select `Go: Install/Update Tools` > select `dlv` and click on `OK`.
3. Open `Command Palette` > select `Debug: Open launch.json`, and use this configuration to get started.
4. Thats it. You can now add a breakpoint and start debugging.

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug",
      "type": "go",
      "request": "launch",
      "port": 8080,
      "mode": "debug",
      "program": "${workspaceFolder}/main.go",
      "env": {},
      "args": []
    }
  ]
}
```
