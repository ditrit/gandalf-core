#! /bin/bash

loadENV(){
    local INVENV=$(python3 -c 'import sys; print ("1" if sys.prefix != sys.base_prefix else "0")')

    while [ $INVENV -eq 0 ]; do
        source ./env/bin/activate || python3 -m virtualenv env || pip3 install virtualenv
        INVENV=$(python3 -c 'import sys; print ("1" if sys.prefix != sys.base_prefix else "0")')
    done
}

detectModule(){
    local curdirs=`ls -d */`
    local delete

    for i in "${!curdirs[@]}"; do [[ $i == "dist/" ]] && unset -v 'curdirs[$i]' ; done
    curdirs=("${curdirs[@]}")

    delete=(build/ dist/ env/ *.egg-info/)
    for del in ${delete[@]}
    do
    curdirs=("${curdirs[@]/$del}")
    done
    
    echo ${curdirs[0]::-2}
}

loadENV

pip3 install -r requirements.txt

module="$(detectModule)"

python3 -m $module