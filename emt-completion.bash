#/usr/bin/env bash
WORDLIST=`emt -completion-list`
complete -W "$WORDLIST" emt
