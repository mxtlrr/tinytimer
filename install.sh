#!/bin/bash
# Script to both build tinytimer as well as install both the man pages and
# itself

echo [...] Building tinytimer
cd src && bash build.sh && cd ..
echo [!] Done

echo [...] Installing both tinytimer and man page
mkdir -p /usr/share/man/man6
cp ./tinytimer.6 /usr/share/man/man6/tinytimer.6 -v
cp ./tinytimer   /usr/bin/tinytimer -v