# Step 1: Common Log locations

The first step in locating logs is to know where to look. In this lesson, we will cover the most common log locations on a Linux server. 

## /var/log

The `/var/log` directory is the most common location for log files on a Linux server. This directory contains log files for various system services, including the kernel, system, and application logs.

Lets take a look at the contents of the `/var/log` directory:

```
ls /var/log
```{{exec}}

## /var/log/syslog

The `/var/log/syslog` file contains messages from the Linux kernel and system services. This file is a good place to start when troubleshooting system issues.

Lets take a look at the contents of the `/var/log/syslog` file:

```
cat /var/log/syslog
```{{exec}}



