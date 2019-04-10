# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/verificationcode/v1/verificationcode.proto](#api/verificationcode/v1/verificationcode.proto)
    - [SendCodeRequest](#juvo.verification_code.v1.SendCodeRequest)
    - [SendCodeResponse](#juvo.verification_code.v1.SendCodeResponse)
    - [VerifyCodeRequest](#juvo.verification_code.v1.VerifyCodeRequest)
    - [VerifyCodeResponse](#juvo.verification_code.v1.VerifyCodeResponse)
  
  
  
    - [VerificationCode](#juvo.verification_code.v1.VerificationCode)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api/verificationcode/v1/verificationcode.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/verificationcode/v1/verificationcode.proto



<a name="juvo.verification_code.v1.SendCodeRequest"></a>

### SendCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| juvo_id | [string](#string) |  | The ID of the user to send the verification code |






<a name="juvo.verification_code.v1.SendCodeResponse"></a>

### SendCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| verification_code | [string](#string) |  |  |






<a name="juvo.verification_code.v1.VerifyCodeRequest"></a>

### VerifyCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| juvo_id | [string](#string) |  | The ID of the user to send the verification code |
| verification_code | [string](#string) |  | The verification code to test |






<a name="juvo.verification_code.v1.VerifyCodeResponse"></a>

### VerifyCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| verified | [bool](#bool) |  | True if the code is correct (just return 200?) |





 

 

 


<a name="juvo.verification_code.v1.VerificationCode"></a>

### VerificationCode


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendCode | [SendCodeRequest](#juvo.verification_code.v1.SendCodeRequest) | [SendCodeResponse](#juvo.verification_code.v1.SendCodeResponse) | Send the user a verification code. Returns xxx if the user has exceeded the limit |
| VerifyCode | [VerifyCodeRequest](#juvo.verification_code.v1.VerifyCodeRequest) | [VerifyCodeResponse](#juvo.verification_code.v1.VerifyCodeResponse) | Verify that the code is correct |

 



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

