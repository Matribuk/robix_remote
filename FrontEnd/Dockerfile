# Utilise l'image officielle de Node.js en version 14
FROM node:14

# Définit le répertoire de travail
WORKDIR /app

# Copie les fichiers de l'application dans le conteneur
COPY package*.json ./
COPY yarn.lock ./
COPY . .

# Installe les dépendances
RUN yarn install

# Build l'application pour la production
RUN yarn build

# Expose le port 3000
EXPOSE 3000

# Démarre l'application
CMD [ "yarn", "start" ]
