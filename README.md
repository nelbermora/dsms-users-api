# DPMS Backend API
Now there is three endpoints

GET ->    /company Shows Information about current logged Company
GET ->    /company/{companyId}/branches Lists all branches of determined Company
GET ->    /branches/{branchId} Shows Information about determined Branch
GET ->    /branches/{branchId}/persons Lists all persons of determined Branch
GET ->    /persons/{personId} Shows User required if exists
POST ->   /persons Creates a new Person
PUT ->    /persons Updates the posted Person
DELETE -> /persons/{personId} Deletes a determined Person


This endpoints allows to Forntend, navigation between company/branch offices/users

## Entities Examples:

### Company:
```json
{
  "id": 1,
  "name": "Testing Company",
  "street": "Baker Street",
  "streetNumber": 221,
  "zipCode": "NW16XE",
  "city": "New Jersey",
  "state": "Demo",
  "country": "Demo Republic",
  "phone": "+54911-123-123",
  "fax": "",
  "email": "company@testing.eu"
}
```
### Branch Office:
```json
{
    "id": 2,
    "companyId": 1,
    "zipCode": "ASXQWE",
    "city": "London",
    "street": "Sheen Street",
    "streetNumber": 774,
    "state": "Demo",
    "country": "Demo Republic",
    "isHeadQuarter": false
 }
```
### Person:
```json
 {
    "id": 1,
    "name": "John Doe",
    "position": "CEO",
    "department": "T&O",
    "street": "LongChamps",
    "streetNumber": 123,
    "zipCode": "CP1022",
    "city": "Cork",
    "phone": "041-123123",
    "mobilePhone": "+123-123-123",
    "fax": "",
    "email": "john@doe.com",
    "branchId": 1,
    "login": "2Factor",
    "status": "Active"
 }
```
