# The Docker host kernel will need the following kernel modules
#  nfs
#  nfsd
#  rpcsec_gss_krb5 (only if Kerberos is used)
# You can manually enable these modules on the Docker host with:
#  modprobe {nfs,nfsd,rpcsec_gss_krb5}
docker start nfs-server ||                                                      \
docker run -d                                                                   \
  -v /exports/data-0001:/exports/data-0001                                      \
  -v /exports/data-0002:/exports/data-0002                                      \
  -e NFS_EXPORT_0='/exports/data-0001                  *(ro,no_subtree_check)'  \
  -e NFS_EXPORT_1='/exports/data-0002                  *(rw,no_subtree_check)'  \
  --privileged                                                                  \
  -p 2099:2049                                                                  \
  --name nfs-server                                                             \
  erichough/nfs-server                                                          