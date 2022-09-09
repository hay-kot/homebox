# CSV Imports

This document outlines the CSV import feature of Homebox and how to use it.

## Quick Start

Using the CSV import is the recommended way for adding items to the database. It is always going to be the fastest way to import any large amount of items and provides the most flexibility when it comes to adding items.

**Limitations**
 - Currently only supports importing items
 - Does not support attachments. Attachments must be uploaded after import

## CSV Reference

| Column Heading    | Type                 | Description                                                                                                                   |
| ----------------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| Location          | String               | This is the location of the item that will be created. These are de-duplicated and won't create another instance when reused. |
| Labels            | `;` Separated String | List of labels to apply to the item seperated by a `;`, can be existing or new                                                |
| Quantity          | Integer              | The quantity of items to create                                                                                               |
| Name              | String               | Name of the item                                                                                                              |
| Description       | String               | Description of the item                                                                                                       |
| Serial Number     | String               | Serial number of the item                                                                                                     |
| Model Number      | String               | Model of the item                                                                                                             |
| Manufacturer      | String               | Manufacturer of the item                                                                                                      |
| Notes             | String               | General notes about the product                                                                                               |
| Purchase From     | String               | Name of the place the item was purchased from                                                                                 |
| Purchase Price    | Float64              |                                                                                                                               |
| Purchase At       | Date                 | Date the item was purchased                                                                                                   |
| Lifetime Warranty | Boolean              | true or false - case insensitive                                                                                              |
| Warranty Expires  | Date                 | Date in the format                                                                                                            |
| Warranty Details  | String               | Details about the warranty                                                                                                    |
| Sold To           | String               | Name of the person the item was sold to                                                                                       |
| Sold At           | Date                 | Date the item was sold                                                                                                        |
| Sold Price        | Float64              |                                                                                                                               |

**Type Key**

| Type   | Format             |
| ------ | ------------------ |
| String | Max 255 Characters |
| Date   | YYYY-MM-DD         |
