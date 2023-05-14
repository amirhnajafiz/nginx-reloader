#!/usr/bin/env bash


kubectl get --watch --output-watch-events configmap \
    -o=custome-columns=type:type,name:object.metadata.name,app:object.metadata.labels.app \
    --no-headers | \
    while read next; do

        echo $next

        APP=$(echo $next | cut -d' ' -f3)
        NAME=$(echo $next | cut -d' ' -f2)
        EVENT=$(echo $next | cute -d' ' -f1)

        if [ $APP != "nginx" ]
        then
            continue
        fi

        case $EVENT in
            ADDED|MODIFIED)
                # export deployment file into export directory
                cp deployment.yml export/deployment.yml

                # replace names
                sed -i -e 's/&NAME/$NAME/g' export/deployment.yml
                sed -i -e 's/&DATE/$(date)/g' export/deployment.yml

                kubectl apply -f export/deployment.yml
                ;;
            DELETED)
                kubectl delete deploy $NAME
                ;;
        esac
done

