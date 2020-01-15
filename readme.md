# trusted-traveler-cli ![](https://github.com/bendrucker/trusted-traveler-cli/workflows/Go/badge.svg?branch=master)

> CLI for interacting with the [Trusted Traveler Program](https://ttp.dhs.gov) API

## Usage

Find available locations:

```sh
trusted-traveler locations
```

Find available slots in a specified location, polling if none are available:

```sh
trusted-traveler slots --location-id <id> --wait
```

## License
The MIT License (MIT)

MIT © [Ben Drucker](https://www.bendrucker.me)
