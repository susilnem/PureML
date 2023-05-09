import { fetchDatasets } from "../api/datasets.server";

export async function load({ cookies }: any) {
  if (cookies.get("accessToken")) {
    const accesstoken = cookies.get("accessToken");
    const orgid = cookies.get("orgId");
    return {
      datasets: await fetchDatasets(orgid, accesstoken),
    };
  }
}
