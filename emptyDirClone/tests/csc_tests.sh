#!/bin/bash

set -o errexit   # abort on nonzero exitstatus
set -o nounset   # abort on unbound variable
set -o pipefail  # don't hide errors within pipes

socket_path="unix://$PWD/test.sock"
plugin="${1:-emptydirclone} --endpoint $socket_path"
csc="${2:-csc} --endpoint $socket_path"

printf '\nRunning csc tests\n'

printf '\nRunning the CSI plugin in background\n'
$plugin > /dev/null 2>&1 &
sleep 2

printf '\nIdentity service\n'
printf '\n==> csc identity plugin-info\n'
$csc identity plugin-info

kill %1
