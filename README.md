# Silly little program to generate random strings

This program generates a random 32 char string, and adds it to your clipboard. Nothing too special about it, but it works like a champ.

# To install and use

* Make sure your $GOPATH is all set up correctly
* `go get ...`
* `go install`
* now, when you run `randomString` your clipboard should have a 32 character random string

## Command line flag options
###### Boolean flags can be omitted to use default value

* -len **_Default: 32_**
    - Determines the length of the string.
    - Example: `randomString -len=20`
* -num **_Default: true_**
    - Determines if string should contain number characters.
    - Example: `randomString -num=false`
* -spec **_Default: true_**
    - Determine if string should contain special characters.
    - Example: `randomString -spec`

###### Examples

cmd: `randomString` = `TEITvYS1GmntMDcMfh9ITRA55iNhuWZU`
cmd: `randomString -len=12` = `hPosLYMYfe6a`
cmd: `randomString -len=10 -num=false` = `HCmuGQnnsJ`
cmd: `randomString -len=23 -spec` = `Leg8T|0jHGt]!WW)*d%gD)[`
