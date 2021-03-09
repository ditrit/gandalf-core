Installation
Installation et utilisation de minik8s:
*vous pouvez retrouver la doc sur https://microk8s.io/*

Lancer l'installation standard:
```
sudo snap install microk8s --classic
```
Regarder le status de Kubernetes:
```
microk8s status --wait-ready

Activer les services voulu
```
# Replace SERVICES by your service
# Par exemple pour avoir le dashboard en proxy remplacer par dashboard dns registry istio 
microk8s enable SERVICES
```

Alias si on utilise aussi les service de Kuectl:
```
alias mkctl="microk8s kubectl"
```
Alias si on n'utilise que microk8s :
```
alias kubectl='microk8s kubectl'
```

Arreter microk8s :
```
microk8s stop
```
Lancer microk8s:
```
microk8s start
```

Vous devez a present etre en capaciter de lancer les commandes standard dans votre cluster microk8s.


Installation et utilisation de Docker:
Pour nos besoins nous allons aussi installer Docker pour cree notre propre image qui contient ce que l'on desire.
```
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo apt-get install docker-ce docker-ce-cli containerd.io
VERSION_STRING=5:20.10.4~3-0~ubuntu-focal
sudo apt-get install docker-ce=${VERSION_STRING} docker-ce-cli=${VERSION_STRING} containerd.io
```
Vous avez maintenant un docker fonctionnelle.


Cree votre propre image docker :

Cree votre Dockerfile avec ce que vous souhaitez faire.

```
sudo docker login
sudo docker build -t edgaunt/gandalf .
sudo docker push edgaunt/gandalf
```

cree votre fichier de fichier de deployment de gandalf en utilisant votre images

Deploy votre gandalf
```
mkctl apply -f gandalf.yaml
```






Les commandes a connaitres :
Deployer un pod :
```
# Toujours privilegier l'utilisation de apply a deploy 
mkctl apply -f <file>.yaml
# OU
mkctl deploy -f <file>.yaml
```
Regarder les pods actif dans le namespace gandalf:
```
mkctl get pods -n gandalf
```

Supprimmer un pod :
```
mkctl delete pods -n gandalf <pod_name>

```
Debug un pod :
```
#Connection SSH au pod quand disponnible
mkctl exec -it <pod_name> -n gandalf -- /bin/sh
#Voir l'etat du pod
mkctl describe pods -n gandalf <pod_name>
#Voir les logs du pod
mkctl logs -n gandalf <pod_name>
```

