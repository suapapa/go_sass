#!/bin/bash

git submodule update
pushd libsass
make shared
sudo make install-shared
sudo ldconfig
