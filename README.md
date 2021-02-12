# stripSn
Utility program for stripping Serial Number from ONU Blacklist. Includes file or terminal i/o options.

`stripSn` [options]
  -a    Append entries to file if exists (default true)
  -h    Show this help
  -i string
        File to read blacklist entries from (default "blackList.txt")
  -o string
        File to write new serial numbers to (default "authList.txt")
  -stdin
        Read blacklist entries from stdin
  -stdout
        Write serial numbers to stdout


Example Input (LindsayBB#show onu black-list):
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
Total: 14

Example Output:
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


