# Hello World: Kubernetes


Language: **Golang**

Cluster management: **Kubernetes, via Minikube**

Containerization: **Docker (+DockerHub)**

Driver/Hypervisor: **xhyve**





### Prep & Installations:

- Be sure [Golang](https://golang.org/doc/install) is installed and Go tools are set up to run our application.
- You’ll need to download [Homebrew](https://docs.brew.sh/Installation.html) to download your driver. In this tutorial we’ll be using [xhyve](https://github.com/mist64/xhyve).
- Install Docker (We'll be using [Docker for Mac](https://docs.docker.com/docker-for-mac/#preferences) for this tutorial)



**Now let’s get into it!**

## Create a Minikube Cluster:

In this tutorial, we’ll be using Minikube to create a cluster locally. If you are not using a Mac, see the [Minikube installation guide](https://github.com/kubernetes/minikube) as the instructions might be different. 

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

So the first thing we did was create a Kubernetes cluster via a VM called Minikube. MiniKube is a popular tool used to run Kubernetes locally. Then we installed a hypervisor (xhyve) for Docker to run on Minikube.


**_Insert diagram w/caption(s)_**

After that, we configured Kubernetes’ command line-tool, **_kubectl_**, to communicate specifically with our minikube cluster. 



**_Insert diagram w/caption(s)_**



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

        docker login --username=yourhubusername --password=iLike2MuvItMuvIt


**RECOMMENDED: USE _docker --password-stdin_ TO LOG IN SAFELY!!**


Now we’ll need to check the image ID:

        docker images


Your output should look something like this:

        REPOSITORY              TAG       IMAGE ID         CREATED           SIZE
        helloworld              v1      056yb71fy406      5 minutes ago    1.076 GB
        monty_grapher          latest    pp58734h67dd     12 minutes ago    1.658 GB
        steph/train            latest    9857j776f554      8 days ago       1.443 GB



Update your image’s tag and the name of your Docker Hub repo:

        docker tag 056yb71fy406 yourhubusername/hello_world:v1


Finally, push the image to your Docker Hub repo:

        docker push yourhubusername/hello_world




### Run the Container

We can test out our container image first by running this command (_be sure to replace ‘yourusername’ with your actual DockerHub username_):

        docker run --p 8080:8080 yourusername/hello_world:v1

Then open a new tab in your terminal and enter: 

        curl http://localhost:8080
        
Lo and behold, there’s our _**‘Hello World’**_ message.



 ### What just happened?..
 

Now we have a full blown application and a Docker image running in the cloud! 

After we downloaded the application, we then created an image (which is an instance of a container) for our application to and it dependencies to live in. We then pushed that image to Docker Hub, Docker’s official container registry. Pushing our container to the cloud gives us the privilege of being able to access that container any given time, even if we tear down our local cluster, or if we want to pull that container to live in separate cluster. After that, we ran the container, binding our local port to the port of the container (_8080:8080_). 



## Deploy

_Remember to stop the container from running by pressing Ctrl-C in the tab where you ran the docker run command._

In Kubernetes, containers are interpreted as objects called [Pods](https://kubernetes.io/docs/concepts/workloads/pods/pod/) (one or more containers in a group). The Pod in our cluster only has one container, the one we just created.

Now how do we manage this Pod? Kubernetes provides a special supervisor for Pods called [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/). Deployments are responsible for monitoring and managing everything from scaling to version control for Pods. They also check and maintain the health of the containers within Pods. 


To create a deployment, we’ll have to use Kubernetes’ kubectl for the following command:

        kubectl run helloworld --image=yourusername/hello_world:v1 --port=8080



Once the Terminal confirms that your deployment has been created, we can view it by running

        kubectl get deployments

Your output should look something like this:

        NAME         DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
        helloworld    1         1         1            1           3m


Now let’s take a look at our Pod:

        kubectl get pods

        NAME                                         READY     STATUS    RESTARTS   AGE
        helloworld-7447bd7d5d-lwnxh   1/1        Running              0            1m




### What just happened?..

Woah! So now our container has become a “Pod”, and we’ve Kubernetes has granted us a manager to keep tabs on our Pods health, scaling and load-balancing, and versioning. So if we had an application with 3 Pods, each of then representing a feature, our deployment would give us control over how and when we can roll out each feature. Powerful stuff huh? Makes microservices that much smoother. Speaking of microservices, let’s move on to the next and final phase of the tutorial.



## Create a Service

In order to make our Pod accessible outside of the cluster, we have to create what’s called a _“Service”_. A Service creates a public IP address for the cluster and presents the individual IP addresses of each Pod as endpoints, allowing allow clients to access and connect to Pods and exposing applications to external traffic. Services also handle load-balancing amongst the Pods as well. 

Go ahead and create a Service by running the command below:

        kubectl expose deployment helloworld --type=LoadBalancer

> Here we use the  **--type=LoadBalancer** flag to indicate that we want our Service to be exposed outside of our cluster.

Now let’s test to see if our Service is accessible:

        minikube service helloworld

> This uses a local IP address that serves our app and opens up a browser displaying our “Hello World” message.


## Updating your Application
 
So now application is exposed, but now what if we need to make changes! Let’s see what our We want to change our message from _“Hello World!!”_ to _“Finally Completed this Tutorial!!”_. 

Let’s go into our source code (_helloworld.go file_) for our application and change it to return our new message.

_Change line 16 in helloworld.go to:_
 
        fmt.Fprint(rw, "Finally Completed this Tutorial!!")
        
Now we want to build a new version of our Docker image:

        docker build -t yourhubusername/hello_world:v2 .
        
Update the image for the Deployment:

        kubectl set image deployment/helloworld helloworld=yourhubusername/hello_world:v2

Now we can check for our updated message:

        minikube service helloworld



## A Clean Finish

Now after all that hard work… **let’s throw it all away!**

You can clean out your cluster simply by using:

        kubectl delete service helloworld
        kubectl delete deployment helloworld

Stop Minikube, then delete it:
        
        minikube stop
        minikube delete






