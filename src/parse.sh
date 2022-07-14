clusterName=ip-10-0-1-22
server=https://10.0.1.22:6443
namespace=jenkins
serviceAccount=jenkins
secretName=$(kubectl --namespace $namespace get serviceAccount $serviceAccount -o jsonpath='{.secrets[0].name}')
ca=$(kubectl --namespace $namespace get secret/$secretName -o jsonpath='{.data.ca\.crt}')
token=$(kubectl --namespace $namespace get secret/$secretName -o jsonpath='{.data.token}' | base64 --decode)
echo "---
apiVersion: v1
kind: Config
clusters:
 - name: ${clusterName}
   cluster:
    certificate-authority-data: ${ca}
    server: ${server}
contexts:
 - name: ${serviceAccount}@${clusterName}
   context:
    cluster: ${clusterName}
    namespace: ${namespace}
    user: ${serviceAccount}
users:
 - name: ${serviceAccount}
  user:
    token: ${token}
current-context: ${serviceAccount}@${clusterName}
"

