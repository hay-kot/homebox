function slugify(text: string) {
  return text
    .toString()
    .toLowerCase()
    .replace(/\s+/g, "-") // Replace spaces with -
    .replace(/[^\w-]+/g, "") // Remove all non-word chars
    .replace(/--+/g, "-") // Replace multiple - with single -
    .replace(/^-+/, "") // Trim - from start of text
    .replace(/-+$/, ""); // Trim - from end of text
}

function idGenerator(): string {
  const id = Math.random().toString(32).substring(2, 6) + Math.random().toString(36).substring(2, 6);
  return slugify(id);
}

/**
 * useFormIds uses the provided label to generate a unique id for the
 * form element. If no label is provided the id is generated using a
 * random string.
 */
export function useFormIds(label: string): string {
  const slug = label ? slugify(label) : idGenerator();
  return `${slug}-${idGenerator()}`;
}

export function useId(): string {
  return idGenerator();
}
