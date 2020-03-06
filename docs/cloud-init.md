#Cloud-init

Kip supports provisioning new Cells through a subset of functionality provided by the popular cloud-init system.  Users can specify a cloud-init file in server.yml and the cloud-init file will be applied when a Cell is booted by Kip.

Kip's cloud-init system provides the following initialization functions:

* Initialize users and set SSH authorized keys
* Set SSH authorized keys for the root user
* Set the hostname
* Write arbitrary files (allowed encodings: plain text, base64, gzip and gzip+base64)
* Execute scripts

##Cloud-init Example

```yaml
users:
  - name: "dbowie"
    passwd: "$6$qhNlkpFW$p.YhGhk1zFd0bQ1Quk/3O42qEtp7vjZ5DB8C/l/VUB..."
    groups:
      - "wheel"
    ssh_authorized_keys:
      - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPU7h8CaYA1VH/CwY3Ah..."

# without a user, ssh_authorized_keys will be added to the root user
ssh_authorized_keys:
  - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0g+ZTxC7weoIJLUafOgrm+h..."

write_files:
  - content: |
        [Unit]
        Description=Socket for the API

        [Socket]
        ListenStream=2375
        Service=docker.service

        [Install]
        WantedBy=sockets.target
    owner: root:root
    path: /root/config
    permissions: "0644"
  - encoding: gzip
    content: !!binary |
        H4sIAIDb/U8C/1NW1E/KzNMvzuBKTc7IV8hIzcnJVyjPL8pJ4QIA6N+MVxsAAAA=
    path: /usr/bin/hello
    permissions: "0755"

hostname: foo.bar.com

runscript: |
  #!/bin/ash
  echo "dbowie   ALL=(ALL)   NOPASSWD: ALL" >> /etc/sudoers.d/dbowie
  apk update
  apk add curl
```

###Reloading and Limits
The cloud-init file can be updated while Kip is running so it's not necessary to restart Kip.  An updated cloud config will only be applied to Cells started after the file has been updated.

### Limitations
The cloud-init file is served to Cells through the EC2 User Data.  User Data is limited to a size of 16KB.  Internally, Kip uses approximately 4KB of the User Data leaving 12K of data for a user's cloud-init file.  If the cloud-init file is too large, Kip will not be able to start Cells.  Currently we do not compress the user data.