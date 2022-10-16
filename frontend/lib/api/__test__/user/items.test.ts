import { faker } from "@faker-js/faker";
import { describe, test, expect } from "vitest";
import { ItemField, LocationOut } from "../../types/data-contracts";
import { AttachmentTypes } from "../../types/non-generated";
import { UserClient } from "../../user";
import { factories } from "../factories";
import { sharedUserClient } from "../test-utils";

describe("user should be able to create an item and add an attachment", () => {
  let increment = 0;
  /**
   * useLocatio sets up a location resource for testing, and returns a function
   * that can be used to delete the location from the backend server.
   */
  async function useLocation(api: UserClient): Promise<[LocationOut, () => Promise<void>]> {
    const { response, data } = await api.locations.create({
      name: `__test__.location.name_${increment}`,
      description: `__test__.location.description_${increment}`,
    });
    expect(response.status).toBe(201);
    increment++;

    const cleanup = async () => {
      const { response } = await api.locations.delete(data.id);
      expect(response.status).toBe(204);
    };

    return [data, cleanup];
  }

  test("user should be able to create an item and add an attachment", async () => {
    const api = await sharedUserClient();
    const [location, cleanup] = await useLocation(api);

    const { response, data: item } = await api.items.create({
      name: "test-item",
      labelIds: [],
      description: "test-description",
      locationId: location.id,
    });
    expect(response.status).toBe(201);

    // Add attachment
    {
      const testFile = new Blob(["test"], { type: "text/plain" });
      const { response } = await api.items.addAttachment(item.id, testFile, "test.txt", AttachmentTypes.Attachment);
      expect(response.status).toBe(201);
    }

    // Get Attachment
    const { response: itmResp, data } = await api.items.get(item.id);
    expect(itmResp.status).toBe(200);

    expect(data.attachments).toHaveLength(1);
    expect(data.attachments[0].document.title).toBe("test.txt");

    const resp = await api.items.deleteAttachment(data.id, data.attachments[0].id);
    expect(resp.response.status).toBe(204);

    api.items.delete(item.id);
    await cleanup();
  });

  test("user should be able to create and delete fields on an item", async () => {
    const api = await sharedUserClient();
    const [location, cleanup] = await useLocation(api);

    const { response, data: item } = await api.items.create({
      name: faker.vehicle.model(),
      labelIds: [],
      description: faker.lorem.paragraph(1),
      locationId: location.id,
    });
    expect(response.status).toBe(201);

    const fields: ItemField[] = [
      factories.itemField(),
      factories.itemField(),
      factories.itemField(),
      factories.itemField(),
    ];

    // Add fields
    const itemUpdate = {
      ...item,
      locationId: item.location.id,
      labelIds: item.labels.map(l => l.id),
      fields,
    };

    const { response: updateResponse, data: item2 } = await api.items.update(item.id, itemUpdate);
    expect(updateResponse.status).toBe(200);

    expect(item2.fields).toHaveLength(fields.length);

    for (let i = 0; i < fields.length; i++) {
      expect(item2.fields[i].name).toBe(fields[i].name);
      expect(item2.fields[i].textValue).toBe(fields[i].textValue);
      expect(item2.fields[i].numberValue).toBe(fields[i].numberValue);
    }

    itemUpdate.fields = [fields[0], fields[1]];

    const { response: updateResponse2, data: item3 } = await api.items.update(item.id, itemUpdate);
    expect(updateResponse2.status).toBe(200);

    expect(item3.fields).toHaveLength(2);
    for (let i = 0; i < item3.fields.length; i++) {
      expect(item3.fields[i].name).toBe(itemUpdate.fields[i].name);
      expect(item3.fields[i].textValue).toBe(itemUpdate.fields[i].textValue);
      expect(item3.fields[i].numberValue).toBe(itemUpdate.fields[i].numberValue);
    }

    cleanup();
  });
});
