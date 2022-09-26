# Zepben license checker
This is a simple program for checking for AGPL/MPL/MIT/Zepben licenses in files.

# Building
Requires golang (tested with 1.14)

    go build -o licence-checker ./...

# Usage

Takes a single argument: The path of the file to check

    ./licence-checker <filepath>

Returns 0 on success and -1 if neither the AGPL license or header snippet did not achieve at least an 80% match.
Should be used on either source files with license headers or COPYING files.
