#!/bin/sh

sleep 2

# elvish

xdotool type --delay 250 "elvish"
xdotool key  --delay 250 Enter

xdotool type --delay 250 "pkill --"
xdotool key  --delay 250 --repeat 5 --repeat-delay 250 Tab
xdotool type --delay 250 "sig"
xdotool key  --delay 250 Enter
xdotool key  --delay 250 --repeat 5 Tab
xdotool key  --delay 250 Escape
xdotool key  --delay 250 Control_L+c

xdotool type --delay 250 "exit"
xdotool key  --delay 250 Enter

# fish

xdotool type --delay 250 "fish"
xdotool key  --delay 250 Enter

xdotool type --delay 250 "pkill --"
xdotool key  --delay 250 --repeat 2 --repeat-delay 250 Tab
xdotool type --delay 250 "sig"
xdotool key  --delay 250 Tab
xdotool key  --delay 250 --repeat 5 Tab
xdotool key  --delay 250 Escape
xdotool key  --delay 250 Control_L+c

xdotool type --delay 250 "exit"
xdotool key  --delay 250 Enter

# xonsh

xdotool type --delay 250 "xonsh"
xdotool key  --delay 250 Enter

xdotool type --delay 250 "pkill --"
xdotool key  --delay 250 --repeat 2 --repeat-delay 250 Tab
xdotool key  --delay 250 Escape
xdotool type --delay 250 "sig"
xdotool key  --delay 250 Tab
xdotool key  --delay 250 --repeat 5 Tab
xdotool key  --delay 250 Escape
xdotool key  --delay 250 Control_L+c

xdotool type --delay 250 "exit"
xdotool key  --delay 250 Enter

sleep 2
