package k8s

const mysqlManifest =`
kind: PersistentVolume
apiVersion: v1
metadata:
  name: mysql-pv-volume
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 20Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data"
---
    apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mysql-pv-claim
spec:
  storageClassName: manual 
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi
apiVersion: v1
kind: Service
metadata:
  name: mysql
spec:
  ports:
  - port: 3306
  selector:
    app: mysql 
  clusterIP: None
---
  apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
spec:
  selector:
    matchLabels:
      app: mysql
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - image: mysql:5.6
        name: mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: password
        ports:
        - containerPort: 3306
          name: mysql
        volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql
      volumes:
      - name: mysql-persistent-storage
        persistentVolumeClaim:
          claimName: mysql-pv-claim
`

const senderManifest =`
---
apiVersion: v1
kind: Pod
metadata:
  name: {{.PodName}}
spec:
  hostNetwork: true
  nodeName: {{.NodeName}}
  containers:
    - name: {{.PodName}}
      image: {{.ContainerImage}}
      imagePullPolicy: IfNotPresent
      command: ["/bin/kokotap_pod"]
      args: [
        "--procprefix=/host",
        "mode",
        "sender",
        "--containerid={{.ContainerID}}",
        "--mirrortype={{.MirrorType}}",
        "--mirrorif={{.MirrorIF}}",
        "--ifname={{.IFName}}",
        "--vxlan-egressip={{.EgressIP}}",
        "--vxlan-ip={{.VXLANIP}}",
        "--vxlan-id={{.VXLANID}}",
        "--vxlan-port={{.VXLANPort}}"
      ]
      securityContext:
        privileged: true
      volumeMounts:
      - name: var-docker
        mountPath: /var/run/docker.sock
      - name: proc
        mountPath: /host/proc
  volumes:
    - name: var-docker
      hostPath:
        path: /var/run/docker.sock
    - name: proc
      hostPath:
        path: /proc
`

const receiverManifest = `
---
apiVersion: v1
kind: Pod
metadata:
  name: {{.PodName}}
spec:
  hostNetwork: true
  nodeName: {{.NodeName}}
  containers:
    - name: {{.PodName}}
      image: {{.ContainerImage}}
      imagePullPolicy: IfNotPresent
      command: ["/bin/kokotap_pod"]
      args: [
        "--procprefix=/host",
        "mode",
        "receiver",
        "--ifname={{.IFName}}",
        "--vxlan-egressip={{.EgressIP}}",
        "--vxlan-ip={{.VXLANIP}}",
        "--vxlan-id={{.VXLANID}}",
        "--vxlan-port={{.VXLANPort}}"
      ]
      securityContext:
        privileged: true
`
