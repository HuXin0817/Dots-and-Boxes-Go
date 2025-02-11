#!/bin/bash

rm -rf "/Applications/Dots and Boxes.app"
fyne package -os darwin
mv "Dots and Boxes.app" "/Applications/Dots and Boxes.app"
