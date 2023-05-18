# DataPowerProfiler
Tool to analysis and summarise the DP cfg file so that it can be shared with engineering without a risk of confidential information being provided.

The Profiler is passed a parameter pointing to the location of a DataPower cfg file.

This file is analysed and a list of objects types from the cfg is curated. Key attributes are also included. At this time the key attributes are only

Action
* type
* aaapolicy

ALL
* admin-enabled


Example command
`./DataPower-profiler-darwin-arm64 -path /path/to/cfg/domain.cfg`

This responds something similar to below

```
--------------------------------------
Loading file :: /path/to/cfg/domain.cfg
--------------------------------------
'action' : 7
'action->type:gatewayscript' : 2
'action->type:results' : 3
'action->type:setvar' : 2
'api-client-identification' : 1
'api-cors' : 1
'api-execute' : 1
'api-final' : 1
'api-rate-limit' : 1
'api-result' : 1
'api-routing' : 1
'api-rule' : 8
'api-security' : 1
'api-security-token-manager' : 1
'api-security-token-manager->admin-state:disabled' : 1
'apiprobe-settings' : 1
'apiprobe-settings->admin-state:disabled' : 1
'assembly' : 4
'assembly-function' : 3
'assembly-function-call' : 3
'assembly-html-page' : 1
'assembly-rate-limit' : 1
'assembly-wsdl' : 1
'compile-options' : 1
'control-list' : 2
'crypto->certificate' : 4
'crypto->ciphers' : 101
'crypto->idcred' : 5
'crypto->key' : 4
'crypto->no-request-client-auth' : 2
'crypto->no-validate-server-cert' : 1
'crypto->protocols' : 3
'crypto->sshdomainclientprofile' : 1
'crypto->sshdomainclientprofile->admin-state:enabled' : 1
'crypto->sshserverprofile' : 1
'crypto->sshserverprofile->admin-state:enabled' : 1
'crypto->ssl-client' : 1
'crypto->ssl-server' : 2
'documentcache' : 5
'domain-availability' : 1
'domain-availability->admin-state:disabled' : 1
'domain-settings' : 1
'domain-settings->admin-state:enabled' : 1
'gateway-peering' : 1
'gateway-peering->admin-state:disabled' : 1
'gateway-peering-manager' : 1
'gateway-peering-manager->admin-state:disabled' : 1
'logging' : 1
'logging->admin-state:disabled' : 1
'matching' : 1
'metadata' : 4
'mpgw' : 2
'nfs-dynamic-mounts' : 1
'nfs-dynamic-mounts->admin-state:disabled' : 1
'no-xml-validate' : 5
'parse-settings' : 1
'policy-attachments' : 3
'rate-limit-config' : 1
'rate-limit-config->admin-state:enabled' : 1
'rule' : 3
'slm-action' : 3
'source-http' : 4
'source-http->admin-state:disabled' : 1
'source-https' : 1
'statistics' : 1
'statistics->admin-state:disabled' : 1
'stylepolicy' : 2
'urlmap' : 1
'user-agent' : 1
'wsm-agent' : 1
'wsm-agent->admin-state:disabled' : 1
'wsm-stylepolicy' : 1
'xmlmgr' : 5
```
