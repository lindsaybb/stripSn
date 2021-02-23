# stripSn

stripSn is an Iskratel PON utility program for stripping Serial Numbers from the ONU Blacklist table in CLI. Also useful from the ONU Whitelist table.

| Flag | Description |
| ------ | ------ |
| -h | Show this help |
| -i | File to read blacklist entries from (default "blackList.txt") |
| -o | File to write new serial numbers to (default "authList.txt") |
| -a | Overwrite entries if file exists (default false) |
| -si| Read blacklist entries from stdin |
| -so | Write serial numbers to stdout |

# Example Input
```sh
Interface  Serial Number     Password              Cause 
---------  ----------------  --------------------  -----------------------
0/1        ISKT019ECFFA                            Serial Number not known
0/1        ISKT019ECE6C                            Serial Number not known
0/1        ISKT019ECE4A                            Serial Number not known
0/1        ISKT020F069A                            Serial Number not known
0/2        ISKT019ECF20                            Serial Number not known
0/2        ISKT019ECECE                            Serial Number not known
0/3        ISKT019ED0B2                            Serial Number not known
0/3        ISKT019ECF0A                            Serial Number not known
0/3        ISKT019ECF26                            Serial Number not known
0/3        ISKT019ECED8                            Serial Number not known
0/4        ISKT019ED03A                            Serial Number not known
0/4        ISKT019ECF14                            Serial Number not known
0/4        ISKT019ECEAC                            Serial Number not known
0/4        ISKT019ECEE8                            Serial Number not known
 ---------------------------
```

# Example Output
```sh
ISKT019ECFFA
ISKT019ECE6C
ISKT019ECE4A
ISKT020F069A
ISKT019ECF20
ISKT019ECECE
ISKT019ED0B2
ISKT019ECF0A
ISKT019ECF26
ISKT019ECED8
ISKT019ED03A
ISKT019ECF14
ISKT019ECEAC
ISKT019ECEE8
```

# Example Input
```sh
Interface State    State   Number       Password   IP Address      MAC Address       Description
--------- -------- ------- ------------ ---------- --------------- ----------------- ------------------
0/1/1     Up       Enable  ISKT019ECFFA            10.5.101.31     64:6E:EA:9E:CF:FA 
0/1/2     Up       Enable  ISKT019ECEC8            10.5.101.211    64:6E:EA:9E:CE:C8 
0/1/3     Up       Enable  ISKT019ECE6C            10.5.101.71     64:6E:EA:9E:CE:6C 
0/1/4     Up       Enable  ISKT019ECE4A            10.5.101.68     64:6E:EA:9E:CE:4A 
0/1/5     Up       Enable  ISKT020F069A            10.5.101.77     F0:68:65:0F:06:9A 
0/1/6     Up       Enable  ISKT61A54F30            10.5.101.74     64:6E:EA:A5:4F:30 
0/2/1     Up       Enable  ISKT019ECF4A            10.5.101.62     64:6E:EA:9E:CF:4A 
0/3/1     Up       Enable  ISKT019ED0B2            10.5.101.112    64:6E:EA:9E:D0:B2 
0/3/2     Up       Enable  ISKT019ECF0A            10.5.101.118    64:6E:EA:9E:CF:0A 
0/4/1     Up       Enable  ISKT019ECF14            10.5.101.104    64:6E:EA:9E:CF:14 
0/4/2     Up       Enable  ISKT019ED03A            10.5.101.103    64:6E:EA:9E:D0:3A 
```

# Example Output
```sh
ISKT019ECFFA
ISKT019ECEC8
ISKT019ECE6C
ISKT019ECE4A
ISKT020F069A
ISKT61A54F30
ISKT019ECF4A
ISKT019ED0B2
ISKT019ECF0A
ISKT019ECF14
ISKT019ED03A
```
