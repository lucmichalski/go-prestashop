_ps2sift()
{
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    if [[ ${cur} == -* ]] ; then
		opts="$(ps2sift --help | ps2sift '\s+(--[\w-]+)(?:\s|=)' --replace '$1' --output-sep ' ')"
		COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        return 0
	fi

	case "${prev}" in
		-t|-T|--type|--no-type|--del-type)
			types="$(sps2ift --list-types | ps2sift '^(\w+)\s+:' --replace '$1' --output-sep ' ')"
			COMPREPLY=( $(compgen -W "${types}" -- ${cur}) )
			return 0;;
	esac

	_filedir
	return 0
}
complete -F _ps2sift ps2sift
