commands.proto contains the message details to control the spark remotely. To build:

**GO**

Install protoc compiler then:

`go get -u github.com/golang/protobuf/protoc-gen-go`

`protoc -I=./ --go_out=./  ./commands.proto`

**Other Issues**

On Linux machines with ModemManager installed, the device may show up as busy for ~15 seconds while modem manager tries to decide if its a modem. To disable this add a custom UDEV rule, [here](https://linux-tips.com/t/prevent-modem-manager-to-capture-usb-serial-devices/284) is a detailed description. **VID** = 0483 **PID** = 5740

For example:

*/etc/udev/rules.d/99-ttyacms.rules* has the line:

`ATTRS{idVendor}=="0483" ATTRS{idProduct}=="5740", ENV{ID_MM_DEVICE_IGNORE}="1"`
