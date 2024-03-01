import { describe, expect, test } from "vitest";
import type { LabelOut } from "../../types/data-contracts";
import type { UserClient } from "../../user";
import { factories } from "../factories";
import { sharedUserClient } from "../test-utils";

describe("locations lifecycle (create, update, delete)", () => {
  /**
   * useLabel sets up a label resource for testing, and returns a function
   * that can be used to delete the label from the backend server.
   */
  async function useLabel(api: UserClient): Promise<[LabelOut, () => Promise<void>]> {
    const { response, data } = await api.labels.create(factories.label());
    expect(response.status).toBe(201);

    const cleanup = async () => {
      const { response } = await api.labels.delete(data.id);
      expect(response.status).toBe(204);
    };
    return [data, cleanup];
  }

  test("user should be able to create a label", async () => {
    const api = await sharedUserClient();

    const labelData = factories.label();

    const { response, data } = await api.labels.create(labelData);

    expect(response.status).toBe(201);
    expect(data.id).toBeTruthy();

    // Ensure we can get the label
    const { response: getResponse, data: getData } = await api.labels.get(data.id);

    expect(getResponse.status).toBe(200);
    expect(getData.id).toBe(data.id);
    expect(getData.name).toBe(labelData.name);
    expect(getData.description).toBe(labelData.description);

    // Cleanup
    const { response: deleteResponse } = await api.labels.delete(data.id);
    expect(deleteResponse.status).toBe(204);
  });

  test("user should be able to update a label", async () => {
    const api = await sharedUserClient();
    const [label, cleanup] = await useLabel(api);

    const labelData = {
      name: "test-label",
      description: "test-description",
      color: "",
    };

    const { response, data } = await api.labels.update(label.id, labelData);
    expect(response.status).toBe(200);
    expect(data.id).toBe(label.id);

    // Ensure we can get the label
    const { response: getResponse, data: getData } = await api.labels.get(data.id);
    expect(getResponse.status).toBe(200);
    expect(getData.id).toBe(data.id);
    expect(getData.name).toBe(labelData.name);
    expect(getData.description).toBe(labelData.description);

    // Cleanup
    await cleanup();
  });

  test("user should be able to delete a label", async () => {
    const api = await sharedUserClient();
    const [label, _] = await useLabel(api);

    const { response } = await api.labels.delete(label.id);
    expect(response.status).toBe(204);

    // Ensure we can't get the label
    const { response: getResponse } = await api.labels.get(label.id);
    expect(getResponse.status).toBe(404);
  });
});
