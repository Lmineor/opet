#!/usr/bin/python

import getopt
import logging
import os
import signal
import subprocess
import sys

def usage():
    print("""
usage: server [options]

common options:
 --certificate-authority        "/etc/sdnc-net-plugin/etcd/ca.crt"
 --client-certificate           "/etc/sdnc-net-plugin/etcd/server.crt"
 --client-key                           "/etc/sdnc-net-plugin/etcd/server.key"
 --endpoints                            "99.0.85.60:2379"
""")

def main():
    def signal_handle(signum, frame):
        LOG.debug("destroy opet container, signum is %s", signum)

    signal.signal(signal.SIGINT, signal_handle)
    signal.signal(signal.SIGTERM, signal_handle)

    stdout_handler = logging.StreamHandler(sys.stdout)

    formatter = logging.Formatter("%(asctime)s - %(filename)s[line:%(lineno)d] - %(levelname)s: %(message)s")
    stdout_handler.setFormatter(formatter)
    LOG = logging.getLogger()
    LOG.setLevel(logging.DEBUG)
    LOG.addHandler(stdout_handler)

    certificate_authority = "/etc/sdnc-net-plugin/etcd/ca.crt"
    client_certificate = "/etc/sdnc-net-plugin/etcd/server.crt"
    client_key = "/etc/sdnc-net-plugin/etcd/server.key"
    endpoints = "99.0.85.60:2379"

    longopts = ["certificate-authority=", "client-certificate=", "client-key=", "endpoints="]

    try:
        opts, args = getopt.getopt(sys.argv[1:], "", longopts)
    except Exception as e:
        LOG.error("failed to parse command opt, exception is %s", e)
        return
    for name, value in opts:
    	if name == "--certificate-authority":
            certificate_authority = "/etc/sdnc-net-plugin/etcd/ca.crt"
        elif name == "--client-certificate":
            client_certificate = "/etc/sdnc-net-plugin/etcd/server.crt"
        elif name == "--client-key":
            client_key = "/etc/sdnc-net-plugin/etcd/server.key"
        elif name == "--endpoints":
            endpoints = "99.0.85.60:2379"

    opet_command = ['/opet/bin/server',
                    "--certificate-authority", certificate_authority,
                    "--client-certificate", client_certificate,
                    "--client-key", client_key,
                    "--endpoints", endpoints
                    ]
    try:
        ret = subprocess.check_output(opet_command, universal_newlines=True)
        LOG.debug("opet_command %, return value is %s", opet_command, ret)
    except Exception as e:
        LOG.error("failed tp execute opet_command %s, exception is %s", opet_command, e)


if __name__ == '__main__':
    sys.exit(main())