#!/bin/bash

git submodule init && git submodule update
pushd libsass
make shared
sudo make install-shared
sudo ldconfig
