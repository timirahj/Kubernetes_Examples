# Hello World: To FaaS or Not to FaaS


Language: Golang
Cluster management: Kubernetes, via Minikube
Containerization: Docker (+DockerHub)
FAAS: Fission
Driver/Hypervisor: Virtualbox





Serverless this, serverless that. What’s the big deal? Well, if you’re  As a developer (especially front-end), the learning curve for Kubernetes can be be super difficult (and annoying). 



### Prep & Installations:

- Be sure [Golang](https://golang.org/doc/install) is installed and Go tools are set up to run our application.
- You’ll need to download [Homebrew](https://docs.brew.sh/Installation.html) to download your driver. In this tutorial we’ll be using [xhyve](https://github.com/mist64/xhyve).
- Install Docker ([Docker for Mac](https://docs.docker.com/docker-for-mac/#preferences) is recommended for Mac users)



**Now let’s get into it!**

## Create a Minikube Cluster:

In this tutorial, we’ll be using Minikube to create a cluster locally. Just as the original example tutorial, this tutorial also assumes that you are using Docker for Mac. If you are on a different platform (Windows, Linux, etc.), see the [Minikube installation guide](https://github.com/kubernetes/minikube) as the instructions might be different. 

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

>(Note that the --vm-driver=xhyve  flag specifies that you are using Docker for Mac)




Next, you’ll need to configure kubectl to communicate specifically to the minikube cluster. In order to do that, we have to set the Minikube context, as such:

        kubectl config use-context minikube

Now let’s check to see if kubectl is all configured to interact with our cluster:

        kubectl cluster-info




### What just happened?..


## Let’s create our Golang application!


Go and download the HelloWorld source code by running the commands below:

        git clone https://github.com/timirahj/kubernetes_examples
        cd kubernetes_examples/helloworld

If you take a look inside the repo, a Dockerfile has already been created. A Dockerfile typically contains all the instructions on how the image is built. However, if you open our Dockerfile, you will notice that it looks a little vague with only two simple commands. Is this Dockerfile complete? Actually, yes! Golang has a variant called “onbuild” which simplifies the build process for our Docker image. When we use the onbuild variant, we’re implying that our image application should be built with generalized instructions as any generic Go application, and the image automatically copies the package source then builds the program and configures it to run upon startup. 


In the next step, we’ll be packaging our application in a Docker container.


## Create our Docker Image

Now let’s build our container image and tag it:

        docker build -t helloworld:v1 .


Let’s double check to see if our build succeeded. If it was, we’ll see our image listed by running this command:

        docker images



### Push your Docker Image to the Cloud

Now we need to push our container to a registry, so we’ll use DockerHub for this tutorial.

If you’re running Docker for Mac, make sure you’re logged into your Docker account and that Docker is running on your machine. You can do that by clicking the Docker icon at the top of your screen. You should see a green light to verify that it’s running. 

([Click here for these instructions using other operating systems](https://docs.docker.com/docker-for-windows/install/).)


Go to [https://hub.docker.com](https://hub.docker.com), log in, then create a repository called hello-world (_ex. timirahj/hello-world_). 


Now let’s log into the Docker Hub from the command line:

        docker login --username=yourhubusername --email=youremail@company.com


Then enter your password when prompted.


Now we’ll need to check the image ID:

        docker images


Your output should look something like this:

        REPOSITORY              TAG       IMAGE ID         CREATED           SIZE
        hello-world              v1      056yb71fy406      3 minutes ago    1.076 GB
        monty_grapher          latest    pp58734h67dd     13 minutes ago    1.658 GB
        steph/train            latest    9857j776f554      3 days ago       1.443 GB



Update your image’s tag and the name of your Docker Hub repo:

        docker tag 056yb71fy406 yourhubusername/hello-world:firstpush


Finally, push the image to your Docker Hub repo:

        docker push yourhubusername/hello-world





 ### What just happened?..






















