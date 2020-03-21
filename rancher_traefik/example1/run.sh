#!/bin/bash

if [ -d /run/secrets/ ]; then
  echo "Loading Secrets into Environments...."
  for file in $(ls /run/secrets/* 2>/dev/null); do
    env_var=$(basename $file | tr '/a-z/' '/A-Z/')
    var_content=$(cat $file)
    export ${env_var}="${var_content}"
  done
fi

# traefik_args=""
while IFS= read -r line
do
  name="${line%%=*}"
  if [[ $name =~ ^TRAEFIK ]]; then
    arg_name=${name/TRAEFIK_/}
    arg_name=$(echo ${arg_name}| tr '[:upper:]' '[:lower:]' | tr '_' '.')
    traefik_args+="--${arg_name}=$(eval echo "\$$name") "
  fi
done < <(env)

# Print out the Confirguration/arguments prior to running
[[ ! -z $DEBUG ]] && echo /traefik ${traefik_args} ${@}
/traefik ${traefik_args} ${@}
