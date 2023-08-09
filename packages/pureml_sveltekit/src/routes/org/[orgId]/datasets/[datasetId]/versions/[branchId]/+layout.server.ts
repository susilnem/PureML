import {
  fetchDatasetBranch,
  fetchDatasetVersions,
} from "../../../../../../api/datasets.server";

export async function load({ params, cookies }: any) {
  if (cookies.get("accessToken")) {
    const orgid = cookies.get("orgId");
    const datasetid = params.datasetId;
    const branchid = params.branchId;
    const accesstoken = cookies.get("accessToken");
    return {
      allBranches: await fetchDatasetBranch(orgid, datasetid, accesstoken),
      versions: await fetchDatasetVersions(
        orgid,
        datasetid,
        branchid,
        accesstoken
      ),
    };
  }
}
