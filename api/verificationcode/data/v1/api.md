# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/verificationcode/data/v1/data.proto](#api/verificationcode/data/v1/data.proto)
    - [GetLastCodeRequest](#juvo.verificationcode.data.v1.GetLastCodeRequest)
    - [GetNewCodeRequest](#juvo.verificationcode.data.v1.GetNewCodeRequest)
    - [InsertCodeRequest](#juvo.verificationcode.data.v1.InsertCodeRequest)
    - [InsertCodeResponse](#juvo.verificationcode.data.v1.InsertCodeResponse)
    - [VerificationCode](#juvo.verificationcode.data.v1.VerificationCode)
  
  
  
    - [VerificationCodeData](#juvo.verificationcode.data.v1.VerificationCodeData)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api/verificationcode/data/v1/data.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/verificationcode/data/v1/data.proto



<a name="juvo.verificationcode.data.v1.GetLastCodeRequest"></a>

### GetLastCodeRequest
Request message for VerificationCodeData.GetLastCode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| juvo_id | [string](#string) |  | The ID of the user for which to fetch a verification code |






<a name="juvo.verificationcode.data.v1.GetNewCodeRequest"></a>

### GetNewCodeRequest
Request message for VerificationCodeData.GetNewCode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| juvo_id | [string](#string) |  | The ID of the user for which to fetch a new verification code |






<a name="juvo.verificationcode.data.v1.InsertCodeRequest"></a>

### InsertCodeRequest
Request message for VerificationCodeData.GetLastCode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| juvo_id | [string](#string) |  | The ID of the user for which to fetch a verification code |
| verification_code | [string](#string) |  |  |






<a name="juvo.verificationcode.data.v1.InsertCodeResponse"></a>

### InsertCodeResponse
Request message for VerificationCodeData.GetLastCode






<a name="juvo.verificationcode.data.v1.VerificationCode"></a>

### VerificationCode
Response message for VerificationCodeData.GetLastCode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| juvo_id | [string](#string) |  | The ID of the user associated with this verification code |
| verification_code | [string](#string) |  | The verification code |





 

 

 


<a name="juvo.verificationcode.data.v1.VerificationCodeData"></a>

### VerificationCodeData


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetLastCode | [GetLastCodeRequest](#juvo.verificationcode.data.v1.GetLastCodeRequest) | [VerificationCode](#juvo.verificationcode.data.v1.VerificationCode) | Returns the last verification code generated for this user |
| InsertCode | [InsertCodeRequest](#juvo.verificationcode.data.v1.InsertCodeRequest) | [InsertCodeResponse](#juvo.verificationcode.data.v1.InsertCodeResponse) | Returns the last verification code generated for this user |

 



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

