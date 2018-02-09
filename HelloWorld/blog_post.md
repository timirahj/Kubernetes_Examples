# Hello World: To FaaS or Not to FaaS


Language: Golang
Cluster management: Kubernetes, via Minikube
Containerization: Docker (+DockerHub)
FAAS: Fission
Driver/Hypervisor: Virtualbox





Serverless this, serverless that. What’s the big deal? Well, if you’re  As a developer (especially front-end), the learning curve for Kubernetes can be be super difficult (and annoying). 



Prep & Installations:

Be sure Golang [link] is installed and Go tools are set up to run our application.
You’ll need to download Homebrew [link] to download your driver. In this tutorial we’ll be using xhyve [link].
Install Docker (Docker for Mac [link] is recommended for Mac users)



Now let’s get into it!

Create a Minikube Cluster:

In this tutorial, we’ll be using Minikube to create a cluster locally. Just as the original example tutorial, this tutorial also assumes that you are using Docker for Mac. If you are on a different platform (Windows, Linux, etc.), see the Minikube installation guide as the instructions [link] might be different. 

Use curl to download the latest release of Minikube:

        curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 && \
        chmod +x minikube && \
        sudo mv minikube /usr/local/bin/

Now we’re gonna use Homebrew to install the driver:

        brew install docker-machine-driver-xhyve
        sudo chown root:wheel $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
        sudo chmod u+s $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve

We need to install Kubernetes’ kubectl command-line tool, which will be our right hand for interacting with our cluster:

        brew install kubectl


Now let’s go ahead and start the Minikube cluster:

        minikube start --vm-driver=xhyve

(Note that the --vm-driver=xhyve  flag specifies that you are using Docker for Mac)




Next, you’ll need to configure kubectl to communicate specifically to the minikube cluster. In order to do that, we have to set the Minikube context, as such:

        kubectl config use-context minikube

Now let’s check to see if kubectl is all configured to interact with our cluster:

        kubectl cluster-info




What just happened?..


Let’s create our Golang application!

Before we go writing any code, we need to make sure we set up a workspace. 

Create a folder on your desktop named ‘goworkspace’. Now inside this folder, we’re gonna create 3 folders -- ‘bin’, ‘pkg’, and ‘src’. Inside the src folder is where our source code files will live. 


 
Now we can create our Go application. Inside your favorite text editor, copy and paste the code below. When you’re done, you can hit Save or Save As to save the file, and name it helloworld.go. Make sure you’ve selected the src folder as the parent folder.


        package main
        
        import (
        "fmt"
        "net/http"
        )
        
        func main() {
          http.HandleFunc("/", HelloServer)
          http.ListenAndServe(":8080", nil)
        }
        
        func HelloServer(rw http.ResponseWriter, req *http.Request) {
          fmt.Fprint(rw, "Hello World!!")
        }



We’ve built our app, now let’s run it by heading into our src directory in the terminal as so (assuming you’re still checked into the Desktop directory):

        cd goworkspace/src
        go run helloworld.go


You may get a prompt like the one below.

[insert photo]
