This is the back-end part of the G.U.R.P.S. Battle Chest.

G.U.R.P.S. is a role-playing system by Steve Jackson Games: http://www.sjgames.com/gurps/

# Development

## Setup

### Dependencies
You need to manually install a few dependencies:
* Docker
* go
* make
* golang-migrate (you can install with `go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`)

### Secrets
You need to provide your own secrets for the DB access. Create the file `secrets/mysql-creds`. In it, put the MySQL admin password, and the db user password in the following format:
```text
MYSQL_ROOT_PASSWORD=<mysql_admin_pwd>
MYSQL_PASSWORD=<pwd_for_gurps_db_user>
```

## Database
The MySQL migration scripts are located in `internal/database/mysql/migrations`
To add a new migration step, use:
```bash
cd internal/database/mysql/migrations/
migrate create -ext sql -dir mysql -seq create_users_table
```

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