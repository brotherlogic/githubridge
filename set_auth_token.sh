kubectl delete secret ghb-auth -n githubridge
kubectl create secret generic ghb-auth --from-literal ghb_auth_token=$1 -n githubridge
