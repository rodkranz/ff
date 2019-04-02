#!/usr/bin/env bash

# Switchable Keys for script.
trace=0      # Enable trace mode of FF
output=0     # Show all output of FF
profiling=0  # Enable profiling mode of FF

# Help message with instructions with options of script
USAGE_HELP="
Usage: $(basename "$0") [OPTIONS]
OPTIONS:
  -t, --trace        Show Last trace log from FF
  -o, --output       Show Last output log from FF
  -p, --profiling    Show Last profiling log from FF
  -h, --help         Show help
"

# if has no parameters it will show output log by default.
if [[ "$#" == 0 ]]; then
    output=1
fi

# Loop to enable/disable options of script
while test -n "$1"
do
	case "$1" in
		-t | --trace         ) trace=1     ;;
		-o | --output        ) output=1    ;;
		-p | --profiling     ) profiling=1 ;;
		-i | --ignore-default) ;;
		-h | --help          )
			echo "$USAGE_HELP"
			exit 0
		;;
		*)
			echo Invalid Option: $1
			exit 1
		;;
	esac

	# next options of args.
	shift
done

# Functions: start options of script/

# Return the last output log  of ff.
getLastOutput() {
    local fileName=$(ls -l ./output/ | tail -1 | cut -d' ' -f13)
    echo "${fileName}"
}

# Return the last trace log of ff.
getLastTrace() {
    local fileName=$(ls -l ./trace/ | tail -1 | cut -d' ' -f13)
    echo "${fileName}"
}

if [[ "$trace" = 1 ]]; then
    go tool trace "./trace/$(getLastTrace)"
fi

if [[ "$output" = 1 ]]; then
    cat "./output/$(getLastOutput)"
fi

if [[ "$profiling" = 1 ]]; then
    echo "profiling mode"
    echo "not implemented yet."
fi
