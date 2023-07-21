kubectl create deployment centari-weaviate --image=us-west1-docker.pkg.dev/centari-dev/weaviate/centari-weaviate
kubectl scale deployment centari-weaviate --replicas=2
