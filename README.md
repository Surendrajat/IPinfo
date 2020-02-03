# IPinfo

Get info about your Public IP, Geo location and VPN/Proxy detection.

## Install

- Clone the repo
  
  ```bash
  git clone https://github.com/surendrajat/ipinfo
  ```

- Build and install the binary

  ```bash
  cd ipinfo && go install .
  ```

## Run `ipinfo`

```bash
➜  ~ ipinfo
IP	124.123.122.111
Geo	📍 Manglore, Karnataka, IN
ASN	[ISP: Atria Convergence Technologies Pvt. Ltd.(actcorp.in)]
Org	[BUSINESS: Beam Telecom Pvt Ltd, Hyderabad(beamtele.com)]
VPN	🔓 none
➜  ~ nordvpn connect
Connecting to India #29 (in29.nordvpn.com)
You are connected to India #29 (in29.nordvpn.com)!
➜  ~ ipinfo
IP	152.34.212.189
Geo	📍 Bengaluru, Karnataka, IN
ASN	[HOSTING: SoftLayer Technologies Inc.(softlayer.com)]
Org	[BUSINESS: TEFINCOM S.A.(nordvpn.com)]
VPN	🔒 ipsec
➜  ~
```
