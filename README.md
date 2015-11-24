# awsenv

AWS environment config loader.

__awsenv__ is a small binary that loads AWS environment variables for an
AWS profile from __~/.aws/credentials__ - useful if you're regularly
switching credentials and using tools like
[Vagrant](https://www.vagrantup.com/). In addition to
`aws_access_key_id` and `aws_secret_access_key`, it will also
optionally load settings for `aws_keyname` and `aws_keypath`.

# installation

If you have Go installed, you can just do:

```shell
go get -u github.com/soniah/awsenv
```

This will automatically download, compile and install the app; putting
an `awsenv` executable in your `$GOPATH/bin`.

Otherwise, download a binary from the [releases
page](https://github.com/soniah/awsenv/releases/latest) and place it in your
$PATH.

# usage

Import variables into your environment by **eval**-ing a
backticked call to **awsenv**.

```shell
eval `awsenv profile-name`
```

For example, if you had the following credential files:

```shell
% cat ~/.aws/credentials
[example1]
aws_access_key_id = DEADBEEFDEADBEEF
aws_secret_access_key = DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF

% cat /var/tmp/credentials
[example2]
aws_access_key_id = DEADBEEFDEADBEEF
aws_secret_access_key = DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF
aws_keyname = 'example2_key'
aws_keypath = "~/.ssh/example2.pem"
```

The following shell commands would import AWS variables into your
environment:

```shell
% eval `awsenv example1`
% env | grep AWS
AWS_ACCESS_KEY_ID=DEADBEEFDEADBEEF
AWS_SECRET_ACCESS_KEY=DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF

% eval `awsenv example2 -f /var/tmp/credentials -v`
AWS_ACCESS_KEY_ID=DEADBEEFDEADBEEF
AWS_SECRET_ACCESS_KEY=DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF
AWS_KEYNAME=example2_key
AWS_KEYPATH=/Users/sonia/.ssh/example2.pem
```

# flags

The accepted flags can be displayed using `-h`:

```
% awsenv -h
Usage:
  awsenv [OPTIONS] Profile

Application Options:
  -v, --verbose   Verbose output
  -f, --filename= Credentials file (~/.aws/credentials)

Help Options:
  -h, --help      Show this help message

Arguments:
  Profile
```

# build

- edits
- `goxc bump`
- git add . ; git commit -m 'foo msg'
- `git push --tags`
- `goxc -bc='linux,!arm darwin'`

