# AGPL Zepben licence checker
This is a simple program for checking for AGPL licences in files.

# Building
Requires golang (tested with 1.14)

    go build

# Usage

Takes a single argument: The path of the file to check

    ./licence_checker <filepath>

Returns 0 on success and -1 if neither the AGPL licence or header snippet did not achieve at least an 80% match.
Should be used on either source files with licence headers or COPYING files.
