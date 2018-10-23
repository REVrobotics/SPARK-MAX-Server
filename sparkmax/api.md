# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [SPARK-MAX-Commands.proto](#SPARK-MAX-Commands.proto)
    - [RequestWire](#sparkmax.RequestWire)
    - [ResponseWire](#sparkmax.ResponseWire)
    - [addressRequest](#sparkmax.addressRequest)
    - [addressResponse](#sparkmax.addressResponse)
    - [burnRequest](#sparkmax.burnRequest)
    - [burnResponse](#sparkmax.burnResponse)
    - [commandLineRequest](#sparkmax.commandLineRequest)
    - [commandLineResponse](#sparkmax.commandLineResponse)
    - [connectRequest](#sparkmax.connectRequest)
    - [connectResponse](#sparkmax.connectResponse)
    - [disconnectRequest](#sparkmax.disconnectRequest)
    - [disconnectResponse](#sparkmax.disconnectResponse)
    - [firmwareRequest](#sparkmax.firmwareRequest)
    - [firmwareResponse](#sparkmax.firmwareResponse)
    - [followerRequest](#sparkmax.followerRequest)
    - [getParameterRequest](#sparkmax.getParameterRequest)
    - [heartbeatRequest](#sparkmax.heartbeatRequest)
    - [listRequest](#sparkmax.listRequest)
    - [listResponse](#sparkmax.listResponse)
    - [parameterListRequest](#sparkmax.parameterListRequest)
    - [parameterListResponse](#sparkmax.parameterListResponse)
    - [parameterResponse](#sparkmax.parameterResponse)
    - [pingRequest](#sparkmax.pingRequest)
    - [pingResponse](#sparkmax.pingResponse)
    - [rootCommand](#sparkmax.rootCommand)
    - [rootResponse](#sparkmax.rootResponse)
    - [setParameterRequest](#sparkmax.setParameterRequest)
    - [setpointRequest](#sparkmax.setpointRequest)
    - [setpointResponse](#sparkmax.setpointResponse)
  
  
  
    - [sparkMaxServer](#sparkmax.sparkMaxServer)
  

- [SPARK-MAX-Types.proto](#SPARK-MAX-Types.proto)
  
    - [configParam](#sparkmax.configParam)
    - [configParamTypes](#sparkmax.configParamTypes)
    - [ctrlType](#sparkmax.ctrlType)
    - [definedFollowerID](#sparkmax.definedFollowerID)
    - [faults](#sparkmax.faults)
    - [followerSignMode](#sparkmax.followerSignMode)
    - [idleMode](#sparkmax.idleMode)
    - [inputMode](#sparkmax.inputMode)
    - [motorType](#sparkmax.motorType)
    - [paramStatus](#sparkmax.paramStatus)
    - [paramType](#sparkmax.paramType)
    - [sensorType](#sparkmax.sensorType)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="SPARK-MAX-Commands.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## SPARK-MAX-Commands.proto



<a name="sparkmax.RequestWire"></a>

### RequestWire
Data format to send over 0mq containing one request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [listRequest](#sparkmax.listRequest) |  |  |
| firmware | [firmwareRequest](#sparkmax.firmwareRequest) |  |  |
| heartbeat | [heartbeatRequest](#sparkmax.heartbeatRequest) |  |  |
| address | [addressRequest](#sparkmax.addressRequest) |  |  |
| setParameter | [setParameterRequest](#sparkmax.setParameterRequest) |  |  |
| getParameter | [getParameterRequest](#sparkmax.getParameterRequest) |  |  |
| setpoint | [setpointRequest](#sparkmax.setpointRequest) |  |  |
| connect | [connectRequest](#sparkmax.connectRequest) |  |  |
| disconnect | [disconnectRequest](#sparkmax.disconnectRequest) |  |  |
| ping | [pingRequest](#sparkmax.pingRequest) |  |  |
| follower | [followerRequest](#sparkmax.followerRequest) |  |  |
| burn | [burnRequest](#sparkmax.burnRequest) |  |  |
| parameterList | [parameterListRequest](#sparkmax.parameterListRequest) |  |  |






<a name="sparkmax.ResponseWire"></a>

### ResponseWire
Data format to recieve over 0mq containting one response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |
| list | [listResponse](#sparkmax.listResponse) |  |  |
| firmware | [firmwareResponse](#sparkmax.firmwareResponse) |  |  |
| address | [addressResponse](#sparkmax.addressResponse) |  |  |
| parameter | [parameterResponse](#sparkmax.parameterResponse) |  |  |
| parameterlist | [parameterListResponse](#sparkmax.parameterListResponse) |  |  |
| setpoint | [setpointResponse](#sparkmax.setpointResponse) |  |  |
| connect | [connectResponse](#sparkmax.connectResponse) |  |  |
| disconnect | [disconnectResponse](#sparkmax.disconnectResponse) |  |  |
| ping | [pingResponse](#sparkmax.pingResponse) |  |  |
| burn | [burnResponse](#sparkmax.burnResponse) |  |  |






<a name="sparkmax.addressRequest"></a>

### addressRequest
Request format for address() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| address | [uint32](#uint32) |  |  |






<a name="sparkmax.addressResponse"></a>

### addressResponse
Response format for address() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentAddress | [uint32](#uint32) |  |  |
| previousAddress | [uint32](#uint32) |  |  |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |






<a name="sparkmax.burnRequest"></a>

### burnRequest
Burn command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| verify | [bool](#bool) |  |  |






<a name="sparkmax.burnResponse"></a>

### burnResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |
| verify | [bool](#bool) |  |  |






<a name="sparkmax.commandLineRequest"></a>

### commandLineRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stdin | [string](#string) |  |  |






<a name="sparkmax.commandLineResponse"></a>

### commandLineResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stderr | [string](#string) |  |  |
| stdout | [string](#string) |  |  |






<a name="sparkmax.connectRequest"></a>

### connectRequest
Request format for connect()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [string](#string) |  |  |






<a name="sparkmax.connectResponse"></a>

### connectResponse
Response format for connect()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connected | [bool](#bool) |  |  |






<a name="sparkmax.disconnectRequest"></a>

### disconnectRequest
Request format for disconnect()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [string](#string) |  |  |






<a name="sparkmax.disconnectResponse"></a>

### disconnectResponse
Response format for disconnect()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connected | [bool](#bool) |  |  |






<a name="sparkmax.firmwareRequest"></a>

### firmwareRequest
Request format for firmware() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| filename | [string](#string) |  |  |






<a name="sparkmax.firmwareResponse"></a>

### firmwareResponse
Response format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |






<a name="sparkmax.followerRequest"></a>

### followerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| followerid | [uint32](#uint32) |  |  |
| followerconfig | [uint32](#uint32) |  |  |






<a name="sparkmax.getParameterRequest"></a>

### getParameterRequest
Request type for Get Parameter()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| parameter | [configParam](#sparkmax.configParam) |  |  |






<a name="sparkmax.heartbeatRequest"></a>

### heartbeatRequest
Request format for heartbeat() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| enable | [bool](#bool) |  |  |






<a name="sparkmax.listRequest"></a>

### listRequest
Request format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| all | [bool](#bool) |  |  |






<a name="sparkmax.listResponse"></a>

### listResponse
Response format for list() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deviceList | [string](#string) | repeated |  |
| deviceDetails | [string](#string) | repeated |  |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |






<a name="sparkmax.parameterListRequest"></a>

### parameterListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |






<a name="sparkmax.parameterListResponse"></a>

### parameterListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parameter | [string](#string) | repeated |  |
| type | [paramType](#sparkmax.paramType) | repeated |  |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |






<a name="sparkmax.parameterResponse"></a>

### parameterResponse
Response type for Set/Get Parameter()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |
| type | [paramType](#sparkmax.paramType) |  |  |
| status | [paramStatus](#sparkmax.paramStatus) |  |  |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |






<a name="sparkmax.pingRequest"></a>

### pingRequest
Response format for ping()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [string](#string) |  |  |






<a name="sparkmax.pingResponse"></a>

### pingResponse
Response format for ping()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |
| connected | [bool](#bool) |  |  |






<a name="sparkmax.rootCommand"></a>

### rootCommand
All RPC services implement this request
keepalive and help are not implemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [string](#string) |  |  |
| keepalive | [bool](#bool) |  |  |
| help | [bool](#bool) |  |  |






<a name="sparkmax.rootResponse"></a>

### rootResponse
All RPC services implement this response
helpString not implemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| helpString | [string](#string) |  |  |
| error | [string](#string) |  |  |






<a name="sparkmax.setParameterRequest"></a>

### setParameterRequest
Request type for Set Parameter()


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| parameter | [configParam](#sparkmax.configParam) |  |  |
| value | [string](#string) |  |  |






<a name="sparkmax.setpointRequest"></a>

### setpointRequest
Request format for Setpoint() command


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [rootCommand](#sparkmax.rootCommand) |  |  |
| setpoint | [float](#float) |  |  |
| enable | [bool](#bool) |  |  |






<a name="sparkmax.setpointResponse"></a>

### setpointResponse
Response format for Setpoint() command
isRunning is not implemented yet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setpoint | [float](#float) |  |  |
| isRunning | [bool](#bool) |  |  |
| root | [rootResponse](#sparkmax.rootResponse) |  |  |





 

 

 


<a name="sparkmax.sparkMaxServer"></a>

### sparkMaxServer
Interface functions for service sparkmax.
All command requests are serialized into a 
RequestWire type before transmission, and
Deserializezd to a ResponseWire on recipt

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Connect | [connectRequest](#sparkmax.connectRequest) | [connectResponse](#sparkmax.connectResponse) | Connect to the requested serial port. If no port is provided, connect to the default port, which is the first port found with a SPARK device. |
| Disconnect | [disconnectRequest](#sparkmax.disconnectRequest) | [disconnectResponse](#sparkmax.disconnectResponse) | Disconnect the serial port |
| Ping | [pingRequest](#sparkmax.pingRequest) | [pingResponse](#sparkmax.pingResponse) | Ping the SPARK to verify connection to the hardware and to this server. |
| List | [listRequest](#sparkmax.listRequest) | [listResponse](#sparkmax.listResponse) | List the serial port information for all connected SPARK devices. |
| Firmware | [firmwareRequest](#sparkmax.firmwareRequest) | [firmwareResponse](#sparkmax.firmwareResponse) | Update the firmware of the device |
| SetParameter | [setParameterRequest](#sparkmax.setParameterRequest) | [parameterResponse](#sparkmax.parameterResponse) | Set a device parameter. The parameter should be configParam type the value is a string in both the request and response. |
| GetParameter | [getParameterRequest](#sparkmax.getParameterRequest) | [parameterResponse](#sparkmax.parameterResponse) | Get a device parameter. The parameter should be configParam type the value returned is a string in both the request and response. The requested value type is also passed to help decode. The type is of type paramType |
| BurnFlash | [burnRequest](#sparkmax.burnRequest) | [burnResponse](#sparkmax.burnResponse) | Make all configuration changes permanent for the next time the device powers on. Note: This writes any values that have changed and can only be called when the device is not enabled. Since this method writes directly to FLASH, avoid calling frequently, as each flash location can be written to a total of 10,000 times in its lifetime. Flash wear leveling is being implemented and should be in the release before kickoff |
| Setpoint | [setpointRequest](#sparkmax.setpointRequest) | [setpointResponse](#sparkmax.setpointResponse) | Send a setpoint command. The value should be native units depending on the curernt control mode (&#43;/- 1.0 for duty cycle control) Setting Enable = true will also send a heartbeat allowing the controller drive the motor. |
| Follow | [followerRequest](#sparkmax.followerRequest) | [rootResponse](#sparkmax.rootResponse) | Set controller to follow another controller. |

 



<a name="SPARK-MAX-Types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## SPARK-MAX-Types.proto


 


<a name="sparkmax.configParam"></a>

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
| kHardLimitRevEn | 53 | sparkusb Limit switch enable, disabled by default @default 0 @type bool |
| kSoftLimitFwdEn | 54 | Soft limit enable, disabled by default @default 0 @type bool |
| kSoftLimitRevEn | 55 | Soft limit enable, disabled by default @default 0 @type bool |
| kRampRate | 56 | Voltage ramp rate active for all control modes in V/s, a value of 0 disable s this feature @default 0 @type float32 @unit V/s |
| kFollowerID | 57 | CAN EXTID of the message with data to follow @default 0 @type uint |
| kFollowerConfig | 58 | Special configuration register for setting up to follow on a repeating mess age (follower mode). CFG[0] to CFG[3] where CFG[0] is the motor output sta rt bit (LSB), CFG[1] is the motor output stop bit (MSB). CFG[0] - CFG[1] d etermines edieness. CFG[2] bits determine sign mode and inverted, CFG[3] s ets a preconfigured controller (0x1A = REV, 0x1B = Talon/Victor style as o f 2018 season) @default 0 @type uint |



<a name="sparkmax.configParamTypes"></a>

### configParamTypes
Type lookup for above, lists what type the controller expects for each parameter

| Name | Number | Description |
| ---- | ------ | ----------- |
| kDefault_t | 0 |  |
| kCanID_t | 1 | Below is auto generated |
| kInputMode_t | 1 |  |
| kMotorType_t | 1 |  |
| kCommAdvance_t | 2 |  |
| kSensorType_t | 1 |  |
| kCtrlType_t | 1 |  |
| kIdleMode_t | 1 |  |
| kInputDeadband_t | 2 |  |
| kFirmwareVer_t | 1 |  |
| kHallOffset_t | 0 |  |
| kPolePairs_t | 1 |  |
| kCurrentChop_t | 2 |  |
| kCurrentLimit_t | 2 |  |
| kP_0_t | 2 |  |
| kI_0_t | 2 |  |
| kD_0_t | 2 |  |
| kF_0_t | 2 |  |
| kIZone_0_t | 2 |  |
| kDFilter_0_t | 2 |  |
| kOutputMin_0_t | 2 |  |
| kOutputMax_0_t | 2 |  |
| kP_1_t | 2 |  |
| kI_1_t | 2 |  |
| kD_1_t | 2 |  |
| kF_1_t | 2 |  |
| kIZone_1_t | 2 |  |
| kDFilter_1_t | 2 |  |
| kOutputMin_1_t | 2 |  |
| kOutputMax_1_t | 2 |  |
| kP_2_t | 2 |  |
| kI_2_t | 2 |  |
| kD_2_t | 2 |  |
| kF_2_t | 2 |  |
| kIZone_2_t | 2 |  |
| kDFilter_2_t | 2 |  |
| kOutputMin_2_t | 2 |  |
| kOutputMax_2_t | 2 |  |
| kP_3_t | 2 |  |
| kI_3_t | 2 |  |
| kD_3_t | 2 |  |
| kF_3_t | 2 |  |
| kIZone_3_t | 2 |  |
| kDFilter_3_t | 2 |  |
| kOutputMin_3_t | 2 |  |
| kOutputMax_3_t | 2 |  |
| kReserved_t | 1 |  |
| kOutputRatio_t | 2 |  |
| kSerialNumberLow_t | 1 |  |
| kSerialNumberMid_t | 1 |  |
| kSerialNumberHigh_t | 1 |  |
| kLimitSwitchFwdPolarity_t | 3 |  |
| kLimitSwitchRevPolarity_t | 3 |  |
| kHardLimitFwdEn_t | 3 |  |
| kHardLimitRevEn_t | 3 |  |
| kSoftLimitFwdEn_t | 3 |  |
| kSoftLimitRevEn_t | 3 |  |
| kRampRate_t | 2 |  |
| kFollowerID_t | 1 |  |
| kFollowerConfig_t | 1 |  |



<a name="sparkmax.ctrlType"></a>

### ctrlType
Control type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| DutyCycle | 0 |  |
| Velocity | 1 |  |
| Voltage | 2 |  |



<a name="sparkmax.definedFollowerID"></a>

### definedFollowerID
Follower ID for pre-defined CAN protocols

| Name | Number | Description |
| ---- | ------ | ----------- |
| FollowerDisabled | 0 |  |
| FollowerCustom | 25 |  |
| FollowerSparkMax | 26 |  |
| FollowerPhoenix | 27 |  |



<a name="sparkmax.faults"></a>

### faults
Faults type enum, these valInvalidIDues match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Brownout | 0 |  |
| Overcurrent | 1 |  |
| Overvoltage | 2 |  |
| MotorFault | 3 |  |
| SensorFault | 4 |  |
| Stall | 5 |  |
| EEPROMCRC | 6 |  |
| CANTX | 7 |  |
| CANRX | 8 |  |
| HasReset | 9 |  |
| DRVFault | 10 |  |
| SoftLimitFwd | 12 |  |
| SoftLimitRev | 13 |  |
| HardLimitFwd | 14 |  |
| HardLimitRev | 15 |  |



<a name="sparkmax.followerSignMode"></a>

### followerSignMode
Follower sign mode

| Name | Number | Description |
| ---- | ------ | ----------- |
| FollowerNoSign | 0 |  |
| FollowerTwosComp | 1 |  |
| FollowerSignMag | 2 |  |



<a name="sparkmax.idleMode"></a>

### idleMode
Idle mode type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Coast | 0 |  |
| Brake | 1 |  |



<a name="sparkmax.inputMode"></a>

### inputMode
Input mode type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| PWM | 0 |  |
| CAN | 1 |  |



<a name="sparkmax.motorType"></a>

### motorType
Motor type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| Brushed | 0 |  |
| Brushless | 1 |  |



<a name="sparkmax.paramStatus"></a>

### paramStatus
Parameter Status returned from get/set parameter

| Name | Number | Description |
| ---- | ------ | ----------- |
| paramOK | 0 |  |
| InvalidID | 1 |  |
| MismatchType | 2 |  |
| AccessMode | 3 |  |
| Invalid | 4 |  |
| NotImplementedDeprecated | 5 |  |



<a name="sparkmax.paramType"></a>

### paramType
Parameter type enum, these values match the SPARK firmware
and are sent as a response in GetParameter() requests

| Name | Number | Description |
| ---- | ------ | ----------- |
| int32 | 0 |  |
| uint32 | 1 |  |
| float32 | 2 |  |
| bool | 3 |  |



<a name="sparkmax.sensorType"></a>

### sensorType
Sensor type enum, these values match the SPARK firmware

| Name | Number | Description |
| ---- | ------ | ----------- |
| NoSensor | 0 |  |
| HallSensor | 1 |  |
| Encoder | 2 |  |
| Sensorless | 3 |  |


 

 

 



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

