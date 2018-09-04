![get-my-invoices](https://www.getmyinvoices.com/wp-content/uploads/2016/04/logo_login.png)

# gogmi is a golang-library for getmyinvoices.com

Here you can find the API-documentation: https://api.getmyinvoices.com/accounts/v1/doc/#

## Already implemented

✔ list supplieres  
✔ get specific supplier  
✔ list invoices  
✘ upload new invoice  
✘ update invoice  
✔ get country list  
✘ add custom supplier  
✘ update custom supplier  
✘ delete custom supplier  

## Getting started

```golang
client := gogmi.GMI{
    APIVersion: "v1",
    APIKey:     "your-API-Key",
}
suppliers, err := client.ListSuppliers()
if err != nil {
    t.Error(err)
}

// do something with suppliers
```

