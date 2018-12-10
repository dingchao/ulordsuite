ulordutil
=======

[![Build Status](http://img.shields.io/travis/ulordsuite/ulordutil.svg)](https://travis-ci.org/ulordsuite/ulordutil)
[![Coverage Status](http://img.shields.io/coveralls/ulordsuite/ulordutil.svg)](https://coveralls.io/r/ulordsuite/ulordutil?branch=master)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ulordsuite/ulordutil)

Package ulordutil provides bitcoin-specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.

This package was developed for ulord, an alternative full-node implementation of
bitcoin which is under active development by Conformal.  Although it was
primarily written for ulord, this package has intentionally been designed so it
can be used as a standalone package for any projects needing the functionality
provided.

## Installation and Updating

```bash
$ go get -u github.com/ulordsuite/ulordutil
```

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the ulordsuite developers.  To
verify the signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

Package ulordutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
