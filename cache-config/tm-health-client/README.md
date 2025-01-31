<!--
    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
-->
<!--

  !!!
      This file is both a Github Readme and manpage!
      Please make sure changes appear properly with man,
      and follow man conventions, such as:
      https://www.bell-labs.com/usr/dmr/www/manintro.html

      A primary goal of t3c is to follow POSIX and LSB standards
      and conventions, so it's easy to learn and use by people
      who know Linux and other *nix systems. Providing a proper
      manpage is a big part of that.
  !!!

-->
# NAME

tm-health-client - Traffic Monitor Health Client service

# SYNOPSIS

tm-health-client [-f config-file]  -h  [-l logging-directory]  -v 

# DESCRIPTION

The tm-health-client command is used to manage **Apache Traffic Server** parents on a
host running **Apache Traffic Server**.  The command should be started by **systemd** 
and run as a service. On startup, the command reads its default configuration file
**/etc/trafficcontrol-cache-config/tm-health-client.json**.  After reading the config
file it polls the configured **Traffic OPs** to obtain a list of **Traffic Monitors**
for the configured **CDN** and begins polling the available **Traffic Monitors** for
Traffic Server cache statuses.

On each polling cycle, defined in the configuration file, the Traffic Server parent
statuses are updated from the Traffic Server **parent.config**, **strategies.yaml** 
files, and the Traffic Server **HostStatus** subsystem.  If **Traffic Monitor** has
determined that a parent utilized by the **Traffic Server** instance is un-healthy or
otherwise unavailable, the tm-health-client will utilize the **Traffic Server** 
**traffic_ctl** tool to mark down the parent host.  If a parent host is marked down 
and **Traffic Monitor** has determined that the marked down host is now available, 
the client will then utilize the **Traffic Server** tool to mark the host back up.

# OPTIONS

-f, -\-config-file=config-file 
  
  Specify the config file to use.  
  Defaults to /etc/trafficcontro-cache-config/tm-health-client.json

-h, -\-help 

  Prints command line usage and exits

-l, -\-logging-dir=logging-directory

  Specify the directory where log files are kept.  The default location
  is **/var/log/trafficcontrol-cache-config/**

-v, -\-verbose

  Logging verbosity.  Errors are logged to the default log file 
  **/var/log/trafficcontrol-cache-config/tm-health-client.log**
  To add Warnings, use -v.  To add Warnings and Informational 
  logging, use -vv.  Finally you may add Debug logging using -vvv.

# CONFIGURATION

The configuration file is a **JSON** file and is looked for by default
at **/etc/trafficcontrol-cache-config/tm-health-client.json**

Sample configuarion file:

```
  {
    "cdn-name": "over-the-top",
    "enable-active-markdowns": false,
    "reason-code": "active",
    "to-credential-file": "/etc/credentials",
    "to-url": "https://tp.cdn.com:443", 
    "to-request-timeout-seconds": "5s",
    "tm-poll-interval-seconds": "60s",
    "trafficserver-config-dir": "/opt/trafficserver/etc/trafficserver",
    "trafficserver-bin-dir": "/opt/trafficserver/bin",
  }
```

### cdn-name 

  The name of the CDN that the Traffic Server host is a member of.

### enable-active-markdowns

  When enabled, the client will actively mark down Traffic Server parents.
  When disabled, the client will only log that it would have marked down
  Traffic Server parents

### reason-code

  Use the reason code **active** or **local** when marking down Traffic Server
  hosts in the Traffic Server **HostStatus** subsystem.

### to-credential-file

  The file where **Traffic Ops** credentials are read.  The file should define the 
  following variables:

  * TO_URL="https://trafficops.cdn.com"
  * TO_USER="touser"
  * TO_PASS="touser_password"

### to-url

  The **Traffic Ops** URL

### to-request-timeout-seconds

  The time in seconds to wait for a query response from both **Traffic Ops** and
  the **Traffic Monitors**

### tm-poll-interval-seconds

  The polling interval in seconds used to update **Traffic Server** parent
  status.

### trafficserver-config-dir

  The location on the host where **Traffic Server** configuration files are 
  located.

### trafficserver-bin-dir

  The location on the host where **Traffic Server** **traffic_ctl** tool may
  be found.

# Files

* /etc/trafficcontrol-cache-config/tm-health-client.json
* /etc/logrotate.d/tm-health-client-logrotate
* /usr/bin/tm-health-client
* /usr/lib/systemd/system/tm-health-client.service
* /var/log/trafficcontrol-cache-config/tm-health-client.json
* Traffic Server **parent.config**
* Traffic Server **strategies.yaml**
* Traffic Server **traffic_ctl** command
  
   



