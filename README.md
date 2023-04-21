# Cloner ChatGPT sur un site local avec OpenAI API

Ce projet permet de cloner ChatGPT, un modèle de traitement du langage naturel, sur un site web local en utilisant l'API d'OpenAI. Le site web permettra aux utilisateurs de discuter avec ChatGPT sur divers sujets.


## Installation

1. Cloner le projet à partir de GitHub
```shell
git clone https://github.com/Matribuk/robix_remote.git
```
2. Créer un compte sur https://openai.com/ et obtenir une clé API
3. Ajouter la clé API d'OpenAI à une variable d'environnement nommée `OPENAI_API_KEY`
#### MACOS :
```shell
export OPENAI_API_KEY="API KEY"
```
#### Linux :
```shell
setenv OPENAI_API_KEY="API KEY"
```
4. Installer Docker et Docker-Compose :
- [Linux](https://docs.docker.com/engine/install/)

- [MacOS](https://docs.docker.com/docker-for-mac/install/)

- [Windows](https://docs.docker.com/docker-for-windows/install/)

- [Docker Compose](https://docs.docker.com/compose/install/)
 
5. Exécuter `docker-compose up` à partir de la racine du projet
```shell
docker-compose up
```

## Utilisation

Une fois que les conteneurs Docker ont démarré, accédez à `http://localhost:3000` dans votre navigateur pour accéder au site web. Vous pouvez maintenant discuter avec ChatGPT en tapant dans la boîte de dialogue.

## Crédits

| [<img src="https://github.com/Matribuk.png?size=85" width=85><br><sub>Antonin Leprest</sub>](https://github.com/Matribuk) |
|:---:|

Ce projet utilise le modèle ChatGPT de OpenAI (https://beta.openai.com/docs/models/gpt-3) pour les conversations.

## Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.
