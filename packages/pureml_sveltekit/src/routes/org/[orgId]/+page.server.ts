import { fetchDatasets } from "../../api/datasets.server";
import { fetchModels } from "../../api/models.server";

export async function load({ cookies }: any) {
  if (cookies.get("accessToken")) {
    const accesstoken = cookies.get("accessToken");
    const orgid = cookies.get("orgId");
    return {
      models: await fetchModels(orgid, accesstoken),
      datasets: await fetchDatasets(orgid, accesstoken),
    };
  }
}
