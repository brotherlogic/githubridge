kubectl delete secret ghb -n githubridge
kubectl create secret generic ghb --from-literal ghb_token=$1 -n githubridge