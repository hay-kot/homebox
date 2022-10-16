# CSV Imports

## Quick Start

Using the CSV import is the recommended way for adding items to the database. It is always going to be the fastest way to import any large amount of items and provides the most flexibility when it comes to adding items.

**Limitations**

 - Currently only supports importing items, locations, and labels
 - Does not support attachments. Attachments must be uploaded after import

**Template**

You can use this snippet as the headers for your CSV. Copy and paste it into your spreadsheet editor of choice and fill in the value.

```csv
Import RefLocation	Labels	Quantity	Name	Description	Insured	Serial Number	Model Number	Manufacturer	Notes	Purchase From	Purchased Price	Purchased Time	Lifetime Warranty	Warranty Expires	Warranty Details	Sold To	Sold Price	Sold Time	Sold Notes
```

!!! tip "Column Order"
    Column headers are just there for reference, the important thing is that the order is correct. You can change the headers to anything you like, this behavior may change in the future.


## CSV Reference

| Column            | Type                 | Description                                                                                                                                                                         |
| ----------------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ImportRef         | String (100)         | Import Refs are unique strings that can be used to deduplicate imports. Before an item is imported, we check the database for a matching ref. If the ref exists, we skip that item. |
| Location          | String               | This is the location of the item that will be created. These are de-duplicated and won't create another instance when reused.                                                       |
| Labels            | `;` Separated String | List of labels to apply to the item separated by a `;`, can be existing or new                                                                                                      |
| Quantity          | Integer              | The quantity of items to create                                                                                                                                                     |
| Name              | String               | Name of the item                                                                                                                                                                    |
| Description       | String               | Description of the item                                                                                                                                                             |
| Insured           | Boolean              | Whether or not the item is insured                                                                                                                                                  |
| Serial Number     | String               | Serial number of the item                                                                                                                                                           |
| Model Number      | String               | Model of the item                                                                                                                                                                   |
| Manufacturer      | String               | Manufacturer of the item                                                                                                                                                            |
| Notes             | String (1000)        | General notes about the product                                                                                                                                                     |
| Purchase From     | String               | Name of the place the item was purchased from                                                                                                                                       |
| Purchase Price    | Float64              |                                                                                                                                                                                     |
| Purchase At       | Date                 | Date the item was purchased                                                                                                                                                         |
| Lifetime Warranty | Boolean              | true or false - case insensitive                                                                                                                                                    |
| Warranty Expires  | Date                 | Date in the format                                                                                                                                                                  |
| Warranty Details  | String               | Details about the warranty                                                                                                                                                          |
| Sold To           | String               | Name of the person the item was sold to                                                                                                                                             |
| Sold At           | Date                 | Date the item was sold                                                                                                                                                              |
| Sold Price        | Float64              |                                                                                                                                                                                     |
| Sold Notes        | String (1000)        |                                                                                                                                                                                     |

**Type Key**

| Type    | Format                                              |
| ------- | --------------------------------------------------- |
| String  | Max 255 Characters unless otherwise specified       |
| Date    | YYYY-MM-DD                                          |
| Boolean | true or false, yes or no, 1 or 0 - case insensitive |
