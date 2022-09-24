import { describe, test, expect } from "vitest";
import { LocationOut } from "../../types/data-contracts";
import { AttachmentTypes } from "../../types/non-generated";
import { UserApi } from "../../user";
import { sharedUserClient } from "../test-utils";

describe("user should be able to create an item and add an attachment", () => {
  let increment = 0;
  /**
   * useLocatio sets up a location resource for testing, and returns a function
   * that can be used to delete the location from the backend server.
   */
  async function useLocation(api: UserApi): Promise<[LocationOut, () => Promise<void>]> {
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
});
