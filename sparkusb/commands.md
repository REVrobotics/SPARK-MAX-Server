# Protocol Documentation
<a name="top"/>

## Table of Contents

- [commands.proto](#commands.proto)
    - [RequestWire](#sparkusb.RequestWire)
    - [ResponseWire](#sparkusb.ResponseWire)
    - [addressRequest](#sparkusb.addressRequest)
    - [addressResponse](#sparkusb.addressResponse)
    - [commandLineRequest](#sparkusb.commandLineRequest)
    - [commandLineResponse](#sparkusb.commandLineResponse)
    - [controlRequest](#sparkusb.controlRequest)
    - [controlResponse](#sparkusb.controlResponse)
    - [firmwareRequest](#sparkusb.firmwareRequest)
    - [firmwareResponse](#sparkusb.firmwareResponse)
    - [heartbeatRequest](#sparkusb.heartbeatRequest)
    - [listRequest](#sparkusb.listRequest)
    - [listResponse](#sparkusb.listResponse)
    - [parameterListRequest](#sparkusb.parameterListRequest)
    - [parameterListResponse](#sparkusb.parameterListResponse)
    - [parameterRequest](#sparkusb.parameterRequest)
    - [parameterResponse](#sparkusb.parameterResponse)
    - [rootCommand](#sparkusb.rootCommand)
    - [rootResponse](#sparkusb.rootResponse)
    - [setpointRequest](#sparkusb.setpointRequest)
    - [setpointResponse](#sparkusb.setpointResponse)
  
    - [configParam](#sparkusb.configParam)
    - [controlMessage](#sparkusb.controlMessage)
    - [ctrlType](#sparkusb.ctrlType)
    - [faults](#sparkusb.faults)
    - [idleMode](#sparkusb.idleMode)
    - [motorType](#sparkusb.motorType)
    - [paramType](#sparkusb.paramType)
    - [sensorType](#sparkusb.sensorType)
    - [stickyFaults](#sparkusb.stickyFaults)
  
  
    - [sparkusb](#sparkusb.sparkusb)
  

- [Scalar Value Types](#scalar-value-types)



<a name="commands.proto"/>
<p align="right"><a href="#top">Top</a></p>

## commands.proto



<a name="sparkusb.RequestWire"/>

### RequestWire
Data format to send over 0mq containing one request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| cmdLine | [commandLineRequest](#sparkusb.commandLineRequest) |  |  |
| list | [listRequest](#sparkusb.listRequest) |  |  |
| firmware | [firmwareRequest](#sparkusb.firmwareRequest) |  |  |
| heartbeat | [heartbeatRequest](#sparkusb.heartbeatRequest) |  |  |
| address | [addressRequest](#sparkusb.addressRequest) |  |  |
| parameter | [parameterRequest](#sparkusb.parameterRequest) |  |  |
| parameterList | [parameterListRequest](#sparkusb.parameterListRequest) |  |  |
| setpoint | [setpointRequest](#sparkusb.setpointRequest) |  |  |
| control | [controlRequest](#sparkusb.controlRequest) |  |  |






<a name="sparkusb.ResponseWire"/>

### ResponseWire
Data format to recieve over 0mq containting one response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |
| cmdLine | [commandLineResponse](#sparkusb.commandLineResponse) |  |  |
| list | [listResponse](#sparkusb.listResponse) |  |  |
| firmware | [firmwareResponse](#sparkusb.firmwareResponse) |  |  |
| address | [addressResponse](#sparkusb.addressResponse) |  |  |
| parameter | [parameterResponse](#sparkusb.parameterResponse) |  |  |
| parameterlist | [parameterListResponse](#sparkusb.parameterListResponse) |  |  |
| setpoint | [setpointResponse](#sparkusb.setpointResponse) |  |  |
| control | [controlResponse](#sparkusb.controlResponse) |  |  |






<a name="sparkusb.addressRequest"/>

### addressRequest
Request format for address() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| address | [uint32](#uint32) |  |  |






<a name="sparkusb.addressResponse"/>

### addressResponse
Response format for address() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentAddress | [uint32](#uint32) |  |  |
| previousAddress | [uint32](#uint32) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.commandLineRequest"/>

### commandLineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stdin | [string](#string) |  |  |






<a name="sparkusb.commandLineResponse"/>

### commandLineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stderr | [string](#string) |  |  |
| stdout | [string](#string) |  |  |






<a name="sparkusb.controlRequest"/>

### controlRequest
Request format for connect() disconnect() and ping()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ctrl | [controlMessage](#sparkusb.controlMessage) |  |  |
| device | [string](#string) |  |  |






<a name="sparkusb.controlResponse"/>

### controlResponse
Response format for connect() disconnect() and ping()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connected | [bool](#bool) |  |  |






<a name="sparkusb.firmwareRequest"/>

### firmwareRequest
Request format for firmware() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| filename | [string](#string) |  |  |






<a name="sparkusb.firmwareResponse"/>

### firmwareResponse
Response format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.heartbeatRequest"/>

### heartbeatRequest
Request format for heartbeat() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| enable | [bool](#bool) |  |  |






<a name="sparkusb.listRequest"/>

### listRequest
Request format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| all | [bool](#bool) |  |  |






<a name="sparkusb.listResponse"/>

### listResponse
Response format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deviceList | [string](#string) | repeated |  |
| deviceDetails | [string](#string) | repeated |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.parameterListRequest"/>

### parameterListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |






<a name="sparkusb.parameterListResponse"/>

### parameterListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parameter | [string](#string) | repeated |  |
| type | [paramType](#sparkusb.paramType) | repeated |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.parameterRequest"/>

### parameterRequest
Request type for Set/Get Parameter()
value is not set to signify a &#39;Get&#39; command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| parameter | [configParam](#sparkusb.configParam) |  |  |
| value | [string](#string) |  |  |






<a name="sparkusb.parameterResponse"/>

### parameterResponse
Response type for Set/Get Parameter()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |
| type | [paramType](#sparkusb.paramType) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.rootCommand"/>

### rootCommand
All RPC services implement this request
keepalive and help are not implemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [string](#string) |  |  |
| keepalive | [bool](#bool) |  |  |
| help | [bool](#bool) |  |  |






<a name="sparkusb.rootResponse"/>

### rootResponse
All RPC services implement this response
helpString not implemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| helpString | [string](#string) |  |  |
| error | [string](#string) |  |  |






<a name="sparkusb.setpointRequest"/>

### setpointRequest
Request format for Setpoint() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| setpoint | [float](#float) |  |  |
| enable | [bool](#bool) |  |  |






<a name="sparkusb.setpointResponse"/>

### setpointResponse
Response format for Setpoint() command
isRunning is not implemented yet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setpoint | [float](#float) |  |  |
| isRunning | [bool](#bool) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |





 


<a name="sparkusb.configParam"/>

### configParam
Parameter ID for set/get parameter fields
these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| CanID | 0 |  |
| InputMode | 1 |  |
| MotorType | 2 |  |
| CommAdv | 3 |  |
| SensorType | 4 |  |
| CtrlType | 5 |  |
| IdleMode | 6 |  |
| InputDeadband | 7 |  |
| FirmwareVersion | 8 |  |



<a name="sparkusb.controlMessage"/>

### controlMessage
Types of control message to send

| Name | Number | Description |
| ---- | ------ | ----------- |
| controlPing | 0 |  |
| controlConnect | 1 |  |
| controlDisconnect | 2 |  |



<a name="sparkusb.ctrlType"/>

### ctrlType
Control type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| DutyCycle | 0 |  |
| Velocity | 1 |  |



<a name="sparkusb.faults"/>

### faults
Faults type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Brownout | 0 |  |
| Overcurrent | 1 |  |
| Overvoltage | 2 |  |
| MotorFault | 3 |  |
| SensorFault | 4 |  |
| Stall | 5 |  |
| EEPROMCRC | 6 |  |



<a name="sparkusb.idleMode"/>

### idleMode
Idle mode type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Coast | 0 |  |
| Brake | 1 |  |



<a name="sparkusb.motorType"/>

### motorType
Motor type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Brushed | 0 |  |
| Brushless | 1 |  |



<a name="sparkusb.paramType"/>

### paramType
Parameter type enum, these values match the SPARK firmware
and are sent as a response in GetParameter() requests

| Name | Number | Description |
| ---- | ------ | ----------- |
| int32 | 0 |  |
| uint32 | 1 |  |
| float32 | 2 |  |



<a name="sparkusb.sensorType"/>

### sensorType
Sensor type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| HallSensor | 0 |  |
| Encoder | 1 |  |
| Sensorless | 2 |  |



<a name="sparkusb.stickyFaults"/>

### stickyFaults
Sticky type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| BrownoutSticky | 0 |  |
| OvercurrentSticky | 1 |  |
| OvervoltageSticky | 2 |  |
| MotorFaultSticky | 3 |  |
| SensorFaultSticky | 4 |  |
| StallSticky | 5 |  |
| EEPROMCRCSticky | 6 |  |


 

 


<a name="sparkusb.sparkusb"/>

### sparkusb
Interface functions for service sparkusb.
All command requests are serialized into a 
RequestWire type before transmission, and
Deserializezd to a ResponseWire on recipt

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Connect | [controlRequest](#sparkusb.controlRequest) | [controlResponse](#sparkusb.controlRequest) | Connect to the requested serial port. If no port is provided, connect to the default port, which is the first port found with a SPARK device. |
| Disconnect | [controlRequest](#sparkusb.controlRequest) | [controlResponse](#sparkusb.controlRequest) | Disconnect the serial port |
| Ping | [controlRequest](#sparkusb.controlRequest) | [controlResponse](#sparkusb.controlRequest) | Ping the SPARK to verify connection to the hardware and to this server. |
| List | [listRequest](#sparkusb.listRequest) | [listResponse](#sparkusb.listRequest) | List the serial port information for all connected SPARK devices. |
| Heartbeat | [heartbeatRequest](#sparkusb.heartbeatRequest) | [rootResponse](#sparkusb.heartbeatRequest) | Send a Heartbeat to the SPARK device (similar to how the roboRIO heartbeat to the device over CAN). This message can be simplified if driving the motor by calling Setpoint({Enable: true}) |
| SetParameter | [parameterRequest](#sparkusb.parameterRequest) | [parameterResponse](#sparkusb.parameterRequest) | Set a device parameter. The parameter should be configParam type the value is a string in both the request and response. |
| GetParameter | [parameterRequest](#sparkusb.parameterRequest) | [parameterResponse](#sparkusb.parameterRequest) | Get a device parameter. The parameter should be configParam type the value returned is a string in both the request and response. The requested value type is also passed to help decode. The type is of type paramType |
| BurnFlash | [rootCommand](#sparkusb.rootCommand) | [rootResponse](#sparkusb.rootCommand) | Make all configuration changes permanent for the next time the device powers on. Note: This writes any values that have changed and can only be called when the device is not enabled. Since this method writes directly to FLASH, avoid calling frequently, as each flash location can be written to a total of 10,000 times in its lifetime. Flash wear leveling is being implemented and should be in the release before kickoff |
| Setpoint | [setpointRequest](#sparkusb.setpointRequest) | [setpointResponse](#sparkusb.setpointRequest) | Send a setpoint command. Right now the value should be from 1023 to -1024 however this will change to native units (&#43;/- 1.0 for duty cycle control). Setting Enable = true will also send a heartbeat allowing the controller drive the motor. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

