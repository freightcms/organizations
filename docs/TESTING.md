# FreightCMS

## Organization Web Application

### Testing

### Creating an organization resource

```GraphQL
mutation {
  createOrganization(dba: "first org", name: "first org", mailingAddress: {
    line1: "123 Fake St.",
    line2: "AB Section",
    line3: "Floor 3",
    region: "Minnesota",
    locale: "Minneapolis",
    countryCode: "US",
    postalCode: "44444"
  }, billingAddress: {
       line1: "123 Fake St.",
    line2: "AB Section",
    line3: "Floor 3",
    region: "Minnesota",
    locale: "Minneapolis",
    countryCode: "US",
    postalCode: "44444"
  }) {
    id
  }
}
```
