import { describe, expect, test } from "vitest";
import type { LocationOut } from "../../types/data-contracts";
import type { UserClient } from "../../user";
import { factories } from "../factories";
import { sharedUserClient } from "../test-utils";

describe("locations lifecycle (create, update, delete)", () => {
  /**
   * useLocatio sets up a location resource for testing, and returns a function
   * that can be used to delete the location from the backend server.
   */
  async function useLocation(api: UserClient): Promise<[LocationOut, () => Promise<void>]> {
    const { response, data } = await api.locations.create(factories.location());
    expect(response.status).toBe(201);

    const cleanup = async () => {
      const { response } = await api.locations.delete(data.id);
      expect(response.status).toBe(204);
    };

    return [data, cleanup];
  }

  test("user should be able to create a location", async () => {
    const api = await sharedUserClient();

    const locationData = factories.location();

    const { response, data } = await api.locations.create(locationData);

    expect(response.status).toBe(201);
    expect(data.id).toBeTruthy();

    // Ensure we can get the location
    const { response: getResponse, data: getData } = await api.locations.get(data.id);

    expect(getResponse.status).toBe(200);
    expect(getData.id).toBe(data.id);
    expect(getData.name).toBe(locationData.name);
    expect(getData.description).toBe(locationData.description);

    // Cleanup
    const { response: deleteResponse } = await api.locations.delete(data.id);
    expect(deleteResponse.status).toBe(204);
  });

  test("user should be able to update a location", async () => {
    const api = await sharedUserClient();
    const [location, cleanup] = await useLocation(api);

    const updateData = {
      id: location.id,
      parentId: location.parent?.id,
      name: "test-location-updated",
      description: "test-description-updated",
    };

    const { response } = await api.locations.update(location.id, updateData);
    expect(response.status).toBe(200);

    // Ensure we can get the location
    const { response: getResponse, data } = await api.locations.get(location.id);
    expect(getResponse.status).toBe(200);

    expect(data.id).toBe(location.id);
    expect(data.name).toBe(updateData.name);
    expect(data.description).toBe(updateData.description);

    await cleanup();
  });

  test("user should be able to delete a location", async () => {
    const api = await sharedUserClient();
    const [location, _] = await useLocation(api);

    const { response } = await api.locations.delete(location.id);
    expect(response.status).toBe(204);

    // Ensure we can't get the location
    const { response: getResponse } = await api.locations.get(location.id);
    expect(getResponse.status).toBe(404);
  });
});
