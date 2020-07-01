## Events

The `poa` module emits the following events:

### **EndBlock**

---

| Type                  | Attribute Key         | Attribute Value           |
| --------------------- | --------------------- | ------------------------- |
| create_validator	    | validator             | {validatorAddress}	      |
| edit_validator 	      | validator             | {validatorAddress}        |
| delete_validator	    | validator             | {validatorAddress}        |
| vote_validator	      | sender                | {senderAddress}	          |

### **Handlers**

---

### MsgCreateValidator

| Type             | Attribute Key | Attribute Value    |
| ---------------- | ------------- | ------------------ |
| create_validator | validator     | {validatorAddress} |
| message          | module        | poa            	  |
| message          | action        | create_validator   |
| message          | sender        | {senderAddress}    |

### MsgEditValidator

| Type           | Attribute Key       | Attribute Value     |
| -------------- | ------------------- | ------------------- |
| edit_validator | validator	         | {validatorAddress}  |
| message        | module              | poa                 |
| message        | action              | edit_validator      |
| message        | sender              | {senderAddress}     |

### MsgRemoveValidator

| Type             | Attribute Key       | Attribute Value       |
| ---------------- | ------------------- | ----------------------|
| delete_validator | validator	         | {validatorAddress}    |
| message          | module              | poa                   |
| message          | action              | delete_validator      |
| message          | sender              | {senderAddress}       |

### MsgCastVote

| Type             | Attribute Key       | Attribute Value       |
| ---------------- | ------------------- | ----------------------|
| vote_validator   | validator	         | {validatorAddress}    |
| message          | module              | poa                   |
| message          | action              | vote_validator        |
| message          | inFavor             | true                  |
| message          | sender              | {senderAddress}       |

