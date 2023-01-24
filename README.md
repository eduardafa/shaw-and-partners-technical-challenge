<div align=center>
	<img src="https://shawandpartners.com/old/wp-content/uploads/2017/02/LOGO_FINAL_SP.png">

# Shaw and Partners - Back-End Technical Challenge
</div>

This repository was created to present a technical test for the job position of BackEnd Developer at Shaw and Partners.

**OBSERVATION: The tests were not finished as I wanted. Some issues were found, but due to lack of time, they can't be resolved.**

## Resources

* **[Visual Studio Code](https://code.visualstudio.com)**
* **[Goland](https://www.jetbrains.com/go/)**
* **[Insomnia](https://insomnia.rest/)**

## Running the Local Application
To run the project on your local machine, just follow the steps below:
### Starting...

You should clone the project repository on your machine and install the necessary packages.
```
git clone https://github.com/eduardafa/shaw-and-partners-technical-challenge.git
```

### Installing Go packages

Open your terminal and type the path where you cloned the project.

```
cd "C:/Users/Username/Documents/.../shaw-and-partners-technical-challenge"
./gomod.sh
```

### .env file

There is a file called `.env.example` in the `config` folder. Please, rename this file to `.env`.


### Running the Application
Start the server with what is indicated below to get the project to run locally.

```
go run ./cmd
```
After the execution, the following message will appear in the terminal:
```
Starting the server with GoLang
Loading .env file
.env file loaded successfully
```

## HTTP Routes

This API is running on `port 8080` and so that all routes can be accessed locally it is necessary to use `http://localhost:8080` before the request endpoints.

### GitHub User
| Feature                                                                                                                 | Method | Route                   |
|-------------------------------------------------------------------------------------------------------------------------|--------|-------------------------|
| Returns a list of GitHub users with an ID greater than the defined ID | GET    | `/api/users?since={number}` |
| Returns the details of a GitHub user | GET    | `/api/users/{username}/details` |
| Returns a list with all GitHub user repositories | GET    | `/api/users/{username}/repos` |

## Testing the application
The endpoints described above can be accessed from the `insomnia` folder.

## Tests
The tests can be run by the following command:
```
cd _tests
go test routes_test.go
go test controllers_test.go
```

## Go installation
If you don't have Go installed on your machine, follow the steps below:
```
# download the latest version o Golang Toolchain
wget https://golang.org/dl/go1.19.4.linux-amd64.tar.gz

# install locally
sudo tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz

# set path
export PATH=$PATH:/usr/local/go/bin

# verify go version
go version

# modify your profile to export the path for all sessions
nano ~/.profile

export PATH=$PATH:/usr/local/go/bin

# reiniciar sessão para que as alterações passem a ter efeito
```

### Update Go global version to latest stable
```
git clone https://github.com/udhos/update-golang
cd update-golang
sudo ./update-golang.sh
```

## Developed by:
<div align=center>
<!--Eduarda França-->      
            <td>
                <a href="https://github.com/eduardafa">
                    <img src="https://avatars.githubusercontent.com/u/54603419?v=4" width="120px;" alt="Eduarda França" style="max-width:100%;">
                </a><br><br>
                <p><b>Eduarda França</b></p><br>
                <a href="https://github.com/eduardafa">
                    <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/github/github-original.svg" width="27px">
                </a>
                <a href="https://www.linkedin.com/in/eduardandrade">
                    <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linkedin/linkedin-original.svg" width="27px">
                </a>
            </td>
</div>
