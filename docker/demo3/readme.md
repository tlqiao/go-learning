build image
docker build -f Dockerfile -t tlqiao/http-basic:0.1 src/http
docker push tlqiao/http-basic:0.1
docker hub login 
user:qiaotl@163.com
pass:Qtl1986926

get image in k8s
create secret firstly
kubectl create secret docker-registry regcred \
--docker-server=https://index.docker.io/v1/ \
--docker-username=tlqiao \
--docker-password=Qtl1986926 \
--docker-email=qiaotl@163.com

connect aws eks cluster from local
vim ~/.aws/credentials
update it to latest credential

aws eks update-kubeconfig --region ap-southeast-1 --name taoli-cluster
execute above command to generate kubeconfig file

export KUBECONFIG=$KUBECONFIG:~/.kube/config
