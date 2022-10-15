const cache = {
  currency: "",
};

export function ResetCurrency() {
  cache.currency = "";
}

export async function useFormatCurrency() {
  if (!cache.currency) {
    const client = useUserApi();

    const { data: group } = await client.group.get();

    if (group) {
      cache.currency = group.currency;
    }
  }

  return (value: number | string) => fmtCurrency(value, cache.currency);
}
