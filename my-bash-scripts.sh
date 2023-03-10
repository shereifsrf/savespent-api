#!/bin/bash

# ke is a function to get into a pod
# $1 is the pod name pattern
function ke() {
    # $1 check if empty
    if [ -z "$1" ]; then
        echo "pod name is empty"
        return
    fi

    # find the pod with $1 and get into it
    pod=$(kubectl get pods | grep $1 | awk '{print $1}')
    echo "redis pod: $pod"
    kubectl exec --stdin --tty $pod -- /bin/bash
}

# kg is kube get all
# getopts is used to parse the arguments
# if no getopts is passed then get all
# getopts will have no arguments
# -a is to get all,pv,pvc,secrets,configmaps
# -p is to get pods
# -sc is to get services,configmaps
function kg() {
    # parse the arguments
    list=""
    while getopts ":apsc" opt; do
        case $opt in
            a)
                echo "get all,pv,pvc,secrets,configmaps"
                list="all,pv,pvc,secrets,configmaps"
                ;;
            p)
                echo "get pods"
                list="pods"
                ;;
            s)
                echo "get secrets"
                # append list if not empty
                if [ -z "$list" ]; then
                    list="secrets"
                else
                    list="$list,secrets"
                fi
                ;;
            c)
                echo "get configmaps"
                # append list if not empty
                if [ -z "$list" ]; then
                    list="configmaps"
                else
                    list="$list,configmaps"
                fi
                ;;
            \?)
                echo "Invalid option: -$OPTARG" >&2
                return
                ;;
        esac
    done


    if [ -z "$1" ]; then
        echo "get all"
        list="all"
    fi

    # get the list
    kubectl get $list    

    # second time call not working
    # so reset the OPTIND
    OPTIND=1
}

# start pods
# $1 can be mysql or mongo or redis
function kstart() {
    # varialble to store type
    type="mysql"
    hasStorage="true"
    hasSecrets="true"
    hasConfigMaps="true"

    if [ "$1" == "mongo" ]; then
        type="mongo"
    elif [ "$1" == "redis" ]; then
        type="redis"
        hasStorage="false"
    elif [ "$1" == "mysql" ]; then
        type="mysql"
    else
        echo "invalid type"
        return
    fi


    echo "selected: $type"

    if [ $hasStorage == "true" ]; then
        echo "creating storage"
        kubectl apply -f k8/${type}/${type}-storage.yml
    fi

    if [ $hasSecrets == "true" ]; then
        echo "creating secrets"
        kubectl apply -f k8/${type}/${type}-secrets.yml
    fi

    if [ $hasConfigMaps == "true" ]; then
        echo "creating configmaps"
        kubectl apply -f k8/${type}/${type}-config.yml
    fi

    echo "creating pod"
    kubectl apply -f k8/${type}/${type}.yml
}

function kdelete() {
    type=""
    hasStorage="true"
    hasSecrets="true"
    hasConfigMaps="true"

    if [ "$1" == "mongo" ]; then
        type="mongo"
    elif [ "$1" == "redis" ]; then
        type="redis"
        $hasStorage="false"
    elif [ "$1" == "mysql" ]; then
        type="mysql" 
    else
        echo "invalid type"
        return
    fi

    echo "selected: $type"

    if [ $hasStorage == "true" ]; then
        echo "deleting storage"
        kubectl delete -f k8/${type}/${type}-storage.yml
    fi

    if [ $hasSecrets == "true" ]; then
        echo "deleting secrets"
        kubectl delete -f k8/${type}/${type}-secrets.yml
    fi

    if [ $hasConfigMaps == "true" ]; then
        echo "deleting configmaps"
        kubectl delete -f k8/${type}/${type}-config.yml
    fi

    echo "deleting pod"
    kubectl delete -f k8/${type}/${type}.yml

    # reset the OPTIND
    OPTIND=1
}

# authenticate redis cli
function redis-cli-auth() {
    redis-cli -a $REDIS_PASSWORD
}

function krecreate() {
    kdelete $1
    kstart $1

    # wait until pod is created
    # find if the pod is running
    # if not running then wait for 1 second
    # if running then get into the pod
    # check if $2 is empty
    echo "waiting for pod to be created"
    while [ -z "$2" ]; do
        status=$(kubectl get pods | grep $1 | awk '{print $3}')
        if [ "$status" == "Running" ]; then
            echo "pod is running"
            ke $1
            break
        fi
        sleep 1
    done
}