#!/bin/bash

echo '// Generated by create-embed.go.sh'
echo
echo 'package fonts'
echo
echo 'import _ "embed"'

find . -type f -iname '*.ttf'| sed 's|^\./||'|while read FONT; do
    VAR_NAME="$(basename "$FONT"|sed -re 's/^x//' -e 's/[^a-zA-Z0-9]/_/g')"
    printf "\n//go:embed %s\nvar %s []byte\n" "$FONT" "$VAR_NAME"
done

echo
