#!/bin/sh
# Eliminando datos guardados del usb
rm -rf datos
#    eliminando registro log.txt
rm log.txt
#    copiando y generando registro
rsync -av --progress /media/usb datos | tee -a log.txt
