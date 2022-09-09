# authn 

WIP

A Windows WebAuthN api warpper. Provide both friendly api and low-level api. 

Target `WebAuthN API` version: 4 

## Usage

We provide a simple cli demo to let you try the basic functions of the `WebAuthN`.

In this demo, we use `go.webauthn.demo.app` as [Relying Party ID](https://w3c.github.io/webauthn/#rp-id).

**Note**: No information will be uploaded. All operation is processing locally.

``` 
PS C:\Users\ink33\authn> .\authn.exe
WebAuthN API Version: 4
Is User Verifying Platform Authenticator Available: true
Select operation:
1: Make Credential
2: Get Assertion
3: Get Platform Credential List
0: Exit
```

New function will be added soon.