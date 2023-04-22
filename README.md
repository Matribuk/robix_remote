# Mr Matribuk presents :
## Clone ChatGPT on a local site with OpenAI API

This project allows you to clone ChatGPT, a natural language processing model, onto a local website using OpenAI's API. The website will allow users to chat with ChatGPT on various topics.


## Installation

1. Clone the project from GitHub
```shell
git clone https://github.com/Matribuk/robix_remote.git
```
2. Create an account on https://openai.com/ and obtain an API key
3. Add the OpenAI API key to an environment variable named `OPENAI_API_KEY`
#### MacOS :
```shell
export OPENAI_API_KEY="API KEY"
```
#### Linux :
```shell
setenv OPENAI_API_KEY="API KEY"
```
4. Install Docker and Docker-Compose:
- [Linux](https://docs.docker.com/engine/install/)

- [MacOS](https://docs.docker.com/docker-for-mac/install/)

- [Windows](https://docs.docker.com/docker-for-windows/install/)

- [Docker Compose](https://docs.docker.com/compose/install/)
 
5. Run `docker-compose up` from the project root directory
```shell
docker-compose up
```

## Usage

Once the Docker containers have started, go to `http://localhost:3000` in your browser to access the website. You can now chat with ChatGPT by typing in the dialogue box.

## Credits

| [<img src="https://github.com/Matribuk.png?size=85" width=85><br><sub>Antonin Leprest</sub>](https://github.com/Matribuk) |
|:---:|

This project uses the ChatGPT model from OpenAI (https://beta.openai.com/docs/models/gpt-3) for conversations.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
