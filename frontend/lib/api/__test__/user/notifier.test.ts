import { faker } from "@faker-js/faker";
import { describe, expect, test } from "vitest";
import { factories } from "../factories";

describe("basic notifier workflows", () => {
  test("user should be able to create, update, and delete a notifier", async () => {
    const { client } = await factories.client.singleUse();

    // Create Notifier
    const result = await client.notifiers.create({
      name: faker.word.words(2),
      url: "discord://" + faker.string.alphanumeric(10),
      isActive: true,
    });

    expect(result.error).toBeFalsy();
    expect(result.status).toBe(201);
    expect(result.data).toBeTruthy();

    const notifier = result.data;

    // Update Notifier with new URL
    {
      const updateData = {
        name: faker.word.words(2),
        url: "discord://" + faker.string.alphanumeric(10),
        isActive: true,
      };

      const updateResult = await client.notifiers.update(notifier.id, updateData);
      expect(updateResult.error).toBeFalsy();
      expect(updateResult.status).toBe(200);
      expect(updateResult.data).toBeTruthy();
      expect(updateResult.data.name).not.toBe(notifier.name);
    }

    // Update Notifier with empty URL
    {
      const updateData = {
        name: faker.word.words(2),
        url: null,
        isActive: true,
      };

      const updateResult = await client.notifiers.update(notifier.id, updateData);
      expect(updateResult.error).toBeFalsy();
      expect(updateResult.status).toBe(200);
      expect(updateResult.data).toBeTruthy();
      expect(updateResult.data.name).not.toBe(notifier.name);
    }

    // Delete Notifier
    {
      const deleteResult = await client.notifiers.delete(notifier.id);
      expect(deleteResult.error).toBeFalsy();
      expect(deleteResult.status).toBe(204);
    }
  });
});
