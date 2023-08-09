import { fetchDatasetReview } from "../../../../../api/datasets.server";

export async function load({ params, cookies }: any) {
  if (cookies.get("accessToken")) {
    const orgid = cookies.get("orgId");
    const datasetName = params.datasetId;
    const accesstoken = cookies.get("accessToken");
    return {
      allCommits: await fetchDatasetReview(orgid, datasetName, accesstoken),
    };
  }
}
