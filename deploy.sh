#!/bin/bash

if ! [ -f ".goxc.local.json" ] ; then
	echo "'.goxc.local.json' not found"
	echo "Generate a 'Personal Access Token' at 'https://github.com/settings/tokens'"
	echo "Then run 'goxc -wlc default publish-github -apikey=123456789012'"
	echo "Exiting..."
	exit 1
fi

read -p "Have you done 'goxc bump'?"
if [[ $REPLY =~ ^[Yy]$ ]] ; then
	goxc -bc='linux,!arm darwin' -tasks-=deb
fi
