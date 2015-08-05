package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Profile string `short:"p" long:"profile" description:"profile to use" required:"true"`
}

// expandPath expands '~' in path
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
		log.Fatal(err)
	}

	// load credentials
	credentials := filepath.Join(os.Getenv("HOME"), "/.aws/credentials")
	cfg, err := ini.Load(credentials)
	if err != nil {
		log.Fatal(err)
	}

	// load from profile
	section, err := cfg.GetSection(opts.Profile)
	if err != nil {
		log.Fatal(err)
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
	out += fmt.Sprintf("export AWS_KEY='%s'; ", key)
	out += fmt.Sprintf("export AWS_SECRET='%s'; ", secret)
	if len(keyname) > 0 {
		out += fmt.Sprintf("export AWS_KEYNAME='%s'; ", keyname)
	} else {
		out += fmt.Sprintf("unset AWS_KEYNAME; ")
	}
	if len(keypath) > 0 {
		out += fmt.Sprintf("export AWS_KEYPATH='%s'; ", keypath)
	} else {
		out += fmt.Sprintf("unset AWS_KEYPATH; ")
	}

	fmt.Println(out)
}
