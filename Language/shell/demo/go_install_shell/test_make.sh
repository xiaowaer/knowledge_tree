set -e

if [ ! -f run.bash ]; then
	echo 'make.bash must be run from $GOROOT/src' 1>&2
	exit 1
fi

if [ "$GOBUILDTIMELOGFILE" != "" ]; then
	echo $(LC_TIME=C date) start make.bash >"$GOBUILDTIMELOGFILE"
fi

# Finally!  Run the build.

verbose=false
vflag=""
if [ "$1" = "-v" ]; then
	verbose=true
	vflag=-v
	shift
fi

goroot_bootstrap_set=${GOROOT_BOOTSTRAP+"true"}
echo   $GOROOT_BOOTSTRAP
if [ -z "$GOROOT_BOOTSTRAP" ]; then
	GOROOT_BOOTSTRAP="$HOME/go1.4"
	for d in sdk/go1.17 go1.17; do
		if [ -d "$HOME/$d" ]; then
			GOROOT_BOOTSTRAP="$HOME/$d"
		fi
	done
fi
export GOROOT_BOOTSTRAP
echo   "GOROOT_BOOTSTRAP:" $GOROOT_BOOTSTRAP
echo   "GOROOT_BOOTSTRAP:/root/go1.4"


export GOROOT="$(cd .. && pwd)"
echo "GOROOT:" $GOROOT
echo "GOROOT: ../  "


#查找go路径
echo $goroot_bootstrap_set

IFS=$'\n'; for go_exe in $(type -ap go); do
	if [ ! -x "$GOROOT_BOOTSTRAP/bin/go" ]; then
		goroot=$(GOROOT='' GOOS='' GOARCH='' "$go_exe" env GOROOT)
		if [ "$goroot" != "$GOROOT" ]; then
			if [ "$goroot_bootstrap_set" = "true" ]; then
				printf 'WARNING: %s does not exist, found %s from env\n' "$GOROOT_BOOTSTRAP/bin/go" "$go_exe" >&2
				printf 'WARNING: set %s as GOROOT_BOOTSTRAP\n' "$goroot" >&2
			fi
			GOROOT_BOOTSTRAP=$goroot
            echo $GOROOT_BOOTSTRAP
		fi
	fi
done; unset IFS

