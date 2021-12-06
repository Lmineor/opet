apiVersion: v1
kind: Namespace
metadata:
  name: opet-ns
---

kind: ConfigMap
apiVersion: v1
metadata:
  name: opet-config
  namespace: opet-ns
data:
  certificate-authority: "/etc/sdnc-net-plugin/etcd/ca.crt"
  client-certificate: "/etc/sdnc-net-plugin/etcd/server.crt"
  client-key: "/etc/sdnc-net-plugin/etcd/server.key"
  endpoints: "99.0.85.60:2379"

---


apiVersion: apps/v1
kind: Deployment
metadata:
  name: opet-app
  namespace: opet-ns
  labels:
    app: opet
spec:
  selector:
    matchLabels:
      app: opet
  replicas: 1
  template:
    metadata:
      labels:
        app: opet
    spec:
      nodeSelector:
        kubernetes.io/hostname: 'k8s-node2'
      volumes:
        - name: opet-etc-dir
          hostPath:
            path: /etc/sdnc-net-plugin/
      containers:
        - name: opet
          image: "opet:v0.1"
          imagePullPolicy: IfNotPresent
          args:
          - --certificate-authority
          - $(ETCD_CA)
          - --client-certificate
          - $(CLIENT_CA)
          - --client-key
          - $(CLIENT_KEY)
          - --endpoints
          - $(EP)
          ports:
          - containerPort: 8080
            name: container
            protocol: TCP
          volumeMounts:
            - mountPath: /etc/sdnc-net-plugin/
              name: opet-etc-dir
              readOnly: false
          env:
            - name: ETCD_CA
              valueFrom:
                configMapKeyRef:
                  name: opet-config
                  key: certificate-authority
            - name: CLIENT_CA
              valueFrom:
                configMapKeyRef:
                  name: opet-config
                  key: client-certificate
            - name: CLIENT_KEY
              valueFrom:
                configMapKeyRef:
                  name: opet-config
                  key: client-key
            - name: EP
              valueFrom:
                configMapKeyRef:
                  name: opet-config
                  key: endpoints


---
apiVersion: v1
kind: Service
metadata:
  name: opet-service
  namespace: opet-ns
  labels:
    app: opet-svc
spec:
  type: NodePort
  ports:
    - port: 8080
      name: container
      protocol: TCP
      targetPort: 8080
      nodePort: 30033
  selector:
    app: opet