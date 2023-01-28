import { Ref } from "vue";
import { TreeItem } from "~~/lib/api/types/data-contracts";

export interface FlatTreeItem {
  id: string;
  name: string;
  display: string;
}

export function flatTree(tree: TreeItem[]): Ref<FlatTreeItem[]> {
  const v = ref<FlatTreeItem[]>([]);

  // turns the nested items into a flat items array where
  // the display is a string of the tree hierarchy separated by breadcrumbs

  function flatten(items: TreeItem[], display: string) {
    for (const item of items) {
      v.value.push({
        id: item.id,
        name: item.name,
        display: display + item.name,
      });
      if (item.children) {
        flatten(item.children, display + item.name + " > ");
      }
    }
  }

  flatten(tree, "");

  return v;
}

export async function useFlatLocations(): Promise<Ref<FlatTreeItem[]>> {
  const api = useUserApi();

  const locations = await api.locations.getTree();

  if (!locations) {
    return ref([]);
  }

  return flatTree(locations.data.items);
}
