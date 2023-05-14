#!/usr/bin/env bash



FILEPATH="export/deployment.yml"



# create export directory on local
if [ ! -d "export" ]; then
    mkdir export
fi



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
                cp deployment.yml "$FILEPATH"

                # replace names
                sed -i -e 's/&NAME/$NAME/g' "$FILEPATH"
                sed -i -e 's/&DATE/$(date)/g' "$FILEPATH"

                kubectl apply -f "$FILEPATH"
                ;;
            DELETED)
                kubectl delete deploy $NAME
                ;;
        esac
done
