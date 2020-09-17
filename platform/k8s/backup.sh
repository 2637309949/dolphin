for ns in $(kubectl get ns --no-headers | cut -d " " -f1)
do
	for n in $(kubectl get -o=name pvc,configmap,serviceaccount,secret,ingress,service,deployment,statefulset,hpa,job,cronjob -n $ns) 
	do
		mkdir -p $(dirname $ns/$n)
		kubectl get -o yaml $n -n $ns > $ns/$n.yaml
	done
done