apt-get udpate
apt-get remove nfs-kernel-server
apt-get install nfs-kernel-server

read -p "ip range(such as 192.168.31.0/24): " input;
iprange=`echo $input | tr '[A-Z]' '[a-z]'`;

mkdir /data/nfs-data
chown nobody:nogroup /data/nfs-data
echo "/opt/ssm-nfs-data $iprange(rw,sync,no_subtree_check,no_root_squash)" >> /etc/exports
sudo systemctl restart nfs-kernel-server

