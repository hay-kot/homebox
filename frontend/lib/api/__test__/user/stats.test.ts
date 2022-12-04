import { faker } from "@faker-js/faker";
import { beforeAll, describe, expect, test } from "vitest";
import { UserClient } from "../../user";
import { factories } from "../factories";

type ImportObj = {
  ImportRef: string;
  Location: string;
  Labels: string;
  Quantity: string;
  Name: string;
  Description: string;
  Insured: boolean;
  SerialNumber: string;
  ModelNumber: string;
  Manufacturer: string;
  Notes: string;
  PurchaseFrom: string;
  PurchasedPrice: number;
  PurchasedTime: string;
  LifetimeWarranty: boolean;
  WarrantyExpires: string;
  WarrantyDetails: string;
  SoldTo: string;
  SoldPrice: number;
  SoldTime: string;
  SoldNotes: string;
};

function toCsv(data: ImportObj[]): string {
  const headers = Object.keys(data[0]).join("\t");
  const rows = data.map(row => {
    return Object.values(row).join("\t");
  });
  return [headers, ...rows].join("\n");
}

function importFileGenerator(entries: number): ImportObj[] {
  const imports: ImportObj[] = [];

  const pick = (arr: string[]) => arr[Math.floor(Math.random() * arr.length)];

  const labels = faker.random.words(5).split(" ").join(";");
  const locations = faker.random.words(3).split(" ");

  const half = Math.floor(entries / 2);

  for (let i = 0; i < entries; i++) {
    imports.push({
      ImportRef: faker.database.mongodbObjectId(),
      Location: pick(locations),
      Labels: labels,
      Quantity: faker.random.numeric(1),
      Name: faker.random.words(3),
      Description: "",
      Insured: faker.datatype.boolean(),
      SerialNumber: faker.random.alphaNumeric(5),
      ModelNumber: faker.random.alphaNumeric(5),
      Manufacturer: faker.random.alphaNumeric(5),
      Notes: "",
      PurchaseFrom: faker.name.fullName(),
      PurchasedPrice: faker.datatype.number(100),
      PurchasedTime: faker.date.past().toDateString(),
      LifetimeWarranty: half > i,
      WarrantyExpires: faker.date.future().toDateString(),
      WarrantyDetails: "",
      SoldTo: faker.name.fullName(),
      SoldPrice: faker.datatype.number(100),
      SoldTime: faker.date.past().toDateString(),
      SoldNotes: "",
    });
  }

  return imports;
}

describe("group related statistics tests", () => {
  const TOTAL_ITEMS = 30;

  let api: UserClient | undefined;
  const imports = importFileGenerator(TOTAL_ITEMS);

  beforeAll(async () => {
    // -- Setup --
    const { client } = await factories.client.singleUse();
    api = client;

    const csv = toCsv(imports);

    const setupResp = await client.items.import(new Blob([csv], { type: "text/csv" }));

    expect(setupResp.status).toBe(204);
  });

  // Write to file system for debugging
  // fs.writeFileSync("test.csv", csv);
  test("Validate Group Statistics", async () => {
    const { status, data } = await api.stats.group();
    expect(status).toBe(200);

    expect(data.totalItems).toEqual(TOTAL_ITEMS);
    expect(data.totalLabels).toEqual(11); // default + new
    expect(data.totalLocations).toEqual(11); // default + new
    expect(data.totalUsers).toEqual(1);
    expect(data.totalWithWarranty).toEqual(Math.floor(TOTAL_ITEMS / 2));
  });

  const labelData: Record<string, number> = {};
  const locationData: Record<string, number> = {};

  for (const item of imports) {
    for (const label of item.Labels.split(";")) {
      labelData[label] = (labelData[label] || 0) + item.PurchasedPrice;
    }

    locationData[item.Location] = (locationData[item.Location] || 0) + item.PurchasedPrice;
  }

  test("Validate Labels Statistics", async () => {
    const { status, data } = await api.stats.labels();
    expect(status).toBe(200);

    for (const label of data) {
      expect(label.total).toEqual(labelData[label.name]);
    }
  });

  test("Validate Locations Statistics", async () => {
    const { status, data } = await api.stats.locations();
    expect(status).toBe(200);

    for (const location of data) {
      expect(location.total).toEqual(locationData[location.name]);
    }
  });

  test("Validate Purchase Over Time", async () => {
    const { status, data } = await api.stats.totalPriceOverTime();
    expect(status).toBe(200);
    expect(data.entries.length).toEqual(TOTAL_ITEMS);
  });
});
