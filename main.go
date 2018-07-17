package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ini/ini"
	"github.com/jessevdk/go-flags"
	ghd "github.com/mitchellh/go-homedir"
)

var opts struct {
	Positional struct {
		Profile string
	} `positional-args:"yes" required:"yes"`
	Verbose  bool   `short:"v" long:"verbose" description:"Verbose output"`
	Filename string `short:"f" long:"filename" description:"Credentials file" default:"~/.aws/credentials"`
}

func main() {

	// parse args
	_, err := flags.Parse(&opts)
	if err != nil {
		typ := err.(*flags.Error).Type
		if typ == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatalln(err)
	}

	// load from profile
	credentialsPath, err := ghd.Expand(opts.Filename)
	if err != nil {
		log.Fatalln(err)
	}
	cfg, err := ini.Load(credentialsPath)
	if err != nil {
		log.Fatalln(err)
	}
	section, err := cfg.GetSection(opts.Positional.Profile)
	if err != nil {
		log.Fatalln(err)
	}

	key := section.Key("aws_access_key_id").String()        // AWS_ACCESS_KEY_ID
	secret := section.Key("aws_secret_access_key").String() // AWS_SECRET_ACCESS_KEY
	keyname := section.Key("aws_keyname").String()          // AWS_KEYNAME
	kp := section.Key("aws_keypath").String()               // AWS_KEYPATH
	keyPath, err := ghd.Expand(kp)
	if err != nil {
		log.Fatalln(err)
	}

	// output
	out := ""
	out += fmt.Sprintf("export AWS_ACCESS_KEY_ID='%s'; ", key)
	out += fmt.Sprintf("export AWS_SECRET_ACCESS_KEY='%s'; ", secret)
	if len(keyname) > 0 {
		out += fmt.Sprintf("export AWS_KEYNAME='%s'; ", keyname)
	} else {
		out += "unset AWS_KEYNAME; "
	}
	if len(keyPath) > 0 {
		out += fmt.Sprintf("export AWS_KEYPATH='%s'; ", keyPath)
	} else {
		out += "unset AWS_KEYPATH; "
	}

	// verbose?
	if opts.Verbose {
		out += "env | grep AWS; "
	}

	fmt.Println(out)
}
