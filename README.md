#A Shipped GoTTY

This sample application demonstrates creating and deploying a Shipped project from an existing GitHub repository. 

 GoTTY - Share your terminal as a web application

[![GitHub release](http://img.shields.io/github/release/yudai/gotty.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/55d0eeff7331453f0801982c.svg?style=flat-square)][wercker]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/yudai/gotty/releases
[wercker]: https://app.wercker.com/project/bykey/03b91f441bebeda34f80e09a9f14126f
[license]: https://github.com/yudai/gotty/blob/master/LICENSE


GoTTY is a simple command line tool that turns your CLI tools into web applications.

![Screenshot](https://raw.githubusercontent.com/yudai/gotty/master/screenshot.gif)


## Usage

```
Usage: gotty [options] <command> [<arguments...>]
```

Run `gotty` with your preferred command as its arguments (e.g. `gotty top`).

By default, GoTTY starts a web server at port 8080. Open the URL on your web browser and you can see the running command as if it were running on your terminal.

### Options

```
--address, -a                                                IP address to listen [$GOTTY_ADDRESS]
--port, -p "8080"                                            Port number to listen [$GOTTY_PORT]
--permit-write, -w                                           Permit clients to write to the TTY (BE CAREFUL) [$GOTTY_PERMIT_WRITE]
--credential, -c                                             Credential for Basic Authentication (ex: user:pass, default disabled) [$GOTTY_CREDENTIAL]
--random-url, -r                                             Add a random string to the URL [$GOTTY_RANDOM_URL]
--random-url-length "8"                                      Random URL length [$GOTTY_RANDOM_URL_LENGTH]
--tls, -t                                                    Enable TLS/SSL [$GOTTY_TLS]
--tls-crt "~/.gotty.key"                                     TLS/SSL crt file path [$GOTTY_TLS_CRT]
--tls-key "~/.gotty.crt"                                     TLS/SSL key file path [$GOTTY_TLS_KEY]
--index                                                      Custom index file [$GOTTY_INDEX]
--title-format "GoTTY - {{ .Command }} ({{ .Hostname }})"    Title format of browser window [$GOTTY_TITLE_FORMAT]
--reconnect                                                  Enable reconnection [$GOTTY_RECONNECT]
--reconnect-time "10"                                        Time to reconnect [$GOTTY_RECONNECT_TIME]
--once                                                       Accept only one client and exit on disconnection [$GOTTY_ONCE]
--config "~/.gotty"                                          Config file path [$GOTTY_CONFIG]
--version, -v                                                print the version

```


## Setting Up on Shipped

### Step 1. Get a GitHub Account
[Git](https://git-scm.com/) is a source and version control system hosted on the web, and [GitHub](https://github.com/) is a repository of Git projects. Shipped stores your projects on GitHub, and you'll need a GitHub account to use Shipped. If you already have a GitHub account, skip to step 2. Otherwise, navigate to the [GitHub home page](https://github.com/) using any modern browser (we recommend Chrome), click on [Sign up](https://github.com/join) and follow the instructions to create a free account. You don't need to do anything more with GitHub at this time; Shipped will do it all for you.

### Step 2 Fork the goTTY project
Fork CiscoCloud/gotty as your repository master project branch.  This sets up goTTY as the sample repository that you'll use for your Shipped project.

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/fork.Png)

### Step 3 Download the Shipped CLI
You can compose a Shipped project either with the Shipped web UI at [http://shipped-cisco.com](http://shipped-cisco.com) or with the Shipped CLI.  However, even if you compose your project online, you'll still need the CLI to bootstrap it to your local laptop, so the first step in using Shipped is to download the CLI.  Download from the appropriate link below:

OS | Download URL
---- | -------------------
linux | https://bintray.com/artifact/download/shippedrepos/shipped-cli/linux/shipped (5.9 MB)
linux (compressed) | https://bintray.com/artifact/download/shippedrepos/shipped-cli/linux/shipped.gz (1.7 MB)
mac | https://bintray.com/artifact/download/shippedrepos/shipped-cli/mac/shipped (6.3 MB)
mac (compressed) | https://bintray.com/artifact/download/shippedrepos/shipped-cli/mac/shipped.gz (1.8 MB)
windows | https://bintray.com/artifact/download/shippedrepos/shipped-cli/windows/shipped.exe (6 MB)
windows (compressed) | https://bintray.com/artifact/download/shippedrepos/shipped-cli/windows/shipped.zip (1.7 MB)

The compressed links contain the same executables as the standard links, but may take less time to download (though you'll need to decompress them after the download). Once you've downloaded an executable, copy it to a directory in your path.   Issue the command:

    shipped -v

to verify you've successfully downloaded and deployed the Shipped CLI.  You should see a response similar to:

> Cisco Shipped 1.0 b103 20150919 001714

### Step 4 Compose and Deploy your Project with the Shipped CLI
This step explains how to compose and deploy your project entirely with the Shipped CLI.  If you'd prefer to use the web UI, skip this step and go to Step 5.

To compose your project with the Shipped CLI, first obtain a terminal window:

| Os Type | Action          |
| ------------------- | -------------------- |
| OS X (Mac)|	Press Command + Space to open Spotlight Search. Type Terminal and double-click the Terminal Application. Menu option Shell -> New Window will open a new window for you. |
| Ubuntu |	Press Ctrl+Alt+T |
| Windows |	Click Start, type "cmd", and press Enter for a normal terminal window, or Ctrl+Shift+Enter for an Administrator terminal window. You will need an Administrator window if the bootstrap process needs to install Vagrant and VirtualBox. If you already have this software installed, you can use a normal terminal window. |

Once you have the terminal window open, enter the following command:

    shipped run create-and-deploy project=MyProject service=gotty framework=express [fast="--fast"]

This command is all you need to create your project and deploy it to the Cloud!  Use the **project** argument to provide a name for your project (you can call it anything you want); specify the **service** and **framework** arguments exactly as shown; and specify the optional **fast="--fast"** argument to suppress building a VM on your laptop for local deployment (which saves two or three minutes of bootstrap time, but doesn't run the application locally).

This command does the following:

* Creates the project and service in the Shipped database
* Sets up a remote GitHub repository in the Cloud
* Sets up a continuous integration (CI) build that rebuilds your project automatically after every commit
* Downloads project source code into a local Git repository on your laptop
* If necessary, installs prerequisite software (Vagrant and Virtualbox)
* Optionally creates a VM on your laptop to build and run the service
* Makes a commit to the Git repository that starts the first build
* Sets up a cloud environment where you can deploy your project
* Waits for the build to complete and deploys the project to the cloud

Once it completes, you'll have your application running both on your laptop and in the cloud.  That's how easy it is to get started with Shipped!

Create-and-deploy is an example of a Shipped CLI script that runs several CLI commands in succession.  You can see the contents of the script with the command:

    shipped script list create-and-deploy

The commands in the script perform the same tasks as the Shipped UI actions described in the next step.  You can get an overview of some CLI commands by following along with the script as you read through the next step.  For more detailed documentation on the CLI, use its help command,

    shipped help

### Step 5. Compose and Deploy Your Project with the Shipped UI
If you've composed your project with the CLI, you can skip this step.  However, if you prefer to compose your project online - or would just like to see how it's done - just use any modern browser (we recommend Chrome) to navigate to the [Shipped welcome page](http://shipped-cisco.com). 

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/home.Png) 
#### Step 5.1 Login to Shipped
Your GitHub account is all you need to login to Shipped, so click on the big green "Sign up with GitHub" button.   

If this is your first project, Shipped automatically pops up a Create New Project form; if not, select Create New Project from the Your Projects dropdown to get the popup:

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/createproject.Png)

Enter a name for your project (we'll use "Sample" for this project) and press Start Composing. Shipped displays a list of development services:

#### Step 5.2 Add a service
The development services listed are examples of what Shipped calls "microservices" - supporting services used by a project. We can choose as many microservices as needed for a project. Shipped installs whatever is needed to deploy the selected microservices on one of the VMs it creates for you. 
 
For this project we'll select **GoLang** which create a golang VM with everything needed to build and run an GoLang project like goTTY.

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/selectservice.Png)

Click on the Select button to the right of this service. Shipped pops up the Service Configuration form that you'll use to specify the GitHub repository where Shipped stores the source code using the service:
 
#### Step 5.3 Configure the Service 
![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/selectrepo.Png)

This form allow us to specify:

| Name | Description          |
| ------------------- | -------------------- |
| Name of GitHub Repository|	The name of the GitHub repository where Shipped stores the source code using this service. Shipped automatically creates the repository if necessary.  *For this example, we'll use the repository you forked in Step 2 and specify* gotty *for the repository name.* |
| GitHub Organization|	The GitHub account owning the repository. This can be your personal account, or the account of a company or organization associated with your account.  *For this example, specify the account where you forked the gotty repository.* |
| Private Public |	The type of repository to create.  Private repositories are available only to specific GitHub users; public repositories are viewable by any web user. You need a paid GitHub account to create a private repository.  *For this example, specify Public, as your fork of a public repository must also be public.*  |
 
#### Step 5.4 Build Your Project
 
Specify "gotty" for the repository name and press Add Service. Shipped re-displays the Compose Your Project form with the repository name for the selected service above the Build Project button.

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/buildproject.Png)

To build your project, press the Build Project button. The button label changes to Building, and a status bar moves across the button while Shipped creates your GitHub repository and stores the description of your project in its database. When it's finished, it pops up the Let's Get Set Up form containing the command to bootstrap the project on your computer:

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/buildlocal.Png)

#### Step 5.5 Bootstrap Your Local Development Environment
Bootstrapping a Shipped project means:

-  Downloads the project's source code to your machine
-  Sets up a local Git repository tied to the cloud-based repository created by Shipped
-  If necessary, installs the prerequisite software Vagrant and VirtualBox
-  Optionally creates the virtual machines that host your development environment

The process is fully automatic; you just need to copy and paste the command presented by Shipped when you created the project.

The bootstrap process runs in a command-line terminal window, so the first step in bootstrapping your project is opening a terminal window. The way you do this depends on your operating system:

| Os Type | Action          |
| ------------------- | -------------------- |
| OS X (Mac)|	Press Command + Space to open Spotlight Search. Type Terminal and double-click the Terminal Application. Menu option Shell -> New Window will open a new window for you. |
| Ubuntu |	Press Ctrl+Alt+T |
| Windows |	Click Start, type "cmd", and press Enter for a normal terminal window, or Ctrl+Shift+Enter for an Administrator terminal window. You will need an Administrator window if the bootstrap process needs to install Vagrant and VirtualBox. If you already have this software installed, you can use a normal terminal window. |

When you completed creating your project in the previous step, Shipped popped up the Let's Get Set Up form containing commands to download the CLI and bootstrap your project:

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/buildlocal.Png)

Since you already installed the Shipped CLI in Step 3, you only need to copy the second command and paste it into your terminal window.  This command looks something like this:

    shipped -t wxkOzQgIgoxSuekHjuyMdMjCIbrKJAcO local bootstrap 65a1fcb4-6141-11e5-befb-0242ac113ce8

If you don't want to create a VM hosting your development environment, append the argument **--fast** to the end of this line.  This saves two or three minutes of bootstrap time, but skips local deployment of the project.  If you choose this option, you can create the VM at a later time by rerunning the bootstrap command without **--fast**.

Press Enter to start the bootstrap process.  While bootstrap runs, the Shipped UI displays a circular animation ticking off each step in the bootstrap process as it happens:

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/buildstatus.Png) 

#### Step 5.6 Run a Build

When you created the project, Shipped set up both a GitHub repository in the cloud and a continuous integration (CI) build, so that any commit to the repository automatically triggers a build.  When you bootstrapped the project, Shipped stored the application's source code in a local Git repository tied to the GitHub repository.  To start a build, all you need to is commit changes to the GitHub repository.  When the bootstrap process completes, the Shipped browser window displays the commit command you need to run your first build:

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/pushbuild.Png)

Once again, select the command by clicking on it and copy and paste it into your terminal window. The command changes the directory to the one containing your new local Git repository and commits the initial copy of the application source to your cloud-based remote Git repository.

This automatically triggers a build, as you can in the event section of your browser window

#### Step 5.7 Create an Environment to Deploy Your Project to the Cloud

The last step in the bootstrap process is deploying your project's application to the Cisco cloud. To do this, click on the Deploy tab at the top left of the screen. Shipped displays the Deploy tab with a message that there are currently no deployed environments:

Create a new Environment by clicking on the New Environment button in the upper right corner of the form.  Shipped displays the New Environment form:

![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/newenv.Png)

Specify a name for your environment and click Add Environment.

#### Step 5.8 Deploy Your Project to the  Environment.
Click on the build in the Select Build column, the environment in the Select Environment column, and then on the Deploy Build button.  Shipped shows the message "Deploying to environment..." under the environment name, and a short time later replaces it with a "Deployed successfully" message:
![](https://github.com/CiscoCloud/gotty/blob/master/shipped/images/deploy.Png)

Congratulations! You've deployed your gotty application to the cloud. Click on the URL in the "Deployed successfully" message to see the application running in its new environment
