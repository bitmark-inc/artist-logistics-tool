# Artist Logistic Server

## Pre-requisite

- golang v1.20+
- node v16+

## Setup

Before set up, please ensure all pre-requisite is installed in you environment.

### Build

By run

```
make
```

- The binary will be located in `<project_root>/bin/` folder.
- The web will be located in `<project_root>/ui/dist`

### Configuration

- Create a new configuration file by copy from `<project_root>/services/logistic-server/config.yaml.sample` to `<project_root>/config.yaml`.
- The `ui_path` is the path for the server to serve static files. We need to set it to the generated webpage. (`ui/dist`)
- The `store.dsn` is the database URI for shippment information.

### DB Migration

After the `store.dsn` is set. We can run

```
make db-migrate
```

to initiate the database schema.

## Run

To start the server, we use

```
make run
```

## APIs

### `HEAD` /claim/:address

Check the claim status for a given address.

#### Header

- `LogisticID` - the logistic id a user is going to claim

#### URL Parameter

- `address` - the blockchain address for this claim

#### Response

- `204` if claimed
- `404` if not found

#### Example

```shell
$ curl -vvvvv --head -H 'LogisticID: test-001' 'http://localhost:8087/api/claim/0xBe680840C298Ad5aE6ea31c54a20C25424b5Bf45'
```

### `POST` /claim

#### Header

- `LogisticID` - the logistic id a user is going to claim
- `Requester` - the requester for a claim. It is an ethereum address here.
- `Authorization` - the signature of current timestamp for the requester address
    - `Bearer web3.eth.personal.sign(keccak256(timestamp_in_string)) + timestamp_string_in_hex`
        ```javascript
            let now = (+new Date()).toString();
            var hash = web3.utils.keccak256(now);
            let signature: string = await web3.eth.personal.sign(
                hash,
                accountNumber
            );
            authToken = signature.slice(2) + Buffer.from(now).toString("hex");
        ```

#### Request Body

| Key  | Description  |
| - | - |
| tokenID | the owned token id for claiming the right |
| information | the shipment information |

```json
{
    "tokenID": "",
    "information": {}
}
```

##### Shipment Information

```json
{
    "firstName": "",
    "lastName": "",
    "email": "",
    "address1": "",
    "address2": "",
    "postcode": "",
    "city": "",
    "state": "",
    "country": "",
    "phoneCountryCode": "",
    "phoneNumber": ""
}
```

#### Response

- `200` if submitted

#### Example

```
curl -vvvvv --head -H 'LogisticID: test-001' http://localhost:8087/api/claim/0xBe680840C298Ad5aE6ea31c54a20C25424b5Bf45
```
