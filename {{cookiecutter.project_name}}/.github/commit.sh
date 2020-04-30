#!/bin/bash

# Validates the commit message according to conventional commits with the angular styleguide
conventional_commit_validator()
{
	declare -a const types=(build ci docs feat fix perf refactor test style)
	readonly const minLength=1
	readonly const maxLength=50
	readonly const unconventionalCommitErrorMsg="Invalid conventional commit message! Valid types are: ${types[@]}. The message length is maximum: $maxLength. Example: 'fix(config): add missing variable'"
	readonly const commitMsg=$(head -1 $1)

	regexp="^([A-Z0-9]+-[0-9]+\s)?(revert:\s)?("

	# Add all commit types to regex string
	for type in "${types[@]}"
	do
		regexp="${regexp}$type|"
	done

	# Support optional scope and add a minimum and maximum length to the regex string
	regexp="${regexp})(\(.+\))?:\s.{$minLength,$maxLength}$"

	if ! (echo $commitMsg | grep -Eq $regexp); then
		echo $unconventionalCommitErrorMsg
		exit 1
	fi
}

# Automatically prepending an issue key retrieved from the start of the current branch name to commit messages.
ticket_prefix()
{
	# check if commit is merge commit or a commit ammend
	if [[ $2 -eq "merge" ]] || [[ $2 -eq "commit" ]]; then
		exit
	fi

	ISSUE_KEY=$(git rev-parse --abbrev-ref HEAD | sed -nr 's,[a-z]+/([A-Z0-9]+-[0-9]+)-.+,\1,p')

	# only prepare commit message if pattern matched and issue key was found
	if [[ ! -z $ISSUE_KEY ]]; then
		sed -i.bak -e "1s/^/$ISSUE_KEY /" $1
		echo "oh noes"
	fi

}

# Converts a yaml file into a markdown table
yaml_table()
{
	local s='[[:space:]]*' w='[a-zA-Z0-9_]*' fs=$(echo @|tr @ '\034')
	sed -ne "s|^\($s\):|\1|" \
		-e "s|^\($s\)\($w\)$s:$s[\"']\(.*\)[\"']$s\$|\1$fs\2$fs\3|p" \
		-e "s|^\($s\)\($w\)$s:$s\(.*\)$s\$|\1$fs\2$fs\3|p"  $1 |
	awk -F$fs '{
		indent = length($1)/2;
		vname[indent] = $2;
		for (i in vname) {if (i > indent) {delete vname[i]}}
		if (length($3) > 0) {
			vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])(".")}
			
			
			printf("`%s%s`|", vn, $2);
			system("cat ./service/config/config.go | grep -i -B 1 " $2  " | grep // | tr -d \"/\n\t\"")
			printf("|`%s`\n", $3);
		}
	}'
}

# Converts a yaml file into a markdown config table
yaml_table_config()
{
	printf "Parameter|Description|Default\n"
	printf '%s\n' '---|---|---'
	yaml_table $1
}

# Returns all input kafka topics as a markdown table
kafka_input()
{
	printf "Topic Config|Description|Default Topic\n"
	printf '%s\n' '---|---|---'
	yaml_table $1 | grep "inputTopics" | sed 's/kafka.inputTopics.//g'
}

# Returns all output kafka topics as a markdown table
kafka_output()
{
	printf "Topic Config|Description|Default Topic\n"
	printf '%s\n' '---|---|---'
	yaml_table $1 | grep "outputTopics" | sed 's/kafka.outputTopics.//g'
}

# Inserts an input string($3) between "START $2" and "END $2" into a file($1)
insert_in_file()
{
	awk -vstart="START $2" -vend="END $2" -vdata="$3" 'BEGIN {f=1} index($0, start) {print;print data;f=0} index($0, end) {f=1} f' $1 > tmp && mv tmp $1
}
