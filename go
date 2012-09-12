#!/bin/bash

cd src/chrome
#\dev\jslint\jsl conf \dev\jslint\jsl.default.conf +recurse -process *.js > c:\tmp\jswarnings.txt
#if %ERRORLEVEL% GTR 0 (
#   call epm c:\tmp\jswarnings.txt
#   echo there were warnings
#)

#if %ERRORLEVEL% GTR 1 (
#   echo Errors - aborting
#   cd ..
#   goto :EOF
#)

#7z a -xr!CVS -tzip -r tamperdata.jar *
zip -r tamperdata.jar *
cd ..
cp chrome/tamperdata.jar .
rm chrome/tamperdata.jar
#7z a -xr!CVS -tzip tamperdata.xpi install.rdf tamperdata.jar chrome.manifest defaults 
zip tamperdata.xpi install.rdf tamperdata.jar chrome.manifest defaults 
cd ..

