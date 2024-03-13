

kubectl port-forward svc/argocd-server -n argocd 8080:443

kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d

argocd app create nginx-app --repo https://github.com/guichafy/xpto.git --path . --dest-server minikube --dest-namespace default