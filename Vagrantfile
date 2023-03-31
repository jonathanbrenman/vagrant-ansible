# ----------------------------------------------------------------------------- #

# app1
Vagrant.configure("2") do |config|
  # VM OS
  config.vm.box = "ubuntu/kinetic64"
  config.vm.hostname = "app1"
  
  # Static IP
  config.vm.network "public_network", ip: "192.168.0.10", bridge: "eno1"
  
  # VM Settings
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "1024"
    vb.name = "app1"
  end
  
  # Provisioning ssh public keys
  config.vm.provision "shell" do |s|
    ssh_pub_key = File.readlines("#{Dir.home}/.ssh/id_rsa.pub").first.strip
    s.inline = <<-SHELL
      echo #{ssh_pub_key} >> /home/vagrant/.ssh/authorized_keys
      echo #{ssh_pub_key} >> /root/.ssh/authorized_keys
    SHELL
  end

  config.vm.provision "ansible" do |ansible|
    ansible.verbose = "v"
    ansible.playbook = "./ansible/playbooks/application.yml"
    ansible.config_file = "./ansible/ansible.cfg"
    ansible.inventory_path = "./ansible/hosts/inventory.ini"
    ansible.limit = "application"
  end
end

# ----------------------------------------------------------------------------- #

# nginx
Vagrant.configure("2") do |config|
  # VM OS
  config.vm.box = "ubuntu/kinetic64"
  config.vm.hostname = "nginx"
  
  # Static IP
  config.vm.network "public_network", ip: "192.168.0.11", bridge: "eno1"
  
  # VM Settings
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "1024"
    vb.name = "nginx"
  end
  
  # Provisioning ssh public keys
  config.vm.provision "shell" do |s|
    ssh_pub_key = File.readlines("#{Dir.home}/.ssh/id_rsa.pub").first.strip
    s.inline = <<-SHELL
      echo #{ssh_pub_key} >> /home/vagrant/.ssh/authorized_keys
      echo #{ssh_pub_key} >> /root/.ssh/authorized_keys
    SHELL
  end

  config.vm.provision "ansible" do |ansible|
    ansible.verbose = "v"
    ansible.playbook = "./ansible/playbooks/nginx.yml"
    ansible.config_file = "./ansible/ansible.cfg"
    ansible.inventory_path = "./ansible/hosts/inventory.ini"
    ansible.limit = "nginx"
  end
end

# ----------------------------------------------------------------------------- #