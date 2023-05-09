import { fetchDatasetByName, fetchDatasets } from "../../../../api/datasets.server";

export async function load({ params, cookies }: any) {
  if (cookies.get("accessToken")) {
    const accesstoken = cookies.get("accessToken");
    const orgid = cookies.get("orgId");
    const datasetid = params.datasetId;
    return {
      datasets: await fetchDatasets(orgid, accesstoken),
      datasetDetails: await fetchDatasetByName(orgid, datasetid, accesstoken),
    };
  }
}
