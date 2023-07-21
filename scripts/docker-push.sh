# Check if logged in
if gcloud auth list --format="value(account)" | grep -q centari; then
    # Perform action when logged in
    docker push us-west1-docker.pkg.dev/centari-dev/weaviate/centari-weaviate
    # Add your desired command for when the user is logged in
else
    # Perform action when not logged in
    echo "You are not logged in to gcloud. Please run the following command:"
    gcloud auth login
    # Add your desired command for when the user is not logged in
fi