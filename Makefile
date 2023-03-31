install:
	python3 -m pip install --user ansible && ansible-galaxy collection install community.docker

ansible-update:
	python3 -m pip install --upgrade --user ansible

start:
	vagrant up

clean:
	vagrant destroy -f