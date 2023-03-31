# <b>*Demo - Vagrant with Ansible*</b>

## <b>*Requirements:*</b>
* *Ansible > 2.9*
* *community.docker*
* *Vagrant*
* *Make*

## <b>*Install Ansible and requirements*</b>
```bash
$ make install
```

## <b>*Need to update ansible?*</b>
```bash
$ make ansible-update
```

## <b>*How to start?*</b>
```bash
$ make start
```
*This command will start vagrant and make the provisioning with ansible.*

## <b>*What we get?*</b>
* *Demo app in golang running on http://192.168.0.10:8080 with docker.*
* *Demo balancer with nginx running on http://192.168.0.11*

*The traffic on http://192.168.0.11 will be redirected to the application.*

## <b>*Destroy all*</b>
*Once you finish to use the demo you can clean all by the following command:*

```bash
$ make clean
```

## <b>Repo scaffolding</b>
* *ansible -> playbooks, config, inventory, files to copy (nginx.conf, app1).*
* *Vagrantfile -> Vagrant configurations.*
* *Makefile -> some commands automation.*

## <b>Project Purpose</b>
*The objective of this project is make use of Vagrant with Ansible creating a simple infrastructure automation using some IaC (Infrastructure As Code). <br>
We will create a Golang API and an nginx balancer that points to our app with only running one command!* <br><br>

#### <b>Author</b>: *Jonathan Brenman - 2023*