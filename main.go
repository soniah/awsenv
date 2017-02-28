package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Positional struct {
		Profile string
	} `positional-args:"yes" required:"yes"`
	Verbose  bool   `short:"v" long:"verbose" description:"Verbose output"`
	Filename string `short:"f" long:"filename" description:"Credentials file" default:"~/.aws/credentials"`
}

// expandPath expands '~' in path
// TODO this will fail when path is '~otheruser'. Investigate:
// http://stackoverflow.com/questions/17609732/expand-tilde-to-home-directory
func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(os.Getenv("HOME"), path[2:])
	}
	return path
}

func main() {

	// parse args
	_, err := flags.Parse(&opts)
	if err != nil {
		typ := err.(*flags.Error).Type
		if typ == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// load from profile
	credentials_path := expandPath(opts.Filename)
	cfg, err := ini.Load(credentials_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	section, err := cfg.GetSection(opts.Positional.Profile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// AWS_KEY
	key := section.Key("aws_access_key_id").String()
	// AWS_SECRET
	secret := section.Key("aws_secret_access_key").String()
	// AWS_KEYNAME
	keyname := section.Key("aws_keyname").String()
	// AWS_KEYPATH
	kp := section.Key("aws_keypath").String()
	keypath := expandPath(kp)

	// output
	out := ""
	out += fmt.Sprintf("export AWS_ACCESS_KEY_ID='%s'; ", key)
	out += fmt.Sprintf("export AWS_SECRET_ACCESS_KEY='%s'; ", secret)
	if len(keyname) > 0 {
		out += fmt.Sprintf("export AWS_KEYNAME='%s'; ", keyname)
	} else {
		out += "unset AWS_KEYNAME; "
	}
	if len(keypath) > 0 {
		out += fmt.Sprintf("export AWS_KEYPATH='%s'; ", keypath)
	} else {
		out += "unset AWS_KEYPATH; "
	}

	// verbose?
	if opts.Verbose {
		out += "env | grep AWS; "
	}

	fmt.Println(out)
}
