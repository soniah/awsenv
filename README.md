# awsenv

AWS environment config loader.

__awsenv__ is a small binary that loads AWS environment variables for an
AWS profile from __~/.aws/credentials__ - useful if you're regularly
switching credentials and using tools like
[Vagrant](https://www.vagrantup.com/). In addition to
`aws_access_key_id` and `aws_secret_access_key`, it will also
optionally load settings for `aws_keyname` and `aws_keypath`.

# usage

Download a binary from the [releases
page](https://github.com/soniah/awsenv/releases/latest) and place it in
your $PATH. Then import variables into your environment by **evaling** a
backticked call to **awsenv**.

```shell
eval `awsenv -p profile-name`
```

# example

For example, if you had the following settings in
`~/.aws/credentials`:

```shell
[example1]
aws_access_key_id = DEADBEEFDEADBEEF
aws_secret_access_key = DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF

[example2]
aws_access_key_id = DEADBEEFDEADBEEF
aws_secret_access_key = DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF
aws_keyname = 'example_key'
aws_keypath = "~/.ssh/example1.pem"
```

The following shell commands would import AWS variables into your
environment:

```shell
eval `awsenv -p example1`

% env | grep AWS
AWS_KEY=DEADBEEFDEADBEEF
AWS_SECRET=DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF

eval `awsenv -p example2`

% env | grep AWS
AWS_KEY=DEADBEEFDEADBEEF
AWS_SECRET=DEADBEEFDEADBEEF1vzfgefDEADBEEFDEADBEEF
AWS_KEYNAME=example_key
AWS_KEYPATH=/Users/sonia/.ssh/example1.pem
```

# development

If you wish to build your own binaries using the
Go build-chain, or for development work:

```shell
go get -u github.com/soniah/awsenv
```
