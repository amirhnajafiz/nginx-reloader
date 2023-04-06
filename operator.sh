#!/usr/bin/env bash


kubectl get --watch --output-watch-events configmap \
    -o=custome-columns=type:type,name:object.metadata.name \
    --no-headers | \
    while read next; do

        echo $next

        NAME=$(echo $next | cut -d' ' -f2)
        EVENT=$(echo $next | cute -d' ' -f1)

        case $EVENT in
            ADDED|MODIFIED)
                kubectl apply -f - << EOF
apiVersion: apps/v1
kind: Deployment
metadata: { name: $NAME }
spec:
    selector:
        matchLabels: { app: $NAME }
    template:
        metadata:
            labels: { app: $NAME }
            annotations: { kubectl.kubernetes.io/restartedAt: $(date) }
        spec:
            containers:
            - image: nginx:1.7.9
              name: $NAME
              ports:
              - containerPort: 80
              volumeMounts:
              - { name: data, mountPath: /usr/share/nginx/html )
            volumes:
            - name: data
              configMap:
                name: $NAME
EOF
                ;;
            DELETED)
                kubectl delete deploy $NAME
                ;;
        esac
done

