# Lungo usage example
Detailed description and step-by-step guide are available at [lungo.dev/docs](https://lungo.dev/docs)

## Intro:
You will need the docker or podman to run it on your machine.

The provided example use MongoDB inside a container, configured in docker-compose.yml and populated with mockData in scripts/mongodb/mongodb-init.js.

## Usage:
- Clone the repo
```
git clone git@github.com:golungo/example.git
```
- Start the Database
```
make db
```
- Run the App
```
make start
```