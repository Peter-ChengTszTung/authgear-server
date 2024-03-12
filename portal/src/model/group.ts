import { LocalValidationError } from "../error/validation";
import { Group } from "../graphql/adminapi/globalTypes.generated";

export interface CreatableGroup
  extends Pick<Group, "key" | "name" | "description"> {}

// Ref: pkg/lib/rolesgroups/key.go
const KEY_REGEX = /^[a-zA-Z_][a-zA-Z0-9:_]*$/;
const MAX_KEY_LENGTH = 40;

export function validateGroup(
  rawInput: CreatableGroup
): [CreatableGroup, LocalValidationError[]] {
  const input = sanitizeGroup(rawInput);
  const errors: LocalValidationError[] = [];
  if (!input.key) {
    errors.push({
      location: "/key",
      messageID: "errors.validation.required",
    });
  } else if (!KEY_REGEX.test(input.key)) {
    errors.push({
      location: "/key",
      messageID: "errors.groups.key.validation.format",
    });
  } else if (input.key.length > MAX_KEY_LENGTH) {
    errors.push({
      location: "/key",
      messageID: "errors.validation.maxLength",
      arguments: { expected: MAX_KEY_LENGTH },
    });
  }

  if (!input.name) {
    errors.push({
      location: "/name",
      messageID: "errors.validation.required",
    });
  }
  return [input, errors];
}

export function sanitizeGroup(input: CreatableGroup): CreatableGroup {
  const description = input.description?.trim();
  const name = input.name?.trim();
  return {
    key: input.key.trim(),
    name: name ? name : null,
    description: description ? description : null,
  };
}

export interface SearchableGroup extends Pick<Group, "id" | "key" | "name"> {}

export function searchGroups<G extends SearchableGroup>(
  groups: G[],
  searchKeyword: string
): G[] {
  if (searchKeyword === "") {
    return groups;
  }
  const keywords = searchKeyword
    .toLowerCase()
    .split(" ")
    .flatMap((keyword) => {
      const trimmedKeyword = keyword.trim();
      if (trimmedKeyword) {
        return [trimmedKeyword];
      }
      return [];
    });
  return groups.filter((group) => {
    const groupID = group.id.toLowerCase();
    const groupKey = group.key.toLowerCase();
    const groupName = group.name?.toLowerCase();
    for (const keyword of keywords) {
      if (groupID === keyword) {
        return true;
      }
      if (groupKey.includes(keyword)) {
        return true;
      }
      if (groupName?.includes(keyword)) {
        return true;
      }
    }
    return false;
  });
}
