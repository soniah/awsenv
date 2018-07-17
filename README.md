# awsenv

AWS environment config loader.

__awsenv__ is a small binary that loads AWS environment variables for an
AWS profile from __~/.aws/credentials__ - useful if you're regularly
switching credentials and using tools like
[Vagrant](https://www.vagrantup.com/). In addition to
`aws_access_key_id` and `aws_secret_access_key`, it will also
optionally load settings for `aws_keyname` and `aws_keypath`.

# Installation

If you have Go installed, you can just do:

```shell
go get -u github.com/soniah/awsenv
```

This will automatically download, compile and install the app; putting
an `awsenv` executable in your `$GOPATH/bin`.

Otherwise, download the [lastest binary](https://github.com/soniah/awsenv/releases/latest) from the releases
page and place it in your $PATH (or on Debian-like machines you can just do `sudo dpkg -i downloaded.deb`).

# Usage

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
# Vagrant Example

In a **Vagrantfile** you could do:

```ruby
override.ssh.username = "ubuntu"                                                
aws.keypair_name = ENV['AWS_KEYNAME']                                           
override.ssh.private_key_path = ENV['AWS_KEYPATH']
```
# Flags

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

# Contributions

Contributions are welcome; here is an example workflow using [hub](https://github.com/github/hub).

1. `go get github.com/soniah/awsenv`
1. `cd $GOPATH/src/github.com/soniah/awsenv`
1. `hub fork`
1. `git co -b dev` (and write some code)
1. `git push -u <your-github-username> dev`
1. `hub pull-request`

# Build

- edits
- `goxc bump`
- `git add . ; git commit -m 'foo msg'`
- `git push ; git push --tags`
- `goxc -bc='linux,!arm darwin'`

Also, copy `.goxc.local.json` to new build machines, or do:

`goxc -wlc default publish-github -apikey=DEADBEEFDEADBEEF`
