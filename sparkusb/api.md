# Protocol Documentation
<a name="top"></a>

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



<a name="commands.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## commands.proto



<a name="sparkusb.RequestWire"></a>

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






<a name="sparkusb.ResponseWire"></a>

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






<a name="sparkusb.addressRequest"></a>

### addressRequest
Request format for address() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| address | [uint32](#uint32) |  |  |






<a name="sparkusb.addressResponse"></a>

### addressResponse
Response format for address() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentAddress | [uint32](#uint32) |  |  |
| previousAddress | [uint32](#uint32) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.commandLineRequest"></a>

### commandLineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stdin | [string](#string) |  |  |






<a name="sparkusb.commandLineResponse"></a>

### commandLineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stderr | [string](#string) |  |  |
| stdout | [string](#string) |  |  |






<a name="sparkusb.controlRequest"></a>

### controlRequest
Request format for connect() disconnect() and ping()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ctrl | [controlMessage](#sparkusb.controlMessage) |  |  |
| device | [string](#string) |  |  |






<a name="sparkusb.controlResponse"></a>

### controlResponse
Response format for connect() disconnect() and ping()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connected | [bool](#bool) |  |  |






<a name="sparkusb.firmwareRequest"></a>

### firmwareRequest
Request format for firmware() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| filename | [string](#string) |  |  |






<a name="sparkusb.firmwareResponse"></a>

### firmwareResponse
Response format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.heartbeatRequest"></a>

### heartbeatRequest
Request format for heartbeat() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| enable | [bool](#bool) |  |  |






<a name="sparkusb.listRequest"></a>

### listRequest
Request format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| all | [bool](#bool) |  |  |






<a name="sparkusb.listResponse"></a>

### listResponse
Response format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deviceList | [string](#string) | repeated |  |
| deviceDetails | [string](#string) | repeated |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.parameterListRequest"></a>

### parameterListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |






<a name="sparkusb.parameterListResponse"></a>

### parameterListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parameter | [string](#string) | repeated |  |
| type | [paramType](#sparkusb.paramType) | repeated |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.parameterRequest"></a>

### parameterRequest
Request type for Set/Get Parameter()
value is not set to signify a &#39;Get&#39; command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| parameter | [configParam](#sparkusb.configParam) |  |  |
| value | [string](#string) |  |  |






<a name="sparkusb.parameterResponse"></a>

### parameterResponse
Response type for Set/Get Parameter()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |
| type | [paramType](#sparkusb.paramType) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |






<a name="sparkusb.rootCommand"></a>

### rootCommand
All RPC services implement this request
keepalive and help are not implemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [string](#string) |  |  |
| keepalive | [bool](#bool) |  |  |
| help | [bool](#bool) |  |  |






<a name="sparkusb.rootResponse"></a>

### rootResponse
All RPC services implement this response
helpString not implemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| helpString | [string](#string) |  |  |
| error | [string](#string) |  |  |






<a name="sparkusb.setpointRequest"></a>

### setpointRequest
Request format for Setpoint() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkusb.rootCommand) |  |  |
| setpoint | [float](#float) |  |  |
| enable | [bool](#bool) |  |  |






<a name="sparkusb.setpointResponse"></a>

### setpointResponse
Response format for Setpoint() command
isRunning is not implemented yet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setpoint | [float](#float) |  |  |
| isRunning | [bool](#bool) |  |  |
| root | [rootResponse](#sparkusb.rootResponse) |  |  |





 


<a name="sparkusb.configParam"></a>

### configParam
Parameter ID for set/get parameter fields
these values match the SPARK firmware
@omit THIS ENUM IS AUTO GENERATED FROM SPREADSHEET

| Name | Number | Description |
| ---- | ------ | ----------- |
| kCanID | 0 | CAN ID @default 0 @type uint |
| kInputMode | 1 | Input mode, this parameter is read only and the input mode is detected by t he firmware automatically, results are %Input Mode% @default 0 @type Input Mode |
| kMotorType | 2 | Motor Type, options are %Motor Type% @default BRUSHED @type Motor Type |
| kCommAdvance | 3 | Electrical degree of offset from the backemf crossing to commutation for se nsorless modes. This is currently not implemented @default 0 @type float32 @unit Degrees |
| kSensorType | 4 | Sensor Type, options are %Sensor Type% @default HALL_SENSOR @type Sensor Type |
| kCtrlType | 5 | Control Type, this is a read only parameter of the currently active control type. Options are %Ctrl Type% @default CTRL_DUTY_CYCLE @type Ctrl Type |
| kIdleMode | 6 | State of the half bridge when the motor controller commands zero output or is disabled. Options are %Idle Mode% @default IDLE_COAST @type Idle Mode |
| kInputDeadband | 7 | Percent of the input which results in zero output @default 0.05 @type float32 @unit Percent |
| kFirmwareVer | 8 | Read only parameter showing the 32-bit firmware version. The first byte is the major build, the second byte is in the minor build, the last two bytes are the build @default 0 @type Firmware |
| kHallOffset | 9 | Electrical offset of the hall sensor compared to the motor phases in degree s. Typically this is either 0, 60, 120 @default 0 @type int @unit Degrees |
| kPolePairs | 10 | Number of pole pairs for the brushless motor. This is the number of poles/2 and can be determined by either counting the number of magents or countin g the number of windings and dividing by 3. This is an important term for speed regulation to properly calculate the speed @default 1 @type uint |
| kCurrentChop | 11 | If the half bridge detects this current limit, it will disable the motor dr iver until the current goes below a certain threshold before enabling. Thi s is a low sophistication &#39;current control&#39; @default 0 @type float32 @unit Amps |
| kCurrentLimit | 12 | If the half bridge detects this current limit, it will disable the motor dr iver and report a sticky fault. The motor driver will not enable again unt il this fault is cleared. @default 0 @type float32 @unit Amps |
| kP_0 | 13 | Perportional gain constant for gain slot 0. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kI_0 | 14 | Integral gain constant for gain slot 0. In cascaded control modes this is t he inner loop gain slot 0. @default 0 @type float32 |
| kD_0 | 15 | Derivative gain constant for gain slot 0. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kF_0 | 16 | Feed Forward gain constant for gain slot 0. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kIZone_0 | 17 | Integrator zone constant for gain slot 0. The PIDF loop integrator will onl y accumulate while the setpoint is within IZone of the target. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kDFilter_0 | 18 | PIDF derivative filter constant for gain slot 0. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kOutputMin_0 | 19 | Max output constant for gain slot 0. This is the max output of the controll er as well as the max integrator value. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kOutputMax_0 | 20 | Min output constant for gain slot 0. This is the min output of the controll er as well as the min integrator value. In cascaded control modes this is the inner loop gain slot 0. @default 0 @type float32 |
| kP_1 | 21 | Perportional gain constant for gain slot 1. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kI_1 | 22 | Integral gain constant for gain slot 1. In cascaded control modes this is t he inner loop gain slot 1. @default 0 @type float32 |
| kD_1 | 23 | Derivative gain constant for gain slot 1. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kF_1 | 24 | Feed Forward gain constant for gain slot 1. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kIZone_1 | 25 | Integrator zone constant for gain slot 1. The PIDF loop integrator will onl y accumulate while the setpoint is within IZone of the target. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kDFilter_1 | 26 | PIDF derivative filter constant for gain slot 1. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kOutputMin_1 | 27 | Max output constant for gain slot 1. This is the max output of the controll er as well as the max integrator value. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kOutputMax_1 | 28 | Min output constant for gain slot 1. This is the min output of the controll er as well as the min integrator value. In cascaded control modes this is the inner loop gain slot 1. @default 0 @type float32 |
| kP_2 | 29 | Perportional gain constant for gain slot 3. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kI_2 | 30 | Integral gain constant for gain slot 3. In cascaded control modes this is t he outer loop gain slot 0. @default 0 @type float32 |
| kD_2 | 31 | Derivative gain constant for gain slot 3. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kF_2 | 32 | Feed Forward gain constant for gain slot 3. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kIZone_2 | 33 | Integrator zone constant for gain slot 3. The PIDF loop integrator will onl y accumulate while the setpoint is within IZone of the target. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kDFilter_2 | 34 | PIDF derivative filter constant for gain slot 3. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kOutputMin_2 | 35 | Max output constant for gain slot 3. This is the max output of the controll er as well as the max integrator value. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kOutputMax_2 | 36 | Min output constant for gain slot 3. This is the min output of the controll er as well as the min integrator value. In cascaded control modes this is the outer loop gain slot 0. @default 0 @type float32 |
| kP_3 | 37 | Perportional gain constant for gain slot 4. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kI_3 | 38 | Integral gain constant for gain slot 4. In cascaded control modes this is t he outer loop gain slot 1. @default 0 @type float32 |
| kD_3 | 39 | Derivative gain constant for gain slot 4. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kF_3 | 40 | Feed Forward gain constant for gain slot 4. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kIZone_3 | 41 | Integrator zone constant for gain slot 4. The PIDF loop integrator will onl y accumulate while the setpoint is within IZone of the target. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kDFilter_3 | 42 | PIDF derivative filter constant for gain slot 4. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kOutputMin_3 | 43 | Max output constant for gain slot 4. This is the max output of the controll er as well as the max integrator value. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kOutputMax_3 | 44 | Min output constant for gain slot 4. This is the min output of the controll er as well as the min integrator value. In cascaded control modes this is the outer loop gain slot 1. @default 0 @type float32 |
| kReserved | 45 | Reserved @default 0 @type uint |
| kOutputRatio | 46 | Simple scalar for all units in all closed loop control modes to scale units to native. Use this to scale the output to things like gear ratios or uni t conversions @default 1 @type float32 |
| kSerialNumberLow | 47 | Low 32-bits of unique 96-bit serial number @default 0 @type uint |
| kSerialNumberMid | 48 | Middle 32-bits of unique 96-bit serial number @default 0 @type uint |
| kSerialNumberHigh | 49 | High 32-bits of unique 96-bit serial number @default 0 @type uint |
| kLimitSwitchFwdPolarity | 50 | Limit switch polarity. Default is Normally Open (1), and can be set to Norm ally Closed (0) @default 1 @type bool |
| kLimitSwitchRevPolarity | 51 | Limit switch polarity. Default is Normally Open (1), and can be set to Norm ally Closed (0) @default 1 @type bool |
| kHardLimitFwdEn | 52 | Limit switch enable, disabled by default @default 0 @type bool |
| kHardLimitRevEn | 53 | Limit switch enable, disabled by default @default 0 @type bool |
| kSoftLimitFwdEn | 54 | Soft limit enable, disabled by default @default 0 @type bool |
| kSoftLimitRevEn | 55 | Soft limit enable, disabled by default @default 0 @type bool |
| kRampRate | 56 | Voltage ramp rate active for all control modes in V/s, a value of 0 disable s this feature @default 0 @type float32 @unit V/s |



<a name="sparkusb.controlMessage"></a>

### controlMessage
Types of control message to send

| Name | Number | Description |
| ---- | ------ | ----------- |
| controlPing | 0 |  |
| controlConnect | 1 |  |
| controlDisconnect | 2 |  |



<a name="sparkusb.ctrlType"></a>

### ctrlType
Control type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| DutyCycle | 0 |  |
| Velocity | 1 |  |



<a name="sparkusb.faults"></a>

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



<a name="sparkusb.idleMode"></a>

### idleMode
Idle mode type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Coast | 0 |  |
| Brake | 1 |  |



<a name="sparkusb.motorType"></a>

### motorType
Motor type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Brushed | 0 |  |
| Brushless | 1 |  |



<a name="sparkusb.paramType"></a>

### paramType
Parameter type enum, these values match the SPARK firmware
and are sent as a response in GetParameter() requests

| Name | Number | Description |
| ---- | ------ | ----------- |
| int32 | 0 |  |
| uint32 | 1 |  |
| float32 | 2 |  |



<a name="sparkusb.sensorType"></a>

### sensorType
Sensor type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| HallSensor | 0 |  |
| Encoder | 1 |  |
| Sensorless | 2 |  |



<a name="sparkusb.stickyFaults"></a>

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


 

 


<a name="sparkusb.sparkusb"></a>

### sparkusb
Interface functions for service sparkusb.
All command requests are serialized into a 
RequestWire type before transmission, and
Deserializezd to a ResponseWire on recipt

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Connect | [controlRequest](#sparkusb.controlRequest) | [controlResponse](#sparkusb.controlResponse) | Connect to the requested serial port. If no port is provided, connect to the default port, which is the first port found with a SPARK device. |
| Disconnect | [controlRequest](#sparkusb.controlRequest) | [controlResponse](#sparkusb.controlResponse) | Disconnect the serial port |
| Ping | [controlRequest](#sparkusb.controlRequest) | [controlResponse](#sparkusb.controlResponse) | Ping the SPARK to verify connection to the hardware and to this server. |
| List | [listRequest](#sparkusb.listRequest) | [listResponse](#sparkusb.listResponse) | List the serial port information for all connected SPARK devices. |
| Heartbeat | [heartbeatRequest](#sparkusb.heartbeatRequest) | [rootResponse](#sparkusb.rootResponse) | Send a Heartbeat to the SPARK device (similar to how the roboRIO heartbeat to the device over CAN). This message can be simplified if driving the motor by calling Setpoint({Enable: true}) |
| SetParameter | [parameterRequest](#sparkusb.parameterRequest) | [parameterResponse](#sparkusb.parameterResponse) | Set a device parameter. The parameter should be configParam type the value is a string in both the request and response. |
| GetParameter | [parameterRequest](#sparkusb.parameterRequest) | [parameterResponse](#sparkusb.parameterResponse) | Get a device parameter. The parameter should be configParam type the value returned is a string in both the request and response. The requested value type is also passed to help decode. The type is of type paramType |
| BurnFlash | [rootCommand](#sparkusb.rootCommand) | [rootResponse](#sparkusb.rootResponse) | Make all configuration changes permanent for the next time the device powers on. Note: This writes any values that have changed and can only be called when the device is not enabled. Since this method writes directly to FLASH, avoid calling frequently, as each flash location can be written to a total of 10,000 times in its lifetime. Flash wear leveling is being implemented and should be in the release before kickoff |
| Setpoint | [setpointRequest](#sparkusb.setpointRequest) | [setpointResponse](#sparkusb.setpointResponse) | Send a setpoint command. Right now the value should be from 1023 to -1024 however this will change to native units (&#43;/- 1.0 for duty cycle control). Setting Enable = true will also send a heartbeat allowing the controller drive the motor. |

 



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

