import { faker } from "@faker-js/faker";
import { LabelCreate, LocationCreate, UserRegistration } from "../../types/data-contracts";

/**
 * Returns a random user registration object that can be
 * used to signup a new user.
 */
function user(): UserRegistration {
  return {
    email: faker.internet.email(),
    password: faker.internet.password(),
    name: faker.name.firstName(),
    token: "",
  };
}

function location(): LocationCreate {
  return {
    name: faker.address.city(),
    description: faker.lorem.sentence(),
  };
}

function label(): LabelCreate {
  return {
    name: faker.lorem.word(),
    description: faker.lorem.sentence(),
    color: faker.internet.color(),
  };
}

export const factories = { user, location, label };
