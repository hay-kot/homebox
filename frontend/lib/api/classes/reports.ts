import { BaseAPI, route } from "../base";

export class ReportsAPI extends BaseAPI {
  async billOfMaterials(): Promise<void> {
    const { data: stream, error } = await this.http.get<ReadableStream>({ url: route("/reporting/bill-of-materials") });

    if (error) {
      return;
    }

    const reader = stream.getReader();
    let data = "";
    while (true) {
      const { done, value } = await reader.read();
      if (done) {
        break;
      }
      data += new TextDecoder("utf-8").decode(value);
    }

    const blob = new Blob([data], { type: "text/tsv" });
    const link = document.createElement("a");
    link.href = window.URL.createObjectURL(blob);
    link.download = "bill-of-materials.tsv";
    link.click();
  }
}
