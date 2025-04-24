This is the back-end part of the G.U.R.P.S. Battle Chest.

G.U.R.P.S. is a role-playing system by Steve Jackson Games: http://www.sjgames.com/gurps/

# Development

## Dependencies
You need the following tools to work on this repo:
* Docker
* go
* make

## Build
You can build the project (using Docker) by using the following command:
```bash
make build
```

## Testing
To simply run the unit tests, use:
```bash
make test
```

To run the tests with coverage:
```bash
make coverage
```
Then you can examine the `cover.html` file for coverage results.