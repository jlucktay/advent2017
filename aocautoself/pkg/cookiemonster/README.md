# About `cookiemonster`

Originally from [here](https://gist.github.com/dacort/bd6a5116224c594b14db).

## Execution

This will give you the session cookie for adventofcode.com:

`go run cookiemonster.go adventofcode.com | grep "\/session:" | cut -d' ' -f 2`

Note that the cookies examined are those of the **Default** profile in Chrome.

If you have not already granted this app access, you will go through a Keychain popup to supply your password.

The password is necessary to decrypt the cookies.
