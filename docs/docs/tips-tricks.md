# Tips and Tricks

## Custom Fields

Custom fields are a great way to add any extra information to your item. The following types are supported:

- [x] Text
- [ ] Integer (Future)
- [ ] Boolean (Future)
- [ ] Timestamp (Future)

Custom fields are appended to the main details section of your item.

!!! tip
    Homebox Custom Fields also have special support for URLs. Provide a URL (`https://google.com`) and it will be automatically converted to a clickable link in the UI. Optionally, you can also use markdown syntax to add a custom text to the button. `[Google](https://google.com)`

## Managing Asset IDs

Homebox provides the option to auto-set asset IDs, this is the default behavior. These can be used for tracking assets with printable tags or labels. You can disable this behavior via a command line flag or ENV variable. See [configuration](../quick-start#env-variables-configuration) for more details.

Example ID: `000-001`

Asset IDs are partially managed by Homebox, but have a flexible implementation to allow for unique use cases. ID's are non-unique at the database level so there is nothing stopping a user from manually setting duplicate IDs for various items. There are two recommended approaches to manage Asset IDs

### 1. Auto Incrementing IDs

This is the default behavior and likely to one to experience the most consistent behavior. Whenever creating or importing an item, that items receives the next available ID. This is the most consistent approach and is recommended for most users.

### 2. Auto Incrementing ID's with Reset

In some cases you may want to skip some items such as consumables, or items that are loosely tracked. In this case, we recommend that you leave auto-incrementing ID's enabled _however_ when you create a new item that you want to skip, you can go to that item and reset the ID to 0. This will remove it from the auto-incrementing sequence and the next item will receive the next available ID.

!!! tip
    If you're migrating from an older version there is a action on the users profile page to assign IDs to all items. This will assign the next available ID to all items in the order of creation. You should _only_ do this once during the migration process. You should be especially cautious of this action if you're using the reset feature described in option number 2

## QR Codes

:octicons-tag-24: 0.7.0

Homebox has a built-in QR code generator that can be used to generate QR codes for your items. This is useful for tracking items with a mobile device. You can generate a QR code for any item by clicking the QR code icon in the top right of the item details page. The same can be done for the Labels and Locations page. Currently support is limited to generating one off QR Codes.

However, the API endpoint is available for generating QR codes on the fly for any item (or any other data) if you provide a valid API key in the query parameters. An example url would look like `/api/v1/qrcode?data=https://homebox.fly.dev/item/{uuid}&access_token={api_key}`. Currently the easiest way to get an API token is to use one from an existing URL of the QR Code in the API key, but this will be improved in the future.

:octicons-tag-24: 0.8.0

In version 0.8.0 We've added a custom label generation. On the tools page, there is now a link to the label-generator page where you can generate labels based on Asset ID for your inventory. These are still in early development, so please provide feedback. There's also more information on the implementation on the label generator page.

[Demo](https://homebox.fly.dev/reports/label-generator)