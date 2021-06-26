#!/bin/bash

LNW="The rest of this chapter will be updated very soon on the ligh­tnov­elp­ub.c­om platform."

grep "$LNW" -r ./novel-grabber/lightnovelworld | \
awk -F":" '{print $1}' | \
xargs rm
